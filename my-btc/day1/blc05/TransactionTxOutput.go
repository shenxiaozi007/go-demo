package blc05

type TxOutput struct {
	Value int64
	//一个锁定脚本(ScriptPubKey)，要花这笔钱，必须要解锁该脚本
	ScriptPubkey string //公钥：先理解为，用户名

}

//判断当前txOutput消费，和指定的address是否一致
func (txOutPut *TxOutput) UnLockWithAddress(address string) bool {
	return txOutPut.ScriptPubkey == address
}