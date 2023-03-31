package cyguy

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

// Relationship 创建关系
func (c *CyGuy) Relationship(name, labels string) *Relationship {
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
	fromNode string
	toNode   string
	info     string
	err      error
}

// Properties 设置节点属性
func (r *Relationship) Properties(obj any) *Relationship {
	var properties string
	properties, r.err = r.getProperties(obj)
	r.info = fmt.Sprintf("%s%s", r.info, properties)
	return r
}

// From 从某个节点来
func (r *Relationship) From(node *Node) *Relationship {

	return r
}

// To 到某个节点去
func (r *Relationship) To(node *Node) *Relationship {

	return r
}

// Create 已经存在节点的关系创建语句
func (r *Relationship) Create() (result string, err error) {
	return result, r.err
}

// getProperties 解析属性
func (r *Relationship) getProperties(obj any) (string, error) {
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
