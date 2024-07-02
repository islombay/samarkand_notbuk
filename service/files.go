package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/config"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Files struct {
	minio      *minio.Client
	bucket     string
	storage    storage.StorageInterface
	log        logs.LoggerInterface
	server_url string
}

func NewFiles(storage storage.StorageInterface, log logs.LoggerInterface, cfg config.Config) *Files {
	endpoint := fmt.Sprintf("%s:%d", cfg.Minio.Host, cfg.Minio.Port)
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: cfg.Minio.SSL,
	})

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	bucketName := "samarkand-notbuk"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if !exists && errBucketExists != nil {
			log.Error(errBucketExists.Error())
			os.Exit(1)
		}
	}
	log.Debug("Successfully created bucket in minio")

	return &Files{
		minio:      minioClient,
		bucket:     bucketName,
		storage:    storage,
		log:        log,
		server_url: cfg.Server.Public,
	}
}

func (srv *Files) Create(ctx context.Context, m models.UploadFile) status.Status {
	file, err := m.File.Open()
	if err != nil {
		srv.log.Error("could not open file", logs.Error(err))
		return status.StatusInternal
	}
	defer file.Close()

	random_id := uuid.NewString()

	contentType := m.File.Header.Get("Content-Type")
	fmt.Println(contentType)

	_, err = srv.minio.PutObject(ctx, srv.bucket, random_id, file, m.File.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		srv.log.Error("could not put object to minio", logs.Error(err))
		return status.StatusInternal
	}

	file.Close()

	file_url := fmt.Sprintf("https://%s/api/v1/files/%s", srv.server_url, random_id)

	m.ID = random_id
	m.FileURL = file_url

	fileObj, err := srv.storage.Files().Create(ctx, m)
	if err != nil {

		// delete object from minio
		if err := srv.minio.RemoveObject(ctx, srv.bucket, random_id, minio.RemoveObjectOptions{}); err != nil {
			srv.log.Error("could not delete object from minio", logs.Error(err))
		}

		return status.StatusInternal
	}

	return status.StatusOk.AddData(fileObj)
}

func (srv *Files) GetByID(ctx context.Context, file_id string) status.Status {
	obj, err := srv.minio.GetObject(ctx, srv.bucket, file_id, minio.GetObjectOptions{})
	if err != nil {
		srv.log.Error("could not get object from minio", logs.Error(err))
		return status.StatusInternal
	}

	objInfo, err := obj.Stat()
	if err != nil {
		srv.log.Error("could not get object info", logs.Error(err))
		return status.StatusInternal
	}

	return status.StatusOk.AddFileObject(obj).AddFileObjectInfo(objInfo)
}

func (srv *Files) GetList(ctx context.Context, p models.Pagination) status.Status {
	pagination, err := srv.storage.Files().GetList(ctx, p)
	if err != nil {
		return status.StatusInternal
	}

	return status.StatusOk.AddData(pagination.Data).AddCount(pagination.Count)
}

func (srv *Files) Delete(ctx context.Context, id string) status.Status {
	if err := srv.storage.Files().Delete(ctx, id); err != nil {
		return status.StatusInternal
	}

	if err := srv.minio.RemoveObject(ctx, srv.bucket, id, minio.RemoveObjectOptions{}); err != nil {
		srv.log.Error("could not remove object from minio", logs.Error(err))
	}

	return status.StatusOk
}
