package binarytree

import (
	"reflect"
	"testing"
)

type testCase struct {
	tree   *Node
	dfs    []interface{}
	bfs    []interface{}
	size   int
	height int
}

var testCases = map[string]testCase{
	"nil tree": testCase{
		tree:   nil,
		dfs:    nil,
		bfs:    nil,
		size:   0,
		height: 0,
	},
	"only empty tree": testCase{
		tree:   &Node{},
		dfs:    []interface{}{nil},
		bfs:    []interface{}{nil},
		size:   1,
		height: 1,
	},
	"only head": testCase{
		tree:   &Node{Data: 1},
		dfs:    []interface{}{1},
		bfs:    []interface{}{1},
		size:   1,
		height: 1,
	},
	"triangle tree": testCase{
		tree:   &Node{Data: 1, Left: &Node{Data: 2}, Right: &Node{Data: 3}},
		dfs:    []interface{}{1, 2, 3},
		bfs:    []interface{}{1, 2, 3},
		size:   3,
		height: 2,
	},
	"line tree": testCase{
		tree: &Node{
			Data: 1,
			Left: &Node{
				Data: 2,
				Left: &Node{
					Data: 3,
				},
			},
		},
		dfs:    []interface{}{1, 2, 3},
		bfs:    []interface{}{1, 2, 3},
		size:   3,
		height: 3,
	},
	"complex tree": testCase{

		tree: &Node{
			Data: 1,
			Left: &Node{
				Data:  2,
				Left:  &Node{Data: 4},
				Right: &Node{Data: 5},
			},
			Right: &Node{Data: 3},
		},
		dfs:    []interface{}{1, 2, 4, 5, 3},
		bfs:    []interface{}{1, 2, 3, 4, 5},
		size:   5,
		height: 3,
	},
}

func TestDFS(t *testing.T) {
	t.Parallel()
	for name, tc := range testCases {
		got := tc.tree.DFS()
		if !reflect.DeepEqual(tc.dfs, got) {
			t.Errorf("wrong dfs for %s. exp: %#v, got: %#v", name, tc.dfs, got)
		}
	}
}

func TestBFS(t *testing.T) {
	t.Parallel()
	for name, tc := range testCases {
		got := tc.tree.BFS()
		if !reflect.DeepEqual(tc.bfs, got) {
			t.Errorf("wrong bfs for %s. exp: %#v, got: %#v", name, tc.bfs, got)
		}
	}
}

func TestSize(t *testing.T) {
	t.Parallel()
	for name, tc := range testCases {
		got := tc.tree.Size()
		if tc.size != got {
			t.Errorf("wrong size for %s. exp: %#v, got: %#v", name, tc.size, got)
		}
	}
}

func TestHeight(t *testing.T) {
	t.Parallel()
	for name, tc := range testCases {
		got := tc.tree.Height()
		if tc.height != got {
			t.Errorf("wrong height for %s. exp: %#v, got: %#v", name, tc.height, got)
		}
	}
}
