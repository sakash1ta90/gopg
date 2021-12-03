package handler

import "fmt"

type DefaultGet interface {
	Get(defaultValue interface{})
}

type JsonString string

func (t *JsonString) Get(defaultValue interface{}) JsonString {
	if t == nil {
		return JsonString(fmt.Sprint(defaultValue))
	}
	return *t
}

type JsonInt int

func (t *JsonInt) Get(defaultValue interface{}) JsonInt {
	if t == nil {
		num, ok := defaultValue.(JsonInt)
		if ok {
			return num
		}
		return 0
	}
	return *t
}
