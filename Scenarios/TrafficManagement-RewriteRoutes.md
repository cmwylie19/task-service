# Traffic Management: Rewriting Routes
_During this scenario we will explore Traffic Management features of Gloo Edge to use write routing rules for Virtual Services that control how traffic gets to your application. It is assumed you are running a Kubernetes cluster in a linux environment._

## Background
Gloo Edge automatically creates Virtual Services and Upstreams.  


***Virtual Services*** define a set of route rules that live under a domain or set of domains. Route rules consist of matchers, which specify the kind of function calls to match (requests and events, are currently supported), and the name of the destination (or destinations) where to route them.

***Upstreams*** Define destinations for routes. Upstreams tell Gloo Edge what to route to.

Typically, the upstream will be your application or Kubernetes Service, and each Service will have a Virtual Service to control the flow of traffic.

## Steps
- install `glooctl`
- install Gloo Edge API Gateway
- deploy the task-service application into Kubernetes
- Create rewrite rules in your Virtual Service to control traffic to your applications using `glooctl`
- Export Virtual Services to yaml for quick deployment and maintaining good infrastructure as code.
- Clean up

### Install `glooctl`
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

### Deploy task-service application to Kubernetes
```
kubectl apply -f k8s/task-service.yaml
```
### Rewrite routes in a Virtual Service using `glooctl`
_Gloo Edge watches for new services to be added to the Kubernetes Cluster. When the application is created, Gloo Edge _automatically_ creates an Upstream for the Kubernetes Service. Remember, upstream is the destination for the routes and in this example the upstream is the task-service itself. Gloo Edge also creates Virtual Services for the Kubernetes Services that control how traffic is routes to the Kubernetes Service._

_By default, Gloo Edge **will not** route traffic until we add routing rules on a Virtual Service._

Lets start be applying a Virtual Service to rewrite our routes. The purpose of this Virtual Service to is to clean up the routes and make them more understandable to humans. We are going to prefix each route with `/api/v1` and give each route a more explicative name.

For instance, before we would execute a POST request like so:
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"test"}' http://localhost:3000/create
```

We are going to write a Virtual Service to rewrite the route to have a new endpoints of `/api/v1/create` and the GET / to get all tasks route to be `/api/v1/tasks`.

```
kubectl apply -f vs-rewrite.yaml
```

Now that we have applied the Virtual Service for our Task Service, lets test out the new routes.

Create a tasks:
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"test"}' $(glooctl proxy url)/api/v1/create
```
You will see that now you have overwritten the old route of `/create` with `/api/v1/create`

Get all tasks:

### Creating Routes on Virtual Servies with `glooctl`
You can also define routes on Virtual Services through the terminal.

### Clean Up
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
