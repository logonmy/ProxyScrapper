package proxy

import "time"
/*
type Proxy struct {
	IP        string
	Port      string
	typeProxy string
}
*/
//URL type contains url to proxy website
type URL string

//Regex type contains regex for different proxy websites
type Regex string

const (
	/*
		Proxy URLs
	*/

	//ProxyListURL url to proxy-list.org site
	ProxyListURL URL = "http://proxy-list.org/english/index.php?p=1"
	//WebanetLabsURL url to webanetlabs.net site
	WebanetLabsURL URL = "https://webanetlabs.net"
	ProxyListMeURL URL = "https://proxylist.me/"

	/*
		Regex
	*/
	IPregex     Regex = `(?P<ip>(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))`
	IPPortregex Regex = `(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|(?P<port>\d{2,5})))`

	ProxyListRegex        Regex = `Proxy\(([^\)]*)\)`
	FreeProxyRegex        Regex = `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})&lt;/td&gt;&lt;td&gt;(\s+|\s*\s*)(\d{2,5})`
	WebanetLabsURLsRegex  Regex = `href\s*=\s*['"]([^'"]*proxylist_at_[^'"]*)['"]`
	CheckProxyURLsRegex   Regex = `href\s*=\s*['"](/archive/\d{4}-\d{2}-\d{2})['"]`
	r                     Regex = `[0-9]+(?:\.[0-9]+){3}:[0-9]+`
	ProxyListMeURLRegex   Regex = `href\s*=\s*['"][^'"]*/?page=(\d+)['"]`
	NumberExtractRegex    Regex = `[-]?\d[\d,]*[\.]?[\d{2}]*`
	ProxyListMeProxyRegex Regex = `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})</a></td>\n[ \t]+\n[ \t]+<td >(\s+|\s*\s*)(\d{2,5})`

	ProxyListPlusComReg Regex =  `(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|(?P<port>\d{2,5}))<)`


DegitRegex Regex = "[0-9]+"

	z = `eval\(unescape\(([^\)]*)\)`
)

//CheckProxyStruct struct to contain the json from check-proxy website
type CheckProxyStruct struct {
	ID             int       `json:"id"`
	LocalID        int       `json:"local_id"`
	ReportID       string    `json:"report_id"`
	Addr           string    `json:"addr"`
	Type           int       `json:"type"`
	Kind           int       `json:"kind"`
	Timeout        int       `json:"timeout"`
	Cookie         bool      `json:"cookie"`
	Referer        bool      `json:"referer"`
	Post           bool      `json:"post"`
	IP             string    `json:"ip"`
	AddrGeoIso     string    `json:"addr_geo_iso"`
	AddrGeoCountry string    `json:"addr_geo_country"`
	AddrGeoCity    string    `json:"addr_geo_city"`
	IPGeoIso       string    `json:"ip_geo_iso"`
	IPGeoCountry   string    `json:"ip_geo_country"`
	IPGeoCity      string    `json:"ip_geo_city"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Skip           bool      `json:"skip"`
	FromCache      bool      `json:"from_cache"`
}

