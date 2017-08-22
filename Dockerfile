FROM golang:1.8

COPY . /go/src/github.com/nstogner/kube-demo
RUN go install github.com/nstogner/kube-demo
