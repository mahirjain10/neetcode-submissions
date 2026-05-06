/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int,0,0)

	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode){
		fmt.Println("traverse hyaah")
		if node == nil{
			return
		}
		traverse(node.Left)
		result = append(result,node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return result;
}
