FROM golang:1.19 AS builder
ARG PANDOC_VER=2.19.2

WORKDIR /tools
RUN mkdir pandoc && \
    wget https://github.com/jgm/pandoc/releases/download/$PANDOC_VER/pandoc-$PANDOC_VER-linux-amd64.tar.gz && \
    tar xvzf pandoc-$PANDOC_VER-linux-amd64.tar.gz --strip-components 1 -C ./pandoc && \
    rm -rf pandoc-$PANDOC_VER-linux-amd64.tar.gz
 
# Copy the code from the host and compile it
WORKDIR /go/src/github.com/nurdsoft/redesign-grp-trust-portal-api
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix nocgo -o /app .

FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates

COPY --from=builder /tools/pandoc/bin /bin
COPY --from=builder /app ./
COPY --from=builder /go/src/github.com/nurdsoft/redesign-grp-trust-portal-api/config.yaml ./config.yaml
COPY --from=builder /go/src/github.com/nurdsoft/redesign-grp-trust-portal-api/migrations ./migrations

EXPOSE 8080

CMD ./app migrate; ./app api
