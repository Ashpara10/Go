package main

import (
	"fmt"
)

type Todo struct {
	ID      string
	title   string
	content string
	user    *User
}
type User struct {
	ID   string
	name string
}

var (
	arr []Todo
  nums []int
)

func main() {
	arr = append(arr, Todo{ID: "1", title: "Hello This is a new       Todo", content: "dummy", user: &User{ID: "2", name: "Ash"}})
  
  nums=append(nums,89)
  nums=append(nums,489)
  nums=append(nums,26)
  nums=append(nums,46)
  for i :=range nums{
    fmt.Println(nums[i]*10/2)
  }
  
	// fmt.Println("Hello", arr)
}


func Hello(name string) string {
	fmt.Printf("Hello %v , You are Welcomed \n", name)
	return name
}
