package main
import "fmt"

type node struct {
	num int
	next *node
	prev *node
}

type myList struct {
	length int
	head *node
	tail *node
	current *node
}
func (l *myList) next() *node {
	if l.current == nil {
		l.current=l.head
		return l.current
	} else {
		l.current=l.current.next
		return l.current
	}
}

func (l *myList) push(n *node) {
	if l.head == nil {
		l.length = 1
		l.head=n
	} else {
		n.next=l.head
		l.head.prev=n
		n.prev=nil
		l.head=n
		l.length=l.length + 1
	}
}
func (l *myList) pop() *node {
	var n *node
	n = l.head
	l.head = l.head.next
	l.length = l.length - 1
	return n
}


func main() {
	var i int =0
	var list myList=myList{length:0,current:nil}

	for i=1;i<=10;i++ {
		tmp:=node{num:i,prev:nil,next:nil}
		list.push(&tmp)
	}

	fmt.Printf("Length of list :: %d\n",list.length)
	fmt.Println("Values in list ::")
	for list.next() != nil {
		fmt.Println(list.current.num)
	}
}
