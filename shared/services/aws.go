package service

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	config "github.com/synergydesigns/stylesblitz-server/shared/config"
)

var conf = config.LoadConfig()

// GetS3SignedURL generates a presigned url to upload files
// to s3
func GetS3SignedURL(key string, expire time.Duration) (string, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.AWSRegion)},
	)

	// Create S3 service client
	svc := s3.New(session)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(conf.AwsS3Bucket),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(expire)

	if err != nil {
		log.Fatalf("Failed to sign request, %v", err)
		return "", err
	}

	return urlStr, nil
}
