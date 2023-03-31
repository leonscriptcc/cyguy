package cyguy

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

// Node 创建节点
func (c *CyGuy) Node(name, labels string) *Node {
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
	obj   any
	name  string
	label string
	info  string
	err   error
}

// Properties 设置节点属性
func (n *Node) Properties(obj any) *Node {
	var properties string
	properties, n.err = n.getProperties(obj)
	n.info = fmt.Sprintf("%s%s", n.info, properties)
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
	ps, err := n.getProperties(obj)
	if err != nil {
		return result, err
	}
	return fmt.Sprintf(`%s(%s) %s %s=%s`, MATCH, n.info, SET, n.name, ps), err
}

// getProperties 解析属性
func (n *Node) getProperties(obj any) (string, error) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return "", errors.New("properties is not Struct")
	}

	buf := bytes.NewBufferString(`{`)

	var (
		tag  string
		kind reflect.Kind
	)

	for k := 0; k < t.NumField(); k++ {
		// 获取标签名
		tag = t.Field(k).Tag.Get("cypher")
		// 忽略的字段
		if tag == "-" {
			continue
		}
		// 设置key
		if tag == "" {
			tag = t.Field(k).Name
		}
		buf.WriteString(tag)
		buf.WriteString(":")

		// 提取字段名称、类型
		kind = v.Field(k).Kind()

		// 获取字段的类型，根据不同的类型配置不同的样式
		if kind >= reflect.Int && kind <= reflect.Uint64 {
			buf.WriteString(fmt.Sprintf("%d", v.Field(k).Int()))
		} else if kind >= reflect.Float32 && kind <= reflect.Float64 {
			buf.WriteString(fmt.Sprintf("%f", v.Field(k).Float()))
		} else if kind == reflect.String {
			buf.WriteString(`"`)
			buf.WriteString(v.Field(k).String())
			buf.WriteString(`"`)
		} else {
			return "", errors.New("illegal filed kind:" + t.Field(k).Name)
		}

		if k != t.NumField()-1 {
			buf.WriteString(`,`)
		}
	}
	buf.WriteString(`}`)
	return buf.String(), nil
}
