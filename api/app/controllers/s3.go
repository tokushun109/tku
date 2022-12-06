package controllers

import (
	"api/config"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Content(keyName *string) (string, error) {
	sess := session.Must(session.NewSession())

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &config.Config.ApiBucketName,
		Key:    keyName,
	})
	url, err := req.Presign(time.Minute * 30)
	if err != nil {
		return "", err
	}
	return url, nil
}

func UploadS3(keyName *string, body io.ReadSeeker) error {
	sess := session.Must(session.NewSession())

	svc := s3.New(sess)

	ctx := context.Background()
	var cancelFn func()
	timeout := 30 * time.Second
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}

	_, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: &config.Config.ApiBucketName,
		Key:    keyName,
		Body:   body,
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			return err
		} else {
			return err
		}
	}
	fmt.Printf("successfully uploaded file to %s/%s\n", config.Config.ApiBucketName, *keyName)
	return nil
}
