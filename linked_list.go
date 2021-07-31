package main

import (
	"fmt"
)

type studentNode struct {
	no   int
	name string
	age  int
	pre  *studentNode
	next *studentNode
}

//给双向链表插入一个结点
//编写第一种插入方法，在双向链表的最后加入
func insertStudentNode(head *studentNode, newStudentNode *studentNode) {
	//思路：
	//1.先找到该链表的最后这个结点
	//2.创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//3.将newStudentNode加入到链表的最后
	temp.next = newStudentNode
	newStudentNode.pre = temp

}

//第二种插入方法，根据no的编号从大到小插入
func sortInsertStudentNode(head *studentNode, newStudentNode *studentNode) {
	//思路
	//1.找到适当的位置
	//2.创建一个辅助结点
	temp := head
	flag := true
	//把newStudentNode.no和temp.next.no作比较
	for {
		if temp.next == nil { //说明已经到链表的最后啦
			break
		} else if temp.next.no > newStudentNode.no {
			//说明newStudentNode就应该插入到temp的后面
			break
		} else if temp.next.no == newStudentNode.no {
			//说明链表中已经有这个no了
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在no=", newStudentNode)
	} else {
		newStudentNode.next = temp.next
		temp.next = newStudentNode
		if temp.next != nil {
			temp.next.pre = newStudentNode
		}
		temp.next = newStudentNode
	}
}

//删除一个结点
func delStudentNode(head *studentNode, id int) {
	temp := head
	flag := false
	//找到要删除的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil { //说明已经到链表的最后了
			break
		} else if temp.next.no == id {
			//说明我们找到这个要删除的结点了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}

		//要删除这个结点即直接略过这个结点，
		//然后这个被略过的结点会变成垃圾结点，会被链表删除
	} else {
		fmt.Println("要删除的id不存在")
	}
}
func modifyStudentNode(head *studentNode, id int, newStudentNode *studentNode) {
	temp := head
	flag := false
	//思路：
	//1.找到要修改的结点的no，和temp.next.no作比较
	for {
		if temp.next == nil { //说明已经到了链表的最后
			break
		} else if temp.next.no == id {
			//说明我们已经找到这个要修改的结点了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = newStudentNode
	} else {
		fmt.Println("要修改的结点不存在")
	}
}

//查
func SearchStudentNode(head *studentNode, id int) {
	temp := head
	flag := false
	//思路：
	//1.找到要查找的结点的no，和temp.next.no作比较
	for {
		if temp.next == nil { //说明已经到了链表的最后
			break
		} else if temp.next.no == id {
			//说明我们已经找到这个要查找的结点了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		fmt.Printf("[%d,%s,%d]\n", temp.next.no,
			temp.next.name, temp.next.age)
	} else {
		fmt.Println("要查找的结点不存在")
	}
}

//显示链表的所有结点信息//迭代器
func listStudentNode(head *studentNode) {
	//1.创建一个辅助结点
	temp := head
	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("这是一个空链表")
		return
	}
	//2.遍历这个链表
	for {
		fmt.Printf("[%d,%s,%d]->", temp.next.no,
			temp.next.name, temp.next.age)
		//判断是否是链表的末尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}

}

func main() {
	//1.先创建一个头结点
	head := &studentNode{}
	//2.创建一个新的studentNode
	stuLisa := &studentNode{
		no:   1,
		name: "Lisa",
		age:  24,
	}
	stuBob := &studentNode{
		no:   2,
		name: "Bob",
		age:  25,
	}
	stuNick := &studentNode{
		no:   3,
		name: "Nick",
		age:  27,
	}
	stuMark := &studentNode{
		no:   4,
		name: "Mark",
		age:  29,
	}
	stuMarket := &studentNode{
		no:   3,
		name: "Market",
		age:  30,
	}
	//3.加入结点
	insertStudentNode(head, stuLisa)
	insertStudentNode(head, stuBob)

	//显示链表
	listStudentNode(head)
	fmt.Println()
	//4.加入结点（第二种方法）
	sortInsertStudentNode(head, stuMark) //no是4
	sortInsertStudentNode(head, stuNick) //no是3
	//显示链表
	listStudentNode(head)
	fmt.Println()
	//5.删除结点
	delStudentNode(head, 4)
	//显示链表
	listStudentNode(head)
	fmt.Println()
	//6.修改结点
	modifyStudentNode(head, 3, stuMarket)
	//显示链表
	listStudentNode(head)
	fmt.Println()
	//查找结点
	SearchStudentNode(head, 3)
	//迭代
	listStudentNode(head)
}
