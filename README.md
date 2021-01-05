# k8s-debug-tools
All debug tools for k8s

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" http_pod_traffic.go
```

```
cat >Dockerfile <<EOF
FROM scratch
LABEL Maintainer="Ansil H"
LABEL Email="xxxxxxx.com"
COPY http_pod_traffic /
CMD ["/http_pod_traffic"]
EOF
```

```
docker build -t registry.ansil.io/library/http_pod_traffic .
```

```
docker push registry.ansil.io/library/http_pod_traffic
```
