package main

import (
	"flag"
	"time"
)

func main() {
	var help string
	flag.StringVar(&help, "h", "", "#help")
	var version string
	flag.StringVar(&version, "v", time.Now().String(), "#version")
	var itemName string
	flag.StringVar(&itemName, "a", " ", "#add new item")
	var listAllItems string
	flag.StringVar(&listAllItems, "l", "", "#list all items")
	var listCompletedItems string
	flag.StringVar(&listCompletedItems, "c", "", "#list completed items")
	var itemNumber int
	flag.IntVar(&itemNumber, "d", 0, "#delete item")
	var itemCompletedNumber int
	flag.IntVar(&itemCompletedNumber, "m", 0, "#mark as complete")
	flag.Parse()

	//showHelps()
	//printVersion(&version)

	product := &list{}
	product.getFromFile()
	product.addItem(itemName)
	product.listAllItems()
	//product.listCompletedItems()
	//product.markAsCompleted(itemCompletedNumber)
	//sproduct.deleteItem(itemNumber)

}
