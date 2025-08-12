APP_NAME := golambdafunctionserverless
BIN_NAME := bootstrap
BUILD_DIR := bin

.PHONY: build package clean deploy invoke

build:
	@echo "==> Building Go binary for linux/arm64"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o $(BIN_NAME) .

package: build
	@echo "==> Packaging $(APP_NAME).zip"
	mkdir -p $(BUILD_DIR)
	cd . && zip -q -j $(BUILD_DIR)/$(APP_NAME).zip $(BIN_NAME)

clean:
	@echo "==> Cleaning"
	rm -f $(BIN_NAME) $(BUILD_DIR)/$(APP_NAME).zip

deploy:
	@echo "==> Deploying"
	sls deploy --stage dev --region eu-central-1 --aws-profile marks-aws

invoke:
	@echo "==> Invoking"
	sls invoke -f listS3 --log  --stage dev --region eu-central-1