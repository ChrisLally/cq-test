// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
)

//go:generate mockgen -package=mocks -destination=../mocks/appflow.go -source=appflow.go AppflowClient
type AppflowClient interface {
	DescribeConnector(context.Context, *appflow.DescribeConnectorInput, ...func(*appflow.Options)) (*appflow.DescribeConnectorOutput, error)
	DescribeConnectorEntity(context.Context, *appflow.DescribeConnectorEntityInput, ...func(*appflow.Options)) (*appflow.DescribeConnectorEntityOutput, error)
	DescribeConnectorProfiles(context.Context, *appflow.DescribeConnectorProfilesInput, ...func(*appflow.Options)) (*appflow.DescribeConnectorProfilesOutput, error)
	DescribeConnectors(context.Context, *appflow.DescribeConnectorsInput, ...func(*appflow.Options)) (*appflow.DescribeConnectorsOutput, error)
	DescribeFlow(context.Context, *appflow.DescribeFlowInput, ...func(*appflow.Options)) (*appflow.DescribeFlowOutput, error)
	DescribeFlowExecutionRecords(context.Context, *appflow.DescribeFlowExecutionRecordsInput, ...func(*appflow.Options)) (*appflow.DescribeFlowExecutionRecordsOutput, error)
	ListConnectorEntities(context.Context, *appflow.ListConnectorEntitiesInput, ...func(*appflow.Options)) (*appflow.ListConnectorEntitiesOutput, error)
	ListConnectors(context.Context, *appflow.ListConnectorsInput, ...func(*appflow.Options)) (*appflow.ListConnectorsOutput, error)
	ListFlows(context.Context, *appflow.ListFlowsInput, ...func(*appflow.Options)) (*appflow.ListFlowsOutput, error)
	ListTagsForResource(context.Context, *appflow.ListTagsForResourceInput, ...func(*appflow.Options)) (*appflow.ListTagsForResourceOutput, error)
}
