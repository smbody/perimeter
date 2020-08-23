# Production environment
FROM golang:alpine AS perimeter_base
RUN apk update && apk upgrade && \
apk add --no-cache bash git openssh
WORKDIR /perimeter
