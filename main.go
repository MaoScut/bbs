package main

import (
	"context"
	"fmt"
	"log"

	bbspb "github.com/MaoScut/bbs/gen/api"
	"github.com/MaoScut/bbs/service/user"
	_ "gorm.io/driver/mysql"
)

func main() {
	s := user.UserService{}
	res, err := s.CreateUser(context.Background(), &bbspb.CreateUserRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.User)
}
