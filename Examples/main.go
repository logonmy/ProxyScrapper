package main

import (
	"fmt"
	"github.com/oucema001/ProxyScrapper/proxy"
)

func main() {
	//r := ProxyListOrgProxies()
	//fmt.Println(r)
	//r := CheckerProxyList()
	//fmt.Println(r)
	//r := BlogSpotProxies()
	//fmt.Println(r)

	//ProxzComProxies()

	//r := AliveProxyProxies()
	//fmt.Println(r)
	//res := proxy.ProxyListDownload()

/*
	proxyListOrg := proxy.ProxyProvider{
		Domain :"http://www.proxylists.net/",
	}
	res := proxyListOrg.GetProxies2()
*/
	//_ = res
	//fmt.Println(res)



/*	prov := proxy.InitProviders()
	var a[]string
	for _,p := range prov {
		var res []string
		if p.GetProxy == nil {

			res = p.GetProxies2()
			if len(res)>0 {
				//fmt.Println(p.Domain)
				//fmt.Println(res)
			}

		}else {
			res = p.GetProxy()
			//fmt.Println(res)
		}
	a = append(a,res...)
	}
fmt.Println(a)*/

	//res := proxy.NntimeComProxies()
	//fmt.Println(res)


	//proxy.GetIPInfo("178.76.129.69");
   //a:=proxy.GetRealIP()
   //fmt.Println(a)
  // proxy.Grab()



  /*	res := proxy.XseoInProxies()
	for _, r := range res {
		fmt.Println(r)
	}

*/

/*
  pr := proxy.InitProviders()
   for _,p := range pr{
	   var res []string
	   if p.GetProxy == nil {
		   res = p.GetProxies2()

	   }else {
		   res =p.GetProxy()
	   }
	if len(res)>0{
	   fmt.Println(p.Domain)
	   fmt.Println(res[0])
	}
   }*/



   /*
	start := time.Now()
	//log.SetOutput(nil)
	res:= proxy.Find(10,[]string{"SOCKS5"},[]string{"DE"},false)
   fmt.Println("res",res)
	for i, r := range res{
		fmt.Println(i,"r : ", r)
		fmt.Println(r.Country.Country.ISOCode)
	}
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)


	*/
//time.Sleep(30*time.Second)

/*a:= proxy.ProxyProvider{
		Domain: "http://www.httptunnel.ge/ProxyListForFree.aspx",
	//	Protocols:[]string{"HTTP", "CONNECT:80", "HTTPS", "CONNECT:25"},
	}

*/	res1 := proxy.ProxyListPlusCom()
for _, r := range res1 {
	fmt.Println(r)
}
_=res1
a := proxy.GetIPInfo("8.8.8.8")
fmt.Println(a)
	}
