# Install Task Service in Kuberbetes
_This document is inteded to guide you through the deployment process in Kubernetes of the Task Service Application_

### Steps
- Create container from Application
- Push container to Dockerhub
- Create a deployment in k8s
- expose the service

#### Create a container of the application
You can build a container of the task-service application by running the command below:
```
docker build -t docker.io/cmwylie19/task-service .
```

#### Log into Dockerhub and push the container
Log into Dockerhub
```
docker login
```

Push the container
```
docker push docker.io/cmwylie19/task-service
```

#### Create a deployment in k8s
```
kubectl create deployment task-service --image=docker.io/cmwylie19/task-service:latest 
```

#### Expose a service in k8s
```
kubectl expose deployment task-service --type=LoadBalancer --port=3000   
```

#### Clean up
```
kubectl delete svc,deployment task-service
```
