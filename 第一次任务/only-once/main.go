package main

import (
	"fmt"
	"slices"
)

func main() {

	// nums := []int{1, 2, 2, 3, 3, 4, 4}
	// reuslt := singleNumber(nums)
	// fmt.Println(reuslt)

	// nums2 := []int{1, 2, 3, 1}
	// result := rob(nums2)
	// fmt.Println(result)

	// node3 := ListNode{3, nil}
	// node2 := ListNode{2, &node3}
	// node1 := ListNode{1, &node2}

	// node6 := ListNode{4, nil}
	// node5 := ListNode{2, &node6}
	// node4 := ListNode{1, &node5}

	// node := mergeTwoList(&node1, &node4)

	// var next *ListNode = node.NextNode
	// for {
	// 	if next != nil {
	// 		fmt.Println(next.Val)
	// 		next = next.NextNode
	// 	} else {
	// 		break
	// 	}
	// }

	// str:="hello"
	// bytes:= []byte (str)
	// reverseString(bytes)
	// fmt.Println(string(bytes))

	// fmt.Println(mySqrt(10))

	nums := []int{1, 1, 2, 2}
	fmt.Println(removeDuplicates(nums))

}

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
// 其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
// 结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素
func singleNumber(nums []int) int {
	countMap := map[int]int{}
	for index := range nums {
		val := nums[index]
		numberCount := countMap[val]
		numberCount++
		countMap[val] = numberCount
	}

	for key, value := range countMap {
		if value == 1 {
			return key
		}
	}

	return 0
}

// 198. 打家劫舍：你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，
// 系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，
// 一夜之内能够偷窃到的最高金额。这道题可以使用动态规划的思想，通过 for 循环遍历数组，
// 利用 if 条件判断来决定是否选择当前房屋进行抢劫，状态转移方程为 dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size+1, size+1)
	dp[0] = 0
	dp[1] = nums[0]

	for k := 2; k <= size; k++ {
		dp[k] = max(dp[k-1], dp[k-2]+nums[k-1])
	}
	return dp[size]
}

// 21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。可以定义一个函数，接收两个链表的头节点作为参数，
// 在函数内部使用双指针法，通过比较两个链表节点的值，将较小值的节点添加到新链表中，直到其中一个链表为空，
// 然后将另一个链表剩余的节点添加到新链表中。
type ListNode struct {
	Val      int
	NextNode *ListNode
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	merge := ListNode{}

	if list1.Val < list2.Val {
		merge.Val = list1.Val
		merge.NextNode = mergeTwoList(list1.NextNode, list2)
	} else {
		merge.Val = list2.Val
		merge.NextNode = mergeTwoList(list1, list2.NextNode)
	}
	return &merge
}

// 46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
// 可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，
// 使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return res
}

//344. 反转字符串：编写一个函数，其作用是将输入的字符串反转过来。

// 输入字符串以字符数组 char[] 的形式给出。不要给另外的数组分配额外的空间，
// 你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
// 可以使用 for 循环和两个指针，一个指向字符串的开头，一个指向字符串的结尾，
// 然后交换两个指针所指向的字符，直到两个指针相遇。
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

// 69. x 的平方根：实现 int sqrt(int x) 函数。计算并返回 x 的平方根，
// 其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
// 可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，
// 直到找到满足条件的平方根或者确定不存在精确的平方根。
func mySqrt(s int) int {
	l, r := 0, s
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= s {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，
// 使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，
// 你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
// 一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			nums[p+1] = nums[q]
			p++
		}
		q++
	}
	return p + 1
}

// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，
// 该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
	ans := [][]int{}
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		temp := ans[len(ans)-1]
		if temp[1] >= intervals[i][0] {
			temp[1] = max(temp[1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}


//430. 扁平化多级双向链表：多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，
// 可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
// 给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
// 可以定义一个结构体来表示链表节点，包含 val、prev、next 和 child 指针，然后使用递归的方法来扁平化链表，
// 先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间

type Node struct {

	    Val int
	    Prev *Node
	    Next *Node
	    Child *Node
 }

 func flatten(root *Node) *Node{
	dfs(root)
	return root
 }

func dfs(root *Node) (last *Node) {
    cur:=root
    for cur != nil {
        next :=cur.Next
        if cur.Child !=nil {
            childLast :=dfs(cur.Child)
            next = cur.Next

            cur.Next = cur.Child
            cur.Child.Prev = cur

            if next !=nil {
                childLast.Next = next
                next.Prev = childLast
            }

            cur.Child = nil
            last = childLast
        } else {
            last = cur
        }
        cur = next
    }
    return 
}


//729. 我的日程安排表 I：实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，
// 则可以存储这个新的日程安排。当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
// 日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end) ，实数 x 的范围为 start <= x < end 。
// 实现 MyCalendar 类：MyCalendar() 初始化日历对象。boolean book(int start, int end) 
// 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true ，否则，返回 false 并且不要将该日程安排添加到日历中。
// 可以定义一个结构体来表示日程安排，包含 start 和 end 字段，然后使用一个切片来存储所有的日程安排，
// 在 book 方法中，遍历切片中的日程安排，判断是否与要添加的日程安排有重叠

type Pair struct{
	start,end int
}

type MyCalendar []Pair

func Constructor() MyCalendar {
    return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
    for _,p:=range *c {
		if p.start< end && start< p.end {
			return false
		}
	}
	*c = append(*c,Pair{start,end})
	return true
}