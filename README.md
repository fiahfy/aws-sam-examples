# aws-sam-sample
> Sample for AWS SAM.


## Requirements
* [Golang](https://golang.org)
* [Docker installed](https://www.docker.com/community-edition)
* aws-cli  
    ```
    brew install awscli
    aws configure
    ```
* aws-sam-cli  
    ```
    brew tap aws/tap
    brew install aws-sam-cli
    ```


## Get Started
```
git clone https://github.com/fiahfy/aws-sam-sample.git
cd $_
make deps
```


## Tasks
### Build
```
make build
```

### Watch
```
make watch
```

### Test
```
make test
```

### Release
```
make package && make deploy
```


## References
* https://github.com/awslabs/aws-lambda-go-api-proxy
* https://github.com/awslabs/serverless-application-model
* https://hackernoon.com/managing-multi-environment-serverless-architecture-using-aws-an-investigation-6cd6501d261e
* https://www.techscore.com/blog/2018/12/07/aws-sam-tips/
* https://www.ketancho.net/entry/2018/01/10/080000


## Issues
* Feature request: make it possible to keep docker container warm  
https://github.com/awslabs/aws-sam-cli/issues/239
* Why are requests so slow?  
https://github.com/awslabs/aws-sam-cli/issues/134
* sam local start-api Missing Authentication Token for root path '/'  
https://github.com/awslabs/aws-sam-cli/issues/437
