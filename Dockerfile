# dynamic config
ARG             BUILD_DATE
ARG             VCS_REF
ARG             VERSION

# build
FROM            golang:1.22.1-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make g++ alsa-lib-dev
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/midcat
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM alpine:3.19
LABEL           org.label-schema.build-date=$BUILD_DATE \
                org.label-schema.name="midcat" \
                org.label-schema.description="" \
                org.label-schema.url="https://moul.io/midcat/" \
                org.label-schema.vcs-ref=$VCS_REF \
                org.label-schema.vcs-url="https://github.com/moul/midcat" \
                org.label-schema.vendor="Manfred Touron" \
                org.label-schema.version=$VERSION \
                org.label-schema.schema-version="1.0" \
                org.label-schema.cmd="docker run -i -t --rm moul/midcat" \
                org.label-schema.help="docker exec -it $CONTAINER midcat --help"
COPY            --from=builder /go/bin/midcat /bin/
RUN             apk add --no-cache gcc alsa-lib
ENTRYPOINT      ["/bin/midcat"]
#CMD             []
