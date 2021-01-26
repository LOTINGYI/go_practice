package main
import("fmt")

func main() {
	// var n1 int32 = 12
	// var n2 int64
	// var n3 int8

	// n2 = int64(n1+20)
	// n3 = int8(n1+20)
	var n1 int32 = 12
	// var n3 int8
	var n4 int8
	n4 = int8(n1) + 127
	// n3 = int8(n1) + 128
	fmt.Println(n4)
}