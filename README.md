# base64-api
A simple API to encode/decode strings to base64.

# Used Tools:
- golang;
- Unit Tests;
- Makefile;
- Gorila/mux;
- Docker;
- DockerHub;
- Shell Script;
- Terraform;
- AWS ECS;
- GitHub Actions;
  
## The API
An swagger DOC was not included, but the methods available are:
- http://<API_URL>/api/encode?string=<STRING_TO_ENCODE>: Return a base64 encode of string query;
- http://<API_URL>/api/decode?string=<STRING_TO_DECODE>: Return a base64 decoded of string query;

## Code
This API was made using golang with gorilla/mux. Simple unit tests have been implemented to assure the correct functionality.
There are provide two methods, one to encode and another to decode strings.
The tests execute some common scenarios, like, just letters, numbers, special characters and so on

## Infra
The API has been containerized, published to Docker Hub and deployed to Amazon ECS using Terraform.

## Execute Locally
To make easy develop, test and deploy the API. An makefile is available. You can execute the following commands:
* `make install`: Install golang libraries;
* `make lint`: Run golang lint;
* `make test`: Run Unit tests;
* `make run`: Run the application;
* `make go`: Run lint, tests and run;
* `make tfInit`: Init terraform;
* `make tfCheck`: Check tf files formatting;
* `make tfPlan`: Plan the infrastructure creation;
* `make tfApply`: Create the infrastructure;
* `make tfDestroy`: Destroy the infrastructure;
* `make tf`: Execute init, plan and check;
* `make tfOut`: Show the API URL;
* `make testInfra`: Test the API endpoint (wait a 200 status code)
## CI/CD
A simple CI/CD pipelie was developed to delivery the API to ECS.
The pipeline, exectue the unit tests, build the docker container, publish it to docker hub and execute the deploy on ECS, making tests if the infra is well, and if not, destroy it.

### The Flow:
- Developer commit changes;
- Git Hub Action "Golang CI" is started:
  - Dependencies are downloaded;
  - Unit tests are executed;
- If everything is working, the git Hub Action "Docker Image CI" is started:
  - The API is containerized using Docker;
  - The Docker is versioned and published to DockerHub;
-  If everything is working, the git Hub Action "Deploy Container" is started:
   - AWS Login;
   - Terraform Cloud Login to storage the terraform state;
   - Terraform Init;
   - Terraform Check formatting;
   - Terraform Plan;
   - Terraform Apply;
   - So, an script try to get a 200 status code in the LB endpoint, if not, Terraform Destroy is executed;  
- If everything is working, the API is available to use

## Next Steps
- Improve unit tests;
- Improve Logs;
- Improve versioning;
- Develop a Swagger to API;
- Develop component test;
- Execute e2e test and promove package to "production";
- Integrate with another tools, like, sonar, clod watch, and etc;
- Improve documentation;
