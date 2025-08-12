package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()
	cfg, _ := config.LoadDefaultConfig(ctx)
	client := s3.NewFromConfig(cfg)
	resp, _ := client.ListBuckets(ctx, &s3.ListBucketsInput{})

	_ = resp
}
