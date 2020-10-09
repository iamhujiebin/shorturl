package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/fx"
)

type ClassData struct {
	ClassId int
}

type ClassObj struct {
	ClassId    int
	StudentIds []int
}

//模仿java8的stream流式api操作。主要是sort/filter/tail,多了一个walk可以中途改变流结构体(生成新的stream)
func main() {
	data := []*ClassData{
		{ClassId: 5},
		{ClassId: 3},
		{ClassId: 2},
		{ClassId: 4},
		{ClassId: 1},
	}
	result := make([]*ClassObj, 0)
	fx.From(func(source chan<- interface{}) {
		for _, item := range data {
			source <- item
		}
	}).Filter(func(item interface{}) bool {
		each := item.(*ClassData)
		if each.ClassId > 1 {
			return true
		}
		return false
	}).Walk(func(item interface{}, pipe chan<- interface{}) { //walk可以"中途"改变流的接口体,比如这里就可以中途把流从ClassData转为ClassObj了
		each := item.(*ClassData)
		students := make([]int, each.ClassId)
		pipe <- &ClassObj{
			ClassId:    each.ClassId,
			StudentIds: students,
		}
	}).Sort(func(a, b interface{}) bool {
		a1, b1 := a.(*ClassObj), b.(*ClassObj)
		return a1.ClassId > b1.ClassId
	}).ForEach(func(item interface{}) {
		o := item.(*ClassObj)
		//result[o.ClassId] = o.StudentIds
		result = append(result, o)
	})
	fmt.Printf("data:%+v", result)
}
