/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AccessLevelFilterKey string

const (
	AccessLevelFilterKey_Account AccessLevelFilterKey = "Account"
	AccessLevelFilterKey_Role    AccessLevelFilterKey = "Role"
	AccessLevelFilterKey_User    AccessLevelFilterKey = "User"
)

type AccessStatus string

const (
	AccessStatus_ENABLED      AccessStatus = "ENABLED"
	AccessStatus_UNDER_CHANGE AccessStatus = "UNDER_CHANGE"
	AccessStatus_DISABLED     AccessStatus = "DISABLED"
)

type ChangeAction string

const (
	ChangeAction_ADD    ChangeAction = "ADD"
	ChangeAction_MODIFY ChangeAction = "MODIFY"
	ChangeAction_REMOVE ChangeAction = "REMOVE"
)

type CopyOption string

const (
	CopyOption_CopyTags CopyOption = "CopyTags"
)

type CopyProductStatus string

const (
	CopyProductStatus_SUCCEEDED   CopyProductStatus = "SUCCEEDED"
	CopyProductStatus_IN_PROGRESS CopyProductStatus = "IN_PROGRESS"
	CopyProductStatus_FAILED      CopyProductStatus = "FAILED"
)

type DescribePortfolioShareType string

const (
	DescribePortfolioShareType_ACCOUNT                     DescribePortfolioShareType = "ACCOUNT"
	DescribePortfolioShareType_ORGANIZATION                DescribePortfolioShareType = "ORGANIZATION"
	DescribePortfolioShareType_ORGANIZATIONAL_UNIT         DescribePortfolioShareType = "ORGANIZATIONAL_UNIT"
	DescribePortfolioShareType_ORGANIZATION_MEMBER_ACCOUNT DescribePortfolioShareType = "ORGANIZATION_MEMBER_ACCOUNT"
)

type EngineWorkflowStatus string

const (
	EngineWorkflowStatus_SUCCEEDED EngineWorkflowStatus = "SUCCEEDED"
	EngineWorkflowStatus_FAILED    EngineWorkflowStatus = "FAILED"
)

type EvaluationType string

const (
	EvaluationType_STATIC  EvaluationType = "STATIC"
	EvaluationType_DYNAMIC EvaluationType = "DYNAMIC"
)

type LastSyncStatus string

const (
	LastSyncStatus_SUCCEEDED LastSyncStatus = "SUCCEEDED"
	LastSyncStatus_FAILED    LastSyncStatus = "FAILED"
)

type OrganizationNodeType string

const (
	OrganizationNodeType_ORGANIZATION        OrganizationNodeType = "ORGANIZATION"
	OrganizationNodeType_ORGANIZATIONAL_UNIT OrganizationNodeType = "ORGANIZATIONAL_UNIT"
	OrganizationNodeType_ACCOUNT             OrganizationNodeType = "ACCOUNT"
)

type PortfolioShareType string

const (
	PortfolioShareType_IMPORTED           PortfolioShareType = "IMPORTED"
	PortfolioShareType_AWS_SERVICECATALOG PortfolioShareType = "AWS_SERVICECATALOG"
	PortfolioShareType_AWS_ORGANIZATIONS  PortfolioShareType = "AWS_ORGANIZATIONS"
)

type PrincipalType string

const (
	PrincipalType_IAM         PrincipalType = "IAM"
	PrincipalType_IAM_PATTERN PrincipalType = "IAM_PATTERN"
)

type ProductSource string

const (
	ProductSource_ACCOUNT ProductSource = "ACCOUNT"
)

type ProductType string

const (
	ProductType_CLOUD_FORMATION_TEMPLATE ProductType = "CLOUD_FORMATION_TEMPLATE"
	ProductType_MARKETPLACE              ProductType = "MARKETPLACE"
	ProductType_TERRAFORM_OPEN_SOURCE    ProductType = "TERRAFORM_OPEN_SOURCE"
	ProductType_TERRAFORM_CLOUD          ProductType = "TERRAFORM_CLOUD"
	ProductType_EXTERNAL                 ProductType = "EXTERNAL"
)

type ProductViewFilterBy string

const (
	ProductViewFilterBy_FullTextSearch  ProductViewFilterBy = "FullTextSearch"
	ProductViewFilterBy_Owner           ProductViewFilterBy = "Owner"
	ProductViewFilterBy_ProductType     ProductViewFilterBy = "ProductType"
	ProductViewFilterBy_SourceProductId ProductViewFilterBy = "SourceProductId"
)

