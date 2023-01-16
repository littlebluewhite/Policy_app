package main

import (
	"app/receipt"
	"app/receiptClass"
	"app/util"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("保險.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)

	receipts := make([]receipt.Receipt, 0, 100)
	for _, item := range records[1:] {
		money, _ := strconv.Atoi(item[1])
		doc := receipt.NewReceipt(util.StrToTime(item[0]), money, item[2])
		receipts = append(receipts, doc)
		fmt.Println(doc)
	}
	fmt.Print("最大金額: $ ")
	var dollar int
	fmt.Scanln(&dollar)
	fmt.Print("幾個部位: ")
	var partial int
	fmt.Scanln(&partial)
	fmt.Print("是否重複: ")
	var isDuplicate bool
	fmt.Scanln(&isDuplicate)
	var duplicate int
	if isDuplicate {
		fmt.Print("重複次數: ")
		fmt.Scanln(&duplicate)

	} else {
		c := receiptClass.NewReceiptClass(receipts, dollar, partial, duplicate)
		NoneDuplicateMethod(c)
	}
}

func duplicateMethod() {

}

func NoneDuplicateMethod(receiptclass receiptClass.ReceiptClass) {
	sort.Slice(receipts, func(i, j int) bool {
		return receipts[i].GetSpend() > receipts[j].GetSpend()
	})
	fmt.Println(receipts)
}
