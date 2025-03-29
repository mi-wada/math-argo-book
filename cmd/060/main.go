package main

import (
	"fmt"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	// N=0の時、必ずターンプレイヤーが負ける。なのでSecondが勝つ。
	// N=1の時、必ずターンプレイヤーが負ける状態N=0でSecondにターンを回せる。なのでFirstが勝つ。
	// N=2の時、必ずターンプレイヤーが負ける状態N=0でSecondにターンを回せる。なのでFirstが勝つ。
	// N=3の時、必ずターンプレイヤーが負ける状態N=0でSecondにターンを回せる。なのでFirstが勝つ。
	// N=4の時、必ずターンプレイヤーが負ける。N=1,2,3のいずれかで相手のターンが来るため。なのでSecondが勝つ。
	// 以降ループする。

	if (n-1)%4 == 3 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
