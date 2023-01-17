package main

import (
	"app/receipt"
	"app/receiptClass"
	"app/util"
	"encoding/csv"
	"fmt"
	"os"
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
		fmt.Print(doc)
	}
	fmt.Print("最大金額: $ ")
	var dollar int
	fmt.Scanln(&dollar)
	fmt.Print("幾個部位: ")
	var partial int
	fmt.Scanln(&partial)
	c := receiptClass.NewReceiptClass(receipts, dollar, partial)
	NoneDuplicateMethod(c)
}

func NoneDuplicateMethod(receiptClass receiptClass.ReceiptClass) {
	receiptClass.Category()
	fmt.Println(receiptClass)
}
