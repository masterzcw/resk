package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Id      int    `json:"id,string"` // ,string以字符串形式序列化
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"` // omitempty如果为空值, 则不会被序列化
	Address string `json:"-"`             // -表示序列化时被忽略
}

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	u := User{
		Id:      12,
		Name:    "wendell",
		Age:     1,
		Address: "滨江区",
	}
	data, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	var u2 User
	err = json.Unmarshal(data, &u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
	fmt.Printf("%+v \n", u)
	fmt.Printf("%+v \n", u2)
}
