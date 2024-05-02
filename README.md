# Go Gin API on AWS
This is an API built with Gin Web Framework (Golang) and deployed on AWS (API Gateway and Lambda) using CDK (Cloud Development Kit).

## How to Build & Deploy
### Lambda Functions

- `# go mod init main`
- `# go mod tidy`
- `# go run main`

### Cloud Infrastructure

- `# npm install`
- `# npm run build`
- `# cdk synth`
- `# cdk deploy`

## Technologies

### Lambda Functions

- [Golang](https://go.dev)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

### Cloud Infrastructure

- [Cloud Development Kit (CDK)](https://docs.aws.amazon.com/cdk/v2/guide/home.html)
- [CloudFormation](https://aws.amazon.com/cloudformation/)
- [API Gateway](https://aws.amazon.com/api-gateway)
- [Lambda](https://aws.amazon.com/lambda)

## Useful Links
- https://djharper.dev/post/2018/01/27/running-go-aws-lambda-functions-locally/
