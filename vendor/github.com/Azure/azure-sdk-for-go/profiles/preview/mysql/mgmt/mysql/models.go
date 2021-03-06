// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package mysql

import original "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"

type CheckNameAvailabilityClient = original.CheckNameAvailabilityClient

func NewCheckNameAvailabilityClient(subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClient(subscriptionID)
}
func NewCheckNameAvailabilityClientWithBaseURI(baseURI string, subscriptionID string) CheckNameAvailabilityClient {
	return original.NewCheckNameAvailabilityClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type ConfigurationsClient = original.ConfigurationsClient

func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}

type DatabasesClient = original.DatabasesClient

func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}

type FirewallRulesClient = original.FirewallRulesClient

func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}

type LocationBasedPerformanceTierClient = original.LocationBasedPerformanceTierClient

func NewLocationBasedPerformanceTierClient(subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClient(subscriptionID)
}
func NewLocationBasedPerformanceTierClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedPerformanceTierClient {
	return original.NewLocationBasedPerformanceTierClientWithBaseURI(baseURI, subscriptionID)
}

type LogFilesClient = original.LogFilesClient

func NewLogFilesClient(subscriptionID string) LogFilesClient {
	return original.NewLogFilesClient(subscriptionID)
}
func NewLogFilesClientWithBaseURI(baseURI string, subscriptionID string) LogFilesClient {
	return original.NewLogFilesClientWithBaseURI(baseURI, subscriptionID)
}

type CreateMode = original.CreateMode

const (
	CreateModeDefault                   CreateMode = original.CreateModeDefault
	CreateModePointInTimeRestore        CreateMode = original.CreateModePointInTimeRestore
	CreateModeServerPropertiesForCreate CreateMode = original.CreateModeServerPropertiesForCreate
)

func PossibleCreateModeValues() []CreateMode {
	return original.PossibleCreateModeValues()
}

type GeoRedundantBackup = original.GeoRedundantBackup

const (
	Disabled GeoRedundantBackup = original.Disabled
	Enabled  GeoRedundantBackup = original.Enabled
)

func PossibleGeoRedundantBackupValues() []GeoRedundantBackup {
	return original.PossibleGeoRedundantBackupValues()
}

type OperationOrigin = original.OperationOrigin

const (
	NotSpecified OperationOrigin = original.NotSpecified
	System       OperationOrigin = original.System
	User         OperationOrigin = original.User
)

func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}

type ServerState = original.ServerState

const (
	ServerStateDisabled ServerState = original.ServerStateDisabled
	ServerStateDropping ServerState = original.ServerStateDropping
	ServerStateReady    ServerState = original.ServerStateReady
)

func PossibleServerStateValues() []ServerState {
	return original.PossibleServerStateValues()
}

type ServerVersion = original.ServerVersion

const (
	FiveFullStopSeven ServerVersion = original.FiveFullStopSeven
	FiveFullStopSix   ServerVersion = original.FiveFullStopSix
)

func PossibleServerVersionValues() []ServerVersion {
	return original.PossibleServerVersionValues()
}

type SkuTier = original.SkuTier

const (
	Basic           SkuTier = original.Basic
	GeneralPurpose  SkuTier = original.GeneralPurpose
	MemoryOptimized SkuTier = original.MemoryOptimized
)

func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}

type SslEnforcementEnum = original.SslEnforcementEnum

const (
	SslEnforcementEnumDisabled SslEnforcementEnum = original.SslEnforcementEnumDisabled
	SslEnforcementEnumEnabled  SslEnforcementEnum = original.SslEnforcementEnumEnabled
)

func PossibleSslEnforcementEnumValues() []SslEnforcementEnum {
	return original.PossibleSslEnforcementEnumValues()
}

type Configuration = original.Configuration
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsCreateOrUpdateFuture = original.ConfigurationsCreateOrUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseProperties = original.DatabaseProperties
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesCreateOrUpdateFuture = original.FirewallRulesCreateOrUpdateFuture
type FirewallRulesDeleteFuture = original.FirewallRulesDeleteFuture
type LogFile = original.LogFile
type LogFileListResult = original.LogFileListResult
type LogFileProperties = original.LogFileProperties
type NameAvailability = original.NameAvailability
type NameAvailabilityRequest = original.NameAvailabilityRequest
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type PerformanceTierListResult = original.PerformanceTierListResult
type PerformanceTierProperties = original.PerformanceTierProperties
type PerformanceTierServiceLevelObjectives = original.PerformanceTierServiceLevelObjectives
type ProxyResource = original.ProxyResource
type Server = original.Server
type ServerForCreate = original.ServerForCreate
type ServerListResult = original.ServerListResult
type ServerProperties = original.ServerProperties
type BasicServerPropertiesForCreate = original.BasicServerPropertiesForCreate
type ServerPropertiesForCreate = original.ServerPropertiesForCreate
type ServerPropertiesForDefaultCreate = original.ServerPropertiesForDefaultCreate
type ServerPropertiesForRestore = original.ServerPropertiesForRestore
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type ServerUpdateParameters = original.ServerUpdateParameters
type ServerUpdateParametersProperties = original.ServerUpdateParametersProperties
type Sku = original.Sku
type StorageProfile = original.StorageProfile
type TrackedResource = original.TrackedResource
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type ServersClient = original.ServersClient

func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