type ProductViewSortBy string

const (
	ProductViewSortBy_Title        ProductViewSortBy = "Title"
	ProductViewSortBy_VersionCount ProductViewSortBy = "VersionCount"
	ProductViewSortBy_CreationDate ProductViewSortBy = "CreationDate"
)

type PropertyKey string

const (
	PropertyKey_OWNER       PropertyKey = "OWNER"
	PropertyKey_LAUNCH_ROLE PropertyKey = "LAUNCH_ROLE"
)

type ProvisionedProductPlanStatus string

const (
	ProvisionedProductPlanStatus_CREATE_IN_PROGRESS  ProvisionedProductPlanStatus = "CREATE_IN_PROGRESS"
	ProvisionedProductPlanStatus_CREATE_SUCCESS      ProvisionedProductPlanStatus = "CREATE_SUCCESS"
	ProvisionedProductPlanStatus_CREATE_FAILED       ProvisionedProductPlanStatus = "CREATE_FAILED"
	ProvisionedProductPlanStatus_EXECUTE_IN_PROGRESS ProvisionedProductPlanStatus = "EXECUTE_IN_PROGRESS"
	ProvisionedProductPlanStatus_EXECUTE_SUCCESS     ProvisionedProductPlanStatus = "EXECUTE_SUCCESS"
	ProvisionedProductPlanStatus_EXECUTE_FAILED      ProvisionedProductPlanStatus = "EXECUTE_FAILED"
)

type ProvisionedProductPlanType string

const (
	ProvisionedProductPlanType_CLOUDFORMATION ProvisionedProductPlanType = "CLOUDFORMATION"
)

type ProvisionedProductStatus_SDK string

const (
	ProvisionedProductStatus_SDK_AVAILABLE        ProvisionedProductStatus_SDK = "AVAILABLE"
	ProvisionedProductStatus_SDK_UNDER_CHANGE     ProvisionedProductStatus_SDK = "UNDER_CHANGE"
	ProvisionedProductStatus_SDK_TAINTED          ProvisionedProductStatus_SDK = "TAINTED"
	ProvisionedProductStatus_SDK_ERROR            ProvisionedProductStatus_SDK = "ERROR"
	ProvisionedProductStatus_SDK_PLAN_IN_PROGRESS ProvisionedProductStatus_SDK = "PLAN_IN_PROGRESS"
)

type ProvisionedProductViewFilterBy string

const (
	ProvisionedProductViewFilterBy_SearchQuery ProvisionedProductViewFilterBy = "SearchQuery"
)

type ProvisioningArtifactGuidance string

const (
	ProvisioningArtifactGuidance_DEFAULT    ProvisioningArtifactGuidance = "DEFAULT"
	ProvisioningArtifactGuidance_DEPRECATED ProvisioningArtifactGuidance = "DEPRECATED"
)

type ProvisioningArtifactPropertyName string

const (
	ProvisioningArtifactPropertyName_Id ProvisioningArtifactPropertyName = "Id"
)

type ProvisioningArtifactType string

const (
	ProvisioningArtifactType_CLOUD_FORMATION_TEMPLATE ProvisioningArtifactType = "CLOUD_FORMATION_TEMPLATE"
	ProvisioningArtifactType_MARKETPLACE_AMI          ProvisioningArtifactType = "MARKETPLACE_AMI"
	ProvisioningArtifactType_MARKETPLACE_CAR          ProvisioningArtifactType = "MARKETPLACE_CAR"
	ProvisioningArtifactType_TERRAFORM_OPEN_SOURCE    ProvisioningArtifactType = "TERRAFORM_OPEN_SOURCE"
	ProvisioningArtifactType_TERRAFORM_CLOUD          ProvisioningArtifactType = "TERRAFORM_CLOUD"
	ProvisioningArtifactType_EXTERNAL                 ProvisioningArtifactType = "EXTERNAL"
)

type RecordStatus string

const (
	RecordStatus_CREATED              RecordStatus = "CREATED"
	RecordStatus_IN_PROGRESS          RecordStatus = "IN_PROGRESS"
	RecordStatus_IN_PROGRESS_IN_ERROR RecordStatus = "IN_PROGRESS_IN_ERROR"
	RecordStatus_SUCCEEDED            RecordStatus = "SUCCEEDED"
	RecordStatus_FAILED               RecordStatus = "FAILED"
)

