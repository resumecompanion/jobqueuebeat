// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateOpenIDConnectProviderThumbprintRequest
type UpdateOpenIDConnectProviderThumbprintInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM OIDC provider resource object for
	// which you want to update the thumbprint. You can get a list of OIDC provider
	// ARNs by using the ListOpenIDConnectProviders operation.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// OpenIDConnectProviderArn is a required field
	OpenIDConnectProviderArn *string `min:"20" type:"string" required:"true"`

	// A list of certificate thumbprints that are associated with the specified
	// IAM OpenID Connect provider. For more information, see CreateOpenIDConnectProvider.
	//
	// ThumbprintList is a required field
	ThumbprintList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateOpenIDConnectProviderThumbprintInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateOpenIDConnectProviderThumbprintInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateOpenIDConnectProviderThumbprintInput"}

	if s.OpenIDConnectProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("OpenIDConnectProviderArn"))
	}
	if s.OpenIDConnectProviderArn != nil && len(*s.OpenIDConnectProviderArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("OpenIDConnectProviderArn", 20))
	}

	if s.ThumbprintList == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThumbprintList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateOpenIDConnectProviderThumbprintOutput
type UpdateOpenIDConnectProviderThumbprintOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateOpenIDConnectProviderThumbprintOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateOpenIDConnectProviderThumbprint = "UpdateOpenIDConnectProviderThumbprint"

// UpdateOpenIDConnectProviderThumbprintRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Replaces the existing list of server certificate thumbprints associated with
// an OpenID Connect (OIDC) provider resource object with a new list of thumbprints.
//
// The list that you pass with this operation completely replaces the existing
// list of thumbprints. (The lists are not merged.)
//
// Typically, you need to update a thumbprint only when the identity provider's
// certificate changes, which occurs rarely. However, if the provider's certificate
// does change, any attempt to assume an IAM role that specifies the OIDC provider
// as a principal fails until the certificate thumbprint is updated.
//
// Trust for the OIDC provider is derived from the provider's certificate and
// is validated by the thumbprint. Therefore, it is best to limit access to
// the UpdateOpenIDConnectProviderThumbprint operation to highly privileged
// users.
//
//    // Example sending a request using UpdateOpenIDConnectProviderThumbprintRequest.
//    req := client.UpdateOpenIDConnectProviderThumbprintRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateOpenIDConnectProviderThumbprint
func (c *Client) UpdateOpenIDConnectProviderThumbprintRequest(input *UpdateOpenIDConnectProviderThumbprintInput) UpdateOpenIDConnectProviderThumbprintRequest {
	op := &aws.Operation{
		Name:       opUpdateOpenIDConnectProviderThumbprint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateOpenIDConnectProviderThumbprintInput{}
	}

	req := c.newRequest(op, input, &UpdateOpenIDConnectProviderThumbprintOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateOpenIDConnectProviderThumbprintRequest{Request: req, Input: input, Copy: c.UpdateOpenIDConnectProviderThumbprintRequest}
}

// UpdateOpenIDConnectProviderThumbprintRequest is the request type for the
// UpdateOpenIDConnectProviderThumbprint API operation.
type UpdateOpenIDConnectProviderThumbprintRequest struct {
	*aws.Request
	Input *UpdateOpenIDConnectProviderThumbprintInput
	Copy  func(*UpdateOpenIDConnectProviderThumbprintInput) UpdateOpenIDConnectProviderThumbprintRequest
}

// Send marshals and sends the UpdateOpenIDConnectProviderThumbprint API request.
func (r UpdateOpenIDConnectProviderThumbprintRequest) Send(ctx context.Context) (*UpdateOpenIDConnectProviderThumbprintResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateOpenIDConnectProviderThumbprintResponse{
		UpdateOpenIDConnectProviderThumbprintOutput: r.Request.Data.(*UpdateOpenIDConnectProviderThumbprintOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateOpenIDConnectProviderThumbprintResponse is the response type for the
// UpdateOpenIDConnectProviderThumbprint API operation.
type UpdateOpenIDConnectProviderThumbprintResponse struct {
	*UpdateOpenIDConnectProviderThumbprintOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateOpenIDConnectProviderThumbprint request.
func (r *UpdateOpenIDConnectProviderThumbprintResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
