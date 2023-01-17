package receiptClass

import (
	"app/receipt"
	"app/util"
	"fmt"
	"sort"
	"strings"
)

type receiptClass struct {
	receipts []receipt.Receipt
	dollar   int
	partial  map[int]int
	target   int
	result   map[string][]receipt.Receipt
}

type ReceiptClass interface {
	Sort()
	GetResult() map[string][]receipt.Receipt
	Category()
}

func NewReceiptClass(receipts []receipt.Receipt, dollar, partial int) ReceiptClass {
	m := make(map[int]int)
	for i := 1; i <= partial; i++ {
		m[i] = 0
	}
	r := make(map[string][]receipt.Receipt)
	return &receiptClass{
		receipts: receipts,
		dollar:   dollar,
		partial:  m,
		target:   1,
		result:   r,
	}
}

func (rc *receiptClass) Sort() {
	sort.Slice(rc.receipts, func(i, j int) bool {
		return rc.receipts[i].GetSpend() > rc.receipts[j].GetSpend()
	})
}

func (rc *receiptClass) String() string {
	var sb strings.Builder
	for key, value := range rc.result {
		sb.WriteString(fmt.Sprintf("分類: %v\n", key))
		total := 0
		for _, r := range value {
			total += r.GetSpend()
		}
		sb.WriteString(fmt.Sprintf("總價: %v\n", total))
		sb.WriteString("收據：\n")
		sb.WriteString(fmt.Sprintf("%v", value))
		sb.WriteString("\n-----------------------------------------------\n")
	}
	return sb.String()
}

func (rc *receiptClass) GetResult() map[string][]receipt.Receipt {
	return rc.result
}

func (rc *receiptClass) Category() {
	rc.Sort()
	cr := util.SliceCopy(rc.receipts)
	index := 0
	for _, r := range cr {
		if r.GetSpend() > rc.dollar {
			key := rc.generateResultKey()
			rc.result[key] = []receipt.Receipt{r}
			index += 1
		} else {
			break
		}
	}
	cr = cr[index:]
	for len(cr) > 0 {
		var remove []int
		var rs []receipt.Receipt
		total := 0
		for i, r := range cr {
			if r.GetSpend() < rc.dollar-total {
				remove = append(remove, i)
				rs = append(rs, r)
				total += r.GetSpend()
			}
		}
		cr = util.SliceRemove(cr, remove)
		key := rc.generateResultKey()
		rc.result[key] = rs
	}
}

func (rc *receiptClass) generateResultKey() string {
	result := fmt.Sprintf("第%v部位--%v號單", rc.target, rc.partial[rc.target]+1)
	rc.partial[rc.target]++
	rc.target++
	if rc.target > len(rc.partial) {
		rc.target -= len(rc.partial)
	}
	return result
}
