package rabbitmq

import (
	"context"
	"strings"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/protox"
	"github.com/go-kratos/kratos/v2/metadata"
)

func Server(msg []byte) (context.Context, context.CancelFunc, []byte, error) {
	var d Data
	if err := protox.Unmarshal(msg, &d); err != nil {
		log.Errorf("[rabbitmq][listener] Unmarshal: %v", err)
		return nil, nil, nil, err
	}
	// metadata
	var md metadata.Metadata = metadata.Metadata{}
	if len(d.Md) > 0 {
		for k, v := range d.Md {
			vals := strings.Split(v, ",")
			for _, val := range vals {
				md.Add(k, val)
			}
		}
	}
	c, cf := context.WithTimeout(context.TODO(), time.Second*10)
	c = metadata.NewServerContext(c, md)
	return c, cf, d.Body, nil
}

func Client(ctx context.Context, msg []byte) []byte {
	d := Data{
		Md:   map[string]string{},
		Body: msg,
	}
	if md, ok := metadata.FromClientContext(ctx); ok {
		md.Range(func(k string, v []string) bool {
			d.Md[k] = strings.Join(v, ",")
			return true
		})
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		md.Range(func(k string, v []string) bool {
			d.Md[k] = strings.Join(v, ",")
			return true
		})
	}
	body, _ := protox.Marshal(&d)
	return body
}
