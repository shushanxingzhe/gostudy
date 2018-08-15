package main

import (
	"time"
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}


func somedoing(s string,message chan  Books) {
	fmt.Println("enter: " + s)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("finish: " + s)
	var book =  Books{title: s, author: "www.runoob.com", subject: "Go 语言教程", book_id: 1}
	message <- book
}

func main() {
	message1 := make(chan Books)
	message2 := make(chan Books)

	fmt.Println(time.Now())
	go somedoing("world",message1)
	go somedoing("hello",message2)
	fmt.Println(time.Now())

	fmt.Println(<-message1);
	fmt.Println(<-message2);
}
