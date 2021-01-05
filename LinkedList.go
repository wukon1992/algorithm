
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 反转一个链表 对应leedcode 206 题 

/**
 * 递归算法
 * 假设列表为：n1->...->nk-1->nk->nk+1->...->nm->null
 	若从节点 nk+1 到nm已经被反转，而我们正处于nk : n1->...->nk-1->nk->nk+1<-...<-nm
	我们希望nk+1的下一个节点指向nk	，所以 nk.next.next = nk 
	要小心的是 n1的下一个必须指向null
 */
func reverseList(head *ListNode) *ListNode {
	// 递归返回结果，
	if(head == nil || head.Next == nil) {
		return head
	}
	newHead = reverseList(head.Next)
	head.Next.Next = head //
	head.Next = nil //n1节点的指向为nil
	return newHead
}

//环形链表 对应leetcode 141题目
/**
 * 快慢指针
 * 我们定义两个指针，一快一满。慢指针每次只移动一步，而快指针每次移动两步。初始时，慢指针在位置 head，而快指针在位置 head.next
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    fastNode := head.Next
    slowNode := head

    for fastNode != nil && slowNode != nil && fastNode.Next != nil{
        if(slowNode == fastNode.Next){
            return true
        }
        fastNode = fastNode.Next.Next
        slowNode = slowNode.Next
    }
    return false
}
// 合并两个有序链表
/**
 * 同时遍历两个链表,比较大小。将小的结点放到哨兵结点tmp下
 */
func mergeTwoLists（l1 *ListNode, l2 *ListNode）*ListNode{
	tmpNode := &ListNode{}
	p := tmpNode //引用
	for l1.Next != nil && l2.Next != nil {
		if l1.Val < l2.Val {
            p.Next = l1
            l1 = l1.Next
        } else {
            p.Next = l2
            l2 = l2.Next
        }
        p = p.Next
	}
	if l1 != nil {
        p.Next = l1
    }
    if l2 != nil {
        p.Next = l2
    }
    return tmpNode.Next
}

// 删除链表第N个节点
/**
 * 快慢指针。放快指针到第N+1个结点的时候慢指针开始走。当快指针到达末尾则慢指针的下一个节点刚好是第N个
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head} //哨兵节点，防止链表长度等于1的时候无法指向Next结点
    first, second := head, dummy
    for i := 0; i < n; i++ {
        first = first.Next
    }
    for ; first != nil; first = first.Next {
        second = second.Next
    }
    second.Next = second.Next.Next
    return dummy.Next
}

