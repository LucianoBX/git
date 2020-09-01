package main

import (
	"fmt"
	"github.com/goEx/cis/service"
	"github.com/goEx/cis/model"

)

// 
type customerView struct {

	// 定义必要字段
	key string // 接受用户输入
	loop bool // 控制主菜单的终止

	//
	customerService *service.CustomerService
}


// 得到用户信息，并调用customerService完成添加
func (cv *customerView) add() {
	fmt.Println("----------------添加客户-----------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 999
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email:= ""
	fmt.Scanln(&email)

	// 构建一个新的Customer实例
	// NOTE：id号由系统直接分配
	customer := model.NewCustomer2(name, gender, age, phone, email)

	// 调用
	if cv.customerService.Add(customer) {
		fmt.Println("---------------添加完成------------------")
	} else {
		fmt.Println("---------------添加失败------------------")
	}
}


// 显示所有用户信息
// 客户信息显示的实现，list调用CustomerService,并显示客户列表
func (cv *customerView) list() {
	
	// 首先获取所有客户信息
	customers := cv.customerService.List()

	// 显示所有用户信息
	fmt.Println("----------------添加客户-----------------")
	fmt.Println("编号	姓名	性别	年龄	电话		邮箱")
	for i:=0; i < len(customers); i++ {
	
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Println("---------------添加完成------------------")
}


// 主菜单显示
func (cv *customerView) mainMenu() {
	for {
	
	fmt.Println("----------------客户信息管理软件-----------------")
	fmt.Println("		1 添加客户")
	fmt.Println("		2 修改客户")
	fmt.Println("		3 删除客户")
	fmt.Println("		4 客户列表")
	fmt.Println("		5 退    出")
	fmt.Println("请选择(1-5):")
	fmt.Println("-------------------------------------------------")

	fmt.Scanln(&cv.key)
	switch cv.key {
	case "1" :
		// 添加客户
		cv.add()
		//fmt.Println("添加客户")

	case "2" :
		fmt.Println("修改客户")

	case "3" :
		cv.delete()
	//	fmt.Println("删除客户")

	case "4" :
		// 客户列表显示
		cv.list()

	case "5" :
		fmt.Println("退    出")
		fmt.Println("你已经退出客户管理软件")
		cv.loop = false
	default :
	fmt.Println("输入有误，请重新输入(1-5):")
	}

	// 退出控制
	if !cv.loop {
		break
	}

	}
}


// 输入id，删除相应客户
func (cv *customerView) delete() {
	fmt.Println("---------------删除客户-------------------")
	fmt.Println("请选择待删除客户编号（-1推出）：")
	id := -1
	fmt.Scanln(&id)
	
	if id == -1 {
		return
	
	}

	fmt.Println("确认是否删除（Y/N）:")

	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" || choice == "yes" {

		// 调用customerService中的Delete方法
		if cv.customerService.Delete(id) {
			fmt.Println("---------------删除完成-------------------")
		} else {
			fmt.Println("---------------客户Id不存在-------------------")

		}
	} else if ( choice == "N" || choice == "n" ) {
		fmt.Println("---------------取消删除-------------------")
	} else {
		fmt.Println("输入错误")
	}
}


func main() {
	fmt.Println("vim-go")

	// 创建一个customerView实例
	customerView := customerView{
		key : "",
		loop : true,
	}

	// 完成对customerView.custoemrServce的初始化
	customerView.customerService = service.NewCustomerService()

	// 显示菜单
	customerView.mainMenu()
}
