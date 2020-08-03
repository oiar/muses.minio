package storage

import (
	"log"
	"context"
	"github.com/minio/minio-go/v6"
)

type Bucket struct {
	bucketName string
	location   string
	objectCache
	strategy string
	configInfo
}

func NewBucket(bucketName, location, strategy, configName, configPath string) *Bucket {
	ctx := context.TODO()
	return &Bucket{
		bucketName: bucketName,
		location:   location,
		objectCache: objectCache{
			ctx,
		},
		strategy: strategy,
		configInfo: configInfo{
			configName: configName,
			configPath: configPath,
		},
	}
}

func (b *Bucket) MakeBucket() error {
	for _, v := range b.getStrategyClients() {
		err := v.client.MakeBucket(b.bucketName, b.location)
		if err != nil {
			return err
		}
	}
	return nil
}

// func (b *Bucket) MakeBucket() error {
// 	for _, v := range b.strategyClients {
// 		err := v.client.MakeBucket(b.bucketName, b.location)
// 		if err != nil {
// 			log.Fatalln(err)
// 			return err
// 		}
// 	}
// 	return nil
// }

func (b *Bucket) CheckBucket(bucketName string) (bool, error) {
	var exists bool
	var err error
	for _, v := range b.getStrategyClients() {
		exists, err = v.client.BucketExists(bucketName)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return exists, nil
}

func (b *Bucket) ListedBucket() []minio.BucketInfo {
	minioClient := b.getStrategyClients()[0].client
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatalln(err)
	}

	return buckets
}

func (b *Bucket) RemoveBucket(bucketName string) error {
	for _, v := range b.getStrategyClients() {
		err := v.client.RemoveBucket(bucketName)
		if err != nil {
			return err
		}
	}

	return nil
}
