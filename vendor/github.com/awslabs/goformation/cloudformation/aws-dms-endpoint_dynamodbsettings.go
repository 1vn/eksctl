package cloudformation

import (
	"encoding/json"
)

// AWSDMSEndpoint_DynamoDbSettings AWS CloudFormation Resource (AWS::DMS::Endpoint.DynamoDbSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html
type AWSDMSEndpoint_DynamoDbSettings struct {

	// ServiceAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-dynamodbsettings.html#cfn-dms-endpoint-dynamodbsettings-serviceaccessrolearn
	ServiceAccessRoleArn *Value `json:"ServiceAccessRoleArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEndpoint_DynamoDbSettings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.DynamoDbSettings"
}

func (r *AWSDMSEndpoint_DynamoDbSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
