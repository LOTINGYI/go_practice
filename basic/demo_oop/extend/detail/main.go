package main


import "fmt"

type A struct{
	Name string
	age int
}


type B struct{
	Name string
	Score float64
}

type C struct{
	A
	B
}

type D struct{
	a A // 有名struct
}

type Goods struct{
	Name string
	Price float64
}

type Brand struct{
	Name string
	Address string
}

type TV struct{
	Goods
	Brand
}

type TV_point struct{
	*Goods
	*Brand
}
func main() {
	// var c C
	// c.Name = "martin"
	// c.A.Name = "tom"
	// c.B.Score = 22.0
	// fmt.Println(c)

	// var d D
	// d.a.Name = "jack"

	tv := TV{
		Goods{"電視機01",5000.0},
		Brand{"海爾","山東"},
	}
	tv2 := TV{
		Goods{
			Name:"電視機02",
			Price:50000.0,
		},
		Brand{
			Name:"海爾",
			Address:"山東",
		},
	}

	tv3 := TV_point{
		&Goods{
			Name:"電視機03",
			Price:7000.0,
		},
		&Brand{
			Name:"海爾",
			Address:"山東",
		},
	}
	fmt.Println(tv)
	fmt.Println(tv2)
	fmt.Println(*tv3.Goods,*tv3.Brand)
}