FUNC_NAME := httpfunction
BUCKET_NAME := marks-test-gcbucket-go-in-the-cloud

.PHONY: deploy invoke

deploy:
	#gcloud config set project <PROJECT_ID>
	#gcloud auth login

	# .env ist optional - hier setzen wir die BUCKET_NAME-Envvar direkt
	gcloud functions deploy $(FUNC_NAME) \
	  --gen2 \
	  --region=europe-west1 \
	  --runtime=go122 \
	  --source=. \
	  --entry-point=HandleRequest \
	  --trigger-http \
	  --allow-unauthenticated \
	  --set-env-vars=BUCKET_NAME=$(BUCKET_NAME)

invoke:
	gcloud functions call $(FUNC_NAME) \
  		--gen2 --region=europe-west1 \
  		--data='{}'