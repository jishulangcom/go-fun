package fun

import (
	"encoding/json"
	"errors"
	"strconv"
)

/*
【名称:】Interface转换string
【参数:】val(interface)
【返回:】string
【备注:】
*/
func InterfaceToStr(val interface{}) string {
	var str string

	if val == nil {
		return str
	}

	switch val.(type) {
	case float64:
		ft := val.(float64)
		str = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := val.(float32)
		str = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := val.(int)
		str = strconv.Itoa(it)
	case uint:
		it := val.(uint)
		str = strconv.Itoa(int(it))
	case int8:
		it := val.(int8)
		str = strconv.Itoa(int(it))
	case uint8:
		it := val.(uint8)
		str = strconv.Itoa(int(it))
	case int16:
		it := val.(int16)
		str = strconv.Itoa(int(it))
	case uint16:
		it := val.(uint16)
		str = strconv.Itoa(int(it))
	case int32:
		it := val.(int32)
		str = strconv.Itoa(int(it))
	case uint32:
		it := val.(uint32)
		str = strconv.Itoa(int(it))
	case int64:
		it := val.(int64)
		str = strconv.FormatInt(it, 10)
	case uint64:
		it := val.(uint64)
		str = strconv.FormatUint(it, 10)
	case string:
		str = val.(string)
	case []byte:
		str = string(val.([]byte))
	default:
		newValue, _ := json.Marshal(val)
		str = string(newValue)
	}

	return str
}

/*
【名称:】interface{}转float64
【参数:】val(interface)
【返回:】float64, error
【备注:】
*/
func InterfaceToFloat64(val interface{}) (f64 float64, err error) {
	switch value := val.(type) {
	case float64:
		return value, nil
	case float32:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case uint64:
		return float64(value), nil
	case uint:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(value), nil
		}
		return 0, errors.New("uint转换为float64失败")
	case int8:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int64:
		return float64(value), nil
	case int:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(value), nil
		}
		return 0, errors.New("int转换为float64失败")
	case bool:
		if value {
			f64 = 1
		}
		return
	case string:
		f64, err = strconv.ParseFloat(value, 64)
		if err == nil {
			return
		}
		return 0, errors.New("字符串转换为float64失败")
	default:
		return 0, errors.New("转换为float64失败")
	}
}

/*
	【名称:】interface{}转int64
	【参数:】val(interface)
	【返回:】int64, error
	【备注:】
*/
func InterfaceToInt64(val interface{}) (int64, error) {
	if IsDouble(val) {
		f64, err := InterfaceToFloat64(val)
		if err != nil {
			return 0, err
		}
		i64 := Float64ToInt64(f64)
		return i64, nil
	}

	if IsNumeric(val) {
		i64, ok := val.(int64)
		if !ok {
			return 0, errors.New("to int64 fail")
		}
		return i64, nil
	}

	return 0, errors.New("not numeric type")
}
