package main

import (
	"fmt"

	"github.com/wafiqpuyol/Go-Learning/auth"
	"github.com/wafiqpuyol/Go-Learning/user"
)

func main() {
	user.CreateUser(user.User{
		Fname: "wafiq",
		Lname: "puyol",
		Age:   34,
	})
	user.CreateUser(user.User{
		Fname: "John",
		Lname: "doe",
		Age:   69,
	})
	user.CreateUser(user.User{
		Fname: "Beth",
		Lname: "lily",
		Age:   26,
	})

	fmt.Println("all users ==> ", user.GetUser())
	fmt.Println(auth.IsUseExist("wafiq"))
}
