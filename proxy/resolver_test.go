package proxy

import (
	"strconv"
	"strings"
	"testing"
)

func BenchmarkGetRealIP(t *testing.B){

	for n := 0; n < t.N; n++ {
		go GetRealIP()
	}
}

func TestGetRealIP(t *testing.T) {
	ip := GetRealIP()
	parts := strings.Split(ip, ".")

	if len(parts) < 4 {
		t.Errorf("%s Not a valid IPv4",ip)
	}

	for _,x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				t.Errorf("%s Not a valid IPv4",ip)
			}
		} else {
			t.Errorf("%s Not a valid IPv4",ip)
		}

	}
}
