package main


import "fmt"


type Usb interface{
	// 聲明兩個沒有實現的方法
	Start()
	Stop()
}

type Phone struct{

}

// 讓Phone實現Usb方法
func (p Phone) Start()  {
	fmt.Println("手機開始工作")
}

func (p Phone) Stop()  {
	fmt.Println("手機停止工作")
}

type Camera struct{

}

func (c Camera) Start()  {
	fmt.Println("相機開始工作")
}

func (c Camera) Stop()  {
	fmt.Println("相機停止工作")
}

type Computer struct{

}

// 接收一個Usb接口類型
// 只要實現Usb接口 (就是指實現了Usb街口聲明所有方法)
func (c Computer) Working(usb Usb)  {
	
	usb.Start()
	usb.Stop()
}
func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}