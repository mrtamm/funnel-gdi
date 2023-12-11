# build stage
FROM golang:1.21.4-alpine AS build-env
RUN apk add make git bash build-base
ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/github.com/ohsu-comp-bio/funnel
COPY go.* .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build make build

# final stage
FROM debian
RUN apt-get update && apt-get install -y --no-install-recommends curl python3-pip && pip3 install htsget --break-system-packages && rm -rf ~/.cache
WORKDIR /opt/funnel
VOLUME /opt/funnel/funnel-work-dir
EXPOSE 8000 9090
ENV PATH="/app:${PATH}"
COPY --from=build-env  /go/src/github.com/ohsu-comp-bio/funnel/funnel /app/

ENTRYPOINT [ "/app/funnel" ]
