# Deploy Application in Kubernetes
_This scenario is created to show you how to deploy the task-service application in Kubernetes. It will take you from building a docker container to exposing the appliaction to the outside internet and clean it up when you are done._

## Prereqs
- Docker
- Kuberentes 
- kubectl

_I am using Docker Desktop with kubernetes enabled._
```
$ kubectl version
Client Version: version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.0", GitCommit:"af46c47ce925f4c4ad5cc8d1fca46c7b77d13b38", GitTreeState:"clean", BuildDate:"2020-12-13T19:50:45Z", GoVersion:"go1.15.5", Compiler:"gc", Platform:"darwin/amd64"}
Server Version: version.Info{Major:"1", Minor:"19", GitVersion:"v1.19.3", GitCommit:"1e11e4a2108024935ecfcb2912226cedeafd99df", GitTreeState:"clean", BuildDate:"2020-10-14T12:41:49Z", GoVersion:"go1.15.2", Compiler:"gc", Platform:"linux/amd64"}
```
## Steps
- Build Container
- Deploy Container to DockerHub
- Deploy in Kubernetes using `kubectl`
- Expose a service in Kubernetes using `kubectl`
- Deploy the application from a yaml manifest

