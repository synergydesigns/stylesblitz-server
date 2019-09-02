package service

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	config "github.com/synergydesigns/stylesblitz-server/shared/config"
)

type AWSService struct {
	session *session.Session
	s3      *s3.S3
	config  *config.Config
}

type AWS interface {
	GetS3SignedURL(key string, contentType string, expire time.Duration) (string, error)
}

func NewAWS(conf *config.Config) *AWSService {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.AWSRegion),
	})

	if err != nil {
		log.Fatal(err)
	}

	aws := AWSService{
		session: session,
		s3:      s3.New(session),
		config:  conf,
	}

	return &aws
}

func (service *AWSService) GetS3SignedURL(key string, contentType string, expire time.Duration) (string, error) {
	req, _ := service.s3.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(service.config.AwsS3Bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	})

	urlStr, err := req.Presign(expire)

	if err != nil {
		log.Printf("Failed to sign request, %v", err)
		return "", err
	}

	return urlStr, nil
}
