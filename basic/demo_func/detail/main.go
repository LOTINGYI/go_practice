package main


import("fmt")

func getSum(n1 int,n2 int) int{
	sum:= n1+n2
	fmt.Println("getSum sum= ",sum)
	return sum
}

func myFun(funvar func(int,int) int, num1 int,num2 int)int  {
	return funvar(num1,num2)
}

func main() {
	// a:= getSum
	// fmt.Printf("a的類型 %T, getSum類型是 %T\n",a,getSum)

	// res:= a(10,40)
	// fmt.Printf("a的類型 %T, res類型是 %T\n",a,res)
	// fmt.Println("res= ",res)

	// res2:=myFun(getSum,50,60)
	// fmt.Println("res2= ",res2)


	/*在Go中myInt和int雖然都是int類型，但是go認為myInt和int是兩個類型*/
	type myInt int

	var num1 myInt
	var num2 int

	num1 = 40
	num2 = int(num1) // 所以要做轉換
	fmt.Println("num2 = ",num2)
}