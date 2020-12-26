[![codecov](https://codecov.io/gh/cmwylie19/task-service/branch/master/graph/badge.svg?token=BRK6V3DOQA)](https://codecov.io/gh/cmwylie19/task-service) ![Node.js CI](https://github.com/cmwylie19/task-service/workflows/Node.js%20CI/badge.svg)

# Task Service

_This repository of scenarios is intended to explain how to use Gloo Edge as an API Gateway and Kubernetes Ingress Controller to manage a simple task api. The scenarios will take you from building and deploying the application from a source to traffic routing and rate limiting scenarios. The first scenario is a provides a background and howto guide on Cloud Native Application Development and Deployment so feel free to skip it if you already know your way around the cluster._

**NOTE:** It may be advantageous to refer back to this `README.md` while you are doing the scenarios to compare and contract how the application works _without_ Gloo Edge to better understand the transformations that are being made at the proxy level.

## Background

Gloo Edge is an API Gateway, and Kubernetes Ingress Controller that is packed with functionality. These scenarios use a simple Task API to teach you how to build, deploy, manage the simple Task API. These scenarios have been tested on Docker for Desktop but should be compatible for native Kubernetes environments like GKE.

## Assumptions

- You have a kubernetes environment with `kubectl`
- You have a unix like environment to execute the commands

## Scenarios

- [Refresher on building a container from a Dockerfile, pushing to quay, and deploying into Kubernetes](https://github.com/cmwylie19/task-service/blob/master/Scenarios/DeployInK8s.md)
- [Using Gloo Edge Traffic Management Features to rewrite routes in a Virtual Service](https://github.com/cmwylie19/task-service/blob/master/Scenarios/TrafficManagement-RewriteRoutes.md)

## Prereqs

- kubectl
- install `glooctl`
- Docker
- Gloo Edge
- Kubernetes v1.11.3+ (I am user Docker Desktop with Kubernetes Enabled)

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

**GET /healthz - check health**

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
