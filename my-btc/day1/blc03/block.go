package blc03

import (
    "bytes"
    "crypto/sha256"
    "encoding/gob"
    "log"
    "strconv"
    "time"
)

//step1:创建Block结构体
type Block struct {
    //高度Height：其实就是区块的编号，第一个区块叫创世区块，高度为0
    Height int64
    //上一个区块的哈希值ProvHash：
    PrevBlockHash []byte
    //交易数据 date 目前先设计为[]byte,后期是Transaction
    Data []byte
    //时间戳
    TimeStamp int64
    //哈希值hash 32字节 64个16进制数
    Hash []byte

    Nonce int64
}

//step2：创建新的区块

func NewBlock(data string, PrevBlockHash []byte, height int64) *Block {
    //创建区块
    block := &Block{height, PrevBlockHash, []byte(data), time.Now().Unix(), nil, 0}

    //
    //设置hash值
    //block.SetHash()
    //调用工作量证明
    pow := NewProofOfwork(block)
    hash, nonce := pow.Run()
    block.Hash = hash
    block.Nonce = nonce
    return block
}

//step3:设置区块的hash
func (block *Block) SetHash() {
    //1 将高度转为字节数组
    heightBytes := IntToHex(block.Height)
    //2.时间戳转为字节数组
    timeString := strconv.FormatInt(block.TimeStamp, 2)
    timeBytes := []byte(timeString)

    //拼接所有属性
    blockBytes := bytes.Join([][]byte{
        heightBytes,
        block.PrevBlockHash,
        block.Data,
        timeBytes}, []byte{})

    //生成hash值
    hash := sha256.Sum256(blockBytes) //数组长度32位

    block.Hash = hash[:]
}

//step6：创建区块链，带有创世区块
func CreateGenesisBlock(data string) *Block {
    return NewBlock(data, make([]byte, 32, 32), 0)
}

// 区块序列化 得到一个字节数组---区块的行为，设计为方法
func (block *Block) Serilalize() []byte {
    //1 创建一个buffer
    var result bytes.Buffer
    //2 创建一个编码器
    encoder := gob.NewEncoder(&result)
    //3 编码----》打包
    err := encoder.Encode(block)
    if err != nil {
        log.Panic(err)
    }
    return result.Bytes()
}

// 反序列化， 得到一个区块----设计为函数
func DeserializeBlock(blockBytes []byte) *Block {
    var block Block
    var reader = bytes.NewReader(blockBytes)
    //1. 创建一个解码器
    decoder := gob.NewDecoder(reader)
    //2 解包
    err := decoder.Decode(&block)
    if err != nil {
        log.Panic(err)
    }
    return &block
}