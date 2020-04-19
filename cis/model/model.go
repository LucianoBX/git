package model

import (
	"fmt"
)

// 声明客户信息,Custom结构
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

// 工厂函数,生成Customer实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string ) Customer {
		return Customer{
			Id : id,
			Name: name,
			Gender: gender,
			Age: age,
			Phone: phone,
			Email: email,
		}
	}

		
// 工厂函数,生成Customer实例
func NewCustomer2(name, gender string, age int,
	phone, email string) Customer {
		return  Customer{
			Name: name,
			Gender: gender,
			Age: age,
			Phone: phone,
			Email: email,
		}

	}


// 返回用户格式化的信息//
func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t\t%v", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}

