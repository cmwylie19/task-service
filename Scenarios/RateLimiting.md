# Rate Limiting     
_API Gateways are the first stop from the outside world before reaching your application services (monoliths, microservices, serverless functions, etc)._

 _Any given application service will need to accept requests from the outside.  The problem is that too many incoming requests at the same time can cause a service or application to crash or become unresponsibe. A deliberate example of this would be a Denial of Service Attack._

 _The traditional way to combat this problem is to write rate limiting rules into the application itself. Image if you have an application with 40-50 microservices, this could be quite cumbersome. Using an API Gateway you can enforce Rate Limits globally, which takes the overhead off the developers and makes it easier to keep up to date._

