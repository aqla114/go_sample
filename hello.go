package main

import (
	"fmt"
)

// Name ですよ。コメントがついていないかたや関数は、go-lintで怒られるっぽい。
type Name string

// ITask is interface。でも、これ明示的に実装しなくていいらしい。キモい。
type ITask interface {
	Finish()
}

// Task ですよ。
type Task struct {
	id     int
	detail string
	done   bool
}

// NewTask id:int, detail: string。基本的にポインタを返す。
func NewTask(id int, detail string) *Task {
	var result = &Task{
		id:     id,
		detail: detail,
		done:   false,
	}
	return result
}

// Finish Task型に生やすメソッド。structは値型だから、ポインタじゃないとプロパティ更新できない。
func (task *Task) Finish() {
	task.done = true
}

// DoFinish interface を引数の型に指定してメソッドを利用できる。
func DoFinish(task ITask) {
	task.Finish()
}

// ExtendedTask 継承ないのでこんな感じで型を埋め込むらしい。
type ExtendedTask struct {
	*Task
	deadline string
}

// NewExtendedTask id:int, detail: string。基本的にポインタを返す。
func NewExtendedTask(id int, detail string, deadline string) (result *ExtendedTask) {
	result = &ExtendedTask{
		deadline: deadline,
		Task:     NewTask(id, detail),
	}
	return
}

var myarr = [4]int{1, 2, 3, 4}
var myslice = []int{1, 2, 3, 4} // array と slice は明確に違う型として扱われる。

func main() {
	const a = 1
	const b = 1.0
	fmt.Println(a == b)

	var name1 Name = "John"
	// var name2 string = "hallo"
	// name1 = name2 これは怒られる。
	// name2 = name1 これも怒られる。
	fmt.Println(name1)

	fmt.Println(myarr, myarr[0:2])
	fmt.Println(append(myslice, 5), myslice, myslice[0:2])

	fmt.Println(sum(1, 2, 3)) // sum(myslice)はダメっぽい。 myslice... で spread してやる必要がある。

	var task = *NewTask(1, "This is detail.")
	fmt.Println(task)
	task.Finish()
	fmt.Println(task)

	var exTask = *NewExtendedTask(2, "This is extended detail.", "TODAY")
	fmt.Println(exTask, exTask.id, exTask.detail, exTask.done)
	exTask.Finish()
	fmt.Println(exTask, exTask.id, exTask.detail, exTask.done)

	fmt.Println(IsString(name1))
}

func sum(nums ...int) (result int) {
	fmt.Println(nums) // sliceが渡ってくる
	for _, n := range nums {
		result += n
	}
	return
}

// IsString 型チェック
func IsString(val interface{}) (result interface{}) {
	s, ok := val.(string)

	fmt.Println("ok :", ok)

	if ok {
		result = s
	} else {
		result = nil
	}
	return
}
