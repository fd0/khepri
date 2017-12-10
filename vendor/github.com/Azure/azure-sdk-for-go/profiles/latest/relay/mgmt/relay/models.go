// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package relay

import original "github.com/Azure/azure-sdk-for-go/services/relay/mgmt/2017-04-01/relay"

type WCFRelaysClient = original.WCFRelaysClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type HybridConnectionsClient = original.HybridConnectionsClient
type AccessRights = original.AccessRights

const (
	Listen AccessRights = original.Listen
	Manage AccessRights = original.Manage
	Send   AccessRights = original.Send
)

type KeyType = original.KeyType

const (
	PrimaryKey   KeyType = original.PrimaryKey
	SecondaryKey KeyType = original.SecondaryKey
)

type ProvisioningStateEnum = original.ProvisioningStateEnum

const (
	Created   ProvisioningStateEnum = original.Created
	Deleted   ProvisioningStateEnum = original.Deleted
	Failed    ProvisioningStateEnum = original.Failed
	Succeeded ProvisioningStateEnum = original.Succeeded
	Unknown   ProvisioningStateEnum = original.Unknown
	Updating  ProvisioningStateEnum = original.Updating
)

type RelaytypeEnum = original.RelaytypeEnum

const (
	HTTP   RelaytypeEnum = original.HTTP
	NetTCP RelaytypeEnum = original.NetTCP
)

type SkuTier = original.SkuTier

const (
	Standard SkuTier = original.Standard
)

type UnavailableReason = original.UnavailableReason

const (
	InvalidName                           UnavailableReason = original.InvalidName
	NameInLockdown                        UnavailableReason = original.NameInLockdown
	NameInUse                             UnavailableReason = original.NameInUse
	None                                  UnavailableReason = original.None
	SubscriptionIsDisabled                UnavailableReason = original.SubscriptionIsDisabled
	TooManyNamespaceInCurrentSubscription UnavailableReason = original.TooManyNamespaceInCurrentSubscription
)

type AccessKeys = original.AccessKeys
type AuthorizationRule = original.AuthorizationRule
type AuthorizationRuleListResult = original.AuthorizationRuleListResult
type AuthorizationRuleProperties = original.AuthorizationRuleProperties
type CheckNameAvailability = original.CheckNameAvailability
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ErrorResponse = original.ErrorResponse
type HybridConnection = original.HybridConnection
type HybridConnectionListResult = original.HybridConnectionListResult
type HybridConnectionProperties = original.HybridConnectionProperties
type Namespace = original.Namespace
type NamespaceListResult = original.NamespaceListResult
type NamespaceProperties = original.NamespaceProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type RegenerateAccessKeyParameters = original.RegenerateAccessKeyParameters
type Resource = original.Resource
type ResourceNamespacePatch = original.ResourceNamespacePatch
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type UpdateParameters = original.UpdateParameters
type WcfRelay = original.WcfRelay
type WcfRelayProperties = original.WcfRelayProperties
type WcfRelaysListResult = original.WcfRelaysListResult
type NamespacesClient = original.NamespacesClient
type OperationsClient = original.OperationsClient

func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func NewWCFRelaysClient(subscriptionID string) WCFRelaysClient {
	return original.NewWCFRelaysClient(subscriptionID)
}
func NewWCFRelaysClientWithBaseURI(baseURI string, subscriptionID string) WCFRelaysClient {
	return original.NewWCFRelaysClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewHybridConnectionsClient(subscriptionID string) HybridConnectionsClient {
	return original.NewHybridConnectionsClient(subscriptionID)
}
func NewHybridConnectionsClientWithBaseURI(baseURI string, subscriptionID string) HybridConnectionsClient {
	return original.NewHybridConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
