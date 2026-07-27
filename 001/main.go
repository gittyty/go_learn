package main

import "fmt"

func main() {
	const password_correct = "123456"
	const MAX_TRIES = 3
	var input string
	var score int
	for attempts := 0; attempts < MAX_TRIES; attempts++ {
		fmt.Println("请输入密码:")
		fmt.Scanln(&input)

		if input == password_correct {
			fmt.Println("登录成功")
			break
		}

		remaining := MAX_TRIES - attempts - 1
		if remaining == 0 {
			fmt.Println("账号已锁定")
		} else {
			fmt.Printf("密码错误，还剩%v次机会\n", remaining)
		}
	}
	if input == password_correct {
		fmt.Println("请输入一个分数:")
		fmt.Scanln(&score)
		switch {
		case score >= 90:
			fmt.Println("A级")
		case score >= 80:
			fmt.Println("B级")
		case score >= 70:
			fmt.Println("C级")
		case score >= 60:
			fmt.Println("D级")
		default:
			fmt.Println("不及格")
		}
	}
}
