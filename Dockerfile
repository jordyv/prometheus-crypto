FROM golang:alpine AS builder

RUN apk --update --no-cache add curl git
COPY . /go/src/github.com/jordyv/prometheus-crypto
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/jordyv/prometheus-crypto
RUN dep ensure
RUN go build -o prometheus-crypto main.go


FROM golang:alpine

COPY --from=builder /go/src/github.com/jordyv/prometheus-crypto/prometheus-crypto /usr/local/bin/prometheus-crypto
EXPOSE 8888
ENTRYPOINT [ "/usr/local/bin/prometheus-crypto", "-listen", ":8888" ]
