package iost

import (
	"testing"

	"encoding/json"

	"github.com/iost-official/go-iost/common"
)

func TestGet(t *testing.T) {
	client := NewClient()
	err := client.Dial("47.244.109.92:30002")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client.ChainInfo())
	t.Log(client.NodeInfo())
	// ...
	t.Log(client.TxByHash("abc"))
	client.Close()
}

func TestSendTx(t *testing.T) {
	client := NewClient()
	err := client.Dial("47.244.109.92:30002")
	if err != nil {
		t.Fatal("error in dial", err)
	}
	defer client.Close()

	tx := NewTx(Config{
		GasRatio:   1,
		GasLimit:   100000,
		Delay:      0,
		Expiration: 90,
	})
	args, err := json.Marshal([]string{"iost", "admin", "admin", "10.000", ""})
	if err != nil {
		t.Fatal(err)
	}
	AddAction(tx, "token.iost", "transfer", string(args))

	kc := NewKeychain("admin")
	kc.AddKey(common.Base58Decode("2yquS3ySrGWPEKywCPzX4RTJugqRh7kJSo5aehsLYPEWkUxBWA39oMrZ7ZxuM4fgyXYs2cPwh5n8aNNpH5x2VyK1"), "active")

	kc.SignTx(tx)

	handler := NewHandler(tx, client)
	hash, err := handler.Send()
	t.Log(hash, err)
}
