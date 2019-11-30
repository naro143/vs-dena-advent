DATE := $(shell date "+%Y%m%d%H%M%S")

.PHONY: deploy cron
deploy:
	gcloud app deploy appengine/app.yaml --project vs-dena-advent --version $(DATE)

cron:
	gcloud app deploy appengine/cron.yaml --project vs-dena-advent --version $(DATE)
