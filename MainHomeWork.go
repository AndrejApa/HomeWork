package main

import (
	"fmt"
	"time"
)

type User struct {
	Id uint
	Username string
	Email string
	BirthDay time.Time

}
type Inmemory struct {
	Db map[int]User
}

func main () {

	fmt.Println("Hello Mihalych")

}