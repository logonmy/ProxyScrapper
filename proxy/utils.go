package proxy

import (
	"archive/tar"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func JSONOutput(res []Proxy) string {
	var jsonData []byte
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	return string(jsonData)
}

func CSVOutput(res []Proxy, writer io.Writer) {
	csvWriter := csv.NewWriter(writer)
	var csvData [][]string
	var row []string
	for _, p := range res {
		port := strconv.Itoa(p.Port)
		var types string
		for _, t := range p.Types {
			types += " " + t
		}
		row = []string{p.Host, port, types}
		csvData = append(csvData, row)
	}
	err := csvWriter.WriteAll(csvData)
	if err != nil {
		log.Println(err)
	}
}

func UpdateGeoDb(licenceKey string) {
	//localDataPath := path.Join(basepath,"proxy","Data")
	fileNameCity := "GeoLite2-City.tar.gz"
	fileNameCountry := "GeoLite2-Country.tar.gz"
	location := []string{"City", "Country"}

	for _, l := range location {
		url := "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-" + l + "&license_key=" + licenceKey + "&suffix=tar.gz"

		dname, _ := ioutil.TempDir("", "proxy-prov")
		var downloadedFile string
		if l == "City" {
			downloadedFile = path.Join(dname, fileNameCity)

		} else {
			downloadedFile = path.Join(dname, fileNameCountry)
		}
		//downloadedFile := dname + "\\" + fileName
		err := downloadFile(downloadedFile, url)
		if err != nil {
			log.Fatal(err)
		}

		defer os.RemoveAll(dname)

		r, err := os.Open(downloadedFile)
		if err != nil {
			log.Fatal(err)
		}
		uncompressedStream, err := gzip.NewReader(r)
		if err != nil {
			log.Fatal(err)
		}
		tarReader := tar.NewReader(uncompressedStream)
		for {

			header, err := tarReader.Next()
			if err == io.EOF {
				// end of tar archive
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if header.Typeflag == tar.TypeReg && strings.HasSuffix(header.Name, ".mmdb") {
				//tempFile := dname+"\\"+header.Name
				tempFile := path.Join(dname, header.Name)
				os.MkdirAll(path.Dir(tempFile), os.ModePerm)
				w, err := os.Create(tempFile)
				if err != nil {
					log.Println(err)
				}
				_, err = io.Copy(w, tarReader)
				if err != nil {
					log.Println(err)
				}
				w.Close()
				name := strings.Split(header.Name, "/")
				err = os.Rename(tempFile, path.Join(basepath, "Data", name[len(name)-1]))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
/*
func Unique(intSlice []Proxy) []Proxy {
	keys := make(map[Proxy]bool)
	var list []Proxy
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
*/

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}