package blc05

import (
    "bytes"
    "encoding/binary"
    "encoding/json"
    "log"
)

/**
将一个int64的整数：转为二进制后，每8bit一个byte。转为[]byte
 */
func IntToHex(num int64) []byte {
    buff := new(bytes.Buffer)

    //将二进制数据写入w
    //func Write(w io.Writer, order ByteOrder, data interface{})
    err := binary.Write(buff, binary.BigEndian, num)

    if err != nil {
        log.Panic(err)
    }
    //转为[]byte并返回
    return buff.Bytes()
}

//json字符串转为[] string数组
func JsonToArray(jsonString string) []string {
    var sArr []string
    if err := json.Unmarshal([]byte(jsonString), &sArr); err != nil {
        log.Panic(err)
    }

    return sArr
}