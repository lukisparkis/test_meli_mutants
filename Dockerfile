FROM golang:1.18 AS base

WORKDIR /go/src/gitlab.com/meli

FROM base AS dev
# application and watches for changes - ONLY development
ARG FRESHER_VERSION=1.2.1

RUN wget -c https://github.com/roger-russel/fresher/releases/download/v${FRESHER_VERSION}/fresher_${FRESHER_VERSION}_Linux_x86_64.tar.gz \
      && tar xzf fresher_${FRESHER_VERSION}_Linux_x86_64.tar.gz -C /go/bin/ \
      && rm -f fresher_*tar.gz

CMD ["fresher", "-c", "./configs/fresher.yaml"]

EXPOSE 33333

##
## Build
##

FROM base AS build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/api .

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go/bin/api /go/bin/api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go/bin/api"]