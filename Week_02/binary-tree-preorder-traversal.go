// @Title  Week_02
// @Description  二叉树前序遍历
// @Author  alan  2020/9/20 22:58
// @Update  alan  2020/9/20 22:58
package Week_02

func preorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}
