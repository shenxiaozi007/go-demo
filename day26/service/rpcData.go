package service

import (
	"bytes"
	"encoding/gob"
)

type RpcData struct {
	//访问的函数
	Name string
	//访问时的参数
	Args []interface{}
}

//编码
func encode(data RpcData) ([]byte, error) {
	//得到字节数组的编码器
	var buf bytes.Buffer

	bufEnc := gob.NewEncoder(&buf)

	//编码器对数据编码
	if err := bufEnc.Encode(data); err != nil {
		return nil, err
	}
	 return buf.Bytes(), nil
}

//解码
func decode(b []byte) (RpcData, error) {
	buf := bytes.NewBuffer(b)

	// 得到字节数组解码器
	bufDec := gob.NewDecoder(buf)

	// v解码器对数据解码
	var data RpcData
	if err := bufDec.Decode(&data); err != nil  {
		return data, err
	}
	return data, nil
}