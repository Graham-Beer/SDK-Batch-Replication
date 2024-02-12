package s3controlFunction

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
)

func CreateManifestGenerator(manifestSourceBucket, accountID string) *types.JobManifestGeneratorMemberS3JobManifestGenerator {
	return &types.JobManifestGeneratorMemberS3JobManifestGenerator{
		Value: types.S3JobManifestGenerator{
			EnableManifestOutput: false,
			SourceBucket:         &manifestSourceBucket,
			ExpectedBucketOwner:  &accountID,
			Filter: &types.JobManifestGeneratorFilter{
				EligibleForReplication: aws.Bool(true),
				ObjectReplicationStatuses: []types.ReplicationStatus{
					types.ReplicationStatusFailed,
				},
			},
		},
	}
}
