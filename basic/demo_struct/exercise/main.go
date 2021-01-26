package main

import "fmt"

type Person struct{
	Name string
}

// 
func test01(p Person)  {
	fmt.Println(p.Name)
}

func test02(p *Person)  {
	fmt.Println(p.Name)
}

func (p Person) test03()  {
	p.Name = "mary"
	fmt.Println("test03() p.Name",p.Name)
}

func (p *Person) test04()  {
	p.Name = "julia"
	fmt.Println("test04() p.Name",p.Name)
}



func main() {
	p:= Person{
		Name:"tom",
	}

	test01(p)
	test02(&p)

	p.test03()
	fmt.Println("main() p.Name",p.Name)
	(&p).test03()
	fmt.Println("main() p.Name",p.Name)
	fmt.Println()
	p.test04()
	fmt.Println("main() p.Name",p.Name)
	(&p).test04()
	fmt.Println("main() p.Name",p.Name)
}