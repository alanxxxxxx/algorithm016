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
