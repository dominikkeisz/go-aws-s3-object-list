package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	bucketName := GetBucketNameFromArg(os.Args)

	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(config, func (o *s3.Options)  {
		o.Region = "us-east-1"
	})

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Objects in %s:\n", bucketName)
	for _, object := range output.Contents {
		fmt.Printf("key=%s size=%d\n", aws.ToString(object.Key), object.Size)
	}
}

func GetBucketNameFromArg(arg []string) string {
	argCount := len(arg[1:])
	if argCount != 1 {
		log.Fatal("Name of your bucket needs to be the only argument!")
	}

	return arg[1]
}
