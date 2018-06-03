FROM golang:1.10 AS builder

ADD . /go/src/github.com/stack-labs-tech/rex-gcp-k8s/demos/kubernetes-autoscaling
WORKDIR /go/src/github.com/stack-labs-tech/rex-gcp-k8s/demos/kubernetes-autoscaling

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /go/bin/kubernetes-autoscaling

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/kubernetes-autoscaling /go/bin/kubernetes-autoscaling

ENTRYPOINT ["/go/bin/kubernetes-autoscaling"]
