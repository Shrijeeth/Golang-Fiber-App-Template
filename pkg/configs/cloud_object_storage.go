package configs

import (
	"errors"
	"os"
	"strconv"

	cloud_storage_object "github.com/Shrijeeth/Golang-Fiber-App-Template/platform/cloud_object_storage"
)

var CloudObjectStorage interface{}

func InitCloudObjectStorage() error {
	cloudObjectStorageType := os.Getenv("CLOUD_OBJECT_STORAGE_TYPE")
	var err error

	switch cloudObjectStorageType {
	case "aws_s3":
		CloudObjectStorage, err = cloud_storage_object.AwsS3Connect()
	case "gcs":
		CloudObjectStorage, err = cloud_storage_object.GCSConnect()
	default:
		CloudObjectStorage, err = nil, errors.New("Invalid Cloud Object Storage Type")
	}

	return err
}

func CloseCloudObjectStorage() error {
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

func IsCloudStorageObjectRequired() bool {
	isCloudStorageObjectRequired, _ := strconv.Atoi(os.Getenv("CLOUD_STORAGE_OBJECT_REQUIRED"))
	return isCloudStorageObjectRequired == 1
}
