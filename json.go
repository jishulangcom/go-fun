package fun

import (
	"bytes"
	"encoding/json"
)

/*
	【名称:】json转map[string]interface{}
	【参数:】b([]byte)
	【返回:】map[string]interface{}, error
	【备注:】解决JSON格式的数据反序列化为map[string]interface{}时，数字变成科学计数法表示的浮点数问题
*/
func JsonToMapStringInterface(b []byte) (map[string]interface{}, error) {
	var err error

	dataMap := make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&dataMap)
	if err != nil {
		return nil, err
	}

	return dataMap, nil
}