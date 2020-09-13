// @Title  Week_01
// @Description  删除数组中的重复项（原地）
// @Author  alan  2020/9/13 15:46
// @Update  alan  2020/9/13 15:bn
package Week_01

import "fmt"

/*
	双指针操作，j为慢指针，i为快指针
	1. 当nums[i] == nums[j] 则 相邻重复 i++
	2. 当nums[i] != nums[j] 则 相邻不等交换元素,因为i比j走的要快所以不管j+1的下一个值是什么都不影响交换
*/
func removeDuplicates(nums []int) int {
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			nums[j+1] = nums[i]
			j++
		}
	}
	return j+1
}
func Test() {
	list := []int{1, 2, 2, 2, 2, 3, 4}
	fmt.Println(removeDuplicates(list))
	fmt.Println(list)
}
