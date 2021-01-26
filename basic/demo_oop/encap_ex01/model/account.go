package model


import "fmt"



// 定義struct
type account struct{
	accountNo string
	pwd string
	balance float64
}

func (account *account) Deposit(money float64, pwd string)  {
	
	// 看下輸入的密碼是否正確
	if pwd != account.pwd {
		fmt.Println("密碼不正確")
		return
	}

	// 存款金額是否正確
	if money <= 0 {
		fmt.Println("輸入金額不正確")
		return
	}

	account.balance+=money
	fmt.Println("存款成功")
}

func (account *account) Withdraw(money float64, pwd string)  {
	
	// 看下輸入的密碼是否正確
	if pwd != account.pwd {
		fmt.Println("密碼不正確")
		return
	}

	// 存款金額是否正確
	if money > account.balance || money <=0 {
		fmt.Println("輸入金額不正確")
		return
	}

	account.balance-=money
	fmt.Println("提款成功")
}

func (account *account) Query(pwd string)  {
	
	// 看下輸入的密碼是否正確
	if pwd != account.pwd {
		fmt.Println("密碼不正確")
		return
	}

	fmt.Printf("你的帳號為=%v 餘額為=%v\n",account.accountNo,account.balance)
}

func NewAccount(accountNo string,pwd string,balance float64) *account{
	if len(accountNo)<6 || len(accountNo) > 10  {
		fmt.Println("帳號長度不對")
		return nil
	}
	if len(pwd)!=6 {
		fmt.Println("密碼長度不對")
		return nil
	}

	if balance < 20 {
		fmt.Println("餘額數目不對")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:pwd,
		balance: balance,
	}
}  