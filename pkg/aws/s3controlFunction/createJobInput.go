package s3controlFunction

import (
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
)

func CreateJobInput(accountID, roleARN string, operation types.JobOperation, priority int32, report types.JobReport, manifestGenerator *types.JobManifestGeneratorMemberS3JobManifestGenerator, jobRunConfirmationRequired bool) *s3control.CreateJobInput {
	return &s3control.CreateJobInput{
		AccountId:            &accountID,
		Operation:            &operation,
		Priority:             &priority,
		Report:               &report,
		RoleArn:              &roleARN,
		ConfirmationRequired: &jobRunConfirmationRequired,
		ManifestGenerator:    manifestGenerator,
	}
}
