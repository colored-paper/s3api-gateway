package router

func (r *Router) Handler() {
	r.engine.GET("/api/v1/health", r.s3Service.Health)
	r.engine.DELETE("/:bucket/:key", r.s3Service.AbortMultipartUpload)
	r.engine.PUT("/:bucket", r.s3Service.CreateBucket)
}
