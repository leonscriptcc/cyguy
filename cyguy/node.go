package cyguy

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

// Node 创建节点
func (c *CyGuy) Node(obj any) *Node {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	method := v.MethodByName("NodeInfo")
	if !method.IsValid() {
		return &Node{err: errors.New("unable to get method：func NodeInfo()(string,string)")}
	}
	results := method.Call([]reflect.Value{})
	//log.Println(results[1].Len())
	node := Node{name: results[0].String(), label: results[1].String(), v: v, t: t}
	return node.setProperties(node)

}

// Node 节点
type Node struct {
	obj        any
	t          reflect.Type
	v          reflect.Value
	name       string
	label      string
	properties string
	err        error
}

// setProperties 设置节点属性
func (n *Node) setProperties(obj any) *Node {
	if n.t.Kind() != reflect.Struct {
		n.err = errors.New("properties is not Struct")
		return n
	}

	buf := bytes.NewBufferString(`{`)

	var (
		tag       string
		kind      reflect.Kind
		filedName string
	)

	for k := 0; k < n.t.NumField(); k++ {
		// 获取标签名
		tag = n.t.Field(k).Tag.Get("cypher")
		// 设置key
		if tag == "" {
			tag = n.t.Field(k).Name
		}
		buf.WriteString(tag)
		buf.WriteString(":")

		// 提取字段名称、类型
		filedName = n.t.Field(k).Name
		kind = n.v.FieldByName(filedName).Kind()

		// 获取字段的类型，根据不同的类型配置不同的样式
		if kind >= reflect.Int && kind <= reflect.Uint64 {
			buf.WriteString(fmt.Sprintf("%d", n.v.FieldByName(filedName).Int()))
		} else if kind >= reflect.Float32 && kind <= reflect.Float64 {
			buf.WriteString(fmt.Sprintf("%f", n.v.FieldByName(filedName).Float()))
		} else if kind == reflect.String {
			buf.WriteString(`"`)
			buf.WriteString(n.v.FieldByName(filedName).String())
			buf.WriteString(`"`)
		} else {
			n.err = errors.New("illegal filed kind:" + filedName)
			return n
		}

		if k != n.t.NumField()-1 {
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
		return fmt.Sprintf(`%s(%s:%s) RETURN %s`, CREATE, n.name, n.label, n.name), err
	}

	// 有属性就拼接上属性
	return fmt.Sprintf(`%s(%s:%s%s) RETURN %s`, CREATE, n.name, n.label, n.properties, n.name), err
}

// Delete 删除节点
func (n *Node) Delete() (result string) {
	return result
}

// DetachDelete 删除节点以及节点的关系
func (n *Node) DetachDelete() (result string) {
	return result
}

// Update 更新节点
func (n *Node) Update() (result string) {
	return result
}
