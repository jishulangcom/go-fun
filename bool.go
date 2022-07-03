package fun

// @title: 变量是否布尔类型
// @param: 变量(interface)
// @return: true/false(bool)
// @description:
// @date: 2022/7/3 21:40
func IsBool(variable interface{}) bool {
	if GetType(variable) == "bool" {
		return true
	}
	return false
}
