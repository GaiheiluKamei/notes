package main

import (
	"fmt"
	"time"
	"math/rand"
	"errors"
) 

func minDistance(word1, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 || n == 0 {
		return 0
	}

	// 申请一个二维数组，标识从word1的第i个字符转化到word2的第j个字符需要的最小步数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// word1边界状态初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	// word2边界状态初始化
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	fmt.Println("m, n: ", m, n)
	// 开始动态规划
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}

	return dp[m][n]
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return min
}

func main() {

	go check()
	time.Sleep(time.Second * 10)
	fmt.Println("main sleeping...")
}

func check() {
	t := time.NewTicker(time.Millisecond*200)
	defer t.Stop()

	timer := time.NewTimer(time.Millisecond * 1200)

	for {
		select {
			// case <- time.After(time.Millisecond * 100):
			case <- timer.C:
				fmt.Println("超时")
				return 
			case <- t.C:
				if check1() {
					fmt.Println("check1: break")
					//return
				}
				fmt.Println("continue")
			default:	
		}
	}
}

func check1() bool {
	t := rand.Intn(3)
	if t == 0 {
		return true
	}

	return false
}

func check2() error {
	ticker := time.NewTicker(time.Millisecond * 10)
	defer ticker.Stop()

	timer := time.After(time.Millisecond * 100)

	for {
		select {
			case <- ticker.C:
				if check3() {
					fmt.Println("done")
					return nil
				}
				fmt.Println("continue")
			case <- timer:
				return errors.New("time out")
		}
	}
}

func check3() bool {
	t := rand.Intn(3)
	if t == 0 {
		return true
	}

	return false
}