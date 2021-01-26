package main

import (
 	"fmt"
)

type Usb interface{
	Say()
	GetName() string
}

type Stu struct{
	Name string
}

func (this *Stu) Say()  {
	fmt.Println("Say() ")
}
func (this *Stu) GetName() string {
	return this.Name
}
func main() {
	var stu Stu = Stu{
		Name:"tim",
	}
	var u Usb = &stu
	u.Say()
	fmt.Println("here ",u.GetName())
}

