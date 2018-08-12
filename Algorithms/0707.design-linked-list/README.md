# [707. Design Linked List](https://leetcode.com/problems/design-linked-list/)

## 题目

Design yourimplementation of the linked list. You can choose to use the singly linked list or the doubly linked list. A node in a singlylinked list should have two attributes: valand next. val is the value of the current node, and nextisapointer/reference to the next node. If you want to use the doubly linked list,you will needone more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement these functions in your linked list class:

- get(index) : Get the value ofthe index-thnode in the linked list. If the index is invalid, return -1.
- addAtHead(val) : Add a node of value valbefore the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
- addAtTail(val) : Append a node of value valto the last element of the linked list.
- addAtIndex(index, val) : Add a node of value valbefore the index-thnode in the linked list.If indexequalsto the length oflinked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
- deleteAtIndex(index) : Deletethe index-thnode in the linked list, if the index is valid.

Example:

```text
MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
linkedList.get(1);            // returns 2
linkedList.deleteAtIndex(1);  // now the linked list is 1->3
linkedList.get(1);         // returns 3
```

Note:

- All values will be in the range of [1, 1000].
- The number of operations will be in the range of[1, 1000].
- Please do not use the built-in LinkedList library.

## 解题思路

见程序注释
