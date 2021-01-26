package service

import (
	"go_code/proj/client_manage/model"
)

// CRUD
type CustomerService struct{
	customers []model.Customer
	// 聲明一個自斷，表示當前切片含有多少客戶
	customerNum int


}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}

	customerService.customerNum = 1
	customer := model.NewCustomer(1,"張三","男",20,"112","test@gmail.com")
	customerService.customers = append(customerService.customers,customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool  {
	// 分配id
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}

func (this *CustomerService) Delete(id int) bool  {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}

func (this *CustomerService) Update(id int,
	name string,gender string, age int, phone string, email string) bool  {
	this.customers[id].Name = name
	this.customers[id].Gender = gender
	this.customers[id].Age = age
	this.customers[id].Phone = phone
	this.customers[id].Email = email
	return true
}

func (this *CustomerService) FindById(id int) int  {
	index := -1

	// 遍歷
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}