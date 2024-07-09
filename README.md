# AWS lambda project

deployed version on ""

**go-lambda**
-- aws infrastructure running `/lambda` go project

**lambda build steps**

- `cd lambda` change to lambda directory
- `make build` to run build script (Makefile with golang specific deoploying option)

```Makefile
build:
	@GOOS=linux GOARCH=amd64 go build -o bootstrap
	@zip function.zip bootstrap
```

**AWS Prequisites**

1. setup aws cli ("https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html")
2. configure aws cli ("https://docs.aws.amazon.com/cli/v1/userguide/cli-chap-configure.html")
3. install aws cdk toolkit ("https://docs.aws.amazon.com/cdk/v2/guide/cli.html")
4. `cdk deploy` to deploy project on aws
5. set JWT_TOKEN environment variable on AWS Lambda

**Testing End Points**

- install curl commandline tool ("https://curl.se/docs/install.html")

1. test `/register` Route 

- CHANGE_USERNAME , CHANGE_PASSWORD  -> change this variable
- `curl -X POST https://6fcucft8dg.execute-api.us-east-1.amazonaws.com/prod/register -H "Contenct-Type: application/json"  -d '{"username": "CHANGE_USERNAME" , "password": "CHANGE_PASSWORD"}' `

2. test `/login` Route

- CHANGE_USERNAME , CHANGE_PASSWORD  -> change this variable
- `curl -X POST https://6fcucft8dg.execute-api.us-east-1.amazonaws.com/prod/login -H "Contenct-Type: application/json"  -d '{"username": "CHANGE_USERNAME" , "password": "CHANGE_PASSWORD"}' `
- Response will return JWT_TOKEN

3. test `/protected` Route

- YOUR_JWT_TOKEN_FROM_LOGIN -> change this variable
- `curl -X GET https://6fcucft8dg.execute-api.us-east-1.amazonaws.com/prod/protected -H "Contenct-Type: application/json" -H "Authorization: Bearer YOUR_JWT_TOKEN_FROM_LOGIN"`

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
