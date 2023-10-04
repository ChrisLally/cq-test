// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"net/http"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

//go:generate mockgen -package=mocks -destination=../mocks/dashboards_api.go -source=dashboards_api.go DashboardsAPIClient
type DashboardsAPIClient interface {
	ListDashboards(context.Context, ...datadogV1.ListDashboardsOptionalParameters) (datadogV1.DashboardSummary, *http.Response, error)
	ListDashboardsWithPagination(context.Context, ...datadogV1.ListDashboardsOptionalParameters) (<-chan datadog.PaginationResult[datadogV1.DashboardSummaryDefinition], func())
}