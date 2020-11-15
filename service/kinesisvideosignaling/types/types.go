// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A structure for the ICE server connection data.
type IceServer struct {

	// A password to login to the ICE server.
	Password *string

	// The period of time, in seconds, during which the username and password are
	// valid.
	Ttl int32

	// An array of URIs, in the form specified in the
	// I-D.petithuguenin-behave-turn-uris
	// (https://tools.ietf.org/html/draft-petithuguenin-behave-turn-uris-03) spec.
	// These URIs provide the different addresses and/or protocols that can be used to
	// reach the TURN server.
	Uris []string

	// A username to login to the ICE server.
	Username *string
}
