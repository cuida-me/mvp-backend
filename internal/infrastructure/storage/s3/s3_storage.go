package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
)

type S3Storage struct {
	region string
	bucket string
}

func NewS3Storage(region, bucket string) *S3Storage {
	return &S3Storage{
		region: region,
		bucket: bucket,
	}
}

func (s *S3Storage) Save(filename, folder string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.region),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	contentType, err := s.getFileContentType(file)
	if err != nil {

	}

	params := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket + "/" + folder),
		Key:         aws.String(filename),
		Body:        file,
		ACL:         aws.String("public-read"),
		ContentType: &contentType,
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return err
	}

	fmt.Printf("Arquivo enviado com sucesso para o bucket %s com chave %s\n", filename, folder)
	return nil
}

func (s *S3Storage) Delete(filename, folder string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.region),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket + "/" + folder),
		Key:    aws.String(filename),
	}

	_, err = svc.DeleteObject(params)
	if err != nil {
		return err
	}

	fmt.Printf("Objeto %s excluído do bucket %s\n", filename, folder)
	return nil
}

func (s *S3Storage) getFileContentType(ouput *os.File) (string, error) {
	buf := make([]byte, 512)

	_, err := ouput.Read(buf)

	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buf)

	return contentType, nil
}
