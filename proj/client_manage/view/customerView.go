package main

import (
	"fmt"
	"go_code/proj/client_manage/service"
	"go_code/proj/client_manage/model"
)


type customerView struct{
	// 接收輸入
	key string
	loop bool

	customerService *service.CustomerService
}
func (this *customerView) exit()  {
	fmt.Println("確認是否刪除(Y/N): ")
	for{
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key =="y" ||this.key == "N" || this.key =="n" {
			break
		}
		fmt.Println("請重新輸入(Y/N)")
	}
	if this.key == "Y" || this.key =="y"  {
		this.loop =false
	}

}


func (cv *customerView) mainMenu()  {
	for{
		fmt.Println("-----------------客戶訊息管理----------------------")
		fmt.Println("                 1. 添加客戶")
		fmt.Println("                 2. 修改客戶")
		fmt.Println("                 3. 刪除客戶")
		fmt.Println("                 4. 客戶列表")
		fmt.Println("                 5. 退出")
	
		fmt.Println("請選擇(1~5) ")
		fmt.Scanln(&cv.key)
	
		switch cv.key {
			case "1":
				cv.add()
			case "2":
				cv.update()
			case "3":
				cv.delete()
			case "4":
				cv.list()
			case "5":
				cv.exit()
			default:
				fmt.Println("輸入有誤")
		}
		if !cv.loop {
			break
		}	
	}
	fmt.Println("你已退出軟件")
	
}


func (this *customerView) list()  {
	customers := this.customerService.List()
	fmt.Println("------------------客戶列表-----------------------")
	fmt.Println("編號\t姓名\t性別\t年齡\t電話\t郵箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------------客戶列表完成----------------------")
}

func (this *customerView) add()  {
	fmt.Println("------------------添加客戶-----------------------")
	fmt.Println("姓名: ")
	name :=""
	fmt.Scanln(&name)
	fmt.Println("性別: ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年齡: ")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("電話: ")
	phone :=""
	fmt.Scanln(&phone)
	fmt.Println("電郵: ")
	email :=""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Add(customer){
		fmt.Println("------------------添加成功-----------------------")
	}else{
		fmt.Println("------------------添加失敗-----------------------")
	}
}
func (this *customerView) update()  {
	fmt.Println("------------------修改客戶-----------------------")
	fmt.Println("請選擇要更改的客戶編號(-1退出): ")
	
	id:=0
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := this.customerService.FindById(id)
	customers := this.customerService.List()
	fmt.Printf("序號:%v\n",customers[index].Id)
	fmt.Printf("姓名:%v",customers[index].Name)
	name :=""
	fmt.Scanln(&name)
	fmt.Printf("性別:%v ",customers[index].Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年齡:%v ",customers[index].Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("電話:%v ",customers[index].Phone)
	phone :=""
	fmt.Scanln(&phone)
	fmt.Printf("電郵:%v ",customers[index].Email)
	email :=""
	fmt.Scanln(&email)
	if this.customerService.Update(index,name,gender,age,phone,email) {
		fmt.Println("------------------修改成功-----------------------")
	}else{
		fmt.Println("------------------修改成功-----------------------")
	}
}
func (this *customerView) delete()  {
	fmt.Println("------------------刪除客戶-----------------------")
	fmt.Println("請選擇帶刪除的客戶編號(-1退出): ")
	id:=0
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("確認是否刪除(Y/N): ")
	choice :=""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		if this.customerService.Delete(id){
			fmt.Println("------------------刪除成功-----------------------")
		}else{
			fmt.Println("------------------刪除失敗-----------------------")
		}
	}
}


func main()  {
	fmt.Println("ok")

	customerView := customerView{
		key:"",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
	
}