package proxy

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Proxy struct {
	Host          string
	Port          int
	Types         []string
	timeout       int
	ExpectedTypes []string
	Country       IPInfo
	IsAnonymous   bool
	Time float64
}

func CreateProxy(host string) *Proxy {
	arr := strings.Split(host, ":")
	if len(arr) < 2 {
		fmt.Println("Split error")
		return nil
	}
	ip := arr[0]
	port, err := strconv.Atoi(arr[1])
	info := GetIPInfo(ip)

	if err != nil {
		log.Println(err)
	}
	if port > 65536 {
		log.Println("Wrong port number")
	}
	proxy := &Proxy{
		Host:    arr[0],
		Port:    port,
		Country: info,
	}
	return proxy
}

func (p *Proxy) getAnonymityLevel(wg sync.WaitGroup, realIP string) {
	//	wg.Add(1)
	b, http := p.HTTPProxyChecker()
	if b {
		//fmt.Println(string(http))
		//p.types = append(p.types,"HTTP")
		if !(strings.Contains(string(http), realIP)) {
			p.IsAnonymous = true
		}
	}
	b2, sock := p.Socks5ProxyChecker()
	if b2 {
		//fmt.Println(string(sock))
		//p.types = append(p.types,"SOCKS5")
		if !(strings.Contains(string(sock), realIP)) {
			p.IsAnonymous = true
		}
	}
}

func (p *Proxy) HTTPProxyChecker() (bool, []byte) {
	port := strconv.Itoa(p.Port)
	prox := fmt.Sprintf("%s:%s", p.Host, port)
	startTs := time.Now()

	c := fasthttp.HostClient{
		Addr: prox,
	}
	statusCode, body, err := c.Get(nil, getRandomHost())
	if err != nil {
		log.Println(err)
		return false, nil
	}
	if strings.Contains(string(body),"400"){
		return false,nil
	}
	if statusCode != http.StatusOK {
		return false, nil
	} else {
		timeDiff := time.Now().UnixNano() - startTs.UnixNano()
		p.Time = float64(timeDiff) / 1e9
		p.Types = append(p.Types, "HTTP")
		return true, body
	}
}

func (p *Proxy) Socks5ProxyChecker() (bool, []byte) {
	port := strconv.Itoa(p.Port)
	prox := fmt.Sprintf("%s:%s", p.Host, port)
	startTs := time.Now()
	pu, err := url.Parse("socks5://" + prox)
	if err!=nil{
		fmt.Println(err)
	}

	d := net.Dialer{
		Timeout:   timeout,
		KeepAlive: timeout,
	}

	dialer, err := proxy.SOCKS5("tcp", prox, nil, &d)
	if err != nil {
		log.Println(err)
	}
	_ = dialer
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			//DialTLS: dialer.Dial,
			Proxy: http.ProxyURL(pu),
		},
	}
	resp, err := client.Get(getRandomHost())
	if err != nil {
		log.Println(err)
		return false, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false, nil
	}
	if resp.StatusCode != http.StatusOK {
		return false, nil
	} else {
		timeDiff := time.Now().UnixNano() - startTs.UnixNano()
		p.Time = float64(timeDiff) / 1e9
		p.Types = append(p.Types, "SOCKS5")
		return true, body
	}
	return false,nil
}

