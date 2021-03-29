FROM --platform=${BUILDPLATFORM} golang:1.16-alpine as base
WORKDIR /src
COPY go.mod /src/
COPY main.go /src/
RUN go mod download > /dev/null

FROM base AS build
ARG GOARCH
ARG GOOS
ARG CGO_ENABLED=0
ARG BUILD_TIME=$(date)
ARG VERSION
RUN --mount=type=cache,target=/root/.cache/go-build GO111MODULE=on CGO_ENABLED=$CGO_ENABLED GOARCH=$GOARCH GOOS=$GOOS \
    go build -o target/mariadb_test -tags $VERSION -ldflags "-s -w -X main.Version=$VERSION -X main.BuildTime=$BUILD_TIME" .

FROM alpine
ARG MARIADB_HOST
ARG MARIADB_ROOT_USER
ARG MARIADB_ROOT_PASSWORD
ARG MARIADB_CLIENT_PORT
COPY --from=build /pkg/src/target/mariadb_test .
RUN uname -a
COPY passwd.minimal /etc/passwd
RUN chown tester mariadb_test
USER tester
CMD ["sh", "-c", "tail -f /dev/null"]