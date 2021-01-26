package main


import "fmt"


type Person struct{
	Name string
}

func (p Person) speak()  {
	fmt.Println(p.Name," 是個好人")
}

func (p Person) test()  {
	p.Name = "jack"
	fmt.Println("test() name= ", p.Name)
}

func (p Person) calc()  {
	res:= 0
	for i := 1; i < 5; i++ {
		res+=i
	}
	fmt.Println(p.Name+ " 計算結果= ",res)
}

func (p Person) calc2(n int)  {
	res:= 0
	for i := 1; i < n; i++ {
		res+=i
	}
	fmt.Println(p.Name+ " 計算結果= ",res)
}

func (p Person) getSum(n1 int,n2 int) int {
	return n1+n2
}

type Circle struct{
	radius float64
}

func (c *Circle) area2() float64  {
	// (*c).radius 等價 c.radius
	// return 3.14 * (*c).radius * (*c).radius
	fmt.Printf("area2() c變量= %p\n",c)
	c.radius = 10
	fmt.Println("area2() c.radius= ",c.radius)
	return 3.14 * c.radius * c.radius
}

func main() {
	
	// var p Person
	// p.Name = "tom"
	// p.test() // 調用方法
	// fmt.Println("main() p.Name= ", p.Name)
	// p.speak()
	// p.calc()
	// p.calc2(10)

	// n1 :=10
	// n2 := 20
	// res:= p.getSum(n1,n2)
	// fmt.Println("res= ",res)

	var c Circle
	fmt.Printf("main() c變量地址= %p\n",&c)
	c.radius = 5.0
	// res := (&c).area2()
	// 編譯器底層做了優化 (&c).area2() 等價 c.area2()
	// 因為編譯器會自動加上 &c
	res:= c.area2()
	fmt.Println("res= ",res)
	fmt.Println("main() c.radius= ",c.radius)
}