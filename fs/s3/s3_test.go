package s3

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"golang.org/x/net/http2"
)

var s *S3Session
var ctx = context.Background()

func TestMain(m *testing.M) {
	var err error
	c := Conf{
		Key:        "AKIA5ACGRTIJ5WAMFTOK",
		Secret:     "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",
		Bucket:     "chat-test.lifeform.cc",
		Domain:     "https://s3.us-east-2.amazonaws.com",
		CloudFront: "https://chat-test.lifeform.cc",
	}
	s, err = New(&c)
	if err != nil {
		return
	}
	ret := m.Run()
	os.Exit(ret)
}

func TestApiCopyFile(t *testing.T) {
	//proxy, err := url.Parse("http://127.0.0.1:7890")
	url, err := s.ApiCopyFile("ipfs.lifeform.cc", "ipfs-avatar.halonft.art/user_upload/86cc1f4b85a9d0940b7858c1685ff273.jpg",
		"bsc_test_stage/user_upload", "debug.png")
	if err != nil {

		t.Fatal(err)
	}
	t.Log(url)
	t.Log("success")
}
func TestUploadImage(t *testing.T) {
	proxy, err := url.Parse("http://127.0.0.1:7890")
	contentType, err := s.UploadImage("https://ipfs.io/ipfs/QmYD9AtzyQPjSa9jfZcZq88gSaRssdhGmKqQifUDjGFfXm/sleepy.png",
		"images/bsc/test.jpg", proxy)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(contentType)
	t.Log("success")
}

func TestUploadBytes(t *testing.T) {
	path, err := s.UploadBytes("metadata", "1.json", []byte(`{"image" : "https://a.com/a.jpg", "avatar" : "https://t.com/t.jpg"}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("upload bytes", path)
}

func TestSignature(t *testing.T) {
	/*funcSignature := []byte("test()")
	hash := crypto.Keccak256Hash(funcSignature)
	t.Log(hash.Hex())*/
	httpClient := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
		Timeout: 120 * time.Second,
	}
	request, err := http.NewRequest("GET", "https://ipfs.io/ipfs/QmYD9AtzyQPjSa9jfZcZq88gSaRssdhGmKqQifUDjGFfXm/sleepy.png", nil)
	if err != nil {
		t.Log(err)
		return
	}
	request.Header.Add("Host", "ipfs.io")
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Postman-Token", "06468734-9e3c-4945-8438-a6b98861d88e")
	resp, respError := httpClient.Do(request)
	if respError != nil || resp == nil || resp.StatusCode/100 > 3 {
		t.Log(respError)
	}
}
