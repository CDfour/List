package main

import (
	"fmt"

	"list/pkg/list"
)

// Проверка каждой функции
func main() {
	myList := list.NewList(1, 2, 3, 4, 5)
	print(myList)

	slice := []interface{}{"one", "two", "three"}
	myList.Assign(slice)
	print(myList)

	i, _ := myList.Back()
	fmt.Println(i)

	i, _ = myList.Front()
	fmt.Println(i)

	myList.Clear()
	print(myList)

	fmt.Println(myList.Empty())

	secondList := list.NewList(1, 2, 3, 4, 1, 2, 3, 4, 0)
	myList.Splice(1, secondList)
	print(myList)
	print(secondList)

	myList.Erase(5)
	print(myList)

	p, _ := myList.Begin()
	list.Advance(&p, 5)
	myList.Erase(p)
	print(myList)

	myList.Insert(8, "one")
	print(myList)

	p, _ = myList.Begin()
	list.Advance(&p, 7)
	myList.Insert(p, "two")
	print(myList)

	p, _ = myList.NewIterator(7)
	fmt.Println(p.Data)

	myList.Pop_back()
	print(myList)

	myList.Pop_front()
	print(myList)

	myList.Push_back(99)
	print(myList)

	myList.Push_front("start")
	print(myList)

	myList.Remove(2)
	print(myList)

	myList.Resize(5)
	print(myList)

	myList.Resize(9)
	print(myList)

	myList.Reverse()
	print(myList)

	fmt.Println(myList.Size())

	myList.Unique()
	print(myList)

	secondList = list.NewList(1, 2, 3)
	myList.Swap(secondList)
	print(myList)
	print(secondList)

	val, _ := myList.GetData(2)
	fmt.Println(val)

	p, _ = myList.End()
	fmt.Println(p.Data)
}

// Выводит данные каждого элемента списка
func print(l *list.List) {
	ptr, err := l.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for ptr != nil {
		fmt.Print(ptr.Data, " ")
		list.Advance(&ptr, 1)
	}

	fmt.Print("\n")
}
