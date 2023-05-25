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
RUN apt-get update && apt-get install -y --no-install-recommends podman curl python3-pip && pip3 install htsget && rm -rf ~/.cache
ADD https://raw.githubusercontent.com/containers/libpod/master/contrib/podmanimage/stable/containers.conf /etc/containers/containers.conf
RUN ln -s /usr/bin/podman /usr/bin/docker
