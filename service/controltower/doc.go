// Code generated by smithy-go-codegen DO NOT EDIT.

// Package controltower provides the API client, operations, and parameter types
// for AWS Control Tower.
//
// These interfaces allow you to apply the AWS library of pre-defined controls to
// your organizational units, programmatically. In AWS Control Tower, the terms
// "control" and "guardrail" are synonyms. To call these APIs, you'll need to know:
//
//   - the controlIdentifier for the control--or guardrail--you are targeting.
//   - the ARN associated with the target organizational unit (OU), which we call
//     the targetIdentifier .
//   - the ARN associated with a resource that you wish to tag or untag.
//
// To get the controlIdentifier for your AWS Control Tower control: The
// controlIdentifier is an ARN that is specified for each control. You can view the
// controlIdentifier in the console on the Control details page, as well as in the
// documentation. The controlIdentifier is unique in each AWS Region for each
// control. You can find the controlIdentifier for each Region and control in the
// Tables of control metadata (https://docs.aws.amazon.com/controltower/latest/userguide/control-metadata-tables.html)
// in the AWS Control Tower User Guide. A quick-reference list of control
// identifers for the AWS Control Tower legacy Strongly recommended and Elective
// controls is given in Resource identifiers for APIs and controls (https://docs.aws.amazon.com/controltower/latest/userguide/control-identifiers.html.html)
// in the Controls reference guide section (https://docs.aws.amazon.com/controltower/latest/userguide/control-identifiers.html)
// of the AWS Control Tower User Guide. Remember that Mandatory controls cannot be
// added or removed. ARN format:
// arn:aws:controltower:{REGION}::control/{CONTROL_NAME} Example:
// arn:aws:controltower:us-west-2::control/AWS-GR_AUTOSCALING_LAUNCH_CONFIG_PUBLIC_IP_DISABLED
// To get the targetIdentifier : The targetIdentifier is the ARN for an OU. In the
// AWS Organizations console, you can find the ARN for the OU on the Organizational
// unit details page associated with that OU. OU ARN format:
// arn:${Partition}:organizations::${MasterAccountId}:ou/o-${OrganizationId}/ou-${OrganizationalUnitId}
// Details and examples
//   - Control API input and output examples with CLI (https://docs.aws.amazon.com/controltower/latest/userguide/control-api-examples-short.html)
//   - Enable controls with CloudFormation (https://docs.aws.amazon.com/controltower/latest/userguide/enable-controls.html)
//   - Control metadata tables (https://docs.aws.amazon.com/controltower/latest/userguide/control-metadata-tables.html)
//   - List of identifiers for legacy controls (https://docs.aws.amazon.com/controltower/latest/userguide/control-identifiers.html)
//   - Controls reference guide (https://docs.aws.amazon.com/controltower/latest/userguide/controls.html)
//   - Controls library groupings (https://docs.aws.amazon.com/controltower/latest/userguide/controls-reference.html)
//   - Creating AWS Control Tower resources with AWS CloudFormation (https://docs.aws.amazon.com/controltower/latest/userguide/creating-resources-with-cloudformation.html)
//
// To view the open source resource repository on GitHub, see
// aws-cloudformation/aws-cloudformation-resource-providers-controltower (https://github.com/aws-cloudformation/aws-cloudformation-resource-providers-controltower)
// Recording API Requests AWS Control Tower supports AWS CloudTrail, a service that
// records AWS API calls for your AWS account and delivers log files to an Amazon
// S3 bucket. By using information collected by CloudTrail, you can determine which
// requests the AWS Control Tower service received, who made the request and when,
// and so on. For more about AWS Control Tower and its support for CloudTrail, see
// Logging AWS Control Tower Actions with AWS CloudTrail (https://docs.aws.amazon.com/controltower/latest/userguide/logging-using-cloudtrail.html)
// in the AWS Control Tower User Guide. To learn more about CloudTrail, including
// how to turn it on and find your log files, see the AWS CloudTrail User Guide.
package controltower
