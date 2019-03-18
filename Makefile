.PHONY: deps clean format test build watch package

deps:
	go get -u github.com/cespare/reflex

clean:
	rm -rf ./main

format:
	go fmt ./...

test:
	go test ./... -v

build:
	GOOS=linux GOARCH=amd64 go build

watch:
	make build
	make -j reflex start-api

reflex:
	reflex -r '\.go' make build

start-api:
	sam local start-api --skip-pull-image --env-vars env.json

package:
	sam package --template-file template.yml --output-template-file packaged.yml --s3-bucket <BUCKET_NAME>

deploy:
	sam deploy --template-file packaged.yml --stack-name <STACK_NAME> --capabilities CAPABILITY_IAM --parameter-overrides Environment=<ENV>
