package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("gin cold start")

		e := gin.Default()

		// Ping handler
		e.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		// Message handler
		e.GET("/message", func(c *gin.Context) {
			c.String(200, os.Getenv("MSG"))
		})

		ginLambda = ginadapter.New(e)
	}

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
