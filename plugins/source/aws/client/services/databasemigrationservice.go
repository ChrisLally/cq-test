// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

//go:generate mockgen -package=mocks -destination=../mocks/databasemigrationservice.go -source=databasemigrationservice.go DatabasemigrationserviceClient
type DatabasemigrationserviceClient interface {
	DescribeAccountAttributes(context.Context, *databasemigrationservice.DescribeAccountAttributesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeAccountAttributesOutput, error)
	DescribeApplicableIndividualAssessments(context.Context, *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error)
	DescribeCertificates(context.Context, *databasemigrationservice.DescribeCertificatesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeCertificatesOutput, error)
	DescribeConnections(context.Context, *databasemigrationservice.DescribeConnectionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeConnectionsOutput, error)
	DescribeConversionConfiguration(context.Context, *databasemigrationservice.DescribeConversionConfigurationInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeConversionConfigurationOutput, error)
	DescribeDataProviders(context.Context, *databasemigrationservice.DescribeDataProvidersInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeDataProvidersOutput, error)
	DescribeEndpointSettings(context.Context, *databasemigrationservice.DescribeEndpointSettingsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointSettingsOutput, error)
	DescribeEndpointTypes(context.Context, *databasemigrationservice.DescribeEndpointTypesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointTypesOutput, error)
	DescribeEndpoints(context.Context, *databasemigrationservice.DescribeEndpointsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEndpointsOutput, error)
	DescribeEngineVersions(context.Context, *databasemigrationservice.DescribeEngineVersionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEngineVersionsOutput, error)
	DescribeEventCategories(context.Context, *databasemigrationservice.DescribeEventCategoriesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *databasemigrationservice.DescribeEventSubscriptionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *databasemigrationservice.DescribeEventsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeEventsOutput, error)
	DescribeExtensionPackAssociations(context.Context, *databasemigrationservice.DescribeExtensionPackAssociationsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeExtensionPackAssociationsOutput, error)
	DescribeFleetAdvisorCollectors(context.Context, *databasemigrationservice.DescribeFleetAdvisorCollectorsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorCollectorsOutput, error)
	DescribeFleetAdvisorDatabases(context.Context, *databasemigrationservice.DescribeFleetAdvisorDatabasesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorDatabasesOutput, error)
	DescribeFleetAdvisorLsaAnalysis(context.Context, *databasemigrationservice.DescribeFleetAdvisorLsaAnalysisInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorLsaAnalysisOutput, error)
	DescribeFleetAdvisorSchemaObjectSummary(context.Context, *databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorSchemaObjectSummaryOutput, error)
	DescribeFleetAdvisorSchemas(context.Context, *databasemigrationservice.DescribeFleetAdvisorSchemasInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeFleetAdvisorSchemasOutput, error)
	DescribeInstanceProfiles(context.Context, *databasemigrationservice.DescribeInstanceProfilesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeInstanceProfilesOutput, error)
	DescribeMetadataModelAssessments(context.Context, *databasemigrationservice.DescribeMetadataModelAssessmentsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeMetadataModelAssessmentsOutput, error)
	DescribeMetadataModelConversions(context.Context, *databasemigrationservice.DescribeMetadataModelConversionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeMetadataModelConversionsOutput, error)
	DescribeMetadataModelExportsAsScript(context.Context, *databasemigrationservice.DescribeMetadataModelExportsAsScriptInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeMetadataModelExportsAsScriptOutput, error)
	DescribeMetadataModelExportsToTarget(context.Context, *databasemigrationservice.DescribeMetadataModelExportsToTargetInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeMetadataModelExportsToTargetOutput, error)
	DescribeMetadataModelImports(context.Context, *databasemigrationservice.DescribeMetadataModelImportsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeMetadataModelImportsOutput, error)
	DescribeMigrationProjects(context.Context, *databasemigrationservice.DescribeMigrationProjectsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeMigrationProjectsOutput, error)
	DescribeOrderableReplicationInstances(context.Context, *databasemigrationservice.DescribeOrderableReplicationInstancesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)
	DescribePendingMaintenanceActions(context.Context, *databasemigrationservice.DescribePendingMaintenanceActionsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error)
	DescribeRecommendationLimitations(context.Context, *databasemigrationservice.DescribeRecommendationLimitationsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeRecommendationLimitationsOutput, error)
	DescribeRecommendations(context.Context, *databasemigrationservice.DescribeRecommendationsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeRecommendationsOutput, error)
	DescribeRefreshSchemasStatus(context.Context, *databasemigrationservice.DescribeRefreshSchemasStatusInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)
	DescribeReplicationConfigs(context.Context, *databasemigrationservice.DescribeReplicationConfigsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationConfigsOutput, error)
	DescribeReplicationInstanceTaskLogs(context.Context, *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error)
	DescribeReplicationInstances(context.Context, *databasemigrationservice.DescribeReplicationInstancesInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	DescribeReplicationSubnetGroups(context.Context, *databasemigrationservice.DescribeReplicationSubnetGroupsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)
	DescribeReplicationTableStatistics(context.Context, *databasemigrationservice.DescribeReplicationTableStatisticsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTableStatisticsOutput, error)
	DescribeReplicationTaskAssessmentResults(context.Context, *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error)
	DescribeReplicationTaskAssessmentRuns(context.Context, *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error)
	DescribeReplicationTaskIndividualAssessments(context.Context, *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error)
	DescribeReplicationTasks(context.Context, *databasemigrationservice.DescribeReplicationTasksInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationTasksOutput, error)
	DescribeReplications(context.Context, *databasemigrationservice.DescribeReplicationsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationsOutput, error)
	DescribeSchemas(context.Context, *databasemigrationservice.DescribeSchemasInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeSchemasOutput, error)
	DescribeTableStatistics(context.Context, *databasemigrationservice.DescribeTableStatisticsInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeTableStatisticsOutput, error)
	ListTagsForResource(context.Context, *databasemigrationservice.ListTagsForResourceInput, ...func(*databasemigrationservice.Options)) (*databasemigrationservice.ListTagsForResourceOutput, error)
}
