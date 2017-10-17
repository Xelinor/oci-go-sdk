// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// CreateRegionSubscriptionRequest wrapper for the CreateRegionSubscription operation
type CreateRegionSubscriptionRequest struct {

	// Request object for activate a new region.
	CreateRegionSubscriptionDetails CreateRegionSubscriptionDetails `contributesTo:"body"`

	// The OCID of the tenancy.
	TenancyID string `mandatory:"true" contributesTo:"path" name:"tenancyId"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

// CreateRegionSubscriptionResponse wrapper for the CreateRegionSubscription operation
type CreateRegionSubscriptionResponse struct {

	// The RegionSubscription instance
	RegionSubscription `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`
}