FROM golang:1.19-buster AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -buildmode=c-shared -o out_exec.so out_exec.go newrelic.go

FROM fluent/fluent-bit
COPY --from=builder /app/out_exec.so ./
COPY docker/default.conf /fluent-bit/etc/fluent-bit.conf
COPY docker/plugins.conf /fluent-bit/etc/plugins.conf
COPY docker/output.conf /fluent-bit/etc/conf.d/output.conf

EXPOSE 2020
EXPOSE 24224
