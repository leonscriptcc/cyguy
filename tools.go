package cyguy

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

// getProperties 解析属性
func getProperties(obj any) (string, error) {
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
