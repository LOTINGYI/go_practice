package utils

import "fmt"


type FamilyAccount struct{
	key string
	loop bool

	// 定義收支明細變量
	balance float64
	money float64
	note string
	details string
	flag bool
}

// Factory Pattern
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:"",
		loop: true,
		balance:10000.0,
		money:0.0,
		note:"",
		flag:false,
		details: "收支\t帳戶金額\t收支金額\t說明",
	}
}

func (this *FamilyAccount) showDetails()  {
	fmt.Println("-----------------當前收支明細------------------")
	if this.flag {
		fmt.Println(this.details)
	}else{
		fmt.Println("當前沒有收支")
	}
}

func (this *FamilyAccount) income()  {
	fmt.Println("本次收入金額: ")
	fmt.Scanln(&this.money)
	fmt.Println("本次收入說明: ")
	fmt.Scanln(&this.note)
	this.balance+=this.money
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance,this.money,this.note)
	this.flag = true
}

func (this *FamilyAccount) pay()  {
	fmt.Println("本次支出金額: ")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("餘額不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出說明: ")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance,this.money,this.note)
	this.flag = true
}


func (this *FamilyAccount) quit()  {
	fmt.Println("你確定要退出嗎")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice =="n" {
			break
		}
		fmt.Println("你的輸入有誤，請輸入y/n?")
	}
	if choice =="y" {
		this.loop = false
	}
}

// 顯示主菜單
func (this *FamilyAccount) MainMenu()  {
	for {
		fmt.Println("---------------家庭收支記帳軟件-------------")
		fmt.Println("                 1. 收支明細					")
		fmt.Println("                 2. 登記收入					")
		fmt.Println("                 3. 登記支出					")
		fmt.Println("                 4. 退出軟件					")
		
		fmt.Print("請選擇1-4  ")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.quit()
			default: 
				fmt.Println("請輸入正確選項")
		}

		if !this.loop {
			break
		}
	}
}