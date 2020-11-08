/* Sample Web Application to test traffic routing
   Author : Ansil H
   Email: ansilh@gmail.com */

// Make sure add below to Pod spec
//  containers:
//      - env:
//        - name: MY_NODE_NAME
//          valueFrom:
//            fieldRef:
//              fieldPath: spec.nodeName

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func main() {
	http.HandleFunc("/", rootDefault)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func rootDefault(w http.ResponseWriter, r *http.Request) {
	var podIp string

	keys, _ := r.URL.Query()["headersize"]
	if keys != nil {
		if len(keys[0]) > 1 {
			length, err := strconv.Atoi(keys[0])
			if err == nil && length > 1 {
				value := RandStringBytesMaskImprSrcUnsafe(length)
				addCookie(w, "AnsiltestCookie", value, 30*time.Minute)
			}
		}
	}

	nodeName := getEnv("MY_NODE_NAME", "NO_ENV")
	podName, _ := os.Hostname()
	addrs, _ := net.LookupIP(podName)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			podIp = (ipv4).String()
		}
	}
	//Dump all headers
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "Header:%s:%s\n", name, value)
		}
	}
	fmt.Fprintf(w, "pod_name:%s\n", podName)
	fmt.Fprintf(w, "pod_ip:%s\n", podIp)
	fmt.Fprintf(w, "node_name:%s\n", nodeName)
}

func addCookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
