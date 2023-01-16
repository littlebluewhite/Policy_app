package receiptClass

import (
	"app/receipt"
	"fmt"
	"sort"
)

type receiptClass struct {
	receipts  []receipt.Receipt
	dollar    int
	partial   int
	duplicate int
	result    map[string][]receipt.Receipt
}

type ReceiptClass interface {
	Sort()
	GetResult() map[string][]receipt.Receipt
}

func NewReceiptClass(receipts []receipt.Receipt, dollar, partial, duplicate int) ReceiptClass {
	return &receiptClass{
		receipts:  receipts,
		dollar:    dollar,
		partial:   partial,
		duplicate: duplicate,
	}
}

func (rc *receiptClass) Sort() {
	sort.Slice(rc.receipts, func(i, j int) bool {
		return rc.receipts[i].GetSpend() > rc.receipts[j].GetSpend()
	})
}

func (rc *receiptClass) String() string {
	return fmt.Sprintln(rc.result)
}

func (rc *receiptClass) GetResult() map[string][]receipt.Receipt {
	return rc.result
}
