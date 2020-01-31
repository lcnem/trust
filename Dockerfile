FROM golang:alpine AS build-env

# Set working directory for the build
WORKDIR /go/src/github.com/lcnem/trust

# Add source files
COPY . .

RUN go install ./cmd/trustd
RUN go install ./cmd/trustcli

# Final image
FROM ubuntu:latest

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

COPY scripts/genesis.json genesis.json
COPY scripts/init.sh init.sh

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/trustd /usr/bin/trustd
COPY --from=build-env /go/bin/trustcli /usr/bin/trustcli

# Run trustd by default, omit entrypoint to ease using container with trustcli
CMD ["trustd"]