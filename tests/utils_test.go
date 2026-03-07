package tests

import (
	"testing"
	"github.com/abdessamad-zgor/dama/utils"
)

func TestList(t *testing.T) {
	list := utils.NewList[int]()

	if list.Length() != 0 {
		t.Errorf("expected length to be 0, found %d", list.Length())
	}

	list.Add(1)
	list.Add(2)

	if list.Length() != 2 {
		t.Errorf("expected length to be 2, found %d", list.Length())
	}

	list.Insert(0, 0)
	items := list.Items()

	if items[0] != 0 {
		t.Errorf("expected value 0, got %d", items[0])
	}
}

func TestTree(t *testing.T) {
	tree := utils.NewTree[int](0)

	tree.AddNode(0, 1)
	tree.AddNode(0, 2)

	node := tree.FindNode(2)
	node.SetValue(4)

	nnode := tree.FindNode(4)
	flattened := tree.Flatten()

	if nnode.GetValue() != 4 {
		t.Errorf("expected value 2, found %d", nnode.Value)
	}
	t.Log("flattened tree: ", flattened)
}
