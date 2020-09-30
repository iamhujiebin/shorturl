package model

type User struct {
	Name   string `o:"find,get,set" c:"姓名"`
	Age    int    `o:"find,get,set" c:"年纪"`
	School string `c:"学校"`
}
