// Code generated by codegen; DO NOT EDIT.
package services

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//go:generate mockgen -package=mocks -destination=../mocks/yunossgosdk_oss.go -source=yunossgosdk_oss.go OssClient
type OssClient interface {
	GetBucketACL(string, ...oss.Option) (oss.GetBucketACLResult, error)
	GetBucketAccessMonitor(string, ...oss.Option) (oss.GetBucketAccessMonitorResult, error)
	GetBucketAccessMonitorXml(string, ...oss.Option) (string, error)
	GetBucketAsyncTask(string, string, ...oss.Option) (oss.AsynFetchTaskInfo, error)
	GetBucketCORS(string, ...oss.Option) (oss.GetBucketCORSResult, error)
	GetBucketCORSXml(string, ...oss.Option) (string, error)
	GetBucketCname(string, ...oss.Option) (string, error)
	GetBucketCnameToken(string, string, ...oss.Option) (oss.GetBucketCnameTokenResult, error)
	GetBucketEncryption(string, ...oss.Option) (oss.GetBucketEncryptionResult, error)
	GetBucketInfo(string, ...oss.Option) (oss.GetBucketInfoResult, error)
	GetBucketInventory(string, string, ...oss.Option) (oss.InventoryConfiguration, error)
	GetBucketInventoryXml(string, string, ...oss.Option) (string, error)
	GetBucketLifecycle(string, ...oss.Option) (oss.GetBucketLifecycleResult, error)
	GetBucketLifecycleXml(string, ...oss.Option) (string, error)
	GetBucketLocation(string, ...oss.Option) (string, error)
	GetBucketLogging(string, ...oss.Option) (oss.GetBucketLoggingResult, error)
	GetBucketPolicy(string, ...oss.Option) (string, error)
	GetBucketQosInfo(string, ...oss.Option) (oss.BucketQoSConfiguration, error)
	GetBucketReferer(string, ...oss.Option) (oss.GetBucketRefererResult, error)
	GetBucketReplication(string, ...oss.Option) (string, error)
	GetBucketReplicationLocation(string, ...oss.Option) (string, error)
	GetBucketReplicationProgress(string, string, ...oss.Option) (string, error)
	GetBucketRequestPayment(string, ...oss.Option) (oss.RequestPaymentConfiguration, error)
	GetBucketResourceGroup(string, ...oss.Option) (oss.GetBucketResourceGroupResult, error)
	GetBucketResourceGroupXml(string, ...oss.Option) (string, error)
	GetBucketStat(string, ...oss.Option) (oss.GetBucketStatResult, error)
	GetBucketStyle(string, string, ...oss.Option) (oss.GetBucketStyleResult, error)
	GetBucketStyleXml(string, string, ...oss.Option) (string, error)
	GetBucketTagging(string, ...oss.Option) (oss.GetBucketTaggingResult, error)
	GetBucketTransferAcc(string, ...oss.Option) (oss.TransferAccConfiguration, error)
	GetBucketVersioning(string, ...oss.Option) (oss.GetBucketVersioningResult, error)
	GetBucketWebsite(string, ...oss.Option) (oss.GetBucketWebsiteResult, error)
	GetBucketWebsiteXml(string, ...oss.Option) (string, error)
	GetBucketWorm(string, ...oss.Option) (oss.WormConfiguration, error)
	GetMetaQueryStatus(string, ...oss.Option) (oss.GetMetaQueryStatusResult, error)
	GetUserQoSInfo(...oss.Option) (oss.UserQoSConfiguration, error)
	ListBucketCname(string, ...oss.Option) (oss.ListBucketCnameResult, error)
	ListBucketInventory(string, string, ...oss.Option) (oss.ListInventoryConfigurationsResult, error)
	ListBucketInventoryXml(string, string, ...oss.Option) (string, error)
	ListBucketStyle(string, ...oss.Option) (oss.GetBucketListStyleResult, error)
	ListBucketStyleXml(string, ...oss.Option) (string, error)
	ListBuckets(...oss.Option) (oss.ListBucketsResult, error)
	ListCloudBoxes(...oss.Option) (oss.ListCloudBoxResult, error)
}
