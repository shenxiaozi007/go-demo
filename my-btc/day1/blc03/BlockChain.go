package blc03

import (
    "fmt"
    "github.com/boltdb/bolt"
    "log"
    "math/big"
    "os"
    "time"
)

//step5:创建区块链
type BlockChain struct {
    //Blocks []*Block //存储有序的区块
    Tip []byte //最后区块的hash值
    DB *bolt.DB //数据库对象
}

//step6：创建区块链，带有创世区块
func CreateBlockChainWithGenesisBlock(data string) *BlockChain {
    //1.先判断数据库是否存在，如果有，从数据库读取
    if dbExists() {
        fmt.Println("数据库已经存在。。。")
        // A： 打开数据库
        db, err := bolt.Open(DBNAME, 0600, nil)
        if err != nil {
            log.Fatal(err)
        }
        //defer db.close
        var blockchain *BlockChain
        //B: 读取数据库
        err = db.View(func(tx *bolt.Tx) error {
            //C: 打开表
            b := tx.Bucket([]byte(BLOCKTABLENAME))
            if b != nil {
                //D: 读取最后一个hash
                hash := b.Get([]byte("1"))
                //E: 创建blockchain
                blockchain = &BlockChain{hash, db}
            }
            return nil
        })
        if err != nil {
            log.Fatal(err)
        }
        return blockchain

    }
    //2.数据库不存在，说明第一次创建，然后存入到数据库中
    fmt.Println("数据库不存在。。")
    //A: 创建创世区块
    genesisBlock := CreateGenesisBlock(data)
    //B: 打开数据库
    db, err := bolt.Open(DBNAME, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    //c: 存入数据库
    err = db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucket([]byte(BLOCKTABLENAME))
        if err != nil {
            log.Panic(err)
        }
        if b != nil {
            err = b.Put(genesisBlock.Hash, genesisBlock.Serilalize())
            if err != nil {
                log.Panic("创世区块存储有误。。。")
            }
            //存储最新区块的hash
            b.Put([]byte("1"), genesisBlock.Hash)
        }
        return nil
    })
    if err != nil {
        log.Panic(err)
    }
    //返回区块对象
    return &BlockChain{genesisBlock.Hash, db}
}

//step7：添加一个新的区块，到区块链中
func (bc *BlockChain) AddBlockToBlockChain(data string/*, height int64, prevHash []byte*/) {
    //创建新区块
    //newBlock := NewBlock(data, prevHash, height)

    //添加到切片中
    //bc.Blocks = append(bc.Blocks, newBlock)
    // 1 更新数据库
    err := bc.DB.Update(func(tx *bolt.Tx) error {
        //2 打开表
        b := tx.Bucket([]byte(BLOCKTABLENAME))
        if b != nil {
            //2 根据最新块的hash读取数据， 并反序列化最后一个区块
            blockBytes := b.Get(bc.Tip)
            lastBlock := DeserializeBlock(blockBytes)
            //
            //3 创建新区块
            newBlock := NewBlock(data, lastBlock.Hash, lastBlock.Height + 1)
            //4 将新的区块序列化并存储
            err := b.Put(newBlock.Hash, newBlock.Serilalize())
            if err != nil {
                log.Panic()
            }
            //5 更新最后一个哈希值， 以及blockchain的tip
            b.Put([]byte("1"), newBlock.Hash)
            bc.Tip = newBlock.Hash
        }
        return nil
    })
    if err != nil {
        log.Panic(err)
    }
}

//判断数据库是否存在
func dbExists() bool {
    if _, err := os.Stat(DBNAME); os.IsNotExist(err) {
        return false
    }
    return true
}

//2.获取一个迭代器的方法
func (bc *BlockChain) Iterator() *BlockChainIterator {
    return &BlockChainIterator{bc.Tip, bc.DB}
}

func (bc *BlockChain) PrintChains() {
    //1 获取迭代器对象
    bcIterator := bc.Iterator()

    var count = 0
    //2 循环迭代
    for {
        block := bcIterator.Next()
        count++
        fmt.Printf("弟%d个区块的信息\n", count)
        //获取当前hash对应的数据，并进行反序列化
        fmt.Printf("\t高度：%d\n", block.Height)
        fmt.Printf("\t上一个区块的hash：%x\n", block.PrevBlockHash)
        fmt.Printf("\t当前的hash：%x\n", block.Hash)
        fmt.Printf("\t数据：%s\n", block.Data)
        fmt.Printf("\t时间：%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"))
        fmt.Printf("\t次数：%d\n", block.Nonce)

        //3.直到父hash值为o
        hashInt := new(big.Int)
        hashInt.SetBytes(block.PrevBlockHash)
        if big.NewInt(0).Cmp(hashInt) == 0 {
            break
        }
    }
}