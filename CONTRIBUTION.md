# CONTRIBUTION GUIDE
This project is both serviced based and modular. Currently, our project is broken into four modular piece 
1. **Lambda**: All lambda functions have been abstracted into the lambda folder.  This folder contains two main folder 
        - graphql: this folder contains everything that has to do with our graphql function.
        - graphqli folder: this folder contains our development web client for testing our functions
2. **Migrations:** The migration files contains everything that has to do with DB configuration. This folder contains two main folder
        - query: every query we need to execute in ```yml``` file
        - utils: custom functions that parse the query file
3. **Shared:** Contains all the services that communicate with the database and other third-party services
Libraries and
4. **Docs**

## SETTING UP DEV ENVIRONMENT:
#### GOLANG:
Our entire project is based on golang, so you might want to download and install ```GO``` locally... set your ```GOPATH``` to point to your project directory. 
We use ```dep``` for managing dependencies, make sure you have the ```dep`` dependency installed. If you don't, checkout the official [documentaion](https://golang.github.io/dep/docs/installation.html) on how to go about it.
Also, you might want to check out these links for more info.
[Installing Golang](https://golang.org/doc/install)
[GOPATH](https://github.com/golang/go/wiki/GOPATH)

#### LAMBDA:
Since lambda functions run on AWS, we use sam cli to simulate that environment locally. If you don't know what sam cli is, you might want to check this [doc](https://docs.aws.amazon.com/lambda/latest/dg/sam-cli-requirements.html) for more info. 
To start working on the lambda functions locally, you need to
1. install sam cli. The GitHub doc is quite comprehensive and should get you up and running click [here](https://github.com/awslabs/aws-sam-cli) to check it out
2. create an AWS account and Setup and AWS profile add it to your environments ``` AWS_ACCESS_KEY_ID=ssdsdsdsdsdssdss
AWS_SECRET_ACCESS_KEY=C4S/dsdsdssddsdsdssds
USER_NAMESPACE=testuer```. You can follow this comprehensive guide on how to set up AWS credentials for lambda [here](https://serverless.com/framework/docs/providers/aws/guide/credentials/)

#### MIGRATIONS:
working with migrations is quite manual as of now. Just move into the query folder and duplicate one entry, make sure you replace the first integer of the filename with the current timestamp. Also, the ID of the query should also use the current timestamp. Add your query, description, and name. 
When you're done.
1. run `make migrate`
You need to run docker before running migration

## DOCKER
this is the simplest and the bestway to work on this project locally. 
Follow the **GOLANG** && **LAMBDA** installation procedure, then.
1. clone project
2. install docker 
3. turn the env.json.example to `env.json`
4. run `docker-compose build`
6. run `docker-compose up`
7. open a new terminal and run `sam local start-api --env-vars env.json`
Any changes you make will trigger an automatic rebuild of the binaries

## NOTE
1. if you're finding graphqli request going past 2s you can quit sam-cli and use this command instead `sam local start-api --skip-pull-image --env-vars=env.json --profile=<aws-credentials>`
2. All changes/features to the shared folder must be tested
3. If you want to run test locally, `exec` into the styleblitz-shared container `docker exec -it styleblitz-shared bash` and run `go test`
