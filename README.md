[![codecov](https://codecov.io/gh/cmwylie19/task-service/branch/master/graph/badge.svg?token=BRK6V3DOQA)](https://codecov.io/gh/cmwylie19/task-service) ![Node.js CI](https://github.com/cmwylie19/task-service/workflows/Node.js%20CI/badge.svg)

# Task Service

_This repository of scenarios is intended to explain how to use **Gloo Edge** as an API Gateway to manage a simple Nodejs express task api._

These scenarios are built around the express application found at `server.js`. 

**Scenario 1**: [Build and Deploy App in Kubernetes](Scenarios/DeployInK8s.md) 
- How to build a container image from the `Dockerfile`
- How to push containers to a container registry
- Create a Kubernetes eployment from the command line using an image
- Create a Kubernetes Service from the command line from a deployment

**Scenario 2**: [Rewriting the Routes using a Virtual Service](Scenarios/TrafficManagement-RewriteRoutes.md)
- How to use `glooctl` to create a virtual service from scratch
- How to create routes with glooctl
- How to add routes to your Virtual Service in the yaml manifest
- How to export and sanitize your code

**Scenario 3**: [Applying Rate Limiting to Protect your app](Scenarios/RateLimiting.md)
- How to install Gloo Edge Enterprise
- How to update Gloo Edge Settings manifest to configure Rate Limiting descriptors
- How to configure Envoy Rate Limiting actions in the `VirtualService`


**Scenario 4**: [Oauth2 with Google](Scenarios/OAUTH2.md)
- Create a callback route in your Virtual Services to forward from the Identity Provider with a `access_token` and `id_token`
- Register you application with Google as the identity provider
- Create a `client_secret` as a kubernetes secret
- Create an `AuthConfig` CR
- Update your VirtualService to use the `AuthConfig` with `ExtAuth`
- Logs from the `AuthServer`

**NOTE:** It may be advantageous to refer back to this `README.md` while you are going through the scenarios to see how the application works _without_ Gloo Edge to better understand the transformations that are being made at the proxy level.

## Background

_Gloo Edge is an API Gateway, and Kubernetes Ingress Controller that is packed with functionality. These scenarios have been tested on Docker for Desktop but should be compatible for native Kubernetes environments like GKE. It is recommended that you do the scenarios in the order that they are listed to get the most out of this repository. If you find any issues please open up an issue in GitHub_


## Scenarios

- [Refresher on building a container from a Dockerfile, pushing to quay, and deploying into Kubernetes](Scenarios/DeployInK8s.md)
- [Using Gloo Edge Traffic Management Features to rewrite routes in a Virtual Service](Scenarios/TrafficManagement-RewriteRoutes.md)
- [Rate limiting Requests to Upstream services using Gloo Edge enterprise](Scenarios/RateLimiting.md)
- [Securing Express endpoints with Google OAuth2 OIDC Flow](Scenarios/OAUTH2.md)

## Prereqs

- kubectl
- Docker Desktop with Kubernetes Enabled

## Basic Usage

_This section describes how to run and interact with the application locally_

### Install

Install the application depencies:

```
npm i
```

### Run the app locally

```
npm start
```

### Endpoints

**GET /check/healthz - Check Health**

```
curl http://localhost:3000/check/healthz
```

**POST /create - Create Task**

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"test"}' http://localhost:3000/create
```

**GET / - Get all tasks**

```
curl http://localhost:3000/
```

**PUT /:id - Update a task by ID**
The updated task goes in the body of the POST request.

```
curl -X PUT -H "Content-Type: application/json" -d '{"name":"tester","complete":"true"}' http://localhost:3000/d29e6d58525
```

**DELETE /:id - Delete a task by ID**

```
curl -X DELETE http://localhost:3000/d29e6d58525
```

**GET /:id - Get a task by ID**

```
curl http://localhost:3000/d29e6d58525
```
