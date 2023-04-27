package cyguy

import (
	"fmt"
)

// Node 创建节点
func (c *CypherGuy) Node(name, labels string) *Node {
	//log.Println(results[1].Len())
	node := Node{
		name:  name,
		label: labels,
		info:  fmt.Sprintf("%s:%s", name, labels),
	}
	return &node

}

// Node 节点
type Node struct {
	obj        any
	name       string
	label      string
	properties string
	info       string
	err        error
}

// Properties 设置节点属性
func (n *Node) Properties(obj any) *Node {
	n.properties, n.err = getProperties(obj)
	n.info = fmt.Sprintf("%s%s", n.info, n.properties)
	return n
}

// Create 创建节点
func (n *Node) Create() (result string, err error) {
	if n.err != nil {
		return result, err
	}

	return fmt.Sprintf(`%s(%s) %s %s`, CREATE, n.info, RETURN, n.name), err

}

// Delete 删除节点
func (n *Node) Delete() (result string, err error) {
	if n.err != nil {
		return result, err
	}

	return fmt.Sprintf(`%s(%s) %s %s`, MATCH, n.info, DELETE, n.name), err
}

// DetachDelete 删除节点以及节点的关系
func (n *Node) DetachDelete() (result string, err error) {
	if n.err != nil {
		return result, err
	}

	return fmt.Sprintf(`%s(%s) %s %s %s`, MATCH, n.info, DELETE, DETACH, n.name), nil
}

// SetLabels 更新标签
func (n *Node) SetLabels(labels string) (result string, err error) {
	if n.err != nil {
		return result, err
	}

	return fmt.Sprintf(`%s(%s) %s %s:%s %s %s:%s`,
		MATCH, n.info, REMOVE, n.name, n.label, SET, n.name, labels), err
}

// SetProperties 更新属性
func (n *Node) SetProperties(obj any) (result string, err error) {
	if n.err != nil {
		return result, err
	}
	// 获取等待更新的obj
	ps, err := getProperties(obj)
	if err != nil {
		return result, err
	}
	return fmt.Sprintf(`%s(%s) %s %s=%s`, MATCH, n.info, SET, n.name, ps), err
}
