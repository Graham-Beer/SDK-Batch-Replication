package s3controlFunction

import (
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
)

func CreateReportConfig(reportBucket string) types.JobReport {
	return types.JobReport{
		Enabled:     true,
		Format:      types.JobReportFormatReportCsv20180820,
		ReportScope: types.JobReportScopeAllTasks,
		Bucket:      &reportBucket,
	}
}
