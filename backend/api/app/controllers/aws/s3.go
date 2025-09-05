package aws

import (
	"api/config"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetS3Content(keyName *string) (string, error) {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)

	presignedReq, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(config.Config.ApiBucketName),
		Key:    keyName,
	}, s3.WithPresignExpires(time.Minute*30))
	if err != nil {
		return "", err
	}
	return presignedReq.URL, nil
}

func UploadS3(keyName *string, body io.ReadSeeker) error {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	s3Client := s3.NewFromConfig(cfg)

	ctx := context.Background()
	var cancelFn func()
	timeout := 30 * time.Second
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}

	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(config.Config.ApiBucketName),
		Key:    keyName,
		Body:   body,
	})

	if err != nil {
		return err
	}
	fmt.Printf("successfully uploaded file to %s/%s\n", config.Config.ApiBucketName, *keyName)
	return nil
}
