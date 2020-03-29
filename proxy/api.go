package proxy

import (
	"fmt"
	"sync"
)

func Gather(limit int)(res []Proxy){
	providers := InitProviders()
	r := make(chan []string, 5)
	var wg sync.WaitGroup
	defer wg.Done()
	wg.Add(len(providers))
	for _, p := range providers {
		go func() {
			r <- getProxiesFromProvider(p)
		}()
	}
	for i:=0;i<len(providers);i++ {
		var arr []string
		arr = <-r
		for _,p := range arr{
			pr := CreateProxy(p)
			if pr == nil {
				break
			}
			fmt.Printf("Found a %s proxy : %s:%d in %v seconds\n",pr.Types,pr.Host,pr.Port,pr.Time)
			//fmt.Printf("%s:%d",pr.Host,pr.Port)
			res = append(res, *pr)
			if len(res) >= limit {
				return res
			}

		}
	}
	return res

}

func Find(limit int, types []string, countries []string, anonymous bool) []Proxy {
	var res []Proxy
	providers := InitProviders()
	r := make(chan []string, 15)
	//pr := make(chan Proxy)
	ip := GetRealIP()
	var wg sync.WaitGroup
	defer wg.Done()
	wg.Add(len(providers))
	for _, p := range providers {
		go func() {
			r <- getProxiesFromProvider(p)
		}()
	}

	res = check(r, limit, types, countries, anonymous, ip)
	return res
}

func check(r chan []string, limit int, types []string, countries []string, anonymous bool, ip string) (res []Proxy) {
	var wg2 sync.WaitGroup
	defer wg2.Done()
	for i := 0; i < len(providers); i++ {
		var arr []string
		arr = <-r
		for _, p := range arr {
			pr := CreateProxy(p)
			if pr == nil {
				break
			}
			go func() {
				wg2.Add(1)
				pr.getAnonymityLevel(wg2, ip)

				if len(countries)> 0 && !containsCountry(countries,pr){
					return
				}

				if len(types)> 0 && !containsType(types,pr) {
					return
				}

				/*if anonymous && pr.IsAnonymous{
					return
				}*/

				if len(pr.Types) > 0 {
					fmt.Printf("Found a %s proxy : %s:%d in %v seconds\n",pr.Types,pr.Host,pr.Port,pr.Time)
					//fmt.Printf("%s:%d\n",pr.Host,pr.Port)
					res = append(res, *pr)
				}
			}()
			if len(res) >= limit {
				return res[:limit]
			}

		}
	}
	return res
}

func containsType(types []string, proxy *Proxy) bool {
	for _, p := range proxy.Types {
		for _, t := range types {
			if p == t {
				return true
			}
		}
	}
	return false

	}

func containsCountry(countries []string,proxy *Proxy)bool{
	for _, p := range countries {
			if p == proxy.Country.Country.ISOCode {
				return true
			}
	}
	return false
}