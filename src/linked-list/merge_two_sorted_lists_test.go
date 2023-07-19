package linked_list

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	// Test case 1: Empty list
	head := mergeTwoLists(nil, nil)
	if head != nil {
		t.Errorf("Expected nil, got %v", head)
	}

	// Test case 2: Single node list
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}

	expected := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	head = mergeTwoLists(node1, node2)
	if !reflect.DeepEqual(head, expected) {
		t.Errorf("Expected %v, got %v", expected, head)
	}

	//// Test case 3: Multiple node list
	//node3 := &ListNode{Val: 3, Next: nil}
	//node2 := &ListNode{Val: 2, Next: node3}
	//node1 := &ListNode{Val: 1, Next: node2}
	//
	//expected := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}
	//head = reverseList(node1)
	//if !reflect.DeepEqual(head, expected) {
	//	t.Errorf("Expected %v, got %v", expected, head)
	//}
}
