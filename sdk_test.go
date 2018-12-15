package iost

import (
	"testing"
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
