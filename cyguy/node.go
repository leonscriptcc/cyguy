package cyguy

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Node 创建节点
func (c *CyGuy) Node(name string, label ...string) *Node {
	return &Node{name: name, label: strings.Join(label, ":")}

}

// Node 节点
type Node struct {
	id         int64
	name       string
	label      string
	properties string
	err        error
}

// SetProperties 设置节点属性
func (n *Node) SetProperties(obj any) *Node {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		n.err = errors.New("properties is not Struct")
		return n
	}

	v := reflect.ValueOf(obj)
	buf := bytes.NewBufferString(`{`)

	var (
		tag       string
		kind      reflect.Kind
		filedName string
	)

	for k := 0; k < t.NumField(); k++ {
		// 获取标签名
		tag = t.Field(k).Tag.Get("cypher")
		// 设置key
		if tag == "" {
			tag = t.Field(k).Name
		}
		buf.WriteString(tag)
		buf.WriteString(":")

		// 提取字段名称、类型
		filedName = t.Field(k).Name
		kind = v.FieldByName(t.Field(k).Name).Kind()

		// 获取字段的类型，根据不同的类型配置不同的样式
		if kind >= reflect.Int && kind <= reflect.Uint64 {
			buf.WriteString(fmt.Sprintf("%d", v.FieldByName(filedName).Int()))
		} else if kind >= reflect.Float32 && kind <= reflect.Float64 {
			buf.WriteString(fmt.Sprintf("%f", v.FieldByName(filedName).Float()))
		} else if kind == reflect.String {
			buf.WriteString(`"`)
			buf.WriteString(v.FieldByName(filedName).String())
			buf.WriteString(`"`)
		} else {
			n.err = errors.New("illegal filed kind:" + filedName)
			return n
		}

		if k != t.NumField()-1 {
			buf.WriteString(`,`)
		}
	}
	buf.WriteString(`}`)
	n.properties = buf.String()
	return n
}

// Create 创建节点
func (n *Node) Create() (result string, err error) {
	if n.err != nil {
		return result, err
	}

	// 没有属性直接返回
	if n.properties == "" {
		return fmt.Sprintf(`%s(%s:%s)`, CREATE, n.name, n.label), err
	}

	// 有属性就拼接上属性
	return fmt.Sprintf(`%s(%s:%s%s)`, CREATE, n.name, n.label, n.properties), err
}

// Delete 删除节点
func (n *Node) Delete() (result string) {
	return result
}

// Update 更新节点
func (n *Node) Update() (result string) {
	return result
}
