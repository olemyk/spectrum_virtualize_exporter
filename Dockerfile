#FROM quay.io/prometheus/golang-builder:1.19-base as builder
FROM quay.io/prometheus/golang-builder:latest as builder

WORKDIR /build

COPY . .
#Usage: builder.sh [args]
#  -i,--import-path arg  : Go import path of the project
#  -p,--platforms arg    : List of platforms (GOOS/GOARCH) to build separated by a space
#  -T,--tests            : Go run tests then exit
RUN go install -v -t --platform linux/arm64 -d ./...
RUN CGO_ENABLED=0 go build -o main .

FROM scratch
WORKDIR /opt/spectrum_virtualize_exporter

COPY --from=builder /build/main .

EXPOSE 9747
CMD ["./main", "-auth-file", "~/spectrum-monitor.yaml", "-extra-ca-cert", "~/tls.crt"]
