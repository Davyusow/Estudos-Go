package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{0, nil}
    current := result
    carry := 0
    for l1 != nil || l2 != nil || carry > 0  {
        sum := carry

        if (l1 != nil) {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        carry = sum / 10
        val := sum % 10

        current.Next = &ListNode{val, nil}
        current = current.Next

    }
    return result.Next
}

func main() {
    l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
    l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
    fmt.Println(addTwoNumbers(l1,l2))

}
