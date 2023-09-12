build:
	go build main.go

build-image:
	docker build . -t gcr.io/zicong-gke-multi-cloud-dev-2/kube-apiserver-throttler:latest

all: build-image
	docker push gcr.io/zicong-gke-multi-cloud-dev-2/kube-apiserver-throttler:latest