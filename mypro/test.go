package main

import "fmt"

func f6() {
	//  Monday Tuesday Wednesday Thursday Friday Saturday Sunday
	var c string
	fmt.Println("输入一个字符：")
	fmt.Scan(&c)

	if c == "S" {
		fmt.Println("输入第二个字符：")
		fmt.Scan(&c)
		if c == "a" {
			fmt.Println("Saturday")
		} else if c == "u" {
			fmt.Println("Sunday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "F" {
		fmt.Println("Friday")
	} else if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("输入第二个字符：")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuesday")
		} else if c == "h" {
			fmt.Println("Thursday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	} else {
		fmt.Println("输入错误")
	}
}

func f1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	f1()
	f6()
	/* s := user.Hello()
	fmt.Printf("s: %v\n", s)
	const (
		a1 = iota //0
		_
		a2 = iota //2
		a3 = 100
		a4 = iota //4
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a4: %v\n", a4) */
	/* var name string = "tom"
	age := 20
	b := true
	fmt.Printf("name: %T\n", name)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("b: %T\n", b)
	a := &b
	fmt.Printf("a: %T\n", a)
	fmt.Printf("a: %v\n", a)
	c := [2]int{1, 2}
	fmt.Printf("c: %v\n", c)
	fmt.Printf("c: %T\n", c)
	d := []int{1, 2, 3}
	fmt.Printf("d: %v\n", d)
	fmt.Printf("d: %T\n", d)
	fmt.Printf("test: %v\n", test)
	fmt.Printf("test: %T\n", test) */
	/* s := "hello world"
	a := 2
	b := 5
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s: %c\n", s[a])//注意字符打印
	fmt.Printf("s: %v\n", s[2:5])
	fmt.Printf("s: %v\n", s[:5])
	fmt.Printf("s: %v\n", s[a:b]) */
}
