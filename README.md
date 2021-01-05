# k8s-debug-tools
All debug tools for k8s

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" http_pod_traffic.go
```
