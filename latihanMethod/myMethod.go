package latihanmethod

import "fmt"

type MyTodo struct {
	ID     int
	Title  string
	IsDone bool
}

func (t MyTodo) StatusText() string {
	if t.IsDone {
		return "Sudah Makan"
	} else {
		return "Belum Makan"
	}
}

func MyMethod() {

	todoKu := MyTodo{ID: 1, Title: "makan", IsDone: false}

	fmt.Println(todoKu)
	fmt.Println(todoKu.StatusText())

	apakahSudahMakan := "belum makan"

	if todoKu.IsDone {
		apakahSudahMakan = "sudah makan"
	}

	fmt.Println(apakahSudahMakan)

}
