# build stage
FROM golang:1.13.8-alpine AS build-env
RUN apk add make git bash build-base
ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"
ADD ./ /go/src/github.com/ohsu-comp-bio/funnel
RUN cd /go/src/github.com/ohsu-comp-bio/funnel && make build

# final stage
FROM debian
WORKDIR /opt/funnel
VOLUME /opt/funnel/funnel-work-dir
EXPOSE 8000 9090
ENV PATH="/app:${PATH}"
COPY --from=build-env  /go/src/github.com/ohsu-comp-bio/funnel/funnel /app/
RUN apt-get update && apt-get -y install --no-install-recommends python3-pip && apt-get clean && pip install htsget && rm -rf ~/.cache
ENTRYPOINT ["/app/funnel"]
