package main

import (
	"fmt"
)
type customErr struct{
	info string
}
func (c customErr) Error() string{
	return fmt.Sprintf("error has occured %v",c.info)
}

func foo(e error){
	//here we are asserting that the incoming error is from customErr
	fmt.Println("err is from inside foo:",e,e.(customErr).info)
}

func main(){

	c:=customErr{
		info:"file open error",

	}
	//fmt.Println("from inside main",c)
	foo(c)


}