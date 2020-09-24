package main

/*
	递归
*/
import "fmt"

func main() {
	fmt.Println(combinations(3, 2))
}

// 获取1~n数字集合中有多少中长度为k的组合
func combinations(n, k int) (res [][]int) {
	tmp := []int{}
	var dfs func(int)
	dfs = func(i int) {
		// 如果剩下的元素+tmp长度小于k则无需继续逻辑
		if len(tmp)+n-i+1 < k {
			return
		}

		// 满足条件添加至结果集
		if len(tmp) == k {
			c := make([]int, k)
			copy(c, tmp)
			res = append(res, c)
			return
		}
		// 两种情况
		// 1.组合中放入元素
		tmp = append(tmp, i)
		dfs(i + 1)
		// 2.从组合中移除元素
		tmp = tmp[:len(tmp)-1]
		dfs(i + 1)
	}
	// 开始递归
	dfs(1)
	return
}
