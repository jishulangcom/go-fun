/*
	【包名:】服务器端相关函数
	【作者:】技术狼
*/
package fun

import "net"


/*
	【名称:】获取服务器端IP
	【参数:】
	【返回:】ip地址(string)
	【备注:】
*/
func Ip_serve() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String()
	}
	return ""

}
