package s3

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-bamboo/pkg/log"
)

type S3Session struct {
	s3         *s3.Client
	bucket     string
	domain     string
	cloudFront string
}

func New(accessKey, secretKey, bucket, domain, cloudFront string) (s3Sess *S3Session, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		//config.WithSharedConfigProfile(opts.profile),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
				SessionToken:    "",
			},
		}),
	)
	if err != nil {
		return
	}
	s3c := s3.NewFromConfig(cfg)
	s3Sess = &S3Session{
		s3:         s3c,
		bucket:     bucket,
		domain:     domain,
		cloudFront: cloudFront,
	}
	return
}

var contentTypeReg = regexp.MustCompile("(video|image|audio)/.+")

func (s3Sess *S3Session) UploadImage(imagePath string, fileName string, proxy *url.URL) (string, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
		Timeout: 2 * time.Second,
	}
	request, err := http.NewRequest("GET", imagePath, nil)
	if err != nil {
		return "", err
	}
	resp, respError := httpClient.Do(request)
	// 默认没有请求到则使用代理进行请求
	if (respError != nil || resp == nil || resp.StatusCode/100 > 3) && proxy != nil {
		proxyURL := http.ProxyURL(proxy)
		httpClientWithProxy := &http.Client{
			Transport: &http.Transport{
				Proxy: proxyURL,
				Dial: (&net.Dialer{
					Timeout:   60 * time.Second,
					Deadline:  time.Now().Add(6 * time.Second),
					KeepAlive: 4 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 4 * time.Second,
			},
			Timeout: 120 * time.Second,
		}
		request, err = http.NewRequest("GET", imagePath, nil)
		if err != nil {
			return "", err
		}
		resp, respError = httpClientWithProxy.Do(request)
	}

	if respError != nil {
		return "", respError
	}
	if resp == nil || resp.StatusCode/100 > 3 {
		return "", errors.New(fmt.Sprintf("Get Resource Error,StatusCode:%d", resp.StatusCode))
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(data)

	if !contentTypeReg.MatchString(contentType) {
		return "", errors.New("Wrong ContentType:" + contentType)
	}
	s := s3Sess
	var objBody *bytes.Reader
	objBody = bytes.NewReader(data)

	contentLength := resp.ContentLength
	if contentLength == -1 {
		contentLength = objBody.Size()
	}

	_, err = s3Sess.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(fileName),
		Body:        objBody,
		ContentType: aws.String(contentType),
		//ContentLength: aws.Int64(contentLength),
		//ContentDisposition: aws.String("attachment"),
	})

	if err != nil {
		return "", err
	}
	return contentType, nil
}

var (
	allowFileExt         = map[string]int{".png": 1, ".PNG": 1, ".jpg": 1, ".JPG": 1, ".jpeg": 1, ".JPEG": 1, ".gif": 1, ".GIF": 1, ".svg": 1}
	allowAdminFileExt    = map[string]int{".png": 1, ".PNG": 1, ".jpg": 1, ".JPG": 1, ".jpeg": 1, ".JPEG": 1, ".gif": 1, ".GIF": 1, ".mp4": 1, ".svg": 1}
	allowLazyMintFileExt = map[string]string{".png": "image", ".PNG": "image", ".jpg": "image", ".JPG": "image",
		".jpeg": "image", ".JPEG": "image",
		".gif": "image", ".GIF": "image",
		".svg": "svg", ".webp": "webp",
		".mp4": "video", ".mp3": "audio",
	}
	NotAllowExt = errors.New("not allow file extenstion")
)

func (s3Sess *S3Session) UploadBytes(dir, fileName string, data []byte) (string, error) {
	if fileName == "" {
		fileName = hex.EncodeToString(md5.New().Sum(data))
	}
	path := fmt.Sprintf("%s/%s", dir, fileName)
	contentType := aws.String(http.DetectContentType(data))
	size := len(data)
	output, err := s3Sess.s3.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket:        &s3Sess.bucket,
			Key:           &path,
			Body:          bytes.NewReader(data),
			ContentType:   contentType,
			ContentLength: int64(size),
		})
	if err != nil {
		return "", err
	}
	log.Info("upload data to s3", output.VersionId)
	return s3Sess.domain + "/" + s3Sess.bucket + "/" + path, nil
}

func (s3Sess *S3Session) ApiCopyFile(dstBucket, srcUrl, dir, filename string) (string, error) {
	fileName := fmt.Sprintf("%s/%s", dir, filename)
	_, err := s3Sess.s3.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(dstBucket),
		Key:        aws.String(fileName),
		CopySource: &srcUrl,
	})
	if err != nil {
		return "", err
	}
	return "https://" + dstBucket + "/" + fileName, nil
}

func (s3Sess *S3Session) ApiUploadBytes(bucket, dir, filename string, data []byte) (string, error) {
	fileName := fmt.Sprintf("%s/%s", dir, filename)
	_, err := s3Sess.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(fileName),
		Body:          bytes.NewReader(data),
		ContentType:   aws.String(http.DetectContentType(data)),
		ContentLength: int64(len(data)),
	})
	if err != nil {
		return "", err
	}
	return "https://" + bucket + "/" + fileName, nil
}

func (s3Sess *S3Session) ApiUpload(file multipart.File, fileHeader *multipart.FileHeader, dir string) (string, string, error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := path.Ext(originFilename)
	if _, ok := allowFileExt[ext]; !ok {
		return "", "", NotAllowExt
	}
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	sh := md5.New()
	sh.Write(buffer)
	imageNameHash := hex.EncodeToString(sh.Sum([]byte("")))
	s := s3Sess
	fileName := fmt.Sprintf("%s/%s%s", dir, imageNameHash, ext)
	_, err := s3Sess.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(fileName),
		Body:          bytes.NewReader(buffer),
		ContentType:   aws.String(http.DetectContentType(buffer)),
		ContentLength: size,
	})
	if err != nil {
		return "", "", err
	}
	return s3Sess.domain + "/" + s3Sess.bucket + "/" + fileName, s3Sess.cloudFront + "/" + fileName, nil
}

func (s3Sess *S3Session) ApiUploadAvatarDoc(file multipart.File, fileHeader *multipart.FileHeader, dir string, userId, docId int64) (string, string, error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := path.Ext(originFilename)
	if _, ok := allowFileExt[ext]; !ok {
		return "", "", NotAllowExt
	}
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	sh := md5.New()
	sh.Write(buffer)
	s := s3Sess
	fileName := fmt.Sprintf("%s/%d/%d%s", dir, userId, docId, ext)
	_, err := s3Sess.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(fileName),
		Body:          bytes.NewReader(buffer),
		ContentType:   aws.String(http.DetectContentType(buffer)),
		ContentLength: size,
	})
	if err != nil {
		return "", "", err
	}
	return s3Sess.domain + "/" + s3Sess.bucket + "/" + fileName, s3Sess.cloudFront + "/" + fileName, nil
}

func (s3Sess *S3Session) AdminUploadResource(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := path.Ext(originFilename)
	if _, ok := allowAdminFileExt[ext]; !ok {
		return "", NotAllowExt
	}
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	sh := md5.New()
	sh.Write(buffer)
	imageNameHash := hex.EncodeToString(sh.Sum([]byte("")))
	s := s3Sess
	fileName := fmt.Sprintf("images/official/%s%s", imageNameHash, ext)
	_, err := s3Sess.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(fileName),
		Body:          bytes.NewReader(buffer),
		ContentType:   aws.String(http.DetectContentType(buffer)),
		ContentLength: size,
	})
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (s3Sess *S3Session) LazyMintUploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	originFilename := filepath.Base(fileHeader.Filename)
	ext := strings.ToLower(path.Ext(originFilename))
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	sh := md5.New()
	sh.Write(buffer)
	imageNameHash := hex.EncodeToString(sh.Sum([]byte("")))
	s := s3Sess
	fileName := fmt.Sprintf("images/lazy_mint/%s%s", imageNameHash, ext)
	_, err := s3Sess.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(fileName),
		Body:          bytes.NewReader(buffer),
		ContentType:   aws.String(http.DetectContentType(buffer)),
		ContentLength: size,
	})
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (s3Sess *S3Session) ResizeImage(imageUrl string, fileName string, width, height int) error {
	return nil
}
