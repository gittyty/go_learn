package main

import "fmt"

type User struct {
	Name     string
	password string
}

var userDB = make(map[string]User)

func (u User) CheckPassword(input string) bool {
	if u.password == input {
		return true
	}
	return false
}

func register() bool {
	fmt.Println("please input the username:")
	var user User
	fmt.Scanln(&user.Name)
	_, ok := userDB[user.Name]
	if ok {
		fmt.Println("username already exist!")
		return false
	}
	fmt.Println("please input the password:")
	fmt.Scanln(&user.password)
	userDB[user.Name] = user
	fmt.Printf("registe success,users number:%v\n", len(userDB))
	return true
}

func login() bool {
	fmt.Println("username:")
	var user_name string
	fmt.Scanln(&user_name)
	user, ok := userDB[user_name]
	if !ok {
		fmt.Println("user not exist,please registe first")
		return false
	}
	for attempts := 0; attempts < 3; attempts++ {
		var pass_word string
		fmt.Println("password:")
		fmt.Scanln(&pass_word)
		if user.CheckPassword(pass_word) {
			fmt.Println("login success,welcome")
			return true
		}
		if attempts == 2 {
			fmt.Println("locked,please start again")
			return false
		}
		fmt.Printf("incorrect password,ramain %v chances\n", 2-attempts)
	}
	return false
}

func main() {
	for {
		fmt.Printf("1.注册\n2.登录\n3.退出\n请选择操作:")
		var userChoose string
		fmt.Scanln(&userChoose)
		switch userChoose {
		case "1":
			register()
		case "2":
			login()
		case "3":
			return
		default:
			fmt.Println("无效选项")
		}
	}

}
