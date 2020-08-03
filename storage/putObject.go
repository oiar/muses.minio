package storage

import (
	"log"
	"os"

	"github.com/minio/minio-go/v6"
)

func (b *Bucket) PutObject(objectName string, object *os.File) error {
	objectStat, err := object.Stat()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	exists, err := b.CheckBucket(b.bucketName)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if b.strategy == "weightStrategy" {
		c := b.saveByWeight()
		if exists {
			_, err = c.PutObject(b.bucketName, objectName, object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
			if err != nil {
				log.Fatalln(err)
				return err
			}
		} else {
			log.Fatalln("Bucket does not exist.")
		}

	}
	if b.strategy == "multiWriteStrategy" {
		for _, v := range b.getStrategyClients() {
			if exists {
				_, err = v.client.PutObject(b.bucketName, objectName, object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
				if err != nil {
					log.Fatalln(err)
					return err
				}
			} else {
				log.Fatalln("Bucket does not exist.")
			}
		}
	}

	if exists {
		buf, err := b.GetObject(objectName)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		b.setCacheObject(buf, objectName)
	}

	return nil
}
