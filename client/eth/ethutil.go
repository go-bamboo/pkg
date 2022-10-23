package eth

import (
	"math/big"
	"net/http"

	"github.com/go-bamboo/pkg/log"
	"github.com/onrik/ethrpc"
	"go.uber.org/zap/zapcore"
)

type Client struct {
	c *ethrpc.EthRPC
}

func New(url string) *Client {
	c := ethrpc.New(url,
		ethrpc.WithHttpClient(http.DefaultClient),
	)
	return &Client{
		c: c,
	}
}

func NewWithLogger(core zapcore.Core, url string) *Client {
	logger := log.NewLogger(core, 1)
	c := ethrpc.New(url,
		ethrpc.WithLogger(logger),
		ethrpc.WithHttpClient(http.DefaultClient),
	)
	return &Client{
		c: c,
	}
}

func (c *Client) CheckHash(hash string) (status int, blocknumber int, err error) {
	receipt, err := c.c.EthGetTransactionReceipt(hash)
	if err != nil { //交易成功
		err = WrapError(err)
		return
	}
	log.Debugw("print receipt", "status", receipt.Status, "gasUsed", receipt.GasUsed, "logsBloom", receipt.LogsBloom)
	// logsBloom
	// 0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	// 0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	if receipt.Status == "0x1" {
		status = 1
	} else if receipt.Status == "0x0" { //交易失败
		status = 2
	}
	blocknumber = receipt.BlockNumber
	return
}

func (c *Client) EthSendRawTransaction(sign string) (hash string, err error) {
	hash, err = c.c.EthSendRawTransaction(sign)
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthBlockNumber() (id int, err error) {
	id, err = c.c.EthBlockNumber()
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthGetTransactionCount(address, block string) (no int, err error) {
	no, err = c.c.EthGetTransactionCount(address, block)
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthGasPrice() (gasPrice big.Int, err error) {
	gasPrice, err = c.c.EthGasPrice()
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthGetTransactionByHash(hash string) (tx *ethrpc.Transaction, err error) {
	tx, err = c.c.EthGetTransactionByHash(hash)
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthGetTransactionReceipt(hash string) (tx *ethrpc.TransactionReceipt, err error) {
	tx, err = c.c.EthGetTransactionReceipt(hash)
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthSendTransaction(transaction ethrpc.T) (hash string, err error) {
	hash, err = c.c.EthSendTransaction(transaction)
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthGetBalance(address, block string) (balance big.Int, err error) {
	balance, err = c.c.EthGetBalance(address, block)
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

func (c *Client) EthGetBlockByNumber(number int, withTransactions bool) (block *ethrpc.Block, err error) {
	block, err = c.c.EthGetBlockByNumber(number, withTransactions)
	if err != nil {
		err = WrapError(err)
		return
	}
	return
}

// ChainId retrieves the current chain ID for transaction replay protection.
// func (c *Client) ChainID() (*big.Int, error) {
// 	var result hexutil.Big
// 	err := c.c.Call(ctx, &result, "eth_chainId")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return (*big.Int)(&result), err
// }
