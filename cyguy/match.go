package cyguy

import "errors"

// Matcher 创建关系
func (c *CyGuy) Matcher() *Matcher {
	return &Matcher{}
}

// Matcher 查询器
type Matcher struct {
	result []byte
	err    error
}

// Nodes 根据node的属性查询节点
func (m *Matcher) Nodes(node *Node) {
	// 搜索的话节点不会有name
	if node.name != "" {
		m.err = errors.New("")
		return
	}

}

// Node 多跳查询-设置节点
func (m *Matcher) Node(node *Node) {

}

// To 多跳查询-设置关系
func (m *Matcher) To(relationship *Relationship) {

}

// What 多跳查询-设置查询项
func (m *Matcher) What() {

}

// Find 生成查询语句
func (m *Matcher) Find() {

}
