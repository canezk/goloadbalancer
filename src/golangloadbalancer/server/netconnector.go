/**
包括如下：
1、http server listner
 */
package server

import (
	"net/http"
	"log"
	"golangloadbalancer/consistenthash"
)

var connsistentHash consistenthash.Consistent

func Listen(port int)  {
	http.Handle("/", http.HandlerFunc(balanceHandler))
	connsistentHash = new(consistenthash.Consistent)
	connsistentHash.ZookeeperUrl = ""
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func balanceHandler(w http.ResponseWriter, req *http.Request)  {
	url := req.URL
	url.Host = connsistentHash.GetServer("", url)

	proxyReq, err := http.NewRequest(req.Method, url.String(), req.Body)
	if err != nil {
		log.Fatal("Proxy error!", err)
	}
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
	defer resp.Body.Close()
}
