// @Title  main
// @Description  请填写文件描述（需要改）
// @Author  alan  2020/9/27 19:42
// @Update  alan  2020/9/27 19:42
package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 3,
	}
	res := permute(nums)
	fmt.Println(res)
	res = per01(nums)
	fmt.Println(res)
}
func permute(nums []int) (res [][]int) {
	if len(nums) == 1 {
		res = append(res, nums)
		return
	}
	for i, num := range nums {
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return
}
func per01(nums []int) (res [][]int) {
	var used = make([]bool, len(nums))
	var tmpList = []int{}
	var backtrace func([]int, []int, []bool)
	backtrace = func(nums []int, tmpList []int, used []bool) {
		// 匹配条件
		if len(tmpList) == len(nums) {
			c := make([]int, len(nums))
			copy(c, tmpList)
			res = append(res, c)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 选择01
			used[i] = true
			tmpList = append(tmpList, nums[i])
			// backtrace
			backtrace(nums, tmpList, used)
			// 选择02
			// 还原操作
			tmpList = tmpList[:len(tmpList)-1]
			used[i] = false
		}
		// 返回结果
		return

	}
	backtrace(nums, tmpList, used)
	return
}
