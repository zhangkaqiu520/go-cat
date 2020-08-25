package laser

import (
	"errors"
	"log"
	"strings"
)

type Node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang 只有最后完整的才有
	part     string  // 路由中的一部分，例如 :lang
	children []*Node // 子节点，例如 [doc, tutorial, intro]
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Insert(pattern string) error {
	if pattern[0] != '/' {
		return errors.New("must start with /")
	}

	parts := strings.Split(pattern[1:], "/")
	n.matchChild(pattern, parts)

	return nil
}

func (n Node) Search(pattern string) error {
	if pattern[0] != '/' {
		return errors.New("must start with /")
	}
	parts := strings.Split(pattern[1:], "/")
	res := n.matchChildren(parts)
	if res != nil {
		log.Println(res)
	}

	return nil
}

func (n Node) matchChildren(parts []string) *Node {
	if len(parts) == 0 {
		return nil
	}
	for _, m := range n.children {
		if m.part == parts[0] {
			if m.pattern != "" && len(parts) == 1 {
				return m
			}
			return m.matchChildren(parts[1:])
		}
	}

	return nil
}

func (n *Node) matchChild(pattern string, parts []string) {
	if len(parts) == 0 {
		return
	}
	for _, m := range n.children {
		if m.part == parts[0] {
			m.matchChild(pattern, parts[1:])
			return
		}
	}

	p := ""
	if len(parts) == 1 {
		p = pattern
	}
	nn := Node{
		p, parts[0], []*Node{},
	}

	n.children = append(n.children, &nn)
	nn.matchChild(pattern, parts[1:])
}

func (n Node) Fetch() {
	for _, v := range n.children {
		log.Println(v.pattern)
		v.Fetch()
	}
}
