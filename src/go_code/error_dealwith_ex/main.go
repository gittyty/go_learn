package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type User struct {
    username string
}

// getUser 验证用户名，非空则返回有效用户
func getUser(u User) (User, error) {
    trimmed := strings.TrimSpace(u.username)
    if trimmed == "" {
        return User{}, fmt.Errorf("用户名不能为空")
    }
    u.username = trimmed
    return u, nil
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    var user User

    // 最多尝试 4 次
    maxAttempts := 4
    for attempt := 1; attempt <= maxAttempts; attempt++ {
        fmt.Print("请输入用户名: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("读取输入失败:", err)
            // 如果是严重错误，可以 return 或 continue，这里直接退出
            return
        }

        username := strings.TrimSpace(input)
        user, err = getUser(User{username: username})
        if err == nil {
            break // 成功，跳出循环
        }

        // 如果还有机会，打印错误并继续
        if attempt < maxAttempts {
            fmt.Printf("错误: %v，还剩 %d 次机会\n", err, maxAttempts-attempt)
        } else {
            fmt.Printf("错误: %v，已达最大尝试次数，程序退出\n", err)
            return
        }
    }

    fmt.Printf("用户信息: %+v\n", user)
    // 后续可以使用 user
}