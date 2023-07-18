# GO Ping Pong Example

Dockerized go api packaged using helm and deployed using **GitOps** approach through ArgoCD. 

## setup

### Dependencies  

create `docker-secret.yaml` containes docker cred that has acess to the repo 

```YAML
apiVersion: v1
kind: Secret
metadata:
name: repo-secret-api
namespace: argocd
type: kubernetes.io/dockerconfigjson
stringData:
.dockerconfigjson: |
    {
    "auths": {
    "https://registry-1.docker.io/v2": {
        "auth": "<DOCKER_LOGIN_TOKEN>"
        }
        }
    }
```

create `repo.yaml` contains argocd repo secret with write access to this repo

```YAML
apiVersion: v1
kind: Secret
metadata:
  name: argocd-repo
  namespace: argocd
  labels:
    argocd.argoproj.io/secret-type: repository
type: Opaque
stringData:
  name: github
  password: <GIT_TOKEN>
  project: default
  type: git
  url: https://github.com/anieo/go-example.git
  username: anieo
```

and update you github action secrets 

* DOCKERHUB_USERNAME : Docker Hub username
* DOCKERHUB_TOKEN : Docker Hub Password

### INFRA

you can deploy a kind cluster locally.

```BASH
$ ./playground/deploy

ali@default::devops-go-api (main)$ ./playground/deploy 
Starting Kind Cluster
Creating cluster "go-cluster" ...
 ✓ Ensuring node image (kindest/node:v1.26.3) 🖼
 ✓ Preparing nodes 📦 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
 ✓ Joining worker nodes 🚜 
Set kubectl context to "kind-go-cluster"
You can now use your cluster with:

kubectl cluster-info --context kind-go-cluster

Thanks for using kind! 😊
deploying test nodes
"argo" already exists with the same configuration, skipping
namespace/argocd created
Release "argocd" does not exist. Installing it now.
NAME: argocd
LAST DEPLOYED: Tue Jul 18 07:33:55 2023
NAMESPACE: argocd
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
In order to access the server UI you have the following options:

1. kubectl port-forward service/argocd-server -n argocd 8080:443

    and then open the browser on http://localhost:8080 and accept the certificate

2. enable ingress in the values file `server.ingress.enabled` and either
      - Add the annotation for ssl passthrough: https://argo-cd.readthedocs.io/en/stable/operator-manual/ingress/#option-1-ssl-passthrough
      - Set the `configs.params."server.insecure"` in the values file and terminate SSL at your ingress: https://argo-cd.readthedocs.io/en/stable/operator-manual/ingress/#option-2-multiple-ingress-objects-and-hosts


After reaching the UI the first time you can login with username: admin and the random password generated during the installation. You can find the password by running:

kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

(You should delete the initial secret afterwards as suggested by the Getting Started Guide: https://argo-cd.readthedocs.io/en/stable/getting_started/#4-login-using-the-cli)
secret/repo-secret-api created
application.argoproj.io/go-example created
application.argoproj.io/argocd-image-updater created
secret/argocd-repo created
ARGOCD ADMIN USERNAME : admin 
ARGOCD ADMIN PASSWORD: R8ld5rkLVf1GUG8k

```

### Infra CleanUp

```BASH
$./playground/delete

Deleting
Deleting cluster "go-cluster" ...
Deleted nodes: ["go-cluster-control-plane" "go-cluster-worker"]
error: current-context is not set
```