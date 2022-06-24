package blc05

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

//step5:创建区块链
type BlockChain struct {
	//Blocks []*Block //存储有序的区块
	Tip []byte   //最后区块的hash值
	DB  *bolt.DB //数据库对象
}

//step6：创建区块链，带有创世区块
func CreateBlockChainWithGenesisBlock(data string) /* *BlockChain */ {
	//1.先判断数据库是否存在，如果有，从数据库读取
	if dbExists() {
		fmt.Println("数据库已经存在。。。")
		return
		/* // A： 打开数据库
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
		   //return blockchain*/

	}
	fmt.Println("创建创世区块：")
	//2.数据库不存在，说明第一次创建，然后存入到数据库中
	fmt.Println("数据库不存在。。")
	//
	txCoinBase := NewCoinBaseTransaction(data)
	//A: 创建创世区块
	genesisBlock := CreateGenesisBlock([]*Transaction{txCoinBase})
	//A: 创建创世区块
	//genesisBlock := CreateGenesisBlock(data)
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
	//return &BlockChain{genesisBlock.Hash, db}
}

//step7：添加一个新的区块，到区块链中
func (bc *BlockChain) AddBlockToBlockChain(txs []*Transaction /*data string, height int64, prevHash []byte*/) {
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
			//newBlock := NewBlock(data, lastBlock.Hash, lastBlock.Height + 1)
			newBlock := NewBlock(txs, lastBlock.Hash, lastBlock.Height+1)
			//4 将新的区块序列化并存储
			err := b.Put(newBlock.Hash, newBlock.Serilalize())
			if err != nil {
				log.Panic(err)
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
		//fmt.Printf("\t数据：%s\n", block.Txs)
		fmt.Println("\t交易：")
		for _, tx := range block.Txs {
			fmt.Printf("\t\t交易ID: %x\n", tx.TxId)
			fmt.Println("\t\tVins:")
			for _, in := range tx.Vins {
				fmt.Printf("\t\t\tTxId:%x\n", in.TxId)
				fmt.Printf("\t\t\tVout:%d\n", in.Vout)
				fmt.Printf("\t\t\tScriptSiq:%s\n", in.ScriptSiq)
			}
			fmt.Println("\t\tVouts:")
			for _, out := range tx.Vouts {
				fmt.Printf("\t\t\t value:%d\n", out.Value)
				fmt.Printf("\t\t\tScripuPubkey:%d\n", out.ScriptPubkey)
			}
		}

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

//新增方法， 用于获取区块链
func GetBlockchainObject() *BlockChain {
	//1.如果数据库不存在。直接返回nil
	//2.读取数据库
	if !dbExists() {
		fmt.Println("数据库不存在，无法获取区块链")
		return nil
	}

	db, err := bolt.Open(DBNAME, 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close
	var blockchain *BlockChain

	//B 读取数据库
	err = db.View(func(tx *bolt.Tx) error {
		//C 打开表
		b := tx.Bucket([]byte(BLOCKTABLENAME))
		if b != nil {
			//D 读取最后一个hash
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

func (bc *BlockChain) MineNewBlock(from, to, amount []string) {
	//1. 新建交易
	//2. 新建区块
	//3. 将区块存入到数据库
	var txs []*Transaction
	for i := 0; i < len(from); i++ {
		amountInt, _ := strconv.ParseInt(from[i], to[i], amount)
	}
}

func (bc *BlockChain) FindSpendTableUtxOs(from string, amount int64, txs []*Transaction) (int64, map[string][]int) {
	/**
	  1.获取所有的UTXO
	  2.遍历UTXO
	  返回值 map[hash]{index}
	*/
	var balance int64
	utxos := bc.
}

//找到所有未花费的交易输出
func (bc *BlockChain) UnUtxOs(address string, txs []*Transaction) []*Utxo {
	/**
	  1. 先遍历打包的交易（参数txs）找出未消费的output
	  2. 遍历数据库， 获取每个块的transaction， 找出未花费的output
	*/
	var unUtxOs []*Utxo                      //未花费
	spentTxOutputs := make(map[string][]int) //存储已经花费

	//1 添加先从txs遍历 查找未花费
	for i := len(txs) - 1; i >= 0; i-- {
		unUtxOs = caculate(txs[i], address, spentTxOutputs, unUtxOs)
	}

	bcIterator := bc.Iterator()

	for {
		block := bcIterator.Next()
		//统计未花费
		//2 获取block中的每个transaction
		for i := len(block.Txs) - 1; i >= 0; i-- {
			unUtxOs = caculate(block.Txs[i], address, spentTxOutputs, unUtxOs)
		}
		//结束迭代
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(hashInt) == 0 {
            break
		}
	}
    return unUtxOs
}

func caculate(tx *Transaction, address string, spentTxOutputs map[string][]int, unUtxOs []*Utxo) []*Utxo {
    //2. 先遍历txInputs 表示花费
    if !tx.IsCoinbaseTransaction() {
        for _, in := range tx.Vins {
            //如果解锁
            if in.UnLockWithAddress(address) {
                key
            }
        }
    }
}