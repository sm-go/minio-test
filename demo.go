package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/encrypt"
)

func main() {
	fmt.Println("test minio")
	endpoint := "127.0.0.1:9000/"
	accessKey := "lgnJnX2QujyiZj0w"
	secretKey := "H3ka494ujRzdsD7bRc0T5KgeYVYOLUZM"

	// Initialize minio client object
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}
	//log.Printf("%#v\n", minioClient) //minioClient is now set up
	bucketName := "toe"
	//found or not
	found, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	if found {
		log.Println("Bucket found")
	} else {
		log.Println("Bucket not found")
	}

	//compose object
	docKey, _ := encrypt.NewSSEC([]byte{1, 2, 3})
	src1 := minio.CopySrcOptions{
		Bucket:               "bucket-one",
		Object:               "object-one",
		VersionID:            "1.0.0",
		MatchETag:            "31624deb84149d2f8ef9c385918b653a",
		NoMatchETag:          "",
		MatchModifiedSince:   time.Time{},
		MatchUnmodifiedSince: time.Time{},
		MatchRange:           false,
		Start:                0,
		End:                  0,
		Encryption:           docKey,
	}
	src2 := minio.CopySrcOptions{
		Bucket:               "bucket-two",
		Object:               "object-two",
		VersionID:            "1.0.0",
		MatchETag:            "31624deb84149d2f8ef9c385918b653a",
		NoMatchETag:          "",
		MatchModifiedSince:   time.Time{},
		MatchUnmodifiedSince: time.Time{},
		MatchRange:           false,
		Start:                0,
		End:                  0,
		Encryption:           docKey,
	}
	srcs := []minio.CopySrcOptions{src1, src2}
	encKey, _ := encrypt.NewSSEC([]byte{8, 9, 0})
	dst := minio.CopyDestOptions{
		Bucket:     "bucket-new",
		Object:     "object-new",
		Encryption: encKey,
	}
	uploadInfo, err := minioClient.ComposeObject(context.Background(), dst, srcs...)
	if err != nil {
		log.Println(err)
	}
	log.Println("Composed object successfully :", uploadInfo)
}
