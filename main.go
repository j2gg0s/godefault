package main

import (
	"fmt"

	v1 "github.com/j2gg0s/kserver/apis/user/v1"
)

func main() {
	user := v1.User{}
	v1.Default(&user)
	fmt.Println(*user.Age)
}
