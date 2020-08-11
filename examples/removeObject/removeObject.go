package main

import (
	"github.com/silverswords/muses.minio/storage"
	"log"
)

func main() {
	b := storage.NewBucketConfig("test", "config.yaml", "../", storage.OtherOptions{})
	exists, err := b.CheckBucket()
	if exists && err != nil {
		log.Println("errors in CheckBucket", err)
	}
	if !exists {
		err = b.MakeBucket()
		if err != nil {
			log.Println("errors in MakeBucket", err)
		}
	}

	err = b.RemoveObject("moon")
	if err != nil {
		log.Println("errors in RemoveObject", err)
	}
}
