FROM golang:alpine AS builder

# docker run -p 1025:1025 -td b1320fb8321a <image>
# docker exec -it <container> sh

RUN addgroup -S zk_auth_service_server_group && \
    adduser -S zk_auth_service_server_user -G zk_auth_service_server_group 

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/zk_auth_lib
COPY ./zk_auth_lib .

WORKDIR $GOPATH/src/zk_auth_service
COPY ./zk_auth_service .

RUN go mod init zk_auth_service
RUN go mod edit -replace zk_auth_service/lib/zk_auth_lib=../zk_auth_lib/go
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/zk_auth_service


FROM golang:alpine

COPY --from=builder /go/bin/zk_auth_service /go/bin/zk_auth_service
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd

# /usr/local/src doesn't exist in Alpine
WORKDIR /usr/local/src
COPY ./scripts/zk-auth-service-entrypoint.sh .
RUN apk update && apk add --no-cache bash postgresql-client curl && rm -rf /var/cache/apk/*

USER zk_auth_service_server_user

ENTRYPOINT ["sh", "/usr/local/src/zk-auth-service-entrypoint.sh"]