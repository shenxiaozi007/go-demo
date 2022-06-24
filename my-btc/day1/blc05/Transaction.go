package blc05

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

//step 1 创建transaction结构体
type Transaction struct {
	//1 交易id
	TxId []byte
	//2 输入
	Vins []*TxInput
	//3 输出
	Vouts []*TxOutput
}

type Utxo struct {
	TxId []byte
	Index int
	Output *TxOutput
}

//transaction 创建
//1, 创世区块创建时的transaction
//2, 转账时产生的transaction
func NewCoinBaseTransaction(address string) *Transaction {
	txInput := &TxInput{[]byte{}, -1, "Genesis Data"}
	txOutput := &TxOutput{10, address}

	txCoinBase := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}

	//设置hash值
	txCoinBase.SetTxId()
	return txCoinBase
}

//设置交易id， 其实就是hash
func (tx *Transaction) SetTxId() {
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	buffBytes := bytes.Join([][]byte{IntToHex(time.Now().Unix()), buff.Bytes()}, []byte{})
	hash := sha256.Sum256(buffBytes)
	tx.TxId = hash[:]
}

func NewSimpleTransaction(from, to string, amount int64, bc *BlockChain, txs []*Transaction)  {

}

//判断当前交易是否是Coinbase交易
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxId) == 0 && tx.Vins[0].Vout == -1
}