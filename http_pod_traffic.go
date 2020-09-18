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
        "net"
        "net/http"
        "os"
)

func getEnv(key, fallback string) string {
        if value, ok := os.LookupEnv(key); ok {
                return value
        }
        return fallback
}

func rootDefault(w http.ResponseWriter, r *http.Request) {
        var podIp string
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
                        fmt.Fprintf(w, "Header:%s:%s\n",name,value)
                }
        }
        fmt.Fprintf(w, "pod_name:%s\n",podName)
        fmt.Fprintf(w, "pod_ip:%s\n",podIp)
        fmt.Fprintf(w, "node_name:%s\n",nodeName)
}

func main() {
        http.HandleFunc("/", rootDefault)
        err := http.ListenAndServe(":80", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
