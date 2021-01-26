package main


import("fmt")


func test(n1 int){
	n1 = n1+1
	fmt.Println("test() n1= ",n1)
}

func getSum(n1 int,n2 int) int{
	sum:= n1+n2
	fmt.Println("getSum sum= ",sum)
	return sum
}

func getSumAndSub(n1 int,n2 int) (int,int){
	sum:=n1+n2
	sub:=n1-n2
	return sum,sub
}

func test02(n1 *int)  {
	fmt.Printf("test02() n1指向的地址 %v\n",n1)
	fmt.Printf("test02() n1的地址 %v\n",&n1)
	*n1 = *n1 + 30
	fmt.Println("test02() n1= ",*n1)
}

func main() {
	n1 := 10

	test(n1)
	fmt.Println("main() n1=",n1)
	test02(&n1)
	fmt.Printf("main() n1的地址 %v\n",&n1)
	fmt.Println("main() n1=",n1)

	sum := getSum(n1,n1)
	fmt.Println("main() sum= ",sum)

	res1,res2 := getSumAndSub(n1,sum)
	fmt.Println("main() res1= ",res1," res2= ",res2)

	_, res3 := getSumAndSub(n1,sum)
	fmt.Println("main() res3= ",res3)
}