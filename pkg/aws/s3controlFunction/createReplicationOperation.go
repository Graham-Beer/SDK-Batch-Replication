package s3controlFunction

import (
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
)

func CreateReplicationOperation() types.JobOperation {
	return types.JobOperation{
		S3ReplicateObject: &types.S3ReplicateObjectOperation{},
	}
}
