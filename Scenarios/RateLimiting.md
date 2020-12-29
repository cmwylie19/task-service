# Rate Limiting     
_API Gateways are the first stop from the outside world before reaching your application services (monoliths, microservices, serverless functions, etc)._

 _Any given application service will need to accept requests from the outside.  The problem is that too many incoming requests at the same time can cause a service or application to crash or become unresponsibe. A deliberate example of this would be a Denial of Service Attack._

 _The traditional way to combat this problem is to write rate limiting rules into the application itself. Image if you have an application with 40-50 microservices, this could be quite cumbersome. Using an API Gateway you can enforce Rate Limits globally, which takes the overhead off the developers and makes it easier to keep up to date._

## Background
We are going to setup rate limiting for our todo app to use Gloo Edge with Envoy's rate-limiting API. _This feature is only available in the enterprise version of Gloo Edge._ You must get an enterprise API Key for this part.

```
glooctl install gateway enterprise --license-key=<license-key>
```

## Steps
- Install Gloo (see Background)
- Deploy `Task-Service` in Kubernetes
- Update the Gloo Edge Settings manifest to configure the rate limiting descriptors
- Configure Envoy rate limiting actions in the `VirtualService`

## Deploy `Task-Service` in Kubernetes
First thing after install Gloo Edge enterprise is to deploy the `Task-Service` app in kubernetes.
```
kubectl apply -f k8s/task-servoce/yaml
```

Once the task-service has a status of `RUNNING` you are done with this step
```
kubectl get pods
```
_or_ if you have `watch` installed on your system
```
watch kubectl get pods
```

## Update the Gloo Edge Settings manifest to configure the reate limiting descriptors
Rate limiting descriptors define an ordered tuple of keys that _have to_ match for the in order for the associated rate limit to be applied [Gloo Edge Docs](https://docs.solo.io/gloo-edge/latest/guides/security/rate_limiting/envoy/). 