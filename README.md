# Task Service
_This lab is intended to explain how to manage a simple task api using Gloo Edge._

## Background
Gloo Edge is an API Gateway, and Kubernetes Ingress Controller that is packed with functionality. 

## Prereqs
- kubectl
- install `glooctl`
- Docker
- Gloo Edge 
- Kubernetes v1.11.3+ (I am user Docker Desktop with Kubernetes Enabled)


### Install glooctl
To install glooctl locally, execute the command below which downloads the binary and installs it in your path.
```
curl -sL https://run.solo.io/gloo/install | sh
export PATH=$HOME/.gloo/bin:$PATH
```
### Install Gloo Edge
Install Gloo Edge using `glooctl`
```
glooctl install gateway
```

## Build Image and Deploy to Dockerhub
```
docker build -t docker.io/cmwylie19/task-service .
docker push docker.io/cmwylie19/task-service
```
## Usage

### Test Health

```
curl http://localhost:3000/healthz
```

### Create Task

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"test"}' http://localhost:3000/create
```

### Get Tasks

```
curl http://localhost:3000/
```

### Update a task

```
curl -X PUT -H "Content-Type: application/json" -d '{"name":"tester","complete":"true"}' http://localhost:3000/d29e6d58525
```

### Delete a task

```
curl -X DELETE http://localhost:3000/d29e6d58525
```


### Get a task

```
curl http://localhost:3000/d29e6d58525
```


#### Clean Up
Delete the virtual service
```
kubectl delete vs default -n gloo-system
```

Delete the upstream
```
kubectl delete upstream -n gloo-system default-task-service-8080
```

Delete the task-service Deployment and Service
```
kubectl delete svc,deployment task-service
```
