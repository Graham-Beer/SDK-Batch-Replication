package main

import (
	"context"
	"fmt"
	"log"

	"AWS-SDK-S3Batch-Operation-job/v2/pkg/aws/s3controlFunction"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
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

	// Create S3 Batch Operations job
	operation := s3controlFunction.CreateReplicationOperation()
	report := s3controlFunction.CreateReportConfig(reportBucket)
	manifestGenerator := s3controlFunction.CreateManifestGenerator(manifestSourceBucket, accountID)
	input := s3controlFunction.CreateJobInput(accountID, roleARN, operation, defaultPriority, report, manifestGenerator, jobRunConfirmationRequired)

	// Execute the job creation
	result, err := s3controlFunction.CreateS3BatchOperationsJob(client, input)
	if err != nil {
		log.Fatalf("error creating S3 Batch Operations job: %v", err)
	}

	fmt.Println("Successfully created S3 Batch Operations job")
	fmt.Println("Job ID:", *result.JobId)
}
