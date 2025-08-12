IMAGE_NAME := golambdafunctionimage
FUNC_NAME := golambdafunctionimage

.PHONY: build push create-lambda invoke invoke-with-logs login update

build:
	@echo "==> Building Docker image"
	docker buildx build --platform linux/amd64 --provenance=false -t $(IMAGE_NAME):latest .
	@echo "==> Tagging the docker image"
	docker tag $(IMAGE_NAME):latest 348737449002.dkr.ecr.eu-central-1.amazonaws.com/$(IMAGE_NAME):latest

login:
	@echo "==> Log in"
	aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 348737449002.dkr.ecr.eu-central-1.amazonaws.com


push:
	@echo "==> Pushing the docker image to AWS ECR"
	docker push 348737449002.dkr.ecr.eu-central-1.amazonaws.com/$(IMAGE_NAME):latest

create-lambda:
	@echo "==> Creating a lambda function"
	aws lambda create-function \
    --function-name $(FUNC_NAME) \
    --package-type Image \
    --code ImageUri=348737449002.dkr.ecr.eu-central-1.amazonaws.com/$(IMAGE_NAME):latest \
    --role arn:aws:iam::348737449002:role/AWSLambdaBasicExecutionRoleInclCWAndS3

update:
	@echo "==> Updating the lambda function"
	aws lambda update-function-code \
    --function-name $(FUNC_NAME) \
    --image-uri 348737449002.dkr.ecr.eu-central-1.amazonaws.com/$(IMAGE_NAME):latest \
    --publish

invoke:
	@echo "==> Invoking the lambda function"
	aws lambda invoke --function-name $(FUNC_NAME) response.json

invoke-with-logs:
	@echo "==> Invoking the lambda function"
	aws lambda invoke --function-name $(FUNC_NAME) --payload '{}' --log-type Tail \
      /tmp/response.json \
      --query 'LogResult' \
      --output text \
    | base64 --decode