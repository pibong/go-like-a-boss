## Create new project on gitLab
* `git clone git@gitlab.com:foo/go-like-a-boss.git`

## Create namespace, service account and ClusterRole in kuberenetes clust
* `vim prepare-k8s-cluster.yml` [prepare-k8s-cluster.yml](/prepare-k8s-cluster.yml)
* `kubectl create -f prepare-k8s-cluster.yml`

## Create a secret to authenticate against GitLab registry
* `kubectl create secret docker-registry gitlabreg --docker-server=registry.gitlab.com --docker-username=foo --docker-password=************ --docker-email=piero.bongiovanni@gmail.com --namespace=go-like-boss-project`
* `kubectl get pods,svc,deployment,rs,secrets,sa -n go-like-boss-project`

```
NAME                         TYPE                                  DATA      AGE
secret/default-token-nnftg   kubernetes.io/service-account-token   3         5m
secret/gitlab-token-sf7sn    kubernetes.io/service-account-token   3         5m
secret/gitlabreg             kubernetes.io/dockerconfigjson        1         40s

NAME                     SECRETS   AGE
serviceaccount/default   1         5m
serviceaccount/gitlab    1         5m
```

## Add kubernetes cluster to the project
* Kubernetes cluster name: gke-cluster-dev
* API endpoint: `kubectl get endpoints | grep kubernetes | awk '{print $2}' | cut -d':' -f1`
* CA Certificate: `kubectl config view --minify=true --flatten | grep certificate-authority-data | awk '{print $2}' | base64 -d`
* Token: `kubectl --namespace=go-like-boss-project describe secret gitlab-token-sf7sn | grep "token:" | awk '{print $2}'`
* Project namespace: go-like-boss-project


## Add files in GitLab repository
* Entry point for the dummy application [main.go](/main.go)	
* Dummy library [mylib](/mylib/multiply.go)
* Unit test file *_test.go* [mylib/multiply_test.go](/mylib/multiply_test.go)
* Script to check teh code coverage [tools/code_coverage.sh](/tools/code_coverage.sh)
* GitLab CI pipeline file [.gitlab-ci.yml](.gitlab-ci.yml)
* Ignore files for git [.gitignore](/.gitignore)
* Dockerfile for building image used by gitlab-runner [Dockerfile](/docker/Dockerfile)

## Commit files
* `git add .gitignore .gitlab-ci.yml Makefile README.md main.go mylib/ prepare-k8s-cluster.yml tools/`
* `$ git status`
* `git commit -m "first commit"`
* `git push`

## Let's do the magic to GitLab and K8S

#### Docker image foo/gocompile_env
**Build a Docker image which is ready for go compiling, testing and code linting**

* `cd docker; docker build --tag 'gocompile_env:0.1' .`
* `docker login https://index.docker.io`
* `docker image tag gocompile_env:0.1 foo/gocompile_env:0.1`
* `docker image push foo/gocompile_env`

