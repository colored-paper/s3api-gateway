package s3

type MultipartUpload struct {
	Bucket   string
	Object   string
	UploadID string
}
