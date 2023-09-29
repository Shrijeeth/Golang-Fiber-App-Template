package cloud_storage_object

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func GCSConnect() (*storage.Client, error) {
	gcsCredentialsFile := os.Getenv("GCS_CREDENTIALS_FILE")
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(gcsCredentialsFile))
	if err != nil {
		return nil, err
	}

	return client, nil
}
