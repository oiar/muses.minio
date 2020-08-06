package main

import (
	"github.com/silverswords/muses.minio/storage"
	"log"
)

func main() {
	b := storage.NewBucket("test", "config.yaml", "../")
	exists, err := b.CheckBucket("test")
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
