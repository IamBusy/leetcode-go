package AddTwoNumbers



//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res ,p1,p2 *ListNode
	p1 = l1
	p2 = l2

	var carry int=0
	var last *ListNode=res

	for ;p1!=nil || p2!=nil || carry!=0;  {

		var tmp = new(ListNode)
		v1 := 0
		v2 := 0
		if p1!=nil{
			v1=p1.Val
			p1 = p1.Next
		}
		if p2!=nil{
			v2=p2.Val
			p2 = p2.Next
		}

		sum := v1+v2+carry

		carry=sum/10

		tmp.Val = sum%10
		tmp.Next = nil


		if last==nil{
			res = tmp
			last = tmp
			continue

		}

		last.Next = tmp
		last = tmp

	}
	return res

}