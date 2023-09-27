package cloud_storage_object

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func AwsS3Connect() (*s3.S3, error) {
	s3Id := os.Getenv("AWS_ACCESS_KEY")
	s3Secret := os.Getenv("AWS_SECRET_KEY")
	s3Endpoint := os.Getenv("AWS_S3_ENDPOINT")
	awsRegion := os.Getenv("AWS_REGION")
	awsProfile := os.Getenv("AWS_PROFILE")

	s3Config := aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3Id, s3Secret, ""),
		Endpoint:         aws.String(s3Endpoint),
		Region:           aws.String(awsRegion),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  s3Config,
		Profile: awsProfile,
	})
	if err != nil {
		return nil, err
	}

	cloudStorageObject := s3.New(sess)
	return cloudStorageObject, nil
}
