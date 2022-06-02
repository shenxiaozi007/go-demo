package main

import (
    "fmt"
    "github.com/huangxinchun/go-demo/my-btc/day1/blc01"
)

func main()  {
    //1 测试
    block := blc01.NewBlock("test", make([]byte, 32, 32), 1)
    fmt.Println(block)

    //2 测试创世区块
    genesisBlock := blc01.CreateGenesisBlock("Genesis Block..")
    fmt.Println(genesisBlock)

    //3 测试区块链
    genesisBlockChain := blc01.CreateBlockChainWithGenesisBlock("Genesis Block..")
    //4 添加新区块
    genesisBlockChain.AddBlockToBlockChain("fuck you", genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Height + 1, genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Hash)
    genesisBlockChain.AddBlockToBlockChain("fuck you c", genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Height + 1, genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Hash)
    genesisBlockChain.AddBlockToBlockChain("fuck you d", genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Height + 1, genesisBlockChain.Blocks[len(genesisBlockChain.Blocks) - 1].Hash)

    for _, v := range genesisBlockChain.Blocks {
        fmt.Println(string(v.Data))
    }

}