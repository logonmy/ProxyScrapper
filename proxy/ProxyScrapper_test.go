package proxy

import (
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestInitProviders(t *testing.T) {
	prov := InitProviders()
	home, err := homedir.Dir()
	if err !=nil{
		log.Println(err)
	}
	f,err := os.Create(home+string(os.PathSeparator)+"proxyBroker.log")
	if err !=nil{
		log.Println(err)
	}
	log.SetOutput(f)
	for _, p := range prov {
		if !p.Active{
			res := getProxiesFromProvider(p)
			if len(res)> 0 {
				t.Logf("provider %s working",p.Domain)
			}
			continue
		}
		res := getProxiesFromProvider(p)
		//res := p.GetProxies2()
		if len(res) < 1 {
			t.Errorf("Provider %s returned empty result", p.Domain)
			continue
		}
		arr := strings.Split(res[0], ":")
		if len(arr) != 2 {
			t.Errorf("Provider %s hosts not formatted ip:port it is %s", p.Domain,res[0])
			continue
		}
		if arr[0] == "" {
			t.Errorf("ip not there in %s domain", p.Domain)
		}

		if arr[1] == "" {
			t.Errorf("port %s not there in %s domain", arr[1],p.Domain)
		}
		_, err := strconv.ParseInt(arr[1], 10, 32)
		if err != nil {
			t.Errorf("Port %s is not an int in %s domain", arr[1],p.Domain)
		}
	}
}
