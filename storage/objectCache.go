package storage

import (
	"sync"

	"github.com/minio/minio-go/v6"
)

type bucketObjectCache struct {
	// mutex is used for handling the concurrent
	// read/write requests for cache.
	sync.RWMutex

	items map[string]*minio.Object
}

func (b *Bucket) cacheGet(objectName string) *minio.Object {
	b.RLock()
	defer b.RUnlock()

	filePath := objectName
	minioObject := b.items[filePath]

	return minioObject
}

func (b *Bucket) cacheSave(objectName string, minioObject *minio.Object) {
	b.Lock()
	defer b.Unlock()

	filePath := objectName
	// minioObject, err := b.GetObject(bucketName, objectName)
	if minioObject != nil {
		b.items[filePath] = minioObject
	}
}

func (b *Bucket) cacheDelete(objectName string) {
	b.Lock()
	defer b.Unlock()

	delete(b.items, objectName)
}