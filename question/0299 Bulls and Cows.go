package question

import (
	"fmt"
	"strings"
)

// 你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：
//写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：
//猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls", 公牛），
//有多少位属于数字猜对了但是位置不对（称为 "Cows", 奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
//给你一个秘密数字secret 和朋友猜测的数字guess ，请你返回对朋友这次猜测的提示。
//提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B表示奶牛。
//请注意秘密数字和朋友猜测的数字都可能含有重复数字。

func getHint(secret string, guess string) string {
	return getHintCore2(secret, guess)
}

// 如果两个位置上的数字相同，则说明是Bulls
// 有多少个非Bulls数字可以通过重新排列转换成Bulls即为Cows
// 主要思想：使用两个map，一个map记录Bulls，另一个记录出现的次数
func getHintCore1(secret string, guess string) string {
	if len(secret) != len(guess) {
		return ""
	}
	// Bulls的数量，即A前面的数字
	bulls := make(map[int]struct{})
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls[i] = struct{}{}
		}
	}
	// Cows
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if _, ok := bulls[i]; !ok {
			m1[secret[i]]++
		}
	}
	for i := 0; i < len(guess); i++ {
		if _, ok := bulls[i]; !ok {
			m2[guess[i]]++
		}
	}
	B := 0
	for b, cnt := range m2 {
		if have, ok := m1[b]; ok {
			B += min(have, cnt)
		}
	}
	transfer := func(n int) []byte {
		if n == 0 {
			return []byte{'0'}
		}
		res := make([]byte, 0, 4)
		for n != 0 {
			res = append(res, byte(n%10)+'0')
			n /= 10
		}
		for i := 0; i < len(res)/2; i++ {
			res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
		}
		return res
	}
	// 最后通过拼接，要注意数字可能由多位
	res := strings.Builder{}
	res.Write(transfer(len(bulls)))
	res.WriteByte('A')
	res.Write(transfer(B))
	res.WriteByte('B')
	return res.String()
}

// 优化方向：只有数字，所以可以用一个计数器和两个长度为10的数组取代map的使用
// 此外可以用Fprintf替代strings.Builder
func getHintCore2(secret string, guess string) string {
	bulls, cows := 0, 0
	cntS, cntG := make([]int, 10), make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(cntS[i], cntG[i])
	}
	return fmt.Sprintf("%dA%dB",bulls,cows)
}
