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

type AWSService struct {
	session *session.Session
	s3      *s3.S3
}

type AWS interface {
	GetS3SignedURL(key string, contentType string, expire time.Duration) (string, error)
}

func NewAWS() *AWSService {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.AWSRegion)},
	)

	if err != nil {
		log.Fatal(err)
	}

	aws := AWSService{
		session: session,
		s3:      s3.New(session),
	}

	return &aws
}

func (service *AWSService) GetS3SignedURL(key string, contentType string, expire time.Duration) (string, error) {
	req, _ := service.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket:              aws.String(conf.AwsS3Bucket),
		Key:                 aws.String(key),
		ResponseContentType: aws.String(contentType),
	})

	urlStr, err := req.Presign(expire)

	if err != nil {
		log.Printf("Failed to sign request, %v", err)
		return "", err
	}

	return urlStr, nil
}
