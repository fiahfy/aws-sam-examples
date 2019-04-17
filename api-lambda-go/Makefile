.PHONY: deps clean format test build watch package deploy

deps:
	go get -u github.com/cespare/reflex

clean:
	rm -rf ./main

format:
	go fmt ./...

test:
	go test ./...

build:
	GOOS=linux GOARCH=amd64 go build

watch:
	make build
	make -j reflex start

reflex:
	reflex -r '\.go' make build

start:
	sam local start-api --skip-pull-image --env-vars env.json

package:
	sam package \
	--template-file template.yml \
	--output-template-file packaged.yml \
	--s3-bucket <bucket_name>

deploy:
	sam deploy \
	--template-file packaged.yml \
	--stack-name <stack_name> \
	--capabilities CAPABILITY_IAM \
	--parameter-overrides Environment=<env>

aws-package:
	aws cloudformation package \
	--template-file template.yml \
	--output-template-file packaged.yml \
	--s3-bucket <bucket_name>

aws-deploy:
	aws cloudformation deploy \
	--template-file packaged.yml \
	--stack-name <stack_name> \
	--capabilities CAPABILITY_IAM \
	--parameter-overrides Environment=<env>
