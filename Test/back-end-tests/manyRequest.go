package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var tr *http.Transport

var (
	remoteURL = "https://impact.ese.ic.ac.uk/ImpactEarth/cgi-bin/crater.cgi?dist=12371&distanceUnits=1&diam=111&diameterUnits=1&pdens=111&pdens_select=0&vel=1111&velocityUnits=1&theta=45&tdens=1000&wdepth=1111&wdepthUnits=1"
	localURL2 = "http://121.36.81.191/cgi-bin/crater.cgi?dist=111&distanceUnits=1&diam=111&diameterUnits=1&pdens=111000&pdens_select=0&vel=111&velocityUnits=1&theta=45&tdens=1000&wdepth=111&wdepthUnits=1"
	localURL  = "http://121.36.81.191:50052/simulator"
	localURL3 = "http://121.36.81.191:50052/simulatorWithRedis"
)

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
	}
}

func GetRe(URL string) {
	client := &http.Client{
		Transport: tr,
		Timeout:   100 * time.Minute,
	}
	content, err := client.Get(URL)
	defer client.CloseIdleConnections()
	if err != nil {
		log.Println(err)
		return
	}
	defer content.Body.Close()

	// body, err := ioutil.ReadAll(content.Body)
	// if err != nil {
	// 	log.Println("error")
	// }
	// fmt.Println(string(body))
}

func Post(args map[string]interface{}, URL string) {
	client := &http.Client{
		Transport: tr,
		Timeout:   100 * time.Minute,
	}

	bytesData, _ := json.Marshal(args)
	req, _ := http.NewRequest("POST", URL, bytes.NewReader(bytesData))

	content, err := client.Do(req)
	defer client.CloseIdleConnections()
	defer content.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	// body, err := ioutil.ReadAll(content.Body)
	// if err != nil {
	// 	log.Println("error")
	// }

	// fmt.Println(string(body))
}

func main() {
	// read the args
	url := flag.String("url", "http://121.36.81.191:50052/simulator", "Post Url")
	fmt.Println(*url)

	impactor_density := flag.Int("impactor_density", 111, "impactor_density")
	impactor_diameter := flag.Int("impactor_diameter", 111, "impactor_diameter")
	impactor_velocity := flag.Int("impactor_velocity", 111, "impactor_velocity")
	impactor_theta := flag.Int("impactor_theta", 45, "impactor_theta")
	target_density := flag.Int("target_density", 111, "target_density")
	target_depth := flag.Int("target_depth", 45, "target_depth")
	target_distance := flag.Int("target_distance", 111, "target_distance")

	requestTimes := flag.Int("requestTimes", 10, "requestTimes")
	requestMethod := flag.String("requestMethod", "post", "requestMethod")

	flag.Parse()

	start := time.Now()
	wg := sync.WaitGroup{}

	fmt.Printf("Run request for %d times to with Redis URL: %s \n", *requestTimes, *url)
	for i := 0; i < *requestTimes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if *requestMethod == "get" {
				GetRe(*url)
			} else if *requestMethod == "post" {
				data := map[string]interface{}{
					"impactor_density":  *impactor_density,
					"impactor_diameter": *impactor_diameter,
					"impactor_velocity": *impactor_velocity,
					"impactor_theta":    *impactor_theta,
					"target_density":    *target_density,
					"target_depth":      *target_depth,
					"target_distance":   *target_distance,
				}
				Post(data, *url)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Time Cost: ", time.Since(start).Milliseconds(), "ms")
	return
}
