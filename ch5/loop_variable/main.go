package main

import "fmt"

func main() {
	fmt.Println("------------------loopVariable--------------------")
	loopVariable()
	fmt.Println("------------------evalLoopVariable--------------------")
	evalLoopVariable()
}

func loopVariable() {
	strings := []string{"a", "b", "c"}
	var printFuncList []func()
	for _, s := range strings {
		s := s // 注意：必要！
		fmt.Printf("s = %s (%p)\n", s, &s)
		printFuncList = append(printFuncList, func() {
			fmt.Printf("s = %s (%p)\n", s, &s)
		})
	}
	fmt.Println()
	for _, printFunc := range printFuncList {
		printFunc()
	}
}

func evalLoopVariable() {
	strings := []string{"a", "b", "c"}
	var printFuncList []func()
	for _, s := range strings {
		fmt.Printf("s = %s (%p)\n", s, &s)
		printFuncList = append(printFuncList, func() {
			fmt.Printf("s = %s (%p)\n", s, &s) // 注意：正しくない！
			// 同じ変数sを捕捉して共有していて最終的にループ処理でsはcに更新されるので、
			// 後続の関数が実行されるときには最後の値cが常に表示される
			// 参照: プログラミング言語Go(p160)
		})
	}
	fmt.Println()
	for _, printFunc := range printFuncList {
		printFunc()
	}
}
