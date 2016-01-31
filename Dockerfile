FROM golang:1.5.3
COPY . src/go_web_server
WORKDIR src/go_web_server
RUN go get && go build
ENV PORT 9999
CMD ["./go_web_server"]
