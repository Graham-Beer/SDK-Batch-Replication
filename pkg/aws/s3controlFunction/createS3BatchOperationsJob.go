package s3controlFunction

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

func CreateS3BatchOperationsJob(client *s3control.Client, input *s3control.CreateJobInput) (*s3control.CreateJobOutput, error) {
	return client.CreateJob(context.Background(), input)
}
