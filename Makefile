BUILD_DT:=$(shell date +%F-%T)
VERSION:=$(shell cat ./VERSION)
GO_LDFLAGS:="-s -w -extldflags \"-static\" -X main.BuildVersion=${VERSION} -X main.BuildDate=$(BUILD_DT)"

.PHONE: all
all:
	export CGO_ENABLED=0 ; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags $(GO_LDFLAGS) -o ".build/prometheus_exporter-linux-amd64/prometheus_exporter" && \
	echo "done"