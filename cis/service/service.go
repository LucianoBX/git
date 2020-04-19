package service

import (
	// "fmt"
	"github.com/goEx/cis/model"

)

// 该CustomerService，完成对customer 的操作，包括
// CRUD
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段表示当前切片有多少个客户
	customerNum int
	
}


//  
func NewCustomerService() *CustomerService {
	// 创建初始示例
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "0775-88669988", "zs@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService

}


// Add方法将客户信息，加入切片中
func (cs *CustomerService) Add(cu model.Customer) bool {
	// 确定一个分配id的规则
	cs.customerNum++
	cu.Id = cs.customerNum
	cs.customers = append(cs.customers, cu)

	// 插入成功
	return true
}

// List 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
	
}


// 根据Id查找客户在切片中对应下标，如果没有该客户，返回-1
func (cs *CustomerService) Delete(id int) bool {
	
	index := cs.FindById(id)

	// 判断有没有这个客户
	if index == -1 {
	return false
	}

	// 切片中删除元素
	// NOTE: "..."不可忽略
	cs.customers = append(cs.customers[:index], cs.customers[index + 1:]...)
	
	return true
}


// 根据Id查找相应客户，没有则返回-1
func (cs *CustomerService) FindById(id int) int {

	index := -1
	// 遍历切片
	for i := 0; i < len(cs.customers); i++ {
		// 找到
		if cs.customers[i].Id == id {
			index = i
		}
	}
	return index
}
