// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// The input for DeleteStack action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeleteStackInput
type DeleteStackInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for this DeleteStack request. Specify this token if you
	// plan to retry requests so that AWS CloudFormation knows that you're not attempting
	// to delete a stack with the same name. You might retry DeleteStack requests
	// to ensure that AWS CloudFormation successfully received them.
	//
	// All events triggered by a given stack operation are assigned the same client
	// request token, which you can use to track operations. For example, if you
	// execute a CreateStack operation with the token token1, then all the StackEvents
	// generated by that operation will have ClientRequestToken set as token1.
	//
	// In the console, stack operations display the client request token on the
	// Events tab. Stack operations that are initiated from the console use the
	// token format Console-StackOperation-ID, which helps you easily identify the
	// stack operation . For example, if you create a stack using the console, each
	// stack event would be assigned the same token in the following format: Console-CreateStack-7f59c3cf-00d2-40c7-b2ff-e75db0987002.
	ClientRequestToken *string `min:"1" type:"string"`

	// For stacks in the DELETE_FAILED state, a list of resource logical IDs that
	// are associated with the resources you want to retain. During deletion, AWS
	// CloudFormation deletes the stack but does not delete the retained resources.
	//
	// Retaining resources is useful when you cannot delete a resource, such as
	// a non-empty S3 bucket, but you want to delete the stack.
	RetainResources []string `type:"list"`

	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM)
	// role that AWS CloudFormation assumes to delete the stack. AWS CloudFormation
	// uses the role's credentials to make calls on your behalf.
	//
	// If you don't specify a value, AWS CloudFormation uses the role that was previously
	// associated with the stack. If no role is available, AWS CloudFormation uses
	// a temporary session that is generated from your user credentials.
	RoleARN *string `min:"20" type:"string"`

	// The name or the unique stack ID that is associated with the stack.
	//
	// StackName is a required field
	StackName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteStackInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.RoleARN != nil && len(*s.RoleARN) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleARN", 20))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeleteStackOutput
type DeleteStackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteStackOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteStack = "DeleteStack"

// DeleteStackRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Deletes a specified stack. Once the call completes successfully, stack deletion
// starts. Deleted stacks do not show up in the DescribeStacks API if the deletion
// has been completed successfully.
//
//    // Example sending a request using DeleteStackRequest.
//    req := client.DeleteStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeleteStack
func (c *Client) DeleteStackRequest(input *DeleteStackInput) DeleteStackRequest {
	op := &aws.Operation{
		Name:       opDeleteStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteStackInput{}
	}

	req := c.newRequest(op, input, &DeleteStackOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteStackRequest{Request: req, Input: input, Copy: c.DeleteStackRequest}
}

// DeleteStackRequest is the request type for the
// DeleteStack API operation.
type DeleteStackRequest struct {
	*aws.Request
	Input *DeleteStackInput
	Copy  func(*DeleteStackInput) DeleteStackRequest
}

// Send marshals and sends the DeleteStack API request.
func (r DeleteStackRequest) Send(ctx context.Context) (*DeleteStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteStackResponse{
		DeleteStackOutput: r.Request.Data.(*DeleteStackOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteStackResponse is the response type for the
// DeleteStack API operation.
type DeleteStackResponse struct {
	*DeleteStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteStack request.
func (r *DeleteStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
