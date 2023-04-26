package main

import (
	"fmt"
)

type Item struct {
	title string
	body  string
}

var database []Item

func GetByTitle(title string) Item {
	var getItem Item
	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}
	return getItem
}

func AddItem(title string, body string) Item {
	newItem := Item{title, body}
	database = append(database, newItem)
	return newItem
}

func EditItem(title string, body string) Item {
	var changed Item
	for idx, val := range database {
		if val.title == title {
			database[idx].body = body
			changed = database[idx]
		}
	}
	return changed
}

func DeleteItem(item Item) Item {
	var del Item
	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
		}
	}
	return del
}

func main() {
	fmt.Println("Initial database: ", database)
	a := Item{"first", "a test item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	AddItem(a.title, a.body)
	AddItem(b.title, b.body)
	AddItem(c.title, c.body)
	fmt.Println("Second database: ", database)

	DeleteItem(b)
	fmt.Println("Third database: ", database)

	EditItem(c.title, "a different third item")
	fmt.Println("Fourth database: ", database)

	fmt.Println("GetByTitle: ", GetByTitle("first"))
	fmt.Println("GetByTitle: ", GetByTitle("third"))

}
