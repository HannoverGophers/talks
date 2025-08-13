PROJECT := mystical-moon-468622-b9
REGION := europe-west1
SERVICE_NAME := multiple-routes


.PHONY: deploy invoke deploy-new-version

deploy:
	gcloud config set project $(PROJECT)
	gcloud config set run/region $(REGION)

	gcloud run deploy $(SERVICE_NAME) \
	--source=. \
	--allow-unauthenticated \
	--set-env-vars=SERVICE_NAME=$(SERVICE_NAME) 