package blc02

import (
    "bytes"
    "math/big"
)

// 0000 0000 0000 0000 1001 0001 0000  .... 0001
//256位Hash里面前面至少有16个零
const TargetBit = 16

//step2：创建pow结构体
type ProofOfWork struct {
    //需要验证的区块
    Block *Block

    //大整数存储 目标哈希
    Target *big.Int
}

//step3: 创建新的工作量证明对象
func NewProofOfwork(block *Block) *ProofOfWork {
    //1.创建一个big对象 0000000.....00001
    /*
       0000 0001
       0010 0000
    */
    target := big.NewInt(1)

    //2.左移256-bits位
    target = target.Lsh(target, 256 - TargetBit)

    return &ProofOfWork{block, target}
}

//step4：返回有效的哈希和nonce值
func (pow *ProofOfWork) Run() ([]byte, int64) {
    //1.将Block的属性拼接成字节数组
    //2.生成Hash
    //3.循环判断Hash的有效性，满足条件，跳出循环结束验证
    nonce := 0

    //var hashInt big.Int //用于存储新生成的hash
    hashInt := new(big.Int)
    var hash [32]byte
    for {
        //获取字节数组
        dataBytes := 
    }
}

//step5：根据block生成一个byte数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
    data := bytes.Join(
        [][]byte{
            pow.Block.PrevBlockHash,
            pow.Block.Data,
            IntToHex(pow.Block.TimeStamp),
            IntToHex(int64(TargetBit)),
            IntToHex(int64(nonce)),
        },
        []byte{},
        )
    return data
}