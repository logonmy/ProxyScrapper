package proxy

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type GetProxy func() []string

type ProxyProvider struct {
	Domain         string
	protocols      []string
	maxConnections int
	maxTries       int
	timeout        int
	GetProxy       GetProxy
	Active			bool
}

type Data struct {
	Type    string `json:"Type"`
	PageIdx int    `json:"PageIdx"`
}

const proxyRegexp = `(?s)(?P<ip>(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|(?P<port>\d{2,5})))`


func ProxyListDownload() []string {
	var res []string
	urls := []string{"https://www.proxy-list.download/api/v1/get?type=http",
		"https://www.proxy-list.download/api/v1/get?type=https",
		"https://www.proxy-list.download/api/v1/get?type=socks4",
		"https://www.proxy-list.download/api/v1/get?type=socks5"}
	res = findOnURLs(urls, proxyRegexp)
	return res
}

func ProxyListPlusCom() []string {
	var res []string
	names := []string{"Fresh-HTTP-Proxy", "SSL", "Socks"}
	url := "http://list.proxylistplus.com/%s-List-%d"
	var urls []string
	for _, name := range names {
		for i := 0; i <= 7; i++ {
			url := fmt.Sprintf(url, name, i)
			urls = append(urls, url)
		}
	}
	prox := `(?s)<td>(?P<ip>(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|(?P<port>\d{2,5})))`

	regex := regexp.MustCompile(string(prox))
	//groupNames := regex.SubexpNames()
	for _, url := range urls {
		page := get(url)
		a := getProxies(page, string(prox))
		//fmt.Println(a)
		_ = a
		for _, p := range regex.FindAllSubmatch(page, -1) {

			proxy := fmt.Sprintf("%s:%s", p[1], p[3])
			res = append(res, proxy)

		}
	}
	return res
}

/*
TODO Website Moved will
*/
func ProxyBNet() []string {
	var res []string
	proxyBUrl := "http://proxyb.net/ajax.php"
	type dataB struct {
		action string
		p      int
		page   string
	}
	data := dataB{
		action: "getProxy",
		p:      5,
		page:   "/anonimnye_proksi_besplatno.html",
	}
	header := "'X-Requested-With': 'XMLHttpRequest'"
	page := post(proxyBUrl, data, header)
	//fmt.Println(string(page))
	_ = page
	return res
}

/*
TODO After captcha
*/
func FreeProxyCz() []string {
	var res []string
	//url := "free-proxy.cz"
	tplURL := "http://free-proxy.cz/en/proxylist/main/date/%d"
	var urls []string
	for i := 0; i <= 15; i++ {
		url := fmt.Sprintf(tplURL, i)
		urls = append(urls, url)
	}
	for _, url := range urls {
		page := get(url)
		_ = page
		//fmt.Println(string(page))
	}
	//fmt.Println(urls)
	return res
}

func MyProxyCom() []string {
	var res []string
	url := "https://www.my-proxy.com/free-proxy-list.html"
	page := get(url)
	regExpStr := `href\s*=\s*['"]([^'"]?free-[^'"]*)['"]`
	regExp := regexp.MustCompile(regExpStr)
	urls := make([]string, 0)
	for _, a := range regExp.FindAllStringSubmatch(string(page), -1) {
		url = fmt.Sprintf("https://www.my-proxy.com/%s", a[1])
		urls = append(urls, url)
	}
	regExp2 := regexp.MustCompile(proxyRegexp)
	for _, url := range urls {
		page := get(url)
		for i, proxy := range regExp2.FindAllStringSubmatch(string(page), -1) {
			if i == 0 {
				continue
			}
			p := fmt.Sprintf("%s:%s", proxy[1], proxy[3])
			res = append(res, p)
		}
	}
	return res
}

//GetProxiesSpysRu TODO
func GetProxiesSpysRu() []string {
	var res []string
	url := "http://spys.one/proxies/"
	expCharToNum := `'([a-z0-9]{32})'`
	expCharToNumReg := regexp.MustCompile(expCharToNum)
	//charTonumMap := make(map[string]string)
	page := get(url)
	type dat struct {
		xf0 string
		xpp int
		xf1 int
	}
	d := expCharToNumReg.FindAllStringSubmatch(string(page), -1)

	dat1 := dat{
		xf0: d[0][0],
		xpp: 3,
		xf1: 4,
	}

	pagePost := post(url, dat1, "")

	_ = pagePost
	//fmt.Println(string(pagePost))

	/*
		for _,i:=range[]int {3,4}{
			dat1 := dat{
				xf0 : d[0][0],
				xpp:3,
				xf1:i,
			}
			pagePost := post(url,dat1)
			//reg := regexp.MustCompile(proxyRegexp)
			fmt.Println(string(pagePost))
			/*for _,j := range  reg.FindAllStringSubmatch(string(pagePost),-1){
				for j,k := range  j{
					fmt.Println(j,k)
				}

		//		fmt.Println(i,j)
			}
		}


	*/

	//pagePost := post(url,dat1)
	//fmt.Println(string(pagePost))
	//e := `[>;]{1}(?P<char>[a-z\d]{4,})=(?P<num>[a-z\d\^]+)`
	/*eo:=regexp.MustCompile(e)
	for i,ch := range  eo.FindAllStringSubmatch(string(pagePost),-1){
	fmt.Println(i,ch)
	}*/

	return res
}

func ProxyNovaProxies() []string {
	var res []string
	var iso []string
	var urls []string
	urlTpl := "https://www.proxynova.com/proxy-server-list/country-%s/"

	url := "https://www.proxynova.com/proxy-server-list/"
	expCountries := `"([a-z]{2})">`
	expCountriesReg := regexp.MustCompile(expCountries)
	page := get(url)
	for _, country := range expCountriesReg.FindAllString(string(page), -1) {
		iso = append(iso, country)
	}
	for _, isoCode := range iso {
		isoCode = strings.TrimSuffix(strings.TrimPrefix(isoCode, "\""), "\">")
		url := fmt.Sprintf(urlTpl, isoCode)
		urls = append(urls, url)
	}
	regExp := regexp.MustCompile(proxyRegexp)
	for _, url := range urls {
		page := get(url)
		for i, proxy := range regExp.FindAllStringSubmatch(string(page), -1) {
			if i == 0 {
				continue
			}
			p := fmt.Sprintf("%s:%s", proxy[1], proxy[3])
			res = append(res, p)
		}
	}
	return res
}

func NntimeComProxies() []string {
	var res []string
	var urls []string
	urlFormat := "http://www.nntime.com/proxy-updated-%02d.htm"
	for i := 1; i < 32; i++ {
		url := fmt.Sprintf(urlFormat, i)
		urls = append(urls, url)
	}

	for _, url := range urls {
		page := get(url)
		expressionPortOnJS := `\b(?P<char>[a-z])=(?P<num>\d);`
		regJs := regexp.MustCompile(expressionPortOnJS)
		portMap := make(map[string]string)
		for _, ch := range regJs.FindAllStringSubmatch(string(page), -1) {
			//fmt.Println(i,j)
			portMap[ch[1]] = ch[2]
		}
		portOnJsExpression := `\(":"\+(?P<chars>[a-z+]+)\)`
		regPortJs := regexp.MustCompile(portOnJsExpression)
		pageCopy := regPortJs.ReplaceAllFunc(page, func(s []byte) []byte {
			z := regPortJs.ReplaceAllString(string(s), "$1")
			var port []byte
			for _, ch := range z {
				a := portMap[string(ch)]
				if ch != '+' {
					port = append(port, []byte(a)...)
				}
			}
			portString := string(port)
			return []byte(portString)
		})
		regProxiesExp := `(?P<ip>(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|(?P<port>\d{2,5})))`
		regExpProxies := regexp.MustCompile(regProxiesExp)
		for i, j := range regExpProxies.FindAllStringSubmatch(string(pageCopy), -1) {
			_ = i
			if strings.Contains(j[0], "script") {
				//fmt.Println(j[1], j[3])
				proxy := fmt.Sprintf("%s:%s",j[1], j[3])
				res = append(res,proxy)
			}
		}
	}
	return res
}

func XseoInProxies() []string {
	var res []string
	url := "http://xseo.in/proxylist"
	page := post(url, "{'submit': 1}", "")
	regCharNum := `\b(?P<char>[a-z])=(?P<num>\d);`
	reg := regexp.MustCompile(regCharNum)
	proxies := reg.FindAllStringSubmatch(string(page), -1)
	exPortJs := `\(""\+(?P<chars>[a-z+]+)\)`
	reg2 := regexp.MustCompile(exPortJs)
	charToNum := make(map[string]string)
	_ = charToNum
	for _, p := range proxies {
		s := strings.Split(strings.TrimSuffix(p[0], ";"), "=")
		charToNum[s[0]] = s[1]
	}
	/*for i, j := range charToNum {
		fmt.Println(i, j)
	}*/

	pageCopy := reg2.ReplaceAllFunc(page, func(s []byte) []byte {
		t := reg2.ReplaceAllString(string(s), `$1`)
		var z string
		var port []byte
		for _, ch := range t {

			if ch != '+' {
				a := charToNum[string(ch)]
				port = append(port, []byte(a)...)
			}

		}
		z = string(port)
		return []byte(z)
	})
//fmt.Println(string(pageCopy))
	reg3 := regexp.MustCompile(`(?P<ip>(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?)).<font(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|(?P<port>\d{2,5})))`)
	//reg3 := regexp.MustCompile(proxyRegexp)
	proxiesArray := reg3.FindAllStringSubmatch(string(pageCopy), -1)
	for _, p := range proxiesArray {
		port, err := strconv.ParseInt(p[3], 10, 64)
		if err != nil {
			log.Println(err)
		}
		proxy := fmt.Sprintf("%s:%d", p[1], port)
		res = append(res, proxy)
	}
	return res
}

/*
func jsPortToNum(port []byte) []byte{

}
*/
//ToolsRosinstrumentCom TODO
func ToolsRosinstrumentCom() []string {
	var res []string
	return res
}

func GatherProxyProxies() []string {
	var res []string
	url := "http://www.gatherproxy.com/sockslist/"
	page := post(url, nil, "")
	//fmt.Println(string(page))
	reg2 := regexp.MustCompile(`(?P<ip>(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?)).*\n(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|'(?P<port>[\d\w]+)'))`)
	proxies := reg2.FindAllStringSubmatch(string(page), -1)
	for _, p := range proxies {
		port, err := strconv.ParseInt(p[3], 16, 64)
		if err != nil {
			log.Println(err)
		}
		proxy := fmt.Sprintf("%s:%d", p[1], port)
		res = append(res, proxy)
	}
	return res
}

func GatherProxyComProxies() []string {
	var res []string
	url := "http://www.gatherproxy.com/proxylist/anonymity/"
	expNumPages := `href="#(\d+)`
	reg := regexp.MustCompile(expNumPages)
	var dataArray []Data
	var temp []string
	typeP := []string{"anonymous", "elite"}
	for _, typ := range typeP {
		data := Data{Type: typ, PageIdx: 1}
		body := post(url, data, "")
		t := reg.FindAllString(string(body), -1)
		temp = append(temp, t...)
		_, max := MinMax(extractNumbers(temp))

		for i := 1; i < max; i++ {
			data1 := Data{Type: typ, PageIdx: i}
			dataArray = append(dataArray, data1)
		}
	}
	var proxies [][]string
	reg2 := regexp.MustCompile(`(?P<ip>(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?)).*\n(.*?(?:(?:(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?))|'(?P<port>[\d\w]+)'))`)
	for _, d := range dataArray {
		page := post(url, d, "")
		f := reg2.FindAllStringSubmatch(string(page), -1)
		proxies = append(proxies, f...)
	}
	for _, p := range proxies {

		//fmt.Println("ip",p[1],"port",p[3])
		port, err := strconv.ParseInt(p[3], 16, 64)
		if err != nil {
			log.Println(err)
		}
		proxy := fmt.Sprintf("%s:%d", p[1], port)
		res = append(res, proxy)
	}
	return res
}

func FoxyToolsRuProxies() []string {
	res := make([]string, 0)
	var urls []string
	for i := 1; i < 6; i++ {
		urls = append(urls, fmt.Sprintf("http://api.foxtools.ru/v2/Proxy.txt?page=%d", i))
	}

	res = findOnURLs(urls, r)
	return res
}

func ProxyListMeProxies() []string {
	//var res [][]string
	var res2 []string
	page := get(string(ProxyListMeURL))
	reg := regexp.MustCompile(string(ProxyListMeURLRegex))
	numberRegex := regexp.MustCompile(string(NumberExtractRegex))
	numbers := make([]int, 0)
	for _, url := range reg.FindAll(page, -1) {
		s := numberRegex.FindString(string(url))
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		numbers = append(numbers, n)
	}
	urls := make([]string, 0)
	_, max := MinMax(numbers)
	for i := 1; i < max; i++ {
		url := fmt.Sprintf("https://proxylist.me/?page=%d", i)
		urls = append(urls, url)
	}
	res2 = findOnURLs(urls[:21], ProxyListMeProxyRegex)
	reg2 := `</a></td>(\s+|\s*\s*)<td >`
	reg3 := regexp.MustCompile(reg2)
	for i := 0; i < len(res2); i++ {
		a := reg3.ReplaceAllString(res2[i], ":")
		res2[i] = a
	}
	return res2
}

//MaxiProxiesCom TODO
func MaxiProxiesCom() []string {
	var res []string
	MaxiProxiesURLregex := `<a href\s*=\s*['"]([^'"]*example[^'"#]*)['"]>`
	reg := regexp.MustCompile(MaxiProxiesURLregex)
	page := get("http://maxiproxies.com/category/proxy-lists/")
	var urls []string
	//fmt.Println(string(page))
	for _, url := range reg.FindAll(page, -1) {
		//fmt.Println(string(url))
		urls = append(urls, string(url))
	}
	_ = urls
	return res
}

func AliveProxyProxies() []string {
	var res []string
	urls := []string{
		"socks5-list",
		"high-anonymity-proxy-list",
		"anonymous-proxy-list",
		"fastest-proxies",
		"us-proxy-list",
		"gb-proxy-list",
		"fr-proxy-list",
		"de-proxy-list",
		"jp-proxy-list",
		"ca-proxy-list",
		"ru-proxy-list",
		"proxy-list-port-80",
		"proxy-list-port-81",
		"proxy-list-port-3128",
		"proxy-list-port-8000",
		"proxy-list-port-8080",
	}
	var proxies [][]byte
	reg := regexp.MustCompile(string(r))
	for _, url := range urls {
		url := fmt.Sprintf("http://www.aliveproxy.com/%s", url)
		page := get(url)
		list := reg.FindAll(page, -1)

		proxies = append(proxies, list...)
	}
	for _, Proxy := range proxies {
		res = append(res, string(Proxy))
	}
	return res
}

//ProxzComProxies TODO
func ProxzComProxies() []string {
	var res []string
	URLregex := `href\s*=\s*['"]([^'"]?proxy_list_high_anonymous_[^'"]*)['"]`
	reg := regexp.MustCompile(URLregex)
	url1 := "http://www.proxz.com/proxy_list_high_anonymous_0.html"
	page := get(url1)
	for _, url := range reg.FindAll(page, -1) {
		fmt.Println(string(url))
	}
	urls := []string{url1}
	m := regexp.MustCompile(z)
	for _, p := range urls {
		page1 := get(p)
		u, err := url.QueryUnescape(string(m.Find(page1)))
		if err != nil {
			fmt.Println(err)
		}
		r := regexp.MustCompile(string(FreeProxyRegex))
		a := r.FindAll([]byte(u), -1)
		for _, z := range a {
			fmt.Println(string(z))
		}
	}

	return res
}

func BlogSpotProxies() []string {
	res := make([]string, 0)
	domains := []string{"sslproxies24.blogspot.com",
		"proxyserverlist-24.blogspot.com",
		"freeschoolproxy.blogspot.com",
		"googleproxies24.blogspot.com"}
	URLregex := `<a href\s*=\s*['"]([^'"]*\.\w+/\d{4}/\d{2}/[^'"#]*)['"]>`
	reg := regexp.MustCompile(URLregex)
	var proxyURLs [][]byte
	for _, domain := range domains {
		url := fmt.Sprintf("http://%s", domain)
		page := get(url)
		list := reg.FindAll(page, -1)
		proxyURLs = append(proxyURLs, list...)
	}
	for _, url := range proxyURLs {
		u := strings.TrimPrefix(strings.TrimSuffix(string(url), "'>"), "<a href='")
		page := get(u)
		a := getProxies(page, string(r))
		res = append(res, a...)
	}
	return res
}

func CheckerProxyList() []string {
	res := make([]string, 0)
	exp := regexp.MustCompile(string(CheckProxyURLsRegex))
	page := get("https://checkerproxy.net/")
	urls := make([]string, 0)
	for _, url := range exp.FindAll(page, -1) {
		url1 := strings.TrimSuffix(strings.TrimPrefix(string(url), "href=\""), "\"")
		urls = append(urls, "https://checkerproxy.net/api"+url1)
	}
	//IPReg := regexp.MustCompile(IPregex)
	proxies := make([]CheckProxyStruct, 0)
	for _, url := range urls {
		proxyPage := get(url)
		//p := getProxies(proxyPage, IPregex)
		//res = append(res, p...)
		err := json.Unmarshal(proxyPage, &proxies)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(proxies)
	}
	for _, proxy := range proxies {
		res = append(res, proxy.Addr)
	}
	return res
}

func WebanetLabsProxies() []string {
	res := make([]string, 0)
	exp := regexp.MustCompile(string(WebanetLabsURLsRegex))
	page := get("https://webanetlabs.net/publ/24")
	//fmt.Println(string(page))
	urls := make([]string, 0)
	for _, url := range exp.FindAll(page, -1) {
		url1 := strings.TrimLeft(strings.TrimRight(string(url), "\""), "href=\"")
		urls = append(urls, string(WebanetLabsURL)+string(url1))
	}
	var proxies [][]byte
	IPReg := regexp.MustCompile(string(proxyRegexp))
	for _, url := range urls {
		page := get(url)
		p := IPReg.FindAll(page, -1)
		proxies = append(proxies, p...)
	}
	for _, proxy := range proxies {
		//fmt.Println(string(proxy))
		res = append(res, string(proxy))
	}

	return res
}

func FreeProxyList() []string {
	res := make([]string, 0)
	urls := []string{"http://www.freeproxylists.com/socks.html",
		"http://www.freeproxylists.com/elite.html",
		"http://www.freeproxylists.com/anonymous.html"}
	exp := `href\s*=\s*['"](?P<t>[^'"]*)/(?P<uts>\d{10})[^'"]*['"]`
	validProxy := regexp.MustCompile(exp)
	proxyLists := make([][]byte, 0)
	for _, url := range urls {
		body := get(url)
		l := validProxy.FindAll(body, -1)
		proxyLists = append(proxyLists, l...)
	}
	for _, list := range proxyLists {
		l := strings.TrimLeft(strings.TrimRight(string(list), "'"), "href='")
		arr := strings.Split(l, "/")
		u := fmt.Sprintf("http://www.freeproxylists.com/load_%s_%s", arr[0], arr[1])
		page := get(u)
		pr := getProxies(page, string(FreeProxyRegex))
		res = append(res, pr...)
	}
	return res
}

/*
ProxyListOrgProxies returns a list of proxies from proxy-list.org
*/
func ProxyListOrgProxies() []string {
	res := make([]string, 0)
	var expr string
	expr = "href\\s*=\\s*['\"]\\./([^'\"]?index\\.php\\?p=\\d+[^'\"]*)['\"]"
	var url string
	url = "http://proxy-list.org/french/index.php?p=1"

	bodyBytes := get(url)
	var validUrl = regexp.MustCompile(expr)

	//fmt.Println(string(validUrl.Find(bodyBytes)))
	//urls := make([]string, 1)
	urls := []string{"/index.php?p=1"}
	for _, a := range validUrl.FindAll(bodyBytes, -1) {
		urls = append(urls, string(a))
	}

	urls = Unique(urls)

	for _, url1 := range urls {
		url1 = strings.TrimLeft(strings.TrimRight(url1, "\""), "href=\".")
		//fmt.Println("url : " ,url1)

		//fmt.Println(url1)
		url1 = "http://proxy-list.org/french" + url1
		bodyBytes := get(url1)
		p := getProxiesProxyListOrg(bodyBytes, string(ProxyListRegex))
		res = append(res, p...)
	}
	return res
}

func getProxiesProxyListOrg(page []byte, regex string) []string {
	ref := regexp.MustCompile(regex)
	proxies := make([]string, 0)
	for _, a := range ref.FindAll(page, -1) {
		b := strings.TrimLeft(strings.TrimRight(string(a), "')"), "Proxy('")
		data, err := base64.StdEncoding.DecodeString(b)
		if err != nil {
			log.Println(err)
		}
		proxies = append(proxies, string(data))
	}
	return proxies
}

func findOnURLs(urls []string, regex Regex) []string {
	reg := regexp.MustCompile(string(regex))
	res := make([]string, 0)
	for _, url := range urls {
		page := get(url)
		proxies := reg.FindAllString(string(page), -1)
		res = append(res, proxies...)
		//fmt.Println(proxies)

	}
	return res
}

func getProxies(page []byte, regex string) []string {
	proxyReg := regexp.MustCompile(regex)
	proxies := make([]string, 0)
	for _, proxy := range proxyReg.FindAll(page, -1) {
		if regex == string(FreeProxyRegex) {
			proxySplit := strings.Replace(string(proxy), "&lt;/td&gt;&lt;td&gt;", ":", 1)
			proxies = append(proxies, proxySplit)
		} else {
			proxies = append(proxies, string(proxy))
		}
	}
	return proxies

}

func get(url string) []byte {
	log.Println(url)
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {
		log.Println(err)
	}
	return body
}

//TO REFRACTOR
func post(url string, data interface{}, header string) []byte {
	d, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	//http.Post(url)
	var strPost = []byte("POST")
	var strRequestURI = []byte(url)
	req := fasthttp.AcquireRequest()
	req.SetBody(d)

	req.Header.SetMethodBytes(strPost)
	req.SetRequestURIBytes(strRequestURI)
	//	req.Header.SetContentType("application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		panic("handle error")
	}
	fasthttp.ReleaseRequest(req)

	body := res.Body()
	//fmt.Println(string(body))
	fasthttp.ReleaseResponse(res)
	return body
}

func MinMax(array []int) (int, int) {
	if len(array) < 1 {
		log.Println("empty array")
		return 0,0
	}
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func extractNumbers(array []string) []int {
	var res []int
	reg := regexp.MustCompile(string(DegitRegex))
	for _, s := range array {
		digit, err := strconv.Atoi(reg.FindString(s))
		if err != nil {
			log.Println(err)
		}
		res = append(res, digit)
	}
	return res
}

var providers []ProxyProvider

func InitProviders() []ProxyProvider {
	providers =[]ProxyProvider{
		ProxyProvider{
			Domain:"http://www.proxylists.net/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "http://ipaddress.com/proxy-list/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "https://www.sslproxies.org/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "https://freshfreeproxylist.wordpress.com/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "http://proxytime.ru/http",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : false,
		},
		ProxyProvider{
			Domain: "https://free-proxy-list.net/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "https://us-proxy.org/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "http://fineproxy.org/eng/fresh-proxies/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : false,
	},
		ProxyProvider{
			Domain: "https://socks-proxy.net/",
			protocols:[]string{"SOCKS4", "SOCKS5"},
			Active : true,
		},
		ProxyProvider{
			Domain: "http://www.httptunnel.ge/ProxyListForFree.aspx",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "http://cn-proxy.com/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		ProxyProvider{
			Domain: "https://hugeproxies.com/home/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
		Active : false,
		},
		ProxyProvider{
			Domain: "http://proxy.rufey.ru/",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
		Active : false,
		},
		ProxyProvider{
			Domain: "https://geekelectronics.org/my-servisy/proxy",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : false,
		},
		ProxyProvider{
			Domain: "http://pubproxy.com/api/proxy?limit=20&format=txt",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			Active : true,
		},
		/*
		With functions
		*/
		ProxyProvider{
			Domain: "https://proxy-list.org",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy:ProxyListOrgProxies,
			Active : true,
		},
		ProxyProvider{
			Domain: "xseo.in",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: XseoInProxies,
			Active : true,
		},
		ProxyProvider{
			Domain: "spys.ru",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: GetProxiesSpysRu,
			Active : false,
		},
		ProxyProvider{
			Domain: "list.proxylistplus.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: ProxyListPlusCom,
			Active : false,
		},
		ProxyProvider{
			Domain: "proxylist.me",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: ProxyListMeProxies,
			Active : true,
		},
		ProxyProvider{
			Domain: "foxtools.ru",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: FoxyToolsRuProxies,
			maxTries:1,
			Active : true,
		},
		ProxyProvider{
			Domain: "gatherproxy.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: GatherProxyComProxies,
			Active : false,
		},
		ProxyProvider{
			Domain: "nntime.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: NntimeComProxies,
			Active : true,
		},
		ProxyProvider{
			Domain: "Blogspot_com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: BlogSpotProxies,
			Active : true,
		},
		ProxyProvider{
			Domain: "gatherproxy.com^sock",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: GatherProxyProxies,
			Active : false,
		},
		ProxyProvider{
			Domain: "tools.rosinstrument.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: ToolsRosinstrumentCom,
		Active : false,
		},
		ProxyProvider{
			Domain: "my-proxy.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: MyProxyCom,
			maxConnections:2,
			Active : true,
		},
		ProxyProvider{
			Domain: "checkerproxy.net",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: CheckerProxyList,
			Active : true,
		},
		ProxyProvider{
			Domain: "aliveproxy.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: AliveProxyProxies,
			Active : true,
		},
		ProxyProvider{
			Domain: "freeproxylists.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: FreeProxyList,
			Active : true,
		},
		ProxyProvider{
			Domain: "webanetlabs.net",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: WebanetLabsProxies,
			Active : true,
		},

		ProxyProvider{
			Domain: "maxiproxies.com",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: MaxiProxiesCom,
			Active : false,
		},
		ProxyProvider{
			Domain: "www.proxy-list.download",
			protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
			GetProxy: ProxyListDownload,
			Active : true,
		},
	}

	return providers
}

func (p ProxyProvider) GetProxies2( ) []string {
	page := get(p.Domain)
	regex := proxyRegexp
	proxyReg := regexp.MustCompile(regex)
	proxies := make([]string, 0)
	for _, p := range proxyReg.FindAllSubmatch(page, -1) {
	if len(p[3]) <1{
		continue
	}
		proxy := fmt.Sprintf("%s:%s", p[1], p[3])
		proxies = append(proxies, proxy)

	}
	return proxies
}

func getProxiesFromProvider(p ProxyProvider) []string{
	var res []string
	if p.GetProxy == nil {
		res = p.GetProxies2()
	}else {
		res =p.GetProxy()
	}
	return res
}