package iost

import (
	"context"

	"github.com/iost-official/go-sdk/pb"
	"google.golang.org/grpc"
)

// Client the gRPC client wrapper
type Client struct {
	asc  rpcpb.ApiServiceClient
	conn *grpc.ClientConn
}

// NewClient ...
func NewClient() *Client {
	return &Client{}
}

// Dial to url, must gRPC API, eg: localhost:30002
func (c *Client) Dial(url string) error {
	var err error
	c.conn, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.asc = rpcpb.NewApiServiceClient(c.conn)
	return nil
}

// Close connection
func (c *Client) Close() {
	c.conn.Close()
}

var er = &rpcpb.EmptyRequest{}

// NodeInfo .
func (c *Client) NodeInfo() (*rpcpb.NodeInfoResponse, error) {
	return c.asc.GetNodeInfo(context.Background(), er)
}

// ChainInfo .
func (c *Client) ChainInfo() (*rpcpb.ChainInfoResponse, error) {
	return c.asc.GetChainInfo(context.Background(), er)
}

// TxByHash .
func (c *Client) TxByHash(hash string) (*rpcpb.TransactionResponse, error) {
	return c.asc.GetTxByHash(context.Background(), &rpcpb.TxHashRequest{
		Hash: hash,
	})
}

// TxReceiptByTxHash .
func (c *Client) TxReceiptByTxHash(hash string) (*rpcpb.TxReceipt, error) {
	return c.asc.GetTxReceiptByTxHash(context.Background(), &rpcpb.TxHashRequest{
		Hash: hash,
	})
}

// BlockByHash .
func (c *Client) BlockByHash(hash string, complete bool) (*rpcpb.BlockResponse, error) {
	return c.asc.GetBlockByHash(context.Background(), &rpcpb.GetBlockByHashRequest{
		Hash:     hash,
		Complete: complete,
	})
}

// BlockByNumber .
func (c *ClienClient) BlockByNumber(number int64, complete bool) (*rpcpb.BlockResponse, error) {
	return c.asc.GetBlockByNumber(context.Background(), &rpcpb.GetBlockByNumberRequest{
		Number:   number,
		Complete: complete,
	})
}

// Account .
func (c *Client) Account(name string, byLongestChain bool) (*rpcpb.Account, error) {
	return c.asc.GetAccount(context.Background(), &rpcpb.GetAccountRequest{
		Name:           name,
		ByLongestChain: byLongestChain,
	})
}

// TokenBalance .
func (c *Client) TokenBalance(account, token string, longest bool) (*rpcpb.GetTokenBalanceResponse, error) {
	return c.asc.GetTokenBalance(context.Background(), &rpcpb.GetTokenBalanceRequest{
		Account:        account,
		Token:          token,
		ByLongestChain: longest,
	})
}

// Contract .
func (c *Client) Contract(id string, longest bool) (*rpcpb.Contract, error) {
	return c.asc.GetContract(context.Background(), &rpcpb.GetContractRequest{
		Id:             id,
		ByLongestChain: longest,
	})
}

// ContractStorage .
func (c *Client) ContractStorage(id, key, field string, longest bool) (*rpcpb.GetContractStorageResponse, error) {
	return c.asc.GetContractStorage(context.Background(), &rpcpb.GetContractStorageRequest{
		Id:             id,
		Key:            key,
		Field:          field,
		ByLongestChain: longest,
	})
}

// SendTransaction .
func (c *Client) SendTransaction(tx *rpcpb.TransactionRequest) (*rpcpb.SendTransactionResponse, error) {
	return c.asc.SendTransaction(context.Background(), tx)
}
