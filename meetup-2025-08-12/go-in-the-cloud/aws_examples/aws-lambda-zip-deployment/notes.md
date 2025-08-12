# command to compile
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go

## Zip file
zip go-in-the-cloud.zip bootstrap

In IAM -> Roles muss man die AWSLambdaBasicExecutionRole anlegen. BZW. die ARN von der Rolle muss dann beim Hochladen der 
Lambda mitgegeben bzw. der Lambda zugewiesen werden.

###
aws lambda create-function --function-name goLambdaFunction \
--runtime provided.al2023 --handler bootstrap \
--architectures arm64 \
--role arn:aws:iam::348737449002:role/AWSLambdaBasicExecutionRoleInclCWAndS3 \
--zip-file fileb:///Users/mark.borukhov/go/src/github.com/gophers-meetup/go-in-the-cloud/aws_examples/aws-lambda-zip-deployment/go-in-the-cloud.zip \
--profile marks-aws
####

lambda ausführen und logs entschlüßeln
aws lambda invoke --function-name goLambdaFunction out --log-type Tail --profile marks-aws \
--query 'LogResult' --output text --cli-binary-format raw-in-base64-out | base64 --decode

update lambda function configuration
aws lambda update-function-configuration --function-name goLambdaFunction \
--role arn:aws:iam::348737449002:role/AWSLambdaBasicExecutionRoleInclCWAndS3 \
--profile marks-aws


