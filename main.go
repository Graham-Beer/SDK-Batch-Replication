package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
)

var (
	roleARN                    string = ""
	accountID                  string = ""
	manifestSourceBucket       string = ""
	reportBucket               string = ""
	defaultPriority            int32  = 10
	jobRunConfirmationRequired bool   = false
)

func main() {
	// Load AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	// Create S3 Control client
	client := s3control.NewFromConfig(cfg)

	// Specify the replication operation for the job
	operation := types.JobOperation{
		S3ReplicateObject: &types.S3ReplicateObjectOperation{},
	}

	// Specify the report configuration for the job
	report := types.JobReport{
		Enabled:     true,
		Format:      types.JobReportFormatReportCsv20180820,
		ReportScope: types.JobReportScopeAllTasks,
		Bucket:      &reportBucket,
	}

	// Specify the manifest generation configuration for the job
	manifestGenerator := &types.JobManifestGeneratorMemberS3JobManifestGenerator{
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

	// Build the input for creating the job
	input := &s3control.CreateJobInput{
		AccountId:            &accountID,
		Operation:            &operation,
		Priority:             &defaultPriority,
		Report:               &report,
		RoleArn:              &roleARN,
		ConfirmationRequired: &jobRunConfirmationRequired,
		ManifestGenerator:    manifestGenerator,
	}

	// Create the job
	result, err := client.CreateJob(context.Background(), input)
	if err != nil {
		log.Fatalf("error creating S3 Batch Operations job: %v", err)
	}

	fmt.Println("Successfully created S3 Batch Operations job")
	fmt.Println("Job ID:", *result.JobId)
}
