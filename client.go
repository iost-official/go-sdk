package iost

import (
	"context"

	"github.com/iost-official/iost.go/pb"
	"google.golang.org/grpc"
)

type Client struct {
	asc  rpcpb.ApiServiceClient
	conn *grpc.ClientConn
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Dial(url string) error {
	var err error
	c.conn, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.asc = rpcpb.NewApiServiceClient(c.conn)
	return nil
}

func (c *Client) Close() {
	c.conn.Close()
}

var er = &rpcpb.EmptyRequest{}

func (c *Client) NodeInfo() (*rpcpb.NodeInfoResponse, error) {
	return c.asc.GetNodeInfo(context.Background(), er)
}
func (c *Client) ChainInfo() (*rpcpb.ChainInfoResponse, error) {
	return c.asc.GetChainInfo(context.Background(), er)
}
func (c *Client) TxByHash(hash string) (*rpcpb.TransactionResponse, error) {
	return c.asc.GetTxByHash(context.Background(), &rpcpb.TxHashRequest{
		Hash: hash,
	})
}
func (c *Client) TxReceiptByTxHash(hash string) (*rpcpb.TxReceipt, error) {
	return c.asc.GetTxReceiptByTxHash(context.Background(), &rpcpb.TxHashRequest{
		Hash: hash,
	})
}
func (c *Client) BlockByHash(hash string, complete bool) (*rpcpb.BlockResponse, error) {
	return c.asc.GetBlockByHash(context.Background(), &rpcpb.GetBlockByHashRequest{
		Hash:     hash,
		Complete: complete,
	})
}
func (c *Client) BlockByNumber(number int64, complete bool) (*rpcpb.BlockResponse, error) {
	return c.asc.GetBlockByNumber(context.Background(), &rpcpb.GetBlockByNumberRequest{
		Number:   number,
		Complete: complete,
	})
}
func (c *Client) Account(name string, byLongestChain bool) (*rpcpb.Account, error) {
	return c.asc.GetAccount(context.Background(), &rpcpb.GetAccountRequest{
		Name:           name,
		ByLongestChain: byLongestChain,
	})
}
func (c *Client) TokenBalance(account, token string, longest bool) (*rpcpb.GetTokenBalanceResponse, error) {
	return c.asc.GetTokenBalance(context.Background(), &rpcpb.GetTokenBalanceRequest{
		Account:        account,
		Token:          token,
		ByLongestChain: longest,
	})
}
func (c *Client) Contract(id string, longest bool) (*rpcpb.Contract, error) {
	return c.asc.GetContract(context.Background(), &rpcpb.GetContractRequest{
		Id:             id,
		ByLongestChain: longest,
	})
}
func (c *Client) ContractStorage(id, key, field string, longest bool) (*rpcpb.GetContractStorageResponse, error) {
	return c.asc.GetContractStorage(context.Background(), &rpcpb.GetContractStorageRequest{
		Id:             id,
		Key:            key,
		Field:          field,
		ByLongestChain: longest,
	})
}
func (c *Client) SendTransaction(tx *rpcpb.TransactionRequest) (*rpcpb.SendTransactionResponse, error) {
	return c.asc.SendTransaction(context.Background(), tx)
}
