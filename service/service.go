package service

type MinioService interface {
	GetFile(bucket, path string) ([]byte, error)
	PutFile(bucket, path string, data []byte) error
	CreateBucket(name string) error
	BucketExists(name string) bool
	GetListBucket() error
}
