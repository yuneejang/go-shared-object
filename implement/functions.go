package implement

import (
	"fmt"
	"net/rpc"
)

type Calc int // RPC 서버에 등록하기 위해 임의의 타입으로 정의

type Args struct { // 매개변수
	A, B int
}

type Reply struct { // 리턴값
	C int
}

func Add(a, b int) int {
	//참고(서버사이드) : http://pyrasis.com/book/GoForTheReallyImpatient/Unit57
	//    (클라이언트) : http://pyrasis.com/book/GoForTheReallyImpatient/Unit57/02
	client, err := rpc.Dial("tcp", "127.0.0.1:7500") // RPC 서버에 연결
	fmt.Println("rpc.Dial(tcp, 127.0.0.1:7500)")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer client.Close() // main 함수가 끝나기 직전에 RPC 연결을 닫음

	//1. 동기 호출
	args := &Args{a, b}
	reply := new(Reply)
	err = client.Call("Calc.Sum", args, reply) // Calc.Sum 함수 호출
	if err != nil {
		fmt.Println(err)
		//	return err
		return -1

	}
	fmt.Println(reply.C)
	/////////////////////////////////////////////////////////////////////////

	// //2. 비동기호출
	// args.A = 4
	// args.B = 9
	// sumCall := client.Go("Calc.Sum", args, reply, nil) // Calc.Sum 함수를 고루틴으로 호출
	// <-sumCall.Done                                     // 함수가 끝날 때까지 대기
	// fmt.Println(reply.C)
	/////////////////////////////////////////////////////////////////////////
	return reply.C
}

//Sub is substract
func Sub(a, b int) int {
	return a - b
}

//Mul is multiply
func Mul(a, b int) int {
	return a * b
}
