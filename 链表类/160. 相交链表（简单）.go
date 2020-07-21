package main

import "leetcode-answer/Types"

//不方便构造初始化的链表，就不用main函数了
 //双指针遍历，末尾交换指针，利用步数的差值来找相交点位置
 //注意不相交的情况，就是p1，p2同时为空的时候
func getIntersectionNode(headA, headB *Types.ListNode) *Types.ListNode {
	if headA==nil||headB==nil{
		return nil
	}
	p1,p2:=headA,headB
	for p1!=p2{
		if p1==nil{
			p1=headB
		}else{
			p1=p1.Next
		}
		if p2==nil{
			p2=headA
		}else{
			p2=p2.Next
		}
	}
	return p1
}