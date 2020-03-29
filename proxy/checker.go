package proxy

import (
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const(
	timeout =time.Duration(5* time.Second)
)

func CheckSOCKS5Proxy(proxyCus Proxy)(err error){
	d:= net.Dialer{
		Timeout:timeout,
		KeepAlive:timeout,
	}
	prox := fmt.Sprintf("%s:%d",proxyCus.Host,proxyCus.Port)
	dialer,_ := proxy.SOCKS5("tcp",prox,nil,&d)
	httpClient := http.Client{
		Timeout:timeout,
		Transport:&http.Transport{
			DisableKeepAlives:true,
			DialTLS:dialer.Dial,
		},
	}
	response,err := httpClient.Get("http://ifconfig.io/ip")
	if err!=nil{
		return err
	}
	defer response.Body.Close()
	io.Copy(ioutil.Discard,response.Body)
	return nil
}
