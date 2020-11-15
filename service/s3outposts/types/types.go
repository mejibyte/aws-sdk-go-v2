// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// S3 on Outposts access points simplify managing data access at scale for shared
// datasets in Amazon S3 on Outposts. S3 on Outposts uses endpoints to connect to
// Outposts buckets so that you can perform actions within your virtual private
// cloud (VPC).
type Endpoint struct {

	// The VPC CIDR committed by this endpoint.
	CidrBlock *string

	// The time the endpoint was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the endpoint.
	EndpointArn *string

	// The network interface of the endpoint.
	NetworkInterfaces []NetworkInterface

	// The ID of the AWS Outpost.
	OutpostsId *string

	// The status of the endpoint.
	Status EndpointStatus
}

// The container for the network interface.
type NetworkInterface struct {

	// The ID for the network interface.
	NetworkInterfaceId *string
}
