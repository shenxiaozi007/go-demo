package service

import (
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

//实现互通的rpc客户端
func (c *Client) callRpc(rpcName string, fPtr interface{}) {
	//通过反射获取fptr未初始化的原型
	fn := reflect.ValueOf(fPtr).Elem()

	// 需要另一个函数。作用是对第一个函数参数操作
	f := func(args []reflect.Value) []reflect.Value {
		// 处理参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}

		//链接
		cliSession := NewSession(c.conn)
		//编码数据
		reqRpc := RpcData{Name: rpcName, Args: inArgs}

		b, err := encode(reqRpc)
		if err != nil {
			panic(b)
		}

		//写数据
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}

		// 服务端过来返回值 读取解析
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}

		//解码
		respRpc, err := decode(respBytes)

		//处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRpc.Args))

		for i, arg := range respRpc.Args {
			// 进行nil转换
			if arg == nil {
				// reflect.Zero()会返回类型的零值的value
				// .out()会返回函数输出的参数类型
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}

		return outArgs
	}
	// 完成原型到函数调用的内部转换
	// 参数1是reflect.Type
	// 参数2 f是函数类型，是对于参数1 fn函数的操作
	// fn是定义，f是具体操作

	v := reflect.MakeFunc(fn.Type(), f)

	fn.Set(v)
}
