package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("好きな数字を入力してね: ")
	nStr, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(nStr[:len(nStr)-1])
	if err != nil {
		fmt.Println("数字を入力してね")
		return
	}

	fmt.Print("さっきの数字より大きい数字を入力してね: ")
	mStr, _ := reader.ReadString('\n')
	m, err := strconv.Atoi(mStr[:len(mStr)-1])
	if err != nil {
		fmt.Println("数字を入力してね")
		return
	}

	if n > m {
		fmt.Println("１回目の数字より大きい数を入力してね")
		return
	}

	// シード値を生成し、nからmまでのランダムな値を生成する
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(m-n+1) + n
	fmt.Printf("%dと%dの間のランダムな数字が生成されました!!\n", n, m)

	// 最大の試行回数と現在の試行回数を定義
	maxAttempts := 5
	attempts := 0

	for attempts < maxAttempts {
		fmt.Print("数字を当てて: ")
		guessStr, _ := reader.ReadString('\n')
		guess, err := strconv.Atoi(guessStr[:len(guessStr)-1])
		if err != nil {
			fmt.Println("数字を入力してね")
			continue
		}

		attempts++

		if guess < target {
			fmt.Println("小さいよ!")
		} else if guess > target {
			fmt.Println("大きいよ!")
		} else {
			fmt.Println("おめでとう!!!あたりだよ!!!!!!")
			return
		}

		if attempts == maxAttempts {
			fmt.Printf("ごめんね、チャレンジ失敗だよ。。。正解は %d でした!\n", target)
		}
	}

}
