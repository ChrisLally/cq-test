// Code generated by codegen; DO NOT EDIT.
package services

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

//go:generate mockgen -package=mocks -destination=../mocks/babacloudsdkgo_services_bssopenapi.go -source=babacloudsdkgo_services_bssopenapi.go BssopenapiClient
type BssopenapiClient interface {
	DescribeCostBudgetsSummary(*bssopenapi.DescribeCostBudgetsSummaryRequest) (*bssopenapi.DescribeCostBudgetsSummaryResponse, error)
	DescribeCostBudgetsSummaryWithCallback(*bssopenapi.DescribeCostBudgetsSummaryRequest, func(*bssopenapi.DescribeCostBudgetsSummaryResponse, error)) <-chan int
	DescribeCostBudgetsSummaryWithChan(*bssopenapi.DescribeCostBudgetsSummaryRequest) (<-chan *bssopenapi.DescribeCostBudgetsSummaryResponse, <-chan error)
	DescribeInstanceAmortizedCostByAmortizationPeriod(*bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodRequest) (*bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodResponse, error)
	DescribeInstanceAmortizedCostByAmortizationPeriodWithCallback(*bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodRequest, func(*bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodResponse, error)) <-chan int
	DescribeInstanceAmortizedCostByAmortizationPeriodWithChan(*bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodRequest) (<-chan *bssopenapi.DescribeInstanceAmortizedCostByAmortizationPeriodResponse, <-chan error)
	DescribeInstanceAmortizedCostByConsumePeriod(*bssopenapi.DescribeInstanceAmortizedCostByConsumePeriodRequest) (*bssopenapi.DescribeInstanceAmortizedCostByConsumePeriodResponse, error)
	DescribeInstanceAmortizedCostByConsumePeriodWithCallback(*bssopenapi.DescribeInstanceAmortizedCostByConsumePeriodRequest, func(*bssopenapi.DescribeInstanceAmortizedCostByConsumePeriodResponse, error)) <-chan int
	DescribeInstanceAmortizedCostByConsumePeriodWithChan(*bssopenapi.DescribeInstanceAmortizedCostByConsumePeriodRequest) (<-chan *bssopenapi.DescribeInstanceAmortizedCostByConsumePeriodResponse, <-chan error)
	DescribeInstanceBill(*bssopenapi.DescribeInstanceBillRequest) (*bssopenapi.DescribeInstanceBillResponse, error)
	DescribeInstanceBillWithCallback(*bssopenapi.DescribeInstanceBillRequest, func(*bssopenapi.DescribeInstanceBillResponse, error)) <-chan int
	DescribeInstanceBillWithChan(*bssopenapi.DescribeInstanceBillRequest) (<-chan *bssopenapi.DescribeInstanceBillResponse, <-chan error)
	DescribePricingModule(*bssopenapi.DescribePricingModuleRequest) (*bssopenapi.DescribePricingModuleResponse, error)
	DescribePricingModuleWithCallback(*bssopenapi.DescribePricingModuleRequest, func(*bssopenapi.DescribePricingModuleResponse, error)) <-chan int
	DescribePricingModuleWithChan(*bssopenapi.DescribePricingModuleRequest) (<-chan *bssopenapi.DescribePricingModuleResponse, <-chan error)
	DescribeProductAmortizedCostByAmortizationPeriod(*bssopenapi.DescribeProductAmortizedCostByAmortizationPeriodRequest) (*bssopenapi.DescribeProductAmortizedCostByAmortizationPeriodResponse, error)
	DescribeProductAmortizedCostByAmortizationPeriodWithCallback(*bssopenapi.DescribeProductAmortizedCostByAmortizationPeriodRequest, func(*bssopenapi.DescribeProductAmortizedCostByAmortizationPeriodResponse, error)) <-chan int
	DescribeProductAmortizedCostByAmortizationPeriodWithChan(*bssopenapi.DescribeProductAmortizedCostByAmortizationPeriodRequest) (<-chan *bssopenapi.DescribeProductAmortizedCostByAmortizationPeriodResponse, <-chan error)
	DescribeProductAmortizedCostByConsumePeriod(*bssopenapi.DescribeProductAmortizedCostByConsumePeriodRequest) (*bssopenapi.DescribeProductAmortizedCostByConsumePeriodResponse, error)
	DescribeProductAmortizedCostByConsumePeriodWithCallback(*bssopenapi.DescribeProductAmortizedCostByConsumePeriodRequest, func(*bssopenapi.DescribeProductAmortizedCostByConsumePeriodResponse, error)) <-chan int
	DescribeProductAmortizedCostByConsumePeriodWithChan(*bssopenapi.DescribeProductAmortizedCostByConsumePeriodRequest) (<-chan *bssopenapi.DescribeProductAmortizedCostByConsumePeriodResponse, <-chan error)
	DescribeResourceCoverageDetail(*bssopenapi.DescribeResourceCoverageDetailRequest) (*bssopenapi.DescribeResourceCoverageDetailResponse, error)
	DescribeResourceCoverageDetailWithCallback(*bssopenapi.DescribeResourceCoverageDetailRequest, func(*bssopenapi.DescribeResourceCoverageDetailResponse, error)) <-chan int
	DescribeResourceCoverageDetailWithChan(*bssopenapi.DescribeResourceCoverageDetailRequest) (<-chan *bssopenapi.DescribeResourceCoverageDetailResponse, <-chan error)
	DescribeResourceCoverageTotal(*bssopenapi.DescribeResourceCoverageTotalRequest) (*bssopenapi.DescribeResourceCoverageTotalResponse, error)
	DescribeResourceCoverageTotalWithCallback(*bssopenapi.DescribeResourceCoverageTotalRequest, func(*bssopenapi.DescribeResourceCoverageTotalResponse, error)) <-chan int
	DescribeResourceCoverageTotalWithChan(*bssopenapi.DescribeResourceCoverageTotalRequest) (<-chan *bssopenapi.DescribeResourceCoverageTotalResponse, <-chan error)
	DescribeResourcePackageProduct(*bssopenapi.DescribeResourcePackageProductRequest) (*bssopenapi.DescribeResourcePackageProductResponse, error)
	DescribeResourcePackageProductWithCallback(*bssopenapi.DescribeResourcePackageProductRequest, func(*bssopenapi.DescribeResourcePackageProductResponse, error)) <-chan int
	DescribeResourcePackageProductWithChan(*bssopenapi.DescribeResourcePackageProductRequest) (<-chan *bssopenapi.DescribeResourcePackageProductResponse, <-chan error)
	DescribeResourceUsageDetail(*bssopenapi.DescribeResourceUsageDetailRequest) (*bssopenapi.DescribeResourceUsageDetailResponse, error)
	DescribeResourceUsageDetailWithCallback(*bssopenapi.DescribeResourceUsageDetailRequest, func(*bssopenapi.DescribeResourceUsageDetailResponse, error)) <-chan int
	DescribeResourceUsageDetailWithChan(*bssopenapi.DescribeResourceUsageDetailRequest) (<-chan *bssopenapi.DescribeResourceUsageDetailResponse, <-chan error)
	DescribeResourceUsageTotal(*bssopenapi.DescribeResourceUsageTotalRequest) (*bssopenapi.DescribeResourceUsageTotalResponse, error)
	DescribeResourceUsageTotalWithCallback(*bssopenapi.DescribeResourceUsageTotalRequest, func(*bssopenapi.DescribeResourceUsageTotalResponse, error)) <-chan int
	DescribeResourceUsageTotalWithChan(*bssopenapi.DescribeResourceUsageTotalRequest) (<-chan *bssopenapi.DescribeResourceUsageTotalResponse, <-chan error)
	DescribeSavingsPlansCoverageDetail(*bssopenapi.DescribeSavingsPlansCoverageDetailRequest) (*bssopenapi.DescribeSavingsPlansCoverageDetailResponse, error)
	DescribeSavingsPlansCoverageDetailWithCallback(*bssopenapi.DescribeSavingsPlansCoverageDetailRequest, func(*bssopenapi.DescribeSavingsPlansCoverageDetailResponse, error)) <-chan int
	DescribeSavingsPlansCoverageDetailWithChan(*bssopenapi.DescribeSavingsPlansCoverageDetailRequest) (<-chan *bssopenapi.DescribeSavingsPlansCoverageDetailResponse, <-chan error)
	DescribeSavingsPlansCoverageTotal(*bssopenapi.DescribeSavingsPlansCoverageTotalRequest) (*bssopenapi.DescribeSavingsPlansCoverageTotalResponse, error)
	DescribeSavingsPlansCoverageTotalWithCallback(*bssopenapi.DescribeSavingsPlansCoverageTotalRequest, func(*bssopenapi.DescribeSavingsPlansCoverageTotalResponse, error)) <-chan int
	DescribeSavingsPlansCoverageTotalWithChan(*bssopenapi.DescribeSavingsPlansCoverageTotalRequest) (<-chan *bssopenapi.DescribeSavingsPlansCoverageTotalResponse, <-chan error)
	DescribeSavingsPlansUsageDetail(*bssopenapi.DescribeSavingsPlansUsageDetailRequest) (*bssopenapi.DescribeSavingsPlansUsageDetailResponse, error)
	DescribeSavingsPlansUsageDetailWithCallback(*bssopenapi.DescribeSavingsPlansUsageDetailRequest, func(*bssopenapi.DescribeSavingsPlansUsageDetailResponse, error)) <-chan int
	DescribeSavingsPlansUsageDetailWithChan(*bssopenapi.DescribeSavingsPlansUsageDetailRequest) (<-chan *bssopenapi.DescribeSavingsPlansUsageDetailResponse, <-chan error)
	DescribeSavingsPlansUsageTotal(*bssopenapi.DescribeSavingsPlansUsageTotalRequest) (*bssopenapi.DescribeSavingsPlansUsageTotalResponse, error)
	DescribeSavingsPlansUsageTotalWithCallback(*bssopenapi.DescribeSavingsPlansUsageTotalRequest, func(*bssopenapi.DescribeSavingsPlansUsageTotalResponse, error)) <-chan int
	DescribeSavingsPlansUsageTotalWithChan(*bssopenapi.DescribeSavingsPlansUsageTotalRequest) (<-chan *bssopenapi.DescribeSavingsPlansUsageTotalResponse, <-chan error)
	DescribeSplitItemBill(*bssopenapi.DescribeSplitItemBillRequest) (*bssopenapi.DescribeSplitItemBillResponse, error)
	DescribeSplitItemBillWithCallback(*bssopenapi.DescribeSplitItemBillRequest, func(*bssopenapi.DescribeSplitItemBillResponse, error)) <-chan int
	DescribeSplitItemBillWithChan(*bssopenapi.DescribeSplitItemBillRequest) (<-chan *bssopenapi.DescribeSplitItemBillResponse, <-chan error)
	GetAccountRelation(*bssopenapi.GetAccountRelationRequest) (*bssopenapi.GetAccountRelationResponse, error)
	GetAccountRelationWithCallback(*bssopenapi.GetAccountRelationRequest, func(*bssopenapi.GetAccountRelationResponse, error)) <-chan int
	GetAccountRelationWithChan(*bssopenapi.GetAccountRelationRequest) (<-chan *bssopenapi.GetAccountRelationResponse, <-chan error)
	GetCloseTrace() bool
	GetCustomerAccountInfo(*bssopenapi.GetCustomerAccountInfoRequest) (*bssopenapi.GetCustomerAccountInfoResponse, error)
	GetCustomerAccountInfoWithCallback(*bssopenapi.GetCustomerAccountInfoRequest, func(*bssopenapi.GetCustomerAccountInfoResponse, error)) <-chan int
	GetCustomerAccountInfoWithChan(*bssopenapi.GetCustomerAccountInfoRequest) (<-chan *bssopenapi.GetCustomerAccountInfoResponse, <-chan error)
	GetCustomerList(*bssopenapi.GetCustomerListRequest) (*bssopenapi.GetCustomerListResponse, error)
	GetCustomerListWithCallback(*bssopenapi.GetCustomerListRequest, func(*bssopenapi.GetCustomerListResponse, error)) <-chan int
	GetCustomerListWithChan(*bssopenapi.GetCustomerListRequest) (<-chan *bssopenapi.GetCustomerListResponse, <-chan error)
	GetEndpointRules(string, string) (string, error)
	GetHTTPSInsecure() bool
	GetHttpProxy() string
	GetHttpsProxy() string
	GetLoggerMsg() string
	GetNoProxy() string
	GetOrderDetail(*bssopenapi.GetOrderDetailRequest) (*bssopenapi.GetOrderDetailResponse, error)
	GetOrderDetailWithCallback(*bssopenapi.GetOrderDetailRequest, func(*bssopenapi.GetOrderDetailResponse, error)) <-chan int
	GetOrderDetailWithChan(*bssopenapi.GetOrderDetailRequest) (<-chan *bssopenapi.GetOrderDetailResponse, <-chan error)
	GetPayAsYouGoPrice(*bssopenapi.GetPayAsYouGoPriceRequest) (*bssopenapi.GetPayAsYouGoPriceResponse, error)
	GetPayAsYouGoPriceWithCallback(*bssopenapi.GetPayAsYouGoPriceRequest, func(*bssopenapi.GetPayAsYouGoPriceResponse, error)) <-chan int
	GetPayAsYouGoPriceWithChan(*bssopenapi.GetPayAsYouGoPriceRequest) (<-chan *bssopenapi.GetPayAsYouGoPriceResponse, <-chan error)
	GetResourcePackagePrice(*bssopenapi.GetResourcePackagePriceRequest) (*bssopenapi.GetResourcePackagePriceResponse, error)
	GetResourcePackagePriceWithCallback(*bssopenapi.GetResourcePackagePriceRequest, func(*bssopenapi.GetResourcePackagePriceResponse, error)) <-chan int
	GetResourcePackagePriceWithChan(*bssopenapi.GetResourcePackagePriceRequest) (<-chan *bssopenapi.GetResourcePackagePriceResponse, <-chan error)
	GetSubscriptionPrice(*bssopenapi.GetSubscriptionPriceRequest) (*bssopenapi.GetSubscriptionPriceResponse, error)
	GetSubscriptionPriceWithCallback(*bssopenapi.GetSubscriptionPriceRequest, func(*bssopenapi.GetSubscriptionPriceResponse, error)) <-chan int
	GetSubscriptionPriceWithChan(*bssopenapi.GetSubscriptionPriceRequest) (<-chan *bssopenapi.GetSubscriptionPriceResponse, <-chan error)
	GetTemplate() string
	QueryAccountBalance(*bssopenapi.QueryAccountBalanceRequest) (*bssopenapi.QueryAccountBalanceResponse, error)
	QueryAccountBalanceWithCallback(*bssopenapi.QueryAccountBalanceRequest, func(*bssopenapi.QueryAccountBalanceResponse, error)) <-chan int
	QueryAccountBalanceWithChan(*bssopenapi.QueryAccountBalanceRequest) (<-chan *bssopenapi.QueryAccountBalanceResponse, <-chan error)
	QueryAccountBill(*bssopenapi.QueryAccountBillRequest) (*bssopenapi.QueryAccountBillResponse, error)
	QueryAccountBillWithCallback(*bssopenapi.QueryAccountBillRequest, func(*bssopenapi.QueryAccountBillResponse, error)) <-chan int
	QueryAccountBillWithChan(*bssopenapi.QueryAccountBillRequest) (<-chan *bssopenapi.QueryAccountBillResponse, <-chan error)
	QueryAccountTransactionDetails(*bssopenapi.QueryAccountTransactionDetailsRequest) (*bssopenapi.QueryAccountTransactionDetailsResponse, error)
	QueryAccountTransactionDetailsWithCallback(*bssopenapi.QueryAccountTransactionDetailsRequest, func(*bssopenapi.QueryAccountTransactionDetailsResponse, error)) <-chan int
	QueryAccountTransactionDetailsWithChan(*bssopenapi.QueryAccountTransactionDetailsRequest) (<-chan *bssopenapi.QueryAccountTransactionDetailsResponse, <-chan error)
	QueryAccountTransactions(*bssopenapi.QueryAccountTransactionsRequest) (*bssopenapi.QueryAccountTransactionsResponse, error)
	QueryAccountTransactionsWithCallback(*bssopenapi.QueryAccountTransactionsRequest, func(*bssopenapi.QueryAccountTransactionsResponse, error)) <-chan int
	QueryAccountTransactionsWithChan(*bssopenapi.QueryAccountTransactionsRequest) (<-chan *bssopenapi.QueryAccountTransactionsResponse, <-chan error)
	QueryAvailableInstances(*bssopenapi.QueryAvailableInstancesRequest) (*bssopenapi.QueryAvailableInstancesResponse, error)
	QueryAvailableInstancesWithCallback(*bssopenapi.QueryAvailableInstancesRequest, func(*bssopenapi.QueryAvailableInstancesResponse, error)) <-chan int
	QueryAvailableInstancesWithChan(*bssopenapi.QueryAvailableInstancesRequest) (<-chan *bssopenapi.QueryAvailableInstancesResponse, <-chan error)
	QueryBill(*bssopenapi.QueryBillRequest) (*bssopenapi.QueryBillResponse, error)
	QueryBillOverview(*bssopenapi.QueryBillOverviewRequest) (*bssopenapi.QueryBillOverviewResponse, error)
	QueryBillOverviewWithCallback(*bssopenapi.QueryBillOverviewRequest, func(*bssopenapi.QueryBillOverviewResponse, error)) <-chan int
	QueryBillOverviewWithChan(*bssopenapi.QueryBillOverviewRequest) (<-chan *bssopenapi.QueryBillOverviewResponse, <-chan error)
	QueryBillToOSSSubscription(*bssopenapi.QueryBillToOSSSubscriptionRequest) (*bssopenapi.QueryBillToOSSSubscriptionResponse, error)
	QueryBillToOSSSubscriptionWithCallback(*bssopenapi.QueryBillToOSSSubscriptionRequest, func(*bssopenapi.QueryBillToOSSSubscriptionResponse, error)) <-chan int
	QueryBillToOSSSubscriptionWithChan(*bssopenapi.QueryBillToOSSSubscriptionRequest) (<-chan *bssopenapi.QueryBillToOSSSubscriptionResponse, <-chan error)
	QueryBillWithCallback(*bssopenapi.QueryBillRequest, func(*bssopenapi.QueryBillResponse, error)) <-chan int
	QueryBillWithChan(*bssopenapi.QueryBillRequest) (<-chan *bssopenapi.QueryBillResponse, <-chan error)
	QueryCashCoupons(*bssopenapi.QueryCashCouponsRequest) (*bssopenapi.QueryCashCouponsResponse, error)
	QueryCashCouponsWithCallback(*bssopenapi.QueryCashCouponsRequest, func(*bssopenapi.QueryCashCouponsResponse, error)) <-chan int
	QueryCashCouponsWithChan(*bssopenapi.QueryCashCouponsRequest) (<-chan *bssopenapi.QueryCashCouponsResponse, <-chan error)
	QueryCommodityList(*bssopenapi.QueryCommodityListRequest) (*bssopenapi.QueryCommodityListResponse, error)
	QueryCommodityListWithCallback(*bssopenapi.QueryCommodityListRequest, func(*bssopenapi.QueryCommodityListResponse, error)) <-chan int
	QueryCommodityListWithChan(*bssopenapi.QueryCommodityListRequest) (<-chan *bssopenapi.QueryCommodityListResponse, <-chan error)
	QueryCostUnit(*bssopenapi.QueryCostUnitRequest) (*bssopenapi.QueryCostUnitResponse, error)
	QueryCostUnitResource(*bssopenapi.QueryCostUnitResourceRequest) (*bssopenapi.QueryCostUnitResourceResponse, error)
	QueryCostUnitResourceWithCallback(*bssopenapi.QueryCostUnitResourceRequest, func(*bssopenapi.QueryCostUnitResourceResponse, error)) <-chan int
	QueryCostUnitResourceWithChan(*bssopenapi.QueryCostUnitResourceRequest) (<-chan *bssopenapi.QueryCostUnitResourceResponse, <-chan error)
	QueryCostUnitWithCallback(*bssopenapi.QueryCostUnitRequest, func(*bssopenapi.QueryCostUnitResponse, error)) <-chan int
	QueryCostUnitWithChan(*bssopenapi.QueryCostUnitRequest) (<-chan *bssopenapi.QueryCostUnitResponse, <-chan error)
	QueryCustomerAddressList(*bssopenapi.QueryCustomerAddressListRequest) (*bssopenapi.QueryCustomerAddressListResponse, error)
	QueryCustomerAddressListWithCallback(*bssopenapi.QueryCustomerAddressListRequest, func(*bssopenapi.QueryCustomerAddressListResponse, error)) <-chan int
	QueryCustomerAddressListWithChan(*bssopenapi.QueryCustomerAddressListRequest) (<-chan *bssopenapi.QueryCustomerAddressListResponse, <-chan error)
	QueryDPUtilizationDetail(*bssopenapi.QueryDPUtilizationDetailRequest) (*bssopenapi.QueryDPUtilizationDetailResponse, error)
	QueryDPUtilizationDetailWithCallback(*bssopenapi.QueryDPUtilizationDetailRequest, func(*bssopenapi.QueryDPUtilizationDetailResponse, error)) <-chan int
	QueryDPUtilizationDetailWithChan(*bssopenapi.QueryDPUtilizationDetailRequest) (<-chan *bssopenapi.QueryDPUtilizationDetailResponse, <-chan error)
	QueryEvaluateList(*bssopenapi.QueryEvaluateListRequest) (*bssopenapi.QueryEvaluateListResponse, error)
	QueryEvaluateListWithCallback(*bssopenapi.QueryEvaluateListRequest, func(*bssopenapi.QueryEvaluateListResponse, error)) <-chan int
	QueryEvaluateListWithChan(*bssopenapi.QueryEvaluateListRequest) (<-chan *bssopenapi.QueryEvaluateListResponse, <-chan error)
	QueryFinancialAccountInfo(*bssopenapi.QueryFinancialAccountInfoRequest) (*bssopenapi.QueryFinancialAccountInfoResponse, error)
	QueryFinancialAccountInfoWithCallback(*bssopenapi.QueryFinancialAccountInfoRequest, func(*bssopenapi.QueryFinancialAccountInfoResponse, error)) <-chan int
	QueryFinancialAccountInfoWithChan(*bssopenapi.QueryFinancialAccountInfoRequest) (<-chan *bssopenapi.QueryFinancialAccountInfoResponse, <-chan error)
	QueryInstanceBill(*bssopenapi.QueryInstanceBillRequest) (*bssopenapi.QueryInstanceBillResponse, error)
	QueryInstanceBillWithCallback(*bssopenapi.QueryInstanceBillRequest, func(*bssopenapi.QueryInstanceBillResponse, error)) <-chan int
	QueryInstanceBillWithChan(*bssopenapi.QueryInstanceBillRequest) (<-chan *bssopenapi.QueryInstanceBillResponse, <-chan error)
	QueryInstanceByTag(*bssopenapi.QueryInstanceByTagRequest) (*bssopenapi.QueryInstanceByTagResponse, error)
	QueryInstanceByTagWithCallback(*bssopenapi.QueryInstanceByTagRequest, func(*bssopenapi.QueryInstanceByTagResponse, error)) <-chan int
	QueryInstanceByTagWithChan(*bssopenapi.QueryInstanceByTagRequest) (<-chan *bssopenapi.QueryInstanceByTagResponse, <-chan error)
	QueryInstanceGaapCost(*bssopenapi.QueryInstanceGaapCostRequest) (*bssopenapi.QueryInstanceGaapCostResponse, error)
	QueryInstanceGaapCostWithCallback(*bssopenapi.QueryInstanceGaapCostRequest, func(*bssopenapi.QueryInstanceGaapCostResponse, error)) <-chan int
	QueryInstanceGaapCostWithChan(*bssopenapi.QueryInstanceGaapCostRequest) (<-chan *bssopenapi.QueryInstanceGaapCostResponse, <-chan error)
	QueryInvoicingCustomerList(*bssopenapi.QueryInvoicingCustomerListRequest) (*bssopenapi.QueryInvoicingCustomerListResponse, error)
	QueryInvoicingCustomerListWithCallback(*bssopenapi.QueryInvoicingCustomerListRequest, func(*bssopenapi.QueryInvoicingCustomerListResponse, error)) <-chan int
	QueryInvoicingCustomerListWithChan(*bssopenapi.QueryInvoicingCustomerListRequest) (<-chan *bssopenapi.QueryInvoicingCustomerListResponse, <-chan error)
	QueryOrders(*bssopenapi.QueryOrdersRequest) (*bssopenapi.QueryOrdersResponse, error)
	QueryOrdersWithCallback(*bssopenapi.QueryOrdersRequest, func(*bssopenapi.QueryOrdersResponse, error)) <-chan int
	QueryOrdersWithChan(*bssopenapi.QueryOrdersRequest) (<-chan *bssopenapi.QueryOrdersResponse, <-chan error)
	QueryPermissionList(*bssopenapi.QueryPermissionListRequest) (*bssopenapi.QueryPermissionListResponse, error)
	QueryPermissionListWithCallback(*bssopenapi.QueryPermissionListRequest, func(*bssopenapi.QueryPermissionListResponse, error)) <-chan int
	QueryPermissionListWithChan(*bssopenapi.QueryPermissionListRequest) (<-chan *bssopenapi.QueryPermissionListResponse, <-chan error)
	QueryPrepaidCards(*bssopenapi.QueryPrepaidCardsRequest) (*bssopenapi.QueryPrepaidCardsResponse, error)
	QueryPrepaidCardsWithCallback(*bssopenapi.QueryPrepaidCardsRequest, func(*bssopenapi.QueryPrepaidCardsResponse, error)) <-chan int
	QueryPrepaidCardsWithChan(*bssopenapi.QueryPrepaidCardsRequest) (<-chan *bssopenapi.QueryPrepaidCardsResponse, <-chan error)
	QueryPriceEntityList(*bssopenapi.QueryPriceEntityListRequest) (*bssopenapi.QueryPriceEntityListResponse, error)
	QueryPriceEntityListWithCallback(*bssopenapi.QueryPriceEntityListRequest, func(*bssopenapi.QueryPriceEntityListResponse, error)) <-chan int
	QueryPriceEntityListWithChan(*bssopenapi.QueryPriceEntityListRequest) (<-chan *bssopenapi.QueryPriceEntityListResponse, <-chan error)
	QueryProductList(*bssopenapi.QueryProductListRequest) (*bssopenapi.QueryProductListResponse, error)
	QueryProductListWithCallback(*bssopenapi.QueryProductListRequest, func(*bssopenapi.QueryProductListResponse, error)) <-chan int
	QueryProductListWithChan(*bssopenapi.QueryProductListRequest) (<-chan *bssopenapi.QueryProductListResponse, <-chan error)
	QueryRIUtilizationDetail(*bssopenapi.QueryRIUtilizationDetailRequest) (*bssopenapi.QueryRIUtilizationDetailResponse, error)
	QueryRIUtilizationDetailWithCallback(*bssopenapi.QueryRIUtilizationDetailRequest, func(*bssopenapi.QueryRIUtilizationDetailResponse, error)) <-chan int
	QueryRIUtilizationDetailWithChan(*bssopenapi.QueryRIUtilizationDetailRequest) (<-chan *bssopenapi.QueryRIUtilizationDetailResponse, <-chan error)
	QueryRedeem(*bssopenapi.QueryRedeemRequest) (*bssopenapi.QueryRedeemResponse, error)
	QueryRedeemWithCallback(*bssopenapi.QueryRedeemRequest, func(*bssopenapi.QueryRedeemResponse, error)) <-chan int
	QueryRedeemWithChan(*bssopenapi.QueryRedeemRequest) (<-chan *bssopenapi.QueryRedeemResponse, <-chan error)
	QueryRelationList(*bssopenapi.QueryRelationListRequest) (*bssopenapi.QueryRelationListResponse, error)
	QueryRelationListWithCallback(*bssopenapi.QueryRelationListRequest, func(*bssopenapi.QueryRelationListResponse, error)) <-chan int
	QueryRelationListWithChan(*bssopenapi.QueryRelationListRequest) (<-chan *bssopenapi.QueryRelationListResponse, <-chan error)
	QueryResellerAvailableQuota(*bssopenapi.QueryResellerAvailableQuotaRequest) (*bssopenapi.QueryResellerAvailableQuotaResponse, error)
	QueryResellerAvailableQuotaWithCallback(*bssopenapi.QueryResellerAvailableQuotaRequest, func(*bssopenapi.QueryResellerAvailableQuotaResponse, error)) <-chan int
	QueryResellerAvailableQuotaWithChan(*bssopenapi.QueryResellerAvailableQuotaRequest) (<-chan *bssopenapi.QueryResellerAvailableQuotaResponse, <-chan error)
	QueryResellerUserAlarmThreshold(*bssopenapi.QueryResellerUserAlarmThresholdRequest) (*bssopenapi.QueryResellerUserAlarmThresholdResponse, error)
	QueryResellerUserAlarmThresholdWithCallback(*bssopenapi.QueryResellerUserAlarmThresholdRequest, func(*bssopenapi.QueryResellerUserAlarmThresholdResponse, error)) <-chan int
	QueryResellerUserAlarmThresholdWithChan(*bssopenapi.QueryResellerUserAlarmThresholdRequest) (<-chan *bssopenapi.QueryResellerUserAlarmThresholdResponse, <-chan error)
	QueryResourcePackageInstances(*bssopenapi.QueryResourcePackageInstancesRequest) (*bssopenapi.QueryResourcePackageInstancesResponse, error)
	QueryResourcePackageInstancesWithCallback(*bssopenapi.QueryResourcePackageInstancesRequest, func(*bssopenapi.QueryResourcePackageInstancesResponse, error)) <-chan int
	QueryResourcePackageInstancesWithChan(*bssopenapi.QueryResourcePackageInstancesRequest) (<-chan *bssopenapi.QueryResourcePackageInstancesResponse, <-chan error)
	QuerySavingsPlansDeductLog(*bssopenapi.QuerySavingsPlansDeductLogRequest) (*bssopenapi.QuerySavingsPlansDeductLogResponse, error)
	QuerySavingsPlansDeductLogWithCallback(*bssopenapi.QuerySavingsPlansDeductLogRequest, func(*bssopenapi.QuerySavingsPlansDeductLogResponse, error)) <-chan int
	QuerySavingsPlansDeductLogWithChan(*bssopenapi.QuerySavingsPlansDeductLogRequest) (<-chan *bssopenapi.QuerySavingsPlansDeductLogResponse, <-chan error)
	QuerySavingsPlansDiscount(*bssopenapi.QuerySavingsPlansDiscountRequest) (*bssopenapi.QuerySavingsPlansDiscountResponse, error)
	QuerySavingsPlansDiscountWithCallback(*bssopenapi.QuerySavingsPlansDiscountRequest, func(*bssopenapi.QuerySavingsPlansDiscountResponse, error)) <-chan int
	QuerySavingsPlansDiscountWithChan(*bssopenapi.QuerySavingsPlansDiscountRequest) (<-chan *bssopenapi.QuerySavingsPlansDiscountResponse, <-chan error)
	QuerySavingsPlansInstance(*bssopenapi.QuerySavingsPlansInstanceRequest) (*bssopenapi.QuerySavingsPlansInstanceResponse, error)
	QuerySavingsPlansInstanceWithCallback(*bssopenapi.QuerySavingsPlansInstanceRequest, func(*bssopenapi.QuerySavingsPlansInstanceResponse, error)) <-chan int
	QuerySavingsPlansInstanceWithChan(*bssopenapi.QuerySavingsPlansInstanceRequest) (<-chan *bssopenapi.QuerySavingsPlansInstanceResponse, <-chan error)
	QuerySettleBill(*bssopenapi.QuerySettleBillRequest) (*bssopenapi.QuerySettleBillResponse, error)
	QuerySettleBillWithCallback(*bssopenapi.QuerySettleBillRequest, func(*bssopenapi.QuerySettleBillResponse, error)) <-chan int
	QuerySettleBillWithChan(*bssopenapi.QuerySettleBillRequest) (<-chan *bssopenapi.QuerySettleBillResponse, <-chan error)
	QuerySkuPriceList(*bssopenapi.QuerySkuPriceListRequest) (*bssopenapi.QuerySkuPriceListResponse, error)
	QuerySkuPriceListWithCallback(*bssopenapi.QuerySkuPriceListRequest, func(*bssopenapi.QuerySkuPriceListResponse, error)) <-chan int
	QuerySkuPriceListWithChan(*bssopenapi.QuerySkuPriceListRequest) (<-chan *bssopenapi.QuerySkuPriceListResponse, <-chan error)
	QuerySplitItemBill(*bssopenapi.QuerySplitItemBillRequest) (*bssopenapi.QuerySplitItemBillResponse, error)
	QuerySplitItemBillWithCallback(*bssopenapi.QuerySplitItemBillRequest, func(*bssopenapi.QuerySplitItemBillResponse, error)) <-chan int
	QuerySplitItemBillWithChan(*bssopenapi.QuerySplitItemBillRequest) (<-chan *bssopenapi.QuerySplitItemBillResponse, <-chan error)
	QueryUserOmsData(*bssopenapi.QueryUserOmsDataRequest) (*bssopenapi.QueryUserOmsDataResponse, error)
	QueryUserOmsDataWithCallback(*bssopenapi.QueryUserOmsDataRequest, func(*bssopenapi.QueryUserOmsDataResponse, error)) <-chan int
	QueryUserOmsDataWithChan(*bssopenapi.QueryUserOmsDataRequest) (<-chan *bssopenapi.QueryUserOmsDataResponse, <-chan error)
}
