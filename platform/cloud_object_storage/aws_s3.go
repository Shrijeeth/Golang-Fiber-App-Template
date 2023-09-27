package cloud_storage_object

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func AwsS3Connect() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{ // TODO: Add correct configs
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		return nil, err
	}

	cloudStorageObject := s3.New(sess)
	return cloudStorageObject, nil
}