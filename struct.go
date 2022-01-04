package fun

import "reflect"

/*
	【名称:】返回空对象
	【参数:】
	【返回:】{}
	【备注:】
*/
func EmptyStruct() interface{} {
	return EmptyStt{}
}

type EmptyStt struct{}

/*
	【名称:】结构体转换为数组
	【参数:】
	【返回:】
	【备注:】 遍历结构体,提取key和val
*/
func StructToArr(data interface{}) []StructInfoStt {
	arr := make([]StructInfoStt, 0)
	t := reflect.TypeOf(data)
	for k := 0; k < t.NumField(); k++ {
		structField := t.Field(k)
		tag := structField.Tag
		item := StructInfoStt{
			Field:    structField.Name,
			DataType: structField.Type.String(),
			Form:     tag.Get("form"),
			Json:     tag.Get("json"),
		}
		arr = append(arr, item)
	}
	return arr
}

type StructInfoStt struct {
	Field    string //字段
	DataType string //数据类型
	Form     string //form
	Json     string //json
}
