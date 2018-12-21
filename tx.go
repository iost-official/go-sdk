package iost

import (
	"strconv"

	"time"

	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-sdk/pb"
)

// Config of tx
type Config struct {
	GasLimit   float64
	GasRatio   float64
	Expiration int64
	Delay      int64
}

// DefaultTxConfig .
var DefaultTxConfig = Config{
	GasLimit:   1000000,
	GasRatio:   1,
	Expiration: time.Second.Nanoseconds() * 90,
	Delay:      0,
}

// NewTx make a tx with config
func NewTx(config Config) *rpcpb.TransactionRequest {
	ret := &rpcpb.TransactionRequest{
		Time:          time.Now().UnixNano(),
		Actions:       []*rpcpb.Action{},
		Signers:       []string{},
		GasLimit:      config.GasLimit,
		GasRatio:      config.GasRatio,
		Expiration:    config.Expiration,
		PublisherSigs: []*rpcpb.Signature{},
		Delay:         config.Delay * 1e9,
		AmountLimit:   []*rpcpb.AmountLimit{},
	}
	return ret
}

// AddAction add calls to a tx
func AddAction(t *rpcpb.TransactionRequest, contractID, abi, args string) {
	t.Actions = append(t.Actions, newAction(contractID, abi, args))
}

// AddApprove add approve to a tx
func AddApprove(t *rpcpb.TransactionRequest, token string, amount float64) {
	t.AmountLimit = append(t.AmountLimit, &rpcpb.AmountLimit{
		Token: token,
		Value: amount,
	})
}

func actionToBytes(a *rpcpb.Action) []byte {
	sn := common.NewSimpleNotation()
	sn.WriteString(a.Contract, true)
	sn.WriteString(a.ActionName, true)
	sn.WriteString(a.Data, true)
	return sn.Bytes()
}

func amountToBytes(a *rpcpb.AmountLimit) []byte {
	sn := common.NewSimpleNotation()
	sn.WriteString(a.Token, true)
	sn.WriteString(strconv.FormatFloat(a.Value, 'f', -1, 64), true)
	return sn.Bytes()
}

func signatureToBytes(s *rpcpb.Signature) []byte {
	sn := common.NewSimpleNotation()
	sn.WriteByte(byte(s.Algorithm), true)
	sn.WriteBytes(s.Signature, true)
	sn.WriteBytes(s.PublicKey, true)
	return sn.Bytes()
}

func txToBytes(t *rpcpb.TransactionRequest) []byte {
	sn := common.NewSimpleNotation()
	sn.WriteInt64(t.Time, true)
	sn.WriteInt64(t.Expiration, true)
	sn.WriteInt64(int64(t.GasRatio*100), true)
	sn.WriteInt64(int64(t.GasLimit*100), true)
	sn.WriteInt64(t.Delay, true)
	sn.WriteStringSlice(t.Signers, true)

	actionBytes := make([][]byte, 0, len(t.Actions))
	for _, a := range t.Actions {
		actionBytes = append(actionBytes, actionToBytes(a))
	}
	sn.WriteBytesSlice(actionBytes, false)

	amountBytes := make([][]byte, 0, len(t.AmountLimit))
	for _, a := range t.AmountLimit {
		amountBytes = append(amountBytes, amountToBytes(a))
	}
	sn.WriteBytesSlice(amountBytes, false)

	signBytes := make([][]byte, 0, len(t.Signatures))
	for _, sig := range t.Signatures {
		signBytes = append(signBytes, signatureToBytes(sig))
	}
	sn.WriteBytesSlice(signBytes, false)

	return sn.Bytes()
}

func newAction(contract string, name string, data string) *rpcpb.Action {
	return &rpcpb.Action{
		Contract:   contract,
		ActionName: name,
		Data:       data,
	}
}
