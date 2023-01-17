package receipt

import (
	"fmt"
	"time"
)

type receipt struct {
	date  time.Time
	spend int
	_type string
	count int
}

type Receipt interface {
	GetTimestamp() int
	GetSpend() int
	GetCount() int
}

func NewReceipt(date time.Time, spend int, _type string) Receipt {
	return &receipt{
		date:  date,
		spend: spend,
		_type: _type,
	}
}

func (r *receipt) String() string {
	return fmt.Sprintf("日期: %v, 花費: %v, 類型: %v\n", r.date.Format("2006/01/02"), r.spend, r._type)
}

func (r *receipt) GetTimestamp() int {
	return int(r.date.Unix())
}

func (r *receipt) GetSpend() int {
	return r.spend
}

func (r *receipt) GetCount() int {
	return r.count
}
