package main


import "fmt"


type Birdable interface{
	Flying()
}

type Fishable interface{
	Swimming()
}

type Monkey struct{
	Name string
}

type LittleMonkey struct{
	Monkey
}


func (this *Monkey) climbing()  {
	fmt.Println(this.Name, "生來會爬樹")
}


func (this *Monkey) Flying()  {
	fmt.Println(this.Name, "學會飛翔")
}

func (this *Monkey) Swimming()  {
	fmt.Println(this.Name, "學會游泳")
}


func main() {
	monkey01 := LittleMonkey{
		Monkey{
			Name:"悟空",
		},
	}

	monkey01.climbing()
	monkey01.Flying()
}