package cyguy

import (
	"errors"
	"fmt"
)

// Matcher 创建关系
func (c *CypherGuy) Matcher() *Matcher {
	return &Matcher{}
}

// Matcher 查询器
type Matcher struct {
	result []byte
	err    error
}

// Nodes 根据node的属性查询节点
func (m *Matcher) Nodes(node *Node) {
	// 生成查询语句
	m.result = []byte(fmt.Sprintf("%s(n:%s %s) %s n",
		MATCH, node.label, node.properties, RETURN))
}

// MultiJumps 多跳查询-设置节点
// MATCH (n)-[:rel3]->(m) RETURN n
func (m *Matcher) MultiJumps(fromNode *Node, relationship Relationship, toNode *Node) *Matcher {
	if fromNode == nil && toNode == nil {
		m.err = errors.New("")
		return m
	}

	if fromNode == nil {

	} else if toNode == nil {

	}
	return m
}

// Find 生成查询语句
func (m *Matcher) Find() string {
	return string(m.result)
}
