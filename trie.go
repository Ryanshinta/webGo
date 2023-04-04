package webGo

import (
	"strings"
)

/*

parameter match: /p/:user/detail  can match->/p/userA/detail  ->/p/userB/detail

wildcard match : /static/*  match all the file/path under the path.

*/

type node struct {
	pattern  string  // routing path for the node
	part     string  // current routing section for node
	children []*node //list of child nodes
	isWild   bool    // is
}

func (n *node) matchSingeChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchAllChildren(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// insert a new routing path into node
// if current node doesn't match the current routing, create a cew child node
// insert the remainging routing into child node recursively
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchSingeChild(part)

	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchAllChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
