package cloudformation

import (
	"encoding/json"
)

// AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource AWS CloudFormation Resource (AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource struct {

	// ReferenceSchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-referenceschema
	ReferenceSchema *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema `json:"ReferenceSchema,omitempty"`

	// S3ReferenceDataSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-s3referencedatasource
	S3ReferenceDataSource *AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource `json:"S3ReferenceDataSource,omitempty"`

	// TableName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-tablename
	TableName *Value `json:"TableName,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource"
}

func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
