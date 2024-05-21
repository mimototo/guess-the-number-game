package main

import (
	"bufio"
	"fmt"

	// "math/rand"
	"os"
	"strconv"
	// "time"
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

	fmt.Print("さっきの入力より大きい数字を入力してね: ")
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
}