type Replacement string

const (
	Replacement_TRUE        Replacement = "TRUE"
	Replacement_FALSE       Replacement = "FALSE"
	Replacement_CONDITIONAL Replacement = "CONDITIONAL"
)

type RequiresRecreation string

const (
	RequiresRecreation_NEVER         RequiresRecreation = "NEVER"
	RequiresRecreation_CONDITIONALLY RequiresRecreation = "CONDITIONALLY"
	RequiresRecreation_ALWAYS        RequiresRecreation = "ALWAYS"
)

type ResourceAttribute string

const (
	ResourceAttribute_PROPERTIES     ResourceAttribute = "PROPERTIES"
	ResourceAttribute_METADATA       ResourceAttribute = "METADATA"
	ResourceAttribute_CREATIONPOLICY ResourceAttribute = "CREATIONPOLICY"
	ResourceAttribute_UPDATEPOLICY   ResourceAttribute = "UPDATEPOLICY"
	ResourceAttribute_DELETIONPOLICY ResourceAttribute = "DELETIONPOLICY"
	ResourceAttribute_TAGS           ResourceAttribute = "TAGS"
)

type ServiceActionAssociationErrorCode string

const (
	ServiceActionAssociationErrorCode_DUPLICATE_RESOURCE ServiceActionAssociationErrorCode = "DUPLICATE_RESOURCE"
	ServiceActionAssociationErrorCode_INTERNAL_FAILURE   ServiceActionAssociationErrorCode = "INTERNAL_FAILURE"
	ServiceActionAssociationErrorCode_LIMIT_EXCEEDED     ServiceActionAssociationErrorCode = "LIMIT_EXCEEDED"
	ServiceActionAssociationErrorCode_RESOURCE_NOT_FOUND ServiceActionAssociationErrorCode = "RESOURCE_NOT_FOUND"
	ServiceActionAssociationErrorCode_THROTTLING         ServiceActionAssociationErrorCode = "THROTTLING"
	ServiceActionAssociationErrorCode_INVALID_PARAMETER  ServiceActionAssociationErrorCode = "INVALID_PARAMETER"
)

type ServiceActionDefinitionKey string

const (
	ServiceActionDefinitionKey_Name       ServiceActionDefinitionKey = "Name"
	ServiceActionDefinitionKey_Version    ServiceActionDefinitionKey = "Version"
	ServiceActionDefinitionKey_AssumeRole ServiceActionDefinitionKey = "AssumeRole"
	ServiceActionDefinitionKey_Parameters ServiceActionDefinitionKey = "Parameters"
)

type ServiceActionDefinitionType string

const (
	ServiceActionDefinitionType_SSM_AUTOMATION ServiceActionDefinitionType = "SSM_AUTOMATION"
)

type ShareStatus string

const (
	ShareStatus_NOT_STARTED           ShareStatus = "NOT_STARTED"
	ShareStatus_IN_PROGRESS           ShareStatus = "IN_PROGRESS"
	ShareStatus_COMPLETED             ShareStatus = "COMPLETED"
	ShareStatus_COMPLETED_WITH_ERRORS ShareStatus = "COMPLETED_WITH_ERRORS"
	ShareStatus_ERROR                 ShareStatus = "ERROR"
)

type SortOrder string

const (
	SortOrder_ASCENDING  SortOrder = "ASCENDING"
	SortOrder_DESCENDING SortOrder = "DESCENDING"
)

type SourceType string

const (
	SourceType_CODESTAR SourceType = "CODESTAR"
)

type StackInstanceStatus string

const (
	StackInstanceStatus_CURRENT    StackInstanceStatus = "CURRENT"
	StackInstanceStatus_OUTDATED   StackInstanceStatus = "OUTDATED"
	StackInstanceStatus_INOPERABLE StackInstanceStatus = "INOPERABLE"
)

type StackSetOperationType string

const (
	StackSetOperationType_CREATE StackSetOperationType = "CREATE"
	StackSetOperationType_UPDATE StackSetOperationType = "UPDATE"
	StackSetOperationType_DELETE StackSetOperationType = "DELETE"
)

type Status string

const (
	Status_AVAILABLE Status = "AVAILABLE"
	Status_CREATING  Status = "CREATING"
	Status_FAILED    Status = "FAILED"
)
