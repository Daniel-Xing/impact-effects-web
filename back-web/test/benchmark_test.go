package test

import (
	"log"
	"net/http"
	"testing"
	"time"
)

var tr *http.Transport

var (
	remoteURL = "https://impact.ese.ic.ac.uk/ImpactEarth/cgi-bin/crater.cgi?dist=12371&distanceUnits=1&diam=111&diameterUnits=1&pdens=111&pdens_select=0&vel=1111&velocityUnits=1&theta=45&tdens=1000&wdepth=1111&wdepthUnits=1"
	localURL  = "http://localhost:50012/impact"
)

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
	}
}

func GetRe() {
	client := &http.Client{
		Transport: tr,
		Timeout:   120 * time.Second,
	}
	_, err := client.Get(remoteURL)
	// defer client.CloseIdleConnections()
	if err != nil {
		log.Println(err)
		return
	}

	// atomic.AddInt64(&count, 1)
}

func BenchmarkImpactEffect(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetRe()
	}
}
