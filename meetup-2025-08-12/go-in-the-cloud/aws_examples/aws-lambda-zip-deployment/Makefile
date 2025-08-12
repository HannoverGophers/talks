APP_NAME := bootstrap
ZIP_NAME := go-in-the-cloud
FUNC_NAME := goLambdaFunction

.PHONY: build package clean deploy invoke zip update-configuration update-function-code

build: clean
	@echo "==> Building Go binary"
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o $(APP_NAME) main.go

zip: build
	@echo "==> Packaging $(APP_NAME).zip"
	zip $(ZIP_NAME).zip $(APP_NAME)

clean:
	@echo "==> Cleaning"
	rm -f $(ZIP_NAME).zip

deploy:
	@echo "==> Deploying"
	aws lambda create-function --function-name $(FUNC_NAME) \
    --runtime provided.al2023 --handler $(APP_NAME) \
    --architectures arm64 \
    --role arn:aws:iam::348737449002:role/AWSLambdaBasicExecutionRoleInclCWAndS3 \
    --zip-file fileb:///Users/mark.borukhov/go/src/github.com/gophers-meetup/talks/meetup-2025-08-12/go-in-the-cloud/aws_examples/aws-lambda-zip-deployment/$(ZIP_NAME).zip \
    --profile marks-aws

invoke:
	@echo "==> Invoking"
	aws lambda invoke --function-name $(FUNC_NAME) out --log-type Tail --profile marks-aws \
    --query 'LogResult' --output text --cli-binary-format raw-in-base64-out | base64 --decode

update-configuration: zip
	aws lambda update-function-configuration --function-name $(FUNC_NAME) \
    --role arn:aws:iam::348737449002:role/AWSLambdaBasicExecutionRoleInclCWAndS3 \
    --profile marks-aws

update-function-code: zip
	aws lambda update-function-code --function-name $(FUNC_NAME) \
    --zip-file fileb:///Users/mark.borukhov/go/src/github.com/gophers-meetup/talks/meetup-2025-08-12/go-in-the-cloud/aws_examples/aws-lambda-zip-deployment/$(ZIP_NAME).zip \
	--profile marks-aws