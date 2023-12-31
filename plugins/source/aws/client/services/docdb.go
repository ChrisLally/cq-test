// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
)

//go:generate mockgen -package=mocks -destination=../mocks/docdb.go -source=docdb.go DocdbClient
type DocdbClient interface {
	DescribeCertificates(context.Context, *docdb.DescribeCertificatesInput, ...func(*docdb.Options)) (*docdb.DescribeCertificatesOutput, error)
	DescribeDBClusterParameterGroups(context.Context, *docdb.DescribeDBClusterParameterGroupsInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameters(context.Context, *docdb.DescribeDBClusterParametersInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterSnapshotAttributes(context.Context, *docdb.DescribeDBClusterSnapshotAttributesInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshots(context.Context, *docdb.DescribeDBClusterSnapshotsInput, ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusters(context.Context, *docdb.DescribeDBClustersInput, ...func(*docdb.Options)) (*docdb.DescribeDBClustersOutput, error)
	DescribeDBEngineVersions(context.Context, *docdb.DescribeDBEngineVersionsInput, ...func(*docdb.Options)) (*docdb.DescribeDBEngineVersionsOutput, error)
	DescribeDBInstances(context.Context, *docdb.DescribeDBInstancesInput, ...func(*docdb.Options)) (*docdb.DescribeDBInstancesOutput, error)
	DescribeDBSubnetGroups(context.Context, *docdb.DescribeDBSubnetGroupsInput, ...func(*docdb.Options)) (*docdb.DescribeDBSubnetGroupsOutput, error)
	DescribeEngineDefaultClusterParameters(context.Context, *docdb.DescribeEngineDefaultClusterParametersInput, ...func(*docdb.Options)) (*docdb.DescribeEngineDefaultClusterParametersOutput, error)
	DescribeEventCategories(context.Context, *docdb.DescribeEventCategoriesInput, ...func(*docdb.Options)) (*docdb.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *docdb.DescribeEventSubscriptionsInput, ...func(*docdb.Options)) (*docdb.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *docdb.DescribeEventsInput, ...func(*docdb.Options)) (*docdb.DescribeEventsOutput, error)
	DescribeGlobalClusters(context.Context, *docdb.DescribeGlobalClustersInput, ...func(*docdb.Options)) (*docdb.DescribeGlobalClustersOutput, error)
	DescribeOrderableDBInstanceOptions(context.Context, *docdb.DescribeOrderableDBInstanceOptionsInput, ...func(*docdb.Options)) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribePendingMaintenanceActions(context.Context, *docdb.DescribePendingMaintenanceActionsInput, ...func(*docdb.Options)) (*docdb.DescribePendingMaintenanceActionsOutput, error)
	ListTagsForResource(context.Context, *docdb.ListTagsForResourceInput, ...func(*docdb.Options)) (*docdb.ListTagsForResourceOutput, error)
}
