package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}
type User1 struct {
	Name string
	Age  int
}

func (u *User1) show() {
	fmt.Println("Type, ", reflect.TypeOf(u))
	fmt.Println("Value, ", reflect.ValueOf(u))
}

//反射小练习，大家可以看下这篇文章
//https://www.jb51.net/article/275637.htm
func main() {
	a := 100.89
	b := "hello"
	u := User{
		Name: "wly",
		Age:  100,
	}
	d := 300
	fmt.Println(reflect.TypeOf(a))             //float64
	fmt.Println(reflect.TypeOf(b))             //string
	fmt.Println(reflect.TypeOf(u))             //main.User
	fmt.Println(reflect.ValueOf(a))            //100.89 这种返回的是副本
	fmt.Println(reflect.ValueOf(b))            //hello
	fmt.Println(reflect.ValueOf(u).NumField()) //{wly}
	if reflect.TypeOf(a).Kind() == reflect.Float64 {
		fmt.Println("ok")
	}
	if reflect.ValueOf(a).Kind() == reflect.Float64 {
		fmt.Println("ok1")
	}
	fmt.Println(reflect.TypeOf(&a).Elem())
	fmt.Println(reflect.ValueOf(&a).Elem())

	fmt.Println(reflect.TypeOf(&u).Elem().NumField()) //利用Elem可以修改对象的值
	fmt.Println(reflect.ValueOf(&u).Elem().NumField())

	var x int = 100
	value := reflect.ValueOf(&x).Elem()
	if value.CanSet() {
		fmt.Println("true")
		value.SetInt(200)
	}
	fmt.Println(x)

	fmt.Println("=========")
	u1 := User1{Name: "ttr", Age: 18}
	u1.show()
	fmt.Println(reflect.TypeOf(u).NumField()) //{wly}
	//通过指针可以获取对应的对象
	v := reflect.ValueOf(&d).Elem()
	if v.Kind() == reflect.Int {
		//获取对应的值然后修改对应的值，ValueOf必须为指针
		v.SetInt(v.Int() + 100)
	}
	fmt.Println(d)

}
