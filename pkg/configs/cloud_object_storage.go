package configs

import (
	"errors"
	"os"

	cloud_storage_object "github.com/Shrijeeth/Golang-Fiber-App-Template/platform/cloud_object_storage"
)

var CloudObjectStorage interface{}

func InitCloudObjectStorage() error {
	cloudObjectStorageType := os.Getenv("CLOUD_OBJECT_STORAGE_TYPE")
	var err error

	switch cloudObjectStorageType {
	case "aws_s3":
		CloudObjectStorage, err = cloud_storage_object.AwsS3Connect()
	default:
		CloudObjectStorage, err = nil, errors.New("Invalid Cloud Object Storage Type")
	}

	return err
}

func CloseCloseObjectStorage() error {
	cloudObjectStorageType := os.Getenv("CLOUD_OBJECT_STORAGE_TYPE")
	var err error

	switch cloudObjectStorageType {
	case "aws_s3":
		err = nil
	default:
		err = errors.New("Invalid Cloud Object Storage Type")
	}

	return err
}