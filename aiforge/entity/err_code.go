package entity

import "fmt"

type ErrCode struct {
	CodeVal    string
	CodeMsg    string
	CodeTrCode string
}

func (e *ErrCode) IsMatch(code interface{}) bool {
	return fmt.Sprint(code) == e.CodeVal
}
