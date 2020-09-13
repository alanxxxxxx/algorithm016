// @Title  Week_01
// @Description  原地旋转数组
// @Author  alan  2020/9/13 17:07
// @Update  alan  2020/9/13 17:07
package Week_01

import "fmt"

/*
	三次翻转
	第一次翻转： 整个数组翻转
	第二次翻转： 翻转第一次翻转的数组前 0 - k%len(num)
	第三次翻转： 翻转第一次翻转的数组后 k%len(num) - len(num)
*/
func rotate(nums []int, k int) {
	reverse(nums)               // 一次翻转
	reverse(nums[:k%len(nums)]) // 二次翻转
	reverse(nums[k%len(nums):]) // 三次翻转
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func RotateArrayTest() {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	fmt.Println(nums)
	rotate(nums, 4)
	fmt.Println(nums)
}
