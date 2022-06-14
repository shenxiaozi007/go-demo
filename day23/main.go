package main

//https://github.com/blockchainGuide/secguide/blob/main/Go%E5%AE%89%E5%85%A8%E6%8C%87%E5%8D%97.md
import (
    "fmt"
    "io"
)

func main1() {
    foo()
}

func foo() {
    var slice = []int{0, 1, 2, 3, 4, 5}
    //在对slice进行操作时，必须判断长度是否合法，防止程序panic
    if len(slice) == 6 {
        fmt.Println(slice[:6])
    }

}

//nil指针判断
//进行指针操作时，必须判断该指针是否为nil，防止程序panic，尤其在进行结构体Unmarshal时

type Packet struct {
    PacketType    uint8
    PacketVersion uint8
    Data          *Data
}

type Data struct {
    Stat uint8
    Len  uint8
    Buf  [8]byte
}

func (p *Packet) UnmarshalBinary(b []byte) error {
    if len(b) < 2 {
        return io.EOF
    }

    p.PacketType = b[0]
    p.PacketVersion = b[1]

    // 若长度等于2，那么不会new Data
    if len(b) > 2 {
        p.Data = new(Data)
    }
    return nil
}

func main2()  {
    packet := new(Packet)
    data := make([]byte, 2)
    if err := packet.UnmarshalBinary(data); err != nil {
        fmt.Println("Failed to unmarshal packet")
        return
    }
    //有可能是空指针
    if packet.Data == nil {
        return
    }

    fmt.Printf("Stat, %v \n", packet.Data.Stat)
}

//必须】整数安全
//
//在进行数字运算操作时，需要做好长度限制，防止外部输入运算导致异常：
//确保无符号整数运算时不会反转
//确保有符号整数运算时不会出现溢出
//确保整型转换时不会出现截断错误
//确保整型转换时不会出现符号错误
//
//以下场景必须严格进行长度限制：
//作为数组索引
//作为对象的长度或者大小
//作为数组的边界（如作为循环计数器）

func overflow(numUser int32)  {
    var numInt int32 = 0
    numInt = numUser + 1

    //可能溢出
    if numInt < 0 {
        fmt.Println("overflow")
        return
    }
    fmt.Printf("%d\n", numInt)
}

func main3()  {
    overflow(2147483647)
}

