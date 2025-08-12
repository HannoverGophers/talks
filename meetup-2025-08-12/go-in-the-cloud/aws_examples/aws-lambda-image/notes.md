docker build
docker buildx build --platform linux/amd64 --provenance=false -t golambdafunctionimage:latest .

aws ecr login command
aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 348737449002.dkr.ecr.eu-central-1.amazonaws.com

tag docker image
docker tag golambdafunctionimage:latest 348737449002.dkr.ecr.eu-central-1.amazonaws.com/golambdafunctionimage:latest

docker push
docker push 348737449002.dkr.ecr.eu-central-1.amazonaws.com/golambdafunctionimage:latest

creete lambda function base on the image
aws lambda create-function \
--function-name golambdafunctionimage \
--package-type Image \
--code ImageUri=348737449002.dkr.ecr.eu-central-1.amazonaws.com/golambdafunctionimage:latest \
--role arn:aws:iam::348737449002:role/AWSLambdaBasicExecutionRoleInclCWAndS3

invoke image
aws lambda invoke --function-name golambdafunctionimage response.json

update function
aws lambda update-function-code \
--function-name golambdafunctionimage \
--image-uri 348737449002.dkr.ecr.eu-central-1.amazonaws.com/golambdafunctionimage:latest \
--publish