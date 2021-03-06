package skeleton

import (
	"context"
	"net/http"

	"github.com/coloered-paper/s3api-gateway/pkg/s3"
	"github.com/labstack/echo/v4"
)

type Actor struct{}

func New(ctx context.Context, cfg s3.Config) *Actor {
	return &Actor{}
}

func (a Actor) Health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (a Actor) AbortMultipartUpload(c echo.Context) error {
	_ = s3.MultipartUpload{
		Bucket:   c.Param("bucket"),
		Object:   c.Param("key"),
		UploadID: c.QueryParams().Get("uploadId"),
	}
	return c.NoContent(http.StatusNoContent)
}

func (a Actor) CompleteMultipartUpload(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) CopyObject(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) CreateBucket(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) CreateMultipartUpload(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucket(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketAnalyticsConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketCors(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketEncryption(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketIntelligentTieringConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketInventoryConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketLifecycle(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketMetricsConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketOwnershipControls(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketPolicy(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketReplication(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketTagging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteBucketWebsite(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteObject(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteObjectTagging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeleteObjects(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) DeletePublicAccessBlock(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketAccelerateConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketAcl(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketAnalyticsConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketCors(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketEncryption(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketIntelligentTieringConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketInventoryConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketLifecycleConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketLocation(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketLogging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketMetricsConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketNotificationConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketOwnershipControls(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketPolicy(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketPolicyStatus(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketReplication(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketRequestPayment(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketTagging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketVersioning(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetBucketWebsite(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObject(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObjectAcl(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObjectAttributes(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObjectLegalHold(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObjectLockConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObjectRetention(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObjectTagging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetObjectTorrent(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) GetPublicAccessBlock(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) HeadBucket(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) HeadObject(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListBucketAnalyticsConfigurations(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListBucketIntelligentTieringConfigurations(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListBucketInventoryConfigurations(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListBucketMetricsConfigurations(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListBuckets(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListMultipartUploads(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListObjectVersions(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListObjects(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListObjectsV2(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) ListParts(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketAccelerateConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketAcl(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketAnalyticsConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketCors(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketEncryption(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketIntelligentTieringConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketInventoryConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketLifecycleConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketLogging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketMetricsConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketNotificationConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketOwnershipControls(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketPolicy(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketReplication(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketRequestPayment(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketTagging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketVersioning(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutBucketWebsite(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutObject(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutObjectAcl(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutObjectLegalHold(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutObjectLockConfiguration(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutObjectRetention(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutObjectTagging(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) PutPublicAccessBlock(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) RestoreObject(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) SelectObjectContent(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) UploadPart(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) UploadPartCopy(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) Wait(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a Actor) WriteGetObjectResponse(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
