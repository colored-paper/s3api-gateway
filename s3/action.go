package s3

import (
	"github.com/labstack/echo/v4"
)

type Config struct {
	Name string
}

type Actor interface {
	AbortMultipartUpload(c echo.Context) error
	CompleteMultipartUpload(c echo.Context) error
	CopyObject(c echo.Context) error
	CreateBucket(c echo.Context) error
	CreateMultipartUpload(c echo.Context) error
	DeleteBucket(c echo.Context) error
	DeleteBucketAnalyticsConfiguration(c echo.Context) error
	DeleteBucketCors(c echo.Context) error
	DeleteBucketEncryption(c echo.Context) error
	DeleteBucketIntelligentTieringConfiguration(c echo.Context) error
	DeleteBucketInventoryConfiguration(c echo.Context) error
	DeleteBucketLifecycle(c echo.Context) error
	DeleteBucketMetricsConfiguration(c echo.Context) error
	DeleteBucketOwnershipControls(c echo.Context) error
	DeleteBucketPolicy(c echo.Context) error
	DeleteBucketReplication(c echo.Context) error
	DeleteBucketTagging(c echo.Context) error
	DeleteBucketWebsite(c echo.Context) error
	DeleteObject(c echo.Context) error
	DeleteObjectTagging(c echo.Context) error
	DeleteObjects(c echo.Context) error
	DeletePublicAccessBlock(c echo.Context) error
	GetBucketAccelerateConfiguration(c echo.Context) error
	GetBucketAcl(c echo.Context) error
	GetBucketAnalyticsConfiguration(c echo.Context) error
	GetBucketCors(c echo.Context) error
	GetBucketEncryption(c echo.Context) error
	GetBucketIntelligentTieringConfiguration(c echo.Context) error
	GetBucketInventoryConfiguration(c echo.Context) error
	GetBucketLifecycleConfiguration(c echo.Context) error
	GetBucketLocation(c echo.Context) error
	GetBucketLogging(c echo.Context) error
	GetBucketMetricsConfiguration(c echo.Context) error
	GetBucketNotificationConfiguration(c echo.Context) error
	GetBucketOwnershipControls(c echo.Context) error
	GetBucketPolicy(c echo.Context) error
	GetBucketPolicyStatus(c echo.Context) error
	GetBucketReplication(c echo.Context) error
	GetBucketRequestPayment(c echo.Context) error
	GetBucketTagging(c echo.Context) error
	GetBucketVersioning(c echo.Context) error
	GetBucketWebsite(c echo.Context) error
	GetObject(c echo.Context) error
	GetObjectAcl(c echo.Context) error
	GetObjectAttributes(c echo.Context) error
	GetObjectLegalHold(c echo.Context) error
	GetObjectLockConfiguration(c echo.Context) error
	GetObjectRetention(c echo.Context) error
	GetObjectTagging(c echo.Context) error
	GetObjectTorrent(c echo.Context) error
	GetPublicAccessBlock(c echo.Context) error
	HeadBucket(c echo.Context) error
	HeadObject(c echo.Context) error
	ListBucketAnalyticsConfigurations(c echo.Context) error
	ListBucketIntelligentTieringConfigurations(c echo.Context) error
	ListBucketInventoryConfigurations(c echo.Context) error
	ListBucketMetricsConfigurations(c echo.Context) error
	ListBuckets(c echo.Context) error
	ListMultipartUploads(c echo.Context) error
	ListObjectVersions(c echo.Context) error
	ListObjects(c echo.Context) error
	ListObjectsV2(c echo.Context) error
	ListParts(c echo.Context) error
	PutBucketAccelerateConfiguration(c echo.Context) error
	PutBucketAcl(c echo.Context) error
	PutBucketAnalyticsConfiguration(c echo.Context) error
	PutBucketCors(c echo.Context) error
	PutBucketEncryption(c echo.Context) error
	PutBucketIntelligentTieringConfiguration(c echo.Context) error
	PutBucketInventoryConfiguration(c echo.Context) error
	PutBucketLifecycleConfiguration(c echo.Context) error
	PutBucketLogging(c echo.Context) error
	PutBucketMetricsConfiguration(c echo.Context) error
	PutBucketNotificationConfiguration(c echo.Context) error
	PutBucketOwnershipControls(c echo.Context) error
	PutBucketPolicy(c echo.Context) error
	PutBucketReplication(c echo.Context) error
	PutBucketRequestPayment(c echo.Context) error
	PutBucketTagging(c echo.Context) error
	PutBucketVersioning(c echo.Context) error
	PutBucketWebsite(c echo.Context) error
	PutObject(c echo.Context) error
	PutObjectAcl(c echo.Context) error
	PutObjectLegalHold(c echo.Context) error
	PutObjectLockConfiguration(c echo.Context) error
	PutObjectRetention(c echo.Context) error
	PutObjectTagging(c echo.Context) error
	PutPublicAccessBlock(c echo.Context) error
	RestoreObject(c echo.Context) error
	SelectObjectContent(c echo.Context) error
	UploadPart(c echo.Context) error
	UploadPartCopy(c echo.Context) error
	Wait(c echo.Context) error
	WriteGetObjectResponse(c echo.Context) error
}
