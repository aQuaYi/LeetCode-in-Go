package problem0707

// MyLinkedList object will be instantiated and called as such:
// obj := Constructor();
// param_1 := obj.Get(index);
// obj.AddAtHead(val);
// obj.AddAtTail(val);
// obj.AddAtIndex(index,val);
// obj.DeleteAtIndex(index);
type MyLinkedList struct {
}

// Constructor initialize your data structure here. */
func Constructor() MyLinkedList {

	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {

	return 0
}

// AddAtHead a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (l *MyLinkedList) AddAtHead(val int) {

}

// AddAtTail append a node of value val to the last element of the linked list.
func (l *MyLinkedList) AddAtTail(val int) {

}

// AddAtIndex add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (l *MyLinkedList) AddAtIndex(index int, val int) {

}

// DeleteAtIndex delete the index-th node in the linked list, if the index is valid.
func (l *MyLinkedList) DeleteAtIndex(index int) {

}
