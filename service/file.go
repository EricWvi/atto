package service

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func UploadFile(name string, content []byte, contentType string) (link string, err error) {
	endpoint := viper.GetString("oss.endpoint")
	accessKeyID := viper.GetString("oss.access")
	secretAccessKey := viper.GetString("oss.secret")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Error(err)
		return
	}

	ctx := context.Background()
	bucketName := viper.GetString("oss.bucket")
	reader := bytes.NewReader(content)

	info, err := minioClient.PutObject(ctx, bucketName, name, reader, int64(len(content)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error(err)
		return
	}
	log.Debug("Successfully uploaded %s with size %d\n", name, info.Size)
	link = viper.GetString("oss.protocol") + "://" + endpoint + "/" + bucketName + "/" + info.Key

	return
}

//func DownloadFile(location int, link string) (content []byte, err error) {
//	root := viper.GetString("localoss.location")
//	path := filepath.Join(root, link)
//	return os.ReadFile(path)
//}
