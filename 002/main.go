package main

import "fmt"

var userList []string
var userDB = make(map[string]string)

func register() bool {
	fmt.Println("请输入注册用户名：")
	var user_name, password string
	fmt.Scanln(&user_name)
	_, ok := userDB[user_name]
	if ok {
		fmt.Println("用户名已存在")
		return false
	}
	fmt.Println("请输入注册密码：")
	fmt.Scanln(&password)
	userList = append(userList, user_name)
	userDB[user_name] = password
	fmt.Printf("注册成功，当前共%d位用户\n", len(userList))
	return true
}

func login() bool {
	var user_name string
	fmt.Println("请输入用户名：")
	fmt.Scanln(&user_name)
	_, ok := userDB[user_name]
	if !ok {
		fmt.Println("用户名不存在")
		return false
	}
	for attempts := 0; attempts < 3; attempts++ {
		var password string
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if userDB[user_name] == password {
			fmt.Println("登录成功，欢迎你")
			return true
		}
		fmt.Printf("密码错误,还剩%d次机会\n", 3-attempts-1)
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
