package main

import (
	"encoding/json"
	"fmt"
	"go-playground-rs/06-strucs-and-methods/foo"
	"time"
)

type MinhaString string

// `json:"name"` => struct tag
type User struct {
	Name          string    `json:"name"`
	ID            uint64    `json:"id"`
	UpdateAt      time.Time `json:"update_at"`
	foo.Profissao           // this is called embedding
}

func (u *User) GetGreetingMessage() {
	fmt.Printf("Welcome, %s!\n", u.Name)
	u.update()
}

func (u *User) UpdateName(name string) {
	u.Name = name
	u.update()
}

func (u *User) update() {
	//u.UpdateAt = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	u.UpdateAt = time.Now()
}

func main() {
	//user := User{"Tadeu", 1, time.Now()}
	//fmt.Println(user)

	user2 := User{
		Name: "Bianca",
		ID:   2,
	}
	fmt.Println(user2)

	user2.GetGreetingMessage()
	fmt.Println(user2)

	time.Sleep(time.Second * 2)

	user2.UpdateName("Bianca Tupinamba")
	fmt.Println(user2)

	time.Sleep(time.Second * 2)

	user2.UpdateProfissaoName("Desenvolvedora") // this is an embedded method
	fmt.Println(user2)

	res, err := json.Marshal(user2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
}
