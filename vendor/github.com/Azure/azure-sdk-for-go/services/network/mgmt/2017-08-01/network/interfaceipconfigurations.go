package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// InterfaceIPConfigurationsClient is the network Client
type InterfaceIPConfigurationsClient struct {
	ManagementClient
}

// NewInterfaceIPConfigurationsClient creates an instance of the InterfaceIPConfigurationsClient client.
func NewInterfaceIPConfigurationsClient(subscriptionID string) InterfaceIPConfigurationsClient {
	return NewInterfaceIPConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInterfaceIPConfigurationsClientWithBaseURI creates an instance of the InterfaceIPConfigurationsClient client.
func NewInterfaceIPConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) InterfaceIPConfigurationsClient {
	return InterfaceIPConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the specified network interface ip configuration.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
// IPConfigurationName is the name of the ip configuration name.
func (client InterfaceIPConfigurationsClient) Get(resourceGroupName string, networkInterfaceName string, IPConfigurationName string) (result InterfaceIPConfiguration, err error) {
	req, err := client.GetPreparer(resourceGroupName, networkInterfaceName, IPConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InterfaceIPConfigurationsClient) GetPreparer(resourceGroupName string, networkInterfaceName string, IPConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigurationName":  autorest.Encode("path", IPConfigurationName),
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceIPConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InterfaceIPConfigurationsClient) GetResponder(resp *http.Response) (result InterfaceIPConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all ip configurations in a network interface
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
func (client InterfaceIPConfigurationsClient) List(resourceGroupName string, networkInterfaceName string) (result InterfaceIPConfigurationListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InterfaceIPConfigurationsClient) ListPreparer(resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InterfaceIPConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InterfaceIPConfigurationsClient) ListResponder(resp *http.Response) (result InterfaceIPConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client InterfaceIPConfigurationsClient) ListNextResults(lastResults InterfaceIPConfigurationListResult) (result InterfaceIPConfigurationListResult, err error) {
	req, err := lastResults.InterfaceIPConfigurationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfaceIPConfigurationsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client InterfaceIPConfigurationsClient) ListComplete(resourceGroupName string, networkInterfaceName string, cancel <-chan struct{}) (<-chan InterfaceIPConfiguration, <-chan error) {
	resultChan := make(chan InterfaceIPConfiguration)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, networkInterfaceName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
