package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	max_num := 100
	rand.Seed(time.Now().UnixNano()) // 设置随机种子为当前时间戳，防止一直生成相同的随机数
	secret_number := rand.Intn(max_num)
	for {
		fmt.Println("Please input your guess!")

		var guess int
		_, err := fmt.Scan(&guess) // 读取用户输入

		if err != nil {
			fmt.Println("Invalid  input. Please enter an integer value!", err)
			continue
		}

		fmt.Println("You guess is", guess)
		if guess > secret_number {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secret_number {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}

}
