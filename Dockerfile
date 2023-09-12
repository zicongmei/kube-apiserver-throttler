FROM golang:1.21
WORKDIR /code/src/kube-apiserver-throttler
ENV GOPATH=/code
ADD main.go .


RUN go mod init
RUN go mod tidy
RUN go build -o /bin/main ./main.go

FROM golang:1.21
COPY --from=0 /bin/main /bin/main
CMD ["/bin/main"]