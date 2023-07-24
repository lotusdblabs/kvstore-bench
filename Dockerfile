FROM golang:1.19-alpine as builder

workdir /rosedblabs

# If you encounter some issues when pulling modules, \
# you can try to use GOPROXY, especially in China.
ENV GOPROXY=https://goproxy.cn

COPY . .
RUN pwd
RUN cd /rosedblabs/cmd/kv-bench && go build . && mv kv-bench ../../


FROM alpine:latest

WORKDIR /rosedblabs
COPY --from=builder /rosedblabs/kv-bench /bin/kv-bench

ENTRYPOINT ["/bin/kv-bench"]
CMD ["--help"]

# Usage:

# build kv-bench image
# docker build -t kv-bench .

# print help
# docker run --rm --name kv-bench -it kv-bench -h