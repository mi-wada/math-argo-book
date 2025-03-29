package main

import (
	"fmt"
	"sort"
	"strconv"
)

var n, b, ans int64

func main() {
	// N B
	fmt.Scan(&n, &b)

	dfs("1")
	dfs("2")
	dfs("3")
	dfs("4")
	dfs("5")
	dfs("6")
	dfs("7")
	dfs("8")
	dfs("9")
	fmt.Println(ans)
}

func f(i int64) int64 {
	var ret int64 = 1
	for i > 0 {
		ret *= i % 10
		i /= 10
	}
	return ret
}

// sは数字のみの文字列である。sの各桁の積がBの和が、sの並び替えであればそれは条件を満たすのでansをインクリメントする。
// ただし、sの長さが11を超える場合は範囲外であるためreturnする。
// この再帰関数はソートされた数字列に対してのみ探索を行う。
func dfs(s string) {
	if len(s) == 11 {
		return
	}

	sInt, _ := strconv.ParseInt(s, 10, 64)
	fPlusB := strconv.FormatInt(f(sInt)+b, 10)
	if eqSorted(s, fPlusB) {
		ans++
	}

	last, _ := strconv.Atoi(string(s[len(s)-1]))
	// 桁を増やす方向への探索
	for i := last; i <= 9; i++ {
		dfs(s + strconv.Itoa(i))
	}
}

// aとbは辞書順にsortした結果同じになるかどうか
func eqSorted(a, b string) bool {
	aRunes := []rune(a)
	sort.Slice(aRunes, func(i, j int) bool { return aRunes[i] < aRunes[j] })
	bRunes := []rune(b)
	sort.Slice(bRunes, func(i, j int) bool { return bRunes[i] < bRunes[j] })

	return string(aRunes) == string(bRunes)
}
