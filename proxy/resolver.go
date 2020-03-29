package proxy

import (
	"github.com/oschwald/maxminddb-golang"
	"github.com/valyala/fasthttp"
	"log"
	"math/rand"
	"net"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

type Location struct {
	Accuracy  int     `maxminddb:"accuracy_radius"`
	Latitude  float64 `maxminddb:"latitude"`
	Longitude float64 `maxminddb:"longitude"`
	TimeZone  string  `maxminddb:"time_zone"`
}

type CountryNameMap struct {
	Chineese   string `maxminddb:"zh-CN"`
	German     string `maxminddb:"de"`
	Spanish    string `maxminddb:"es"`
	French     string `maxminddb:"fr"`
	Japanese   string `maxminddb:"ja"`
	Portuguese string `maxminddb:"pt-BR"`
	Russian    string `maxminddb:"ru"`
}

type Place struct {
	GeoNameID         int            `maxminddb:"geoname_id"'`
	ISOCode           string         `maxminddb:"iso_code"`
	IsInEuropeanUnion bool           `maxminddb:"is_in_european_union"'`
	//Name              CountryNameMap `maxminddb:"names"`
}

type Country Place

type Continent Place

type RegisteredCountry Place

type IPInfo struct {
	Country           Country           `maxminddb:"country"`
	Continent         Continent         `maxminddb:"continent"`
	RegisteredCountry RegisteredCountry `maxminddb:"registered_country"`
	Location          Location          `maxminddb:"location"`
}

func CheckIsIP(host string) bool {
	return net.ParseIP(host) != nil
}

var IPHosts []string
var basePath string

func init(){
	IPHosts = []string{"https://wtfismyip.com/text",
		"http://api.ipify.org/",
		"http://ipinfo.io/ip",
		"http://ipv4.icanhazip.com/",
		"http://myexternalip.com/raw",
		"http://ipinfo.io/ip",
		"http://ifconfig.io/ip"}
		_, b, _, _ = runtime.Caller(0)
		basePath   = filepath.Dir(b)


}

func GetIPInfo(host string) IPInfo {
	countryDb := path.Join(basepath,"data","GeoLite2-Country.mmdb")
	db,err := maxminddb.Open(countryDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip := net.ParseIP(host)
	var record IPInfo
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Println(err)
	}
	return record
}

func getRandomIpHost()string{
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return IPHosts[r.Intn(len(IPHosts))]
}

/*
func GetRealIP() string{
	var ip string
	var wg sync.WaitGroup
	//defer wg.Done()
	for _, host := range IPHosts {
		wg.Add(1)
		_,body ,err:= fasthttp.Get(nil,host)
		if err!=nil{
			log.Println(err)
		}
		if CheckIsIP(string(body)){
			ip =string(body)
			fmt.Println(string(body))
			wg.Done()
			break
		}
	}
	return ip
}
*/

func GetRealIP() string{
	var ip string
	p := make(chan string)
	for _, host := range IPHosts {
		go func(h string )  {
			_, body, err := fasthttp.Get(nil, h)
			if err != nil {
				log.Println(err)
			}
			p <- string(body)
		}(host)
		 ip:=<- p
		if CheckIsIP(ip) {
			break
		}
	}
	return ip
}

func GetAllIps(content []byte) []string{
	var res []string
	reg := `(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?)`
	regex := regexp.MustCompile(reg)
	for _, p := range regex.FindAll(content,-1){
		res = append(res,string(p))
	}
	return res
}

func getRandomHost()string{
	rand.Seed(time.Now().Unix())
	return IPHosts[rand.Intn(len(IPHosts))]
}