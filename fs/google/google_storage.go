package google

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/go-bamboo/pkg/log"
)

type Client struct {
	c      *Conf
	client *storage.Client
}

func New(c *Conf) *Client {
	client, err := storage.NewClient(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		c:      c,
		client: client,
	}
}

func (c *Client) Close() {
	c.client.Close()
}

// SaveWebSiteImage 转存网络上的视频或图片
func (c *Client) SaveWebSiteImage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", err
	}
	fileNameSuffix := url[strings.LastIndex(url, "."):]
	if len(fileNameSuffix) > 5 {
		//url后缀长度大于5，基本不可能是后缀名了
		contentType := resp.Header.Get("content-type")
		if strings.HasPrefix(contentType, "image") {
			switch contentType {
			case "image/jpg":
				fileNameSuffix = ".jpg"
			case "image/jpeg":
				fileNameSuffix = ".jpg"
			case "image/png":
				fileNameSuffix = ".png"
			case "image/gif":
				fileNameSuffix = ".gif"
			case "image/bmp":
				fileNameSuffix = ".bmp"
			case "image/svg":
				fileNameSuffix = ".svg"
			case "image/webp":
				fileNameSuffix = ".webp"
			case "image/tiff":
				fileNameSuffix = ".tiff"
			case "video/mp4":
				fileNameSuffix = ".mp4"
			case "video/mpeg4":
				fileNameSuffix = ".mp4"
			case "video/webm":
				fileNameSuffix = ".webm"
			case "application/ogg":
				fileNameSuffix = ".ogg"
			default:
				fileNameSuffix = ""
			}
		} else {
			return "", errors.New("Resource is not a image")
		}
	}
	// 用URL的hash作为文件名
	fileName := fmt.Sprintf("%x", md5.Sum([]byte(url))) + fileNameSuffix
	data, err := ioutil.ReadAll(resp.Body)
	return c.UploadImage(data, fileName), nil
}

// UploadImage 保存图片到google storage上
func (c *Client) UploadImage(file []byte, fileName string) string {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("storage.NewClient: %v", err)
		return ""
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.Bucket(c.c.Bucket).Object(fileName).NewWriter(ctx)
	_, err = wc.Write(file)
	if err := wc.Close(); err != nil {
		fmt.Printf("Writer.Close: %v", err)
		return ""
	}
	return fmt.Sprintf("/%s/%s", c.c.Bucket, fileName)
}
