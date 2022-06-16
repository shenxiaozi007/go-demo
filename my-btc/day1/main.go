package main

import (
    "crypto/sha256"
    "fmt"
    "github.com/huangxinchun/go-demo/my-btc/day1/blc02"
    "math/big"
)

func main()  {
  /*  //1 测试
    block := blc02.NewBlock("test", make([]byte, 32, 32), 1)
    fmt.Println(block)

    //2 测试创世区块
    genesisBlock := blc02.CreateGenesisBlock("Genesis Block..")
    fmt.Println(genesisBlock)

    //3 测试区块链
    genesisBlockChain := blc02.CreateBlockChainWithGenesisBlock("Genesis Block..")
    //4 添加新区块
    genesisBlockChain.AddBlockToBlockChain("fuck you", genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Height + 1, genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Hash)
    genesisBlockChain.AddBlockToBlockChain("fuck you c", genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Height + 1, genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Hash)
    genesisBlockChain.AddBlockToBlockChain("fuck you d", genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Height + 1, genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Hash)

    for _, v := range genesisBlockChain.Blocks {
        fmt.Println(string(v.Data))
    }*/

    // 检测pow
    //1.创建一个big对象 0000000.....00001
    target := big.NewInt(1)

    fmt.Printf("0x%x\n", target)

    //2.左移256-bits位
    target = target.Lsh(target, 256 - blc02.TargetBit)
    fmt.Printf("0x%x\n",target) //61

    s1 := "helloWorld"
    hash := sha256.Sum256([]byte(s1))
    fmt.Printf("0x%x\n",hash)

    // 测试添加新区块
    blockChain := blc02.CreateBlockChainWithGenesisBlock("Genesis Block..")
    blockChain.AddBlockToBlockChain("Send 100RMB To Wangergou",blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
    blockChain.AddBlockToBlockChain("Send 300RMB To lixiaohua",blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
    blockChain.AddBlockToBlockChain("Send 500RMB To rose",blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
    fmt.Println(blockChain)
    pow := blc02.NewProofOfwork(blockChain.Blocks[0])
    fmt.Printf("%v \n", pow.IsValid())
}