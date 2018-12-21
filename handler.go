package iost

import (
	"errors"
	"time"

	"github.com/iost-official/go-sdk/pb"
)

// Handler handle tx with polling
type Handler struct {
	client    *Client
	tx        *rpcpb.TransactionRequest
	ChSuccess chan *rpcpb.TxReceipt
	ChFailed  chan error
	ChPending chan *rpcpb.SendTransactionResponse
}

// NewHandler make a handler
func NewHandler(tx *rpcpb.TransactionRequest, client *Client) *Handler {
	return &Handler{
		client:    client,
		tx:        tx,
		ChSuccess: make(chan *rpcpb.TxReceipt),
		ChFailed:  make(chan error),
		ChPending: make(chan *rpcpb.SendTransactionResponse),
	}
}

// Send sync send and get the tx hash
func (h *Handler) Send() (string, error) {
	res, err := h.client.SendTransaction(h.tx)
	if err != nil {
		return "", err
	}
	return res.Hash, nil
}

// SendAndListen polling result after send
func (h *Handler) SendAndListen(interval time.Duration, times int) {
	res, err := h.client.SendTransaction(h.tx)
	if err != nil {
		h.ChFailed <- err
		return
	}
	for i := 0; i < times; i++ {
		tr, err := h.client.TxReceiptByTxHash(res.Hash)
		if err != nil {
			continue
		}
		if tr.StatusCode != 0 {
			h.ChFailed <- errors.New(tr.Message)
			return
		}
		h.ChSuccess <- tr
		return
	}
}
