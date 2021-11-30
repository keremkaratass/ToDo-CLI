package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

//Shows ToDo's rules.
const rules string = "|todo -v(#version)|,|todo -l(#list all item)|,|todo -c(#list completed items)|,|todo -a Buy Milk(#add new item)|,|todo -m TODO-ID(#mark as complete)|,|todo -m TODO-ID(#delete item)|"

//Item struct represents a ToDO item
type item struct {
	IdNumber         int
	OrderName        string
	Done             bool
	OrderDescription time.Time
}

//List represents a list of ToDo items
type list []item

//idNumber will start at 0. This is a variable for idNumber.
var number int

//This function adds items
func (l *list) addItem(itemName string) {
	insert := item{
		IdNumber:         number,
		OrderName:        itemName,
		Done:             false,
		OrderDescription: time.Now(),
	}
	number++ // After the first item has been added, this will be 1.   +1 inrease in each item.
	*l = append(*l, insert)
	l.saveToFile()

}

//A function written to keep in a file when a item is added
func (l *list) saveToFile() error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("liste", js, 0666)
}

//decodes
//the JSON data and parses it into a List
func (l *list) getFromFile() error {
	file, err := ioutil.ReadFile("liste")
	if err != nil {
		return err
	}
	if len(file) == 0 {
		return nil
	}
	err = json.Unmarshal(file, l)
	if err != nil {
		return err
	}
	number = len(*l)
	return nil
}

//Delete method to remove an item from list
func (l *list) deleteItem(itemNumber int) error {
	ls := *l
	if itemNumber < 0 || itemNumber > len(ls) {
		return fmt.Errorf("Item %d does not exist", itemNumber)
	}
	*l = append(ls[:itemNumber-1], ls[itemNumber:]...)
	return nil
}

//markAdCompleted method marks a ToDo item as completed by
//setting done=true
func (l *list) markAsCompleted(itemNumber int) error {
	ls := *l
	if itemNumber < 0 || itemNumber > len(ls) {
		return fmt.Errorf("İtem %d does not exist", itemNumber)

	}
	ls[itemNumber-1].Done = true
	return nil
}

//Shows a list of completed items.
func (l *list) listCompletedItems() {
	for _, item := range *l {
		if item.Done {
			fmt.Println("|idNumber|:", item.IdNumber, "|orderName|:", item.OrderName, "|done|:", item.Done, "|orderDescription|:", item.OrderDescription)
		}
	}
}

//Shows the entire item list.
func (l *list) listAllItems() {
	for _, item := range *l {
		fmt.Println("|idNumber|:", item.IdNumber, "|orderName|:", item.OrderName, "|done|:", item.Done, "|orderDescription|:", item.OrderDescription)
	}
}

//Show the version
func printVersion(pointerToVersion *string) {
	fmt.Println("Version", *pointerToVersion)
}

//Show the help info
func showHelps() {
	fmt.Println(rules)

}
