package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/valyala/fasthttp"
	"gitlab.33.cn/chain33/chain33/common/merkle"
)

type LatestBlock struct {
	Hash       string   `json:"hash"`
	Time       int64    `json:"time"`
	BlockIndex uint64   `json:"block_index"`
	Height     uint64   `json:"height"`
	TxIndexes  []uint64 `json:"txIndexes"`
}

type Header struct {
	Hash         string  `json:"hash"`
	Ver          uint64  `json:"ver"`
	PrevBlock    string  `json:"prev_block"`
	MerkleRoot   string  `json:"mrkl_root"`
	Time         int64   `json:"time"`
	Bits         int64   `json:"bits"`
	Fee          float64 `json:"fee,omitempty"`
	Nonce        int64   `json:"nonce"`
	TxNum        uint64  `json:"n_tx"`
	Size         uint64  `json:"size"`
	BlockIndex   uint64  `json:"block_index"`
	MainChain    bool    `json:"main_chain"`
	Height       uint64  `json:"height"`
	ReceivedTime int64   `json:"received_time"`
	RelayedBy    string  `json:"relayed_by"`
}

type Block struct {
	Header
	Tx []TransactionDetails `json:"tx"`
}

type TransactionDetails struct {
	LockTime  int64     `json:"lock_time"`
	Ver       uint64    `json:"ver"`
	Size      uint64    `json:"size"`
	Inputs    []TxInput `json:"inputs"`
	Weight    int64     `json:"weight"`
	Time      int64     `json:"time"`
	TxIndex   uint64    `json:"tx_index"`
	VinSz     uint64    `json:"vin_sz"`
	Hash      string    `json:"hash"`
	VoutSz    uint64    `json:"vout_sz"`
	RelayedBy string    `json:"relayed_by"`
	Outs      []TxOut   `json:"out"`
}

type TxInput struct {
	Sequence int64  `json:"sequence"`
	Witness  string `json:"witness"`
	Script   string `json:"script"`
}

type TxOut struct {
	Spent   bool   `json:"spent"`
	TxIndex uint64 `json:"tx_index"`
	Type    int    `json:"type"`
	Address string `json:"addr"`
	Value   uint64 `json:"value"`
	N       uint64 `json:"n"`
	Script  string `json:"script"`
}

type Blocks struct {
	Blocks []Block `json:"blocks"`
}

func main() {
	client := &fasthttp.Client{TLSConfig: &tls.Config{InsecureSkipVerify: true}}
	// url := "https://blockchain.info/rawblock/000000000000000000223cbf6c473a4dedf6ebe19d17879eafc41f38ad1fdf1d"
	// url := "https://blockchain.info/latestblock"
	url := "https://blockchain.info/block-height/111110?format=json"
	status, body, err := client.Get(nil, url)
	if err != nil {
		fmt.Printf("error %#v", err)
	}
	fmt.Printf("status %#v\n", status)
	fmt.Printf("data %#v\n", string(body))
	var blocks = Blocks{}
	json.Unmarshal(body, &blocks)
	latestBlock := blocks.Blocks[0]

	fmt.Printf("latestBlock %#v\n", latestBlock)
	fmt.Printf("latestBlock merkeroot %#v\n", latestBlock.MerkleRoot)

	txs := make([][]byte, 0, len(latestBlock.Tx))

	for _, tx := range latestBlock.Tx {
		fmt.Println(tx.Hash)
		hash, err := chainhash.NewHashFromStr(tx.Hash)
		if err != nil {
			fmt.Println(err)
		}
		txs = append(txs, hash.CloneBytes())
	}
	fmt.Printf("%v\n", txs)

	rootHash := merkle.GetMerkleRoot(txs)

	fmt.Printf("roothash %#v\n", rootHash)

	merkeRoot, err := chainhash.NewHash(rootHash)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("roothash = %#v\n", merkeRoot.String())

	if merkeRoot.String() == latestBlock.MerkleRoot {
		fmt.Printf("------")
	}
	// fmt.Printf("data %#v\n", latestBlock)
}
