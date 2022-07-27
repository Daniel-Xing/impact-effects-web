package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var count int64 = 0

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
		Timeout:   60 * time.Second, // 超时加在这里，是每次调用的超时
	}
	_, err := client.Get("http://localhost:50012/impact")
	// defer client.CloseIdleConnections()
	if err != nil {
		log.Println(err)
		return
	}

	// atomic.AddInt64(&count, 1)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}

	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetRe()
		}()
	}

	wg.Wait()

	fmt.Println(time.Since(start), count)
	return
}
