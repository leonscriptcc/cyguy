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
	result   []byte
	err      error
	fromNode *Node
	toNode   *Node
	r        *Relationship
}

// Nodes 根据node的属性查询节点
func (m *Matcher) Nodes(node *Node) {
	// 生成查询语句
	m.result = []byte(fmt.Sprintf("%s(n:%s %s) %s n",
		MATCH, node.label, node.properties, RETURN))
}

// Find 生成查询语句
func (m *Matcher) Find() string {
	return string(m.result)
}

func (m *Matcher) FromNode(node *Node) *Matcher {
	m.fromNode = node
	return m
}

func (m *Matcher) ToNode(relationship *Relationship, node *Node) *Matcher {
	m.r = relationship
	m.toNode = node
	return m
}

// MultiJumps 多跳查询-设置节点
// MATCH (n)-[:rel3*1..]->(m:label) RETURN n
func (m *Matcher) MultiJumps() (result string, err error) {
	// 关系不能为空
	if m.r == nil {
		err = errors.New("")
		return "", err
	}

	// 不能同时为空
	if m.fromNode == nil && m.toNode == nil {
		err = errors.New("")
		return "", err
	}

	// 组装查询语句
	if m.fromNode == nil {
		result = fmt.Sprintf("%s (n)-[r:%s*1..]->(m:%s) %s n,m",
			MATCH, m.r.label, m.toNode.label, RETURN)
	} else if m.toNode == nil {
		result = fmt.Sprintf("%s (n:%s)-[r:%s*1..]->(m) %s n,m",
			MATCH, m.r.fromNode.label, m.r.label, RETURN)
	}

	return result, nil

}
