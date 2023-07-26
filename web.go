package fun

import "net/url"

/*
	【名称:】获取URL主域名
	【参数:】
	【返回:】*time.Location
	【备注:】
*/
func UrlDomain(link string) (string) {
	if link == "" {
		return ""
	}
	r, err := url.Parse(link)
	if err != nil {
		return ""
	}

	domain := r.Hostname()
	if domain != "" {
		domain = StrToLower(domain)
		domain = StrReplace(domain, "www.", "")
	}

	return domain
}