package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func LongRunningHandler(ctx context.Context) string {
	deadline, _ := ctx.Done()
	for {
		select {
		case <-time.Until(deadline).Truncate(100 * time.Millisecond):
			return "Finished before timing out."
		default:
			log.Print("hello!")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	lambda.Start(LongRunningHandler)
}
