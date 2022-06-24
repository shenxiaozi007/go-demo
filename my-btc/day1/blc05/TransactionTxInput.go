package blc05

type TxInput struct {
	//1.交易id
	TxId []byte
	//2.存储txoutput的vout里面的索引
	Vout int
	//3.用户名
	ScriptSiq string
}

func (TxInput *TxInput) UnLockWithAddress(address string) bool {
	return TxInput.ScriptSiq == address
}