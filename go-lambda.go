package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GoLambdaStackProps struct {
	awscdk.StackProps
}

func NewGoLambdaStack(scope constructs.Construct, id string, props *GoLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	myfunction := awslambda.NewFunction(stack, jsii.String("newUsernameValidate"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	table := awsdynamodb.NewTable(stack, jsii.String("myUserTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("userTable"),
	})

	table.GrantReadWriteData(myfunction)

	api := awsapigateway.NewRestApi(stack, jsii.String("exampleAPI"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "X-Amz-Date", "Authorization", "X-Api-Key"),
			AllowMethods: jsii.Strings("GET", "POST", "PUT", "DELETE", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
		DeployOptions: &awsapigateway.StageOptions{
			LoggingLevel:     awsapigateway.MethodLoggingLevel_INFO,
			DataTraceEnabled: jsii.Bool(true),
		},
		EndpointConfiguration: &awsapigateway.EndpointConfiguration{
			Types: &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
		},
		CloudWatchRole: jsii.Bool(true),
	})

	integration := awsapigateway.NewLambdaIntegration(myfunction, nil)

	// define routes
	registerRoute := api.Root().AddResource(jsii.String("register"), nil)
	registerRoute.AddMethod(jsii.String("POST"), integration, nil)

	loginRoute := api.Root().AddResource(jsii.String("login"), nil)
	loginRoute.AddMethod(jsii.String("POST"), integration, nil)

	//protected route
	protectedRoute := api.Root().AddResource(jsii.String("protected"), nil)
	protectedRoute.AddMethod(jsii.String("GET"), integration, nil)
	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("GoLambdaQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGoLambdaStack(app, "GoLambdaStack", &GoLambdaStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
