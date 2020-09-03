package gfUtils

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gdb"
)

type Err struct {
	errs []error
}

func NewErr() *Err {
	return &Err{make([]error, 0)}
}

func (e *Err) Append(err ...error) {
	for _, v := range err {
		if v != nil {
			e.errs = append(e.errs, v)
		}
	}
}

func (e *Err) IsEmpty() bool {
	if len(e.errs) == 0 {
		return true
	}

	return false
}

func (e *Err) HandleEmptyRecord(rec gdb.Record, tag string) {
	if rec.IsEmpty() {
		e.Append(errors.New(fmt.Sprintf("数据不存在：[%s]", tag)))
	}
}

func (e *Err) Errs() []error {
	return e.errs
}

func (e *Err) String() string {
	s := ""
	for _, v := range e.errs {
		s += fmt.Sprintf("%s\n", v.Error())
	}
	return s
}
