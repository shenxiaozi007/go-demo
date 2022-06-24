package blc04

import (
    "github.com/boltdb/bolt"
    "log"
)

type BlockChainIterator struct {
    CurrentHash []byte //当前区块的hash
    Db *bolt.DB
}

//获取区块
func (blockChainIterator *BlockChainIterator) Next() *Block {
    block := new(Block)
    //1.打开数据库并读取
    err := blockChainIterator.Db.View(func(tx *bolt.Tx) error {
        //2, 打开数据库
        b := tx.Bucket([]byte(BLOCKTABLENAME))
        if b != nil {
            //3. 根据当前hash获取数据并反序列化
            blockBytes := b.Get(blockChainIterator.CurrentHash)
            block = DeserializeBlock(blockBytes)
            // 更新当前的hash
            blockChainIterator.CurrentHash = block.PrevBlockHash
        }
        return nil
    })

    if err != nil {
        log.Panic(err)
    }
    return block
}