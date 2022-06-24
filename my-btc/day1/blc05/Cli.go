package blc05

import (
    "flag"
    "fmt"
    "log"
    "os"
)

type CLi struct {

}

func isValidArgs()  {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }
}

func printUsage()  {
    fmt.Println("Usage:")
    fmt.Println("\t createblockchain -data DATA -- 创建创世区块")
    fmt.Println("\t addblock -data Data -- 交易数据")
    fmt.Println("\t printchain -- 输出信息")
}

func (cli *CLi) addBlock(data string) {
    bc := GetBlockchainObject()
    if bc == nil {
        fmt.Println("没有创世区块，无法添加。。")
        os.Exit(1)
    }
    defer bc.DB.Close()
    bc.AddBlockToBlockChain(data)
}

func (cli *CLi) printChains() {
    bc := GetBlockchainObject()
    if bc == nil {
        fmt.Println("没有区块可以打印")
        os.Exit(1)
    }
    defer bc.DB.Close()
    bc.PrintChains()
}

func (cli *CLi) createGenesisBlockchain(data string) {
    //fmt.println
    CreateBlockChainWithGenesisBlock(data)
}

func (cli *CLi) Run() {
    //判断命令行参数的长度
    isValidArgs()

    //1. 创建flagest标签对象
    addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)

    printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

    createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
    fmt.Println(os.Args[2:])
    //2. 设置标签后的参数
    flagAddBlockData := addBlockCmd.String("data","helloworld..","交易数据")

    flagCreateBlockChainData := createBlockChainCmd.String("data","Genesis block data..","创世区块交易数据")
    
    //3. 解析
    switch os.Args[1] {
    case "addblock":
        err := addBlockCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "printchain":
        err := printChainCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "createblockchain":
        err := createBlockChainCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    default:
        printUsage()
        os.Exit(1)
    }

    if addBlockCmd.Parsed() {
        if *flagAddBlockData == "" {
            printUsage()
            os.Exit(1)
        }
        cli.addBlock(*flagAddBlockData)
    }

    if printChainCmd.Parsed() {
        cli.printChains()
    }

    if createBlockChainCmd.Parsed() {
        if *flagCreateBlockChainData == "" {
            printUsage()
            os.Exit(1)
        }
        cli.createGenesisBlockchain(*flagCreateBlockChainData)
    }
}