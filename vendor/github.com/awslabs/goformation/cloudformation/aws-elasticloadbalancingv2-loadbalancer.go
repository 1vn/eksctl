package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSElasticLoadBalancingV2LoadBalancer AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::LoadBalancer)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html
type AWSElasticLoadBalancingV2LoadBalancer struct {

	// IpAddressType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-ipaddresstype
	IpAddressType *Value `json:"IpAddressType,omitempty"`

	// LoadBalancerAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes
	LoadBalancerAttributes []AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute `json:"LoadBalancerAttributes,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-name
	Name *Value `json:"Name,omitempty"`

	// Scheme AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-scheme
	Scheme *Value `json:"Scheme,omitempty"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-securitygroups
	SecurityGroups []*Value `json:"SecurityGroups,omitempty"`

	// SubnetMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-subnetmappings
	SubnetMappings []AWSElasticLoadBalancingV2LoadBalancer_SubnetMapping `json:"SubnetMappings,omitempty"`

	// Subnets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-subnets
	Subnets []*Value `json:"Subnets,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-tags
	Tags []Tag `json:"Tags,omitempty"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-type
	Type *Value `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2LoadBalancer) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::LoadBalancer"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSElasticLoadBalancingV2LoadBalancer) MarshalJSON() ([]byte, error) {
	type Properties AWSElasticLoadBalancingV2LoadBalancer
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSElasticLoadBalancingV2LoadBalancer) UnmarshalJSON(b []byte) error {
	type Properties AWSElasticLoadBalancingV2LoadBalancer
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSElasticLoadBalancingV2LoadBalancer(*res.Properties)
	}

	return nil
}

// GetAllAWSElasticLoadBalancingV2LoadBalancerResources retrieves all AWSElasticLoadBalancingV2LoadBalancer items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticLoadBalancingV2LoadBalancerResources() map[string]AWSElasticLoadBalancingV2LoadBalancer {
	results := map[string]AWSElasticLoadBalancingV2LoadBalancer{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSElasticLoadBalancingV2LoadBalancer:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ElasticLoadBalancingV2::LoadBalancer" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSElasticLoadBalancingV2LoadBalancer{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSElasticLoadBalancingV2LoadBalancerWithName retrieves all AWSElasticLoadBalancingV2LoadBalancer items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticLoadBalancingV2LoadBalancerWithName(name string) (AWSElasticLoadBalancingV2LoadBalancer, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSElasticLoadBalancingV2LoadBalancer:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ElasticLoadBalancingV2::LoadBalancer" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSElasticLoadBalancingV2LoadBalancer{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSElasticLoadBalancingV2LoadBalancer{}, errors.New("resource not found")
}
