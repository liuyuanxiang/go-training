package main

import (
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

// 模拟sql no rows error
var ErrNoRows = errors.New("no rows")


func main(){
	err := daoHandle()
	fmt.Printf("main:%+v\n",err)

}

// 数据库操作
func daoHandle() error{
	return xerrors.Wrapf(ErrNoRows,"没有找到记录")
}
