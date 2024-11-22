// 707. Design Linked List
// https://leetcode.com/problems/design-linked-list/

type Node struct {
	Next *Node
	Val  int
}
type MyLinkedList struct {
	head *Node
	tail *Node
}

func Constructor() MyLinkedList {
	ll := new(MyLinkedList)
	return *ll
}

func (this *MyLinkedList) Get(index int) int {

	curr := this.head
	i := 0
	for curr != nil {
		if i == index {
			return curr.Val
		}
		i++
		curr = curr.Next
	}
	return -1
}

func (this *MyLinkedList) printList() {
    curr := this.head
    for curr != nil {
        fmt.Printf("%d, ", curr.Val)
        curr = curr.Next
    }
    fmt.Println()
}
func (this *MyLinkedList) AddAtHead(val int) {

	node := new(Node)
	node.Val = val
	node.Next = this.head
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
    // this.printList()
}

func (this *MyLinkedList) AddAtTail(val int) {

	node := new(Node)
	node.Val = val
	if this.tail == nil {
		this.tail = node
	} else {
		this.tail.Next = node
		this.tail = node
	}
	if this.head == nil {
		this.head = node
	}
    // this.printList()
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	i := 0
	node := new(Node)
	node.Val = val
	curr := this.head
	var prev *Node

	for curr != nil {
		if i == index {
			if prev != nil {
				prev.Next = node
			}
			node.Next = curr
			if this.head == curr {
				this.head = node
			}
			if this.tail == nil {
				this.tail = node
			}
            // this.printList()
			return
		}
		i++
		prev = curr
		curr = curr.Next
	}
	if i == index {
		if this.tail != nil {
			this.tail.Next = node
		}
		if this.head == nil {
			this.head = node
		}
		this.tail = node
	}
    // this.printList()
}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	var prev *Node
	curr := this.head

	i := 0
	for curr != nil {
		if i == index {
			if prev != nil {
				prev.Next = curr.Next
			}
			if this.head == curr {
				this.head = curr.Next
			}
			if this.tail == curr {
				this.tail = prev
			}

		}
		prev = curr
		curr = curr.Next
		i++
	}
    // this.printList()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
