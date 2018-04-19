# Golang Lambda Functions [🤖](https://github.com/aws/aws-lambda-go)

```go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s! 💩", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
```



### Valid Function Signatures of the Handler

* handler *must* be a function
* between 0-2 arguments; if two arguments the first is always context.context
  * context.Context = runtime information of lambda func
* returns 0-2 arguments; if 1 return value = error; if 2 value, error

```
func ()
func () error
func (TIn), error
func () (TOut, error)
func (context.Context) error
func (context.Context, TIn) error
func (context.Context) (TOut, error)
func (context.Context, TIn) (TOut, error)
```

### Global State to Maximize Performance

* its **best practice to declare & modify global variables** that are independent of the handler code
* a single instance of the Lambda function will never handle multiple events simultaneously
* you can safely change global state; changes will require a new execution context and will not introduce locking/unstable behavior from function invocations directed at the previous execution context

### `context.Context`

* Contains runtime information
  * How much time is remaining before my Lambda dies? (timeouts)
  * Where is Lambda writing its logs? (Cloudwatch log group)
  * e.g. Api Gateway API Key, CognitoPoolID, Function ARN, Name, Version, Memory Limit... 
  * some magic with AWS X-Ray SDK (performance tracing)...

### Logging

* there is a `import "log"` module:
  * `log.Print("Hey Mr. Log!")` logging always to cloudwatch
  * automagically adds a time stamp
* but anything that writes to `stdout` or `stderr` will be written to cloudwatch

### Errors

* custom error handling to raise an exception
  * but custom errors must import the errors module

```go
package main
 
import (
        "errors"
        "github.com/aws/aws-lambda-go/lambda"
)
 
func OnlyErrors() error {
  		// panic(errors.New("Something went wrong"))
        return errors.New("something went wrong!")
}
 
func main() {
        lambda.Start(OnlyErrors)
```

* Lambdas can fail for reasons beyond your control (network outage!)
  * Go `panics`; Lambda attempts to capture the error & write it to stderr; will also be written into cloudwatch
  * Lambda will be recreated automatically

### Environment Variables

* as simple as `os.Getenv("NAME")`
* there are some env vars set!
  * `AWS_ACCESS_KEY`, `AWS_SECRET_KEY` (easy for wroking with aws-go-sdk)
  * `AWS_REGION`, `TZ` (current local tme)

### [Creating a Deployment Package](https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html)

```bash
GOOS=linux go build lambda_handler.go
zip handler.zip ./lambda_handler
# --handler is the path to the executable inside the .zip
aws lambda create-function \
  --region region \
  --function-name lambda-handler \
  --memory 128 \
  --role arn:aws:iam::account-id:role/execution_role \
  --runtime go1.x \
  --zip-file fileb://path-to-your-zip-file/handler.zip \
  --handler lambda-handler
```

or… use [Apex](http://apex.run/)!

> Apex lets you **build, deploy, and manage AWS Lambda functions with ease**. With Apex you can use languages that are not natively supported by AWS Lambda, such as Golang, through the use of a Node.js shim injected into the build. A variety of workflow related tooling is provided for **testing functions, rolling back deploys, viewing metrics, tailing logs, hooking into the build system and more.**



# Drawbacks

* Dynamic vs. Compiled
  * I can't see my code in AWS Console 🤖
* bad integration /w Cloudformation (building + packaging needs to be done separatly 💩)

# Performance

[Source 1](https://hackernoon.com/aws-lambda-go-vs-node-js-performance-benchmark-1c8898341982)

[Source 2](https://github.com/yunspace/serverless-golang)

**Fibonacci**



```
Max requests:        1000
Concurrency level:   5
Agent:               keepalive
Requests per second: 10
                  Node.js     Node.js (rec.)  Go
Mean latency:     76.9 ms     407.8 ms       75.3 ms
 50%              73          392            67  
 90%              95          492            91 
 95%             101          526           109
 99%             201          709           226
100%             630          814           562 (longest request)
```





**S3 & Dynamo interaction**

1. Grabs a ~50kb image from S3.
2. Writes its LastModified timestamp to a DynamoDB table.

([Go Code](https://gist.github.com/tnolet/8b7614c6fa87b9322d7e0a86995866bc#file-s3dynamo-go) | [Node Code](https://gist.github.com/tnolet/a56c338581f95a1a8b462791c8464d5b#file-s3dynamo-js))

```
Max requests:        1000
Concurrency level:   5
Agent:               keepalive
Requests per second: 10
                 Node.js    Go
Mean latency:    252.2 ms   109.7 ms
 50%             203        91
 90%             384        151
 95%             478        197
 99%             894        435
100%            8103       1133(longest request)
```



# Monitoring

### AWS-XRay

* [AWS X-Ray](https://aws.amazon.com/de/xray/)
* [github.com/aws/aws-xray-sdk-go](https://github.com/aws/aws-xray-sdk-go)

![AWS-X-Ray for Go](https://github.com/aws/aws-xray-sdk-go/raw/master/images/example.png?raw=true)

### Sparta

* [gosparta.io](http://gosparta.io/)
* [Monitoring /w Sparta](https://medium.com/@mweagle/spartagrafana-serverless-monitoring-f86ca6da79ed)
* [Build an S3 website with API Gateway and AWS Lambda for Go using Sparta](https://read.acloud.guru/go-aws-lambda-building-an-html-website-with-api-gateway-and-lambda-for-go-using-sparta-5e6fe79f63ef)
* ​

![Sparta Lambda Call Count Monitoring](https://cdn-images-1.medium.com/max/800/1*dx0yISRPvSVe-NV9zKIgMQ.gif)

!(Lambda Call Count Different Instances)[https://cdn-images-1.medium.com/max/1000/1*-zIZJUdU5INy4Ms4R1P4tg.png]