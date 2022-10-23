package binance

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/tools"

	"github.com/ethereum/go-ethereum/event"
	"github.com/gorilla/websocket"
)

type Service interface {
	Unmarshal(data []byte) (interface{}, error)
}

// Client 单个 websocket 信息
type Client struct {
	APIKey    string
	SecretKey string
	BaseURL   string
	Id        string
	RpcId     int64
	conn      *websocket.Conn
	lock      sync.RWMutex
	message   chan []byte // 订阅数据
	feed      event.Feed
	ctx       context.Context
	ctxCancel context.CancelFunc
	wg        sync.WaitGroup
	unpack    map[int64]Service
}

func New(opts ...Option) *Client {
	defaultOpts := options{}
	for _, o := range opts {
		o(&defaultOpts)
	}
	cCtx, cCancel := context.WithCancel(context.TODO())
	return &Client{
		APIKey:    defaultOpts.apiKey,
		SecretKey: defaultOpts.secretKey,
		BaseURL:   defaultOpts.baseURL,
		Id:        tools.GetUUID(),
		message:   make(chan []byte, 256),
		ctx:       cCtx,
		ctxCancel: cCancel,
		unpack:    map[int64]Service{},
	}
}

func (c *Client) Send(data []byte, id int64, callback Service) {
	c.unpack[id] = callback
	c.message <- data
}

func (c *Client) Subscribe(ch interface{}) event.Subscription {
	return c.feed.Subscribe(ch)
}

func (c *Client) Start() error {
	c.wg.Add(3)
	go c.watchConn()
	go c.read()
	go c.write()
	return nil
}

func (c *Client) Stop() error {
	c.ctxCancel()
	c.wg.Wait()
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			log.Errorf("client [%s] disconnect err: %s", c.Id, err)
			return err
		}
	}
	return nil
}

func (c *Client) watchConn() {
	defer func() {
		c.wg.Done()
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("client watchConn call panic: %v\n%s\n", err, buf)
			log.Errorf("%s", pl)
		}
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// try conn
			c.lock.Lock()
			if c.conn == nil {
				conn, _, err := websocket.DefaultDialer.Dial(c.BaseURL, nil)
				if err != nil {
					log.Errorf("err = %v", err)
					c.lock.Unlock()
					time.Sleep(1 * time.Second)
					continue
				} else {
					c.conn = conn
					c.lock.Unlock()
				}
			} else {
				c.lock.Unlock()
			}

		}
	}
}

// 读信息，从 websocket 连接直接读取数据
func (c *Client) read() {
	defer func() {
		c.wg.Done()
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("client read call panic: %v\n%s\n", err, buf)
			log.Errorf("%s", pl)
		}
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			c.lock.RLock()
			if c.conn == nil {
				time.Sleep(1 * time.Second)
				continue
			}
			c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
			messageType, message, err := c.conn.ReadMessage()
			c.lock.RUnlock()
			if err != nil {
				log.Errorf("err = %v", err)
				time.Sleep(1 * time.Second)
				continue
			}
			if messageType == websocket.CloseMessage {
				c.conn = nil
				time.Sleep(1 * time.Second)
				continue
			}
			// c.log.Infof("client [%s] receive message: %s", c.Id, string(message))
			ret, err := c.UnpackTrade(message)
			if err != nil {
				log.Errorf("err = %v", err)
				time.Sleep(1 * time.Second)
				continue
			}
			c.feed.Send(ret)
		}
	}
}

// 写信息，从 channel 变量 Send 中读取数据写入 websocket 连接
func (c *Client) write() {
	defer func() {
		c.wg.Done()
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("client write call panic: %v\n%s\n", err, buf)
			log.Errorf("%s", pl)
		}
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		case message, ok := <-c.message:
			if !ok {
				c.lock.RLock()
				if c.conn == nil {
					c.lock.RUnlock()
					return
				}
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Errorf("err = %v", err)
				}
				c.lock.RUnlock()
				return
			}
			log.Infof("client [%s] write message: %s", c.Id, string(message))
			c.lock.RLock()
			if c.conn == nil {
				time.Sleep(1 * time.Second)
				continue
			}
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			c.lock.RUnlock()
			if err != nil {
				log.Errorf("client [%s] write message err: %s", c.Id, err)
			}
		}
	}
}
