package cyguy

import (
	"errors"
	"fmt"
)

// Relationship 创建关系
func (c *CypherGuy) Relationship(name, labels string) *Relationship {
	relationship := Relationship{
		name:  name,
		label: labels,
		info:  fmt.Sprintf("%s:%s", name, labels),
	}
	return &relationship
}

// Relationship 关系
type Relationship struct {
	obj      any
	name     string
	label    string
	fromNode *Node
	toNode   *Node
	info     string
	err      error
}

// Properties 设置节点属性
func (r *Relationship) Properties(obj any) *Relationship {
	var properties string
	properties, r.err = getProperties(obj)
	r.info = fmt.Sprintf("%s%s", r.info, properties)
	return r
}

// From 从某个节点来
func (r *Relationship) From(node *Node) *Relationship {
	r.fromNode = node
	return r
}

// To 到某个节点去
func (r *Relationship) To(node *Node) *Relationship {
	r.toNode = node
	return r
}

// Create 已经存在节点的关系创建语句
func (r *Relationship) Create() (result string, err error) {
	if r.err != nil {
		return result, err
	}

	// 判断节点信息
	if r.fromNode == nil || r.toNode == nil {
		err = errors.New("missing node information")
		return result, err
	}

	// 组装语句
	result = fmt.Sprintf("%s(%s),(%s) %s (%s)-[%s]->(%s) %s %s",
		MATCH, r.fromNode.info, r.toNode.info, CREATE, r.fromNode.name, r.info, r.toNode.name, RETURN, r.name)
	return result, r.err
}

// Delete 删除关系
func (r *Relationship) Delete() (result string, err error) {
	if r.err != nil {
		return result, err
	}

	// 判断节点信息
	if r.fromNode == nil || r.toNode == nil {
		err = errors.New("missing node information")
		return result, err
	}

	// 组装语句
	result = fmt.Sprintf("%s (%s)-[%s]-%s %s %s",
		MATCH, r.fromNode.info, r.info, r.toNode.info, DELETE, r.name)

	return result, r.err
}

// SetLabels 更新标签
func (r *Relationship) SetLabels(labels string) (result string, err error) {
	if r.err != nil {
		return result, err
	}

	// 判断节点信息
	if r.fromNode == nil || r.toNode == nil {
		err = errors.New("missing node information")
		return result, err
	}

	// 组装语句
	result = fmt.Sprintf("%s (%s)-[%s]-%s %s %s",
		MATCH, r.fromNode.info, r.info, r.toNode.info, DELETE, r.name)

	return result, r.err
}

// SetProperties 更新属性
func (r *Relationship) SetProperties(obj any) (result string, err error) {
	if r.err != nil {
		return result, err
	}

	// 判断节点信息
	if r.fromNode == nil || r.toNode == nil {
		err = errors.New("missing node information")
		return result, err
	}

	// 获取等待更新的obj
	ps, err := getProperties(obj)

	// 组装语句
	result = fmt.Sprintf("%s (%s)-[%s]-%s %s %s=%s",
		MATCH, r.fromNode.info, r.info, r.toNode.info, SET, r.name, ps)

	return result, r.err
}
