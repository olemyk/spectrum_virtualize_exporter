# FROM quay.io/prometheus/golang-builder:1.19-base as builder

# Replaced the builder image to support serveral architectures. did not get the golang-builder to work with Apple M1 arm64v8.
# amd64, arm32v5, arm32v6, arm32v7, arm64v8, i386, mips64le, ppc64le, s390x, windows-amd64
FROM golang:alpine AS builder

# Using a trusted docker image like golang:alpine is not always enough for security. People can intercept your request to provide a modified docker image. The best solution : using digest  # golang alpine 3.16
#FROM golang@sha256:5cc70c23fa75dd6d5b1ed628891933fbf984385b0492f80c81b0920ffc4c20db as builder

LABEL stage=gobuilder

# Create appuser
ENV USER=appuser
ENV UID=10001
#See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /build

COPY . .

RUN go get -v -t -d ./...

# Using go mod with go 1.11
RUN go mod download
RUN go mod verify
# Build the binary
RUN CGO_ENABLED=0 go build -o main .

FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /opt/spectrum_virtualize_exporter

# Copy our static executable
COPY --from=builder /build/main .

# Use an unprivileged user.
USER appuser:appuser

# Port on which the service will be exposed.
EXPOSE 9747
CMD ["./main", "-auth-file", "/config/virtualize-monitor.yaml", "-extra-ca-cert", "~/tls.crt"]