package main


import "fmt"



// 定義struct
type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

func (account *Account) Deposit(money float64, pwd string)  {
	
	// 看下輸入的密碼是否正確
	if pwd != account.Pwd {
		fmt.Println("密碼不正確")
		return
	}

	// 存款金額是否正確
	if money <= 0 {
		fmt.Println("輸入金額不正確")
		return
	}

	account.Balance+=money
	fmt.Println("存款成功")
}

func (account *Account) Withdraw(money float64, pwd string)  {
	
	// 看下輸入的密碼是否正確
	if pwd != account.Pwd {
		fmt.Println("密碼不正確")
		return
	}

	// 存款金額是否正確
	if money > account.Balance || money <=0 {
		fmt.Println("輸入金額不正確")
		return
	}

	account.Balance-=money
	fmt.Println("提款成功")
}

func (account *Account) Query(pwd string)  {
	
	// 看下輸入的密碼是否正確
	if pwd != account.Pwd {
		fmt.Println("密碼不正確")
		return
	}

	fmt.Printf("你的帳號為=%v 餘額為=%v\n",account.AccountNo,account.Balance)
}

func main() {
	account:= Account{
		AccountNo: "gs11111",
		Pwd: "66666",
		Balance: 1000.0,
	}	

	account.Query("66666")
	account.Deposit(2005.0,"66666") 
	account.Query("66666")
	account.Withdraw(34000,"66666")
}