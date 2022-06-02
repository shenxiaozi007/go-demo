package blc02

import (
    "bytes"
    "crypto/sha256"
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
}

//step2：创建新的区块

func NewBlock(data string, PrevBlockHash []byte, height int64) *Block {
    //创建区块
    block := &Block{height, PrevBlockHash, []byte(data), time.Now().Unix(), nil}

    //设置hash值
    block.SetHash()
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
