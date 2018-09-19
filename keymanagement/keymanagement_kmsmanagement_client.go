// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//KmsManagementClient a client for KmsManagement
type KmsManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewKmsManagementClientWithConfigurationProvider Creates a new default KmsManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewKmsManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client KmsManagementClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = KmsManagementClient{BaseClient: baseClient}
	client.BasePath = "20180608"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *KmsManagementClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "kms", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *KmsManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *KmsManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateKey Creates a new key.
func (client KmsManagementClient) CreateKey(ctx context.Context, request CreateKeyRequest) (response CreateKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createKey, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateKeyResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateKeyResponse")
	}
	return
}

// createKey implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) createKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys")
	if err != nil {
		return nil, err
	}

	var response CreateKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateKeyVersion Generates new cryptographic material for a key. Key must be in an `ENABLED` state to be
// rotated.
func (client KmsManagementClient) CreateKeyVersion(ctx context.Context, request CreateKeyVersionRequest) (response CreateKeyVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createKeyVersion, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateKeyVersionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateKeyVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateKeyVersionResponse")
	}
	return
}

// createKeyVersion implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) createKeyVersion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys/{keyId}/keyVersions")
	if err != nil {
		return nil, err
	}

	var response CreateKeyVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableKey Disables a key to make it unavailable for encryption
// or decryption.
func (client KmsManagementClient) DisableKey(ctx context.Context, request DisableKeyRequest) (response DisableKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.disableKey, policy)
	if err != nil {
		if ociResponse != nil {
			response = DisableKeyResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableKeyResponse")
	}
	return
}

// disableKey implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) disableKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys/{keyId}/actions/disable")
	if err != nil {
		return nil, err
	}

	var response DisableKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableKey Enables a key to make it available for encryption or
// decryption.
func (client KmsManagementClient) EnableKey(ctx context.Context, request EnableKeyRequest) (response EnableKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.enableKey, policy)
	if err != nil {
		if ociResponse != nil {
			response = EnableKeyResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableKeyResponse")
	}
	return
}

// enableKey implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) enableKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/keys/{keyId}/actions/enable")
	if err != nil {
		return nil, err
	}

	var response EnableKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetKey Gets information about the specified key.
func (client KmsManagementClient) GetKey(ctx context.Context, request GetKeyRequest) (response GetKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getKey, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetKeyResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetKeyResponse")
	}
	return
}

// getKey implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) getKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keys/{keyId}")
	if err != nil {
		return nil, err
	}

	var response GetKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetKeyVersion Gets information about the specified key version.
func (client KmsManagementClient) GetKeyVersion(ctx context.Context, request GetKeyVersionRequest) (response GetKeyVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getKeyVersion, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetKeyVersionResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetKeyVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetKeyVersionResponse")
	}
	return
}

// getKeyVersion implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) getKeyVersion(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keys/{keyId}/keyVersions/{keyVersionId}")
	if err != nil {
		return nil, err
	}

	var response GetKeyVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListKeyVersions Lists all key versions for the specified key.
func (client KmsManagementClient) ListKeyVersions(ctx context.Context, request ListKeyVersionsRequest) (response ListKeyVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listKeyVersions, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListKeyVersionsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListKeyVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListKeyVersionsResponse")
	}
	return
}

// listKeyVersions implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) listKeyVersions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keys/{keyId}/keyVersions")
	if err != nil {
		return nil, err
	}

	var response ListKeyVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListKeys Lists the keys in the specified vault and compartment.
func (client KmsManagementClient) ListKeys(ctx context.Context, request ListKeysRequest) (response ListKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listKeys, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListKeysResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListKeysResponse")
	}
	return
}

// listKeys implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) listKeys(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/keys")
	if err != nil {
		return nil, err
	}

	var response ListKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateKey Updates the properties of a key. Specifically, you can
// only update the `displayName` property. Furthermore, the
// key must in an `ACTIVE` or `CREATING` state.
func (client KmsManagementClient) UpdateKey(ctx context.Context, request UpdateKeyRequest) (response UpdateKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateKey, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateKeyResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateKeyResponse")
	}
	return
}

// updateKey implements the OCIOperation interface (enables retrying operations)
func (client KmsManagementClient) updateKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/keys/{keyId}")
	if err != nil {
		return nil, err
	}

	var response UpdateKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
