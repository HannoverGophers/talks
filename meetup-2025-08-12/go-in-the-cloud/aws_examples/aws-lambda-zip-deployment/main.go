package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configuration struct {
	BucketName string
}

const msgprefix = "<<Zip execution>>: "

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context) error {
	cfg := config()
	s3c := s3Client()

	log.Printf("%sBucket-Name is: %s", msgprefix, cfg.BucketName)

	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(cfg.BucketName),
		Delimiter: aws.String("/"),
	}

	paginator := s3.NewListObjectsV2Paginator(s3c, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			handleError(err, "get next page from paginator")
		}

		for _, object := range page.Contents {
			if object.Key == nil {
				handleError(err, "key is nil, cannot process s3 object list")
			}

			log.Printf("%sLIST OBJECT: Name %q", msgprefix, *object.Key)
		}
	}

	return nil
}

func s3Client() *s3.Client {
	// Initialize the S3 client outside of the handler, during the init phase
	cfg, err := awscfg.LoadDefaultConfig(context.TODO())
	if err != nil {
		handleError(err, "unable to load SDK config")
	}

	return s3.NewFromConfig(cfg)
}

func config() *Configuration {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("No config file found. Using environment variables.")
		}
	}

	cfg := &Configuration{
		BucketName: viper.GetString("BUCKET_NAME"),
	}

	return cfg
}

func handleError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %+v", msg, err)
		os.Exit(1)
	}
}
