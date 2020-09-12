// @Title  Week_01
// @Description  移动零 leetcode[283]
// @Author  alan  2020/9/12 16:02
// @Update  alan  2020/9/12 16:02
package Week_01

/*
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[j] != 0 { // 如果 j 不等于 0 则右移一位
			j++
			continue
		}
		if nums[j] == 0 && nums[i] != 0 {
			nums[j] = nums[i]
			nums[i] = 0
			j++
		}
	}
	return
}
