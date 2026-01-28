package s3_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-bamboo/pkg/fs"
	. "github.com/go-bamboo/pkg/fs/s3"
)

var s fs.FileStorage
var ctx = context.Background()

func TestMain(m *testing.M) {
	var err error
	c := fs.Conf{}
	s, err = New(&c)
	if err != nil {
		return
	}
	ret := m.Run()
	os.Exit(ret)
}

//func TestApiCopyFile(t *testing.T) {
//	uri, err := s.CopyObject(ctx, "ipfs-v2.halonft.art", "haloworld/L2", "0.jpg", "https://s3.us-east-2.amazonaws.com/ipfs-avatar.halonft.art/avatar_doc/random/haloworld/image/0.jpg")
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(uri)
//	t.Log("success")
//}
//
//func TestUploadImage(t *testing.T) {
//	proxy, err := url.Parse("http://127.0.0.1:7890")
//	contentType, err := s.UploadImage("https://ipfs.io/ipfs/QmYD9AtzyQPjSa9jfZcZq88gSaRssdhGmKqQifUDjGFfXm/sleepy.png",
//		"images/bsc/test.jpg", proxy)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(contentType)
//	t.Log("success")
//}
//
//func TestUploadBytes(t *testing.T) {
//	path, err := s.UploadBytes(context.TODO(), "2039.json", []byte(`{
//    "attributes": [
//        {
//            "trait_type": "platform",
//            "value": "web"
//        },
//        {
//            "trait_type": "class",
//            "value": "avatar"
//        },
//        {
//            "trait_type": "Mood",
//            "value": "100"
//        },
//        {
//            "trait_type": "gender",
//            "value": "male"
//        },
//        {
//            "trait_type": "Energy",
//            "value": "100"
//        },
//        {
//            "trait_type": "workAt",
//            "value": "1683122523"
//        },
//        {
//            "trait_type": "Health",
//            "value": "100"
//        },
//        {
//            "trait_type": "HealthMax",
//            "value": "100"
//        },
//        {
//            "trait_type": "ip",
//            "value": "halo"
//        },
//        {
//            "trait_type": "Ability",
//            "value": "1"
//        }
//    ],
//    "description": "Role in HALOWORLD, the founding citizen.",
//    "image": "https://ipfs-v2.halonft.art/bsc_v2/haloworld/L1/image/2039.jpg",
//    "animation_url": "https://ipfs-v2.halonft.art/bsc_v2/haloworld/L1/gif/2039.gif",
//    "modified_url": "https://ipfs-v2.halonft.art/bsc_v2/haloworld/L1/res/2039.main",
//    "original_url": "https://ipfs-v2.halonft.art/bsc_v2/haloworld/L1/original/original_2039.pak",
//    "token_id": 2039,
//    "name": "Origines Citizen"
//}`))
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("upload bytes", path)
//}
//
//func TestSignature(t *testing.T) {
//	/*funcSignature := []byte("test()")
//	hash := crypto.Keccak256Hash(funcSignature)
//	t.Log(hash.Hex())*/
//	httpClient := http.Client{
//		Transport: &http2.Transport{
//			AllowHTTP: true,
//			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
//				return net.Dial(network, addr)
//			},
//		},
//		Timeout: 120 * time.Second,
//	}
//	request, err := http.NewRequest("GET", "https://ipfs.io/ipfs/QmYD9AtzyQPjSa9jfZcZq88gSaRssdhGmKqQifUDjGFfXm/sleepy.png", nil)
//	if err != nil {
//		t.Log(err)
//		return
//	}
//	request.Header.Add("Host", "ipfs.io")
//	request.Header.Add("Cache-Control", "no-cache")
//	request.Header.Add("Postman-Token", "06468734-9e3c-4945-8438-a6b98861d88e")
//	resp, respError := httpClient.Do(request)
//	if respError != nil || resp == nil || resp.StatusCode/100 > 3 {
//		t.Log(respError)
//	}
//}
