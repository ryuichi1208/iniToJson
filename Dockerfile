FROM golang:1.12.7-alpine AS builder
LABEL maintainer="ryuichi1208 <ryucrosskey@gmail.com>"
WORKDIR /app
RUN apk add --no-cache make && \
    rm -rf /var/cache/apk/* && \
    mkdir src
COPY ./src ./src
#COPY ./test.ini .
COPY ./Makefile .
RUN make

#COPY ./test.ini .
#ENTRYPOINT ["./main"]

# ==========

FROM busybox
WORKDIR /app
COPY --from=builder /app/main .
COPY ./test.ini .
ENTRYPOINT ["./main"]
CMD [""]
