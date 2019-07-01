FROM golang:alpine as BUILDER
WORKDIR /home/tmp
RUN apk --no-cache add ca-certificates git
RUN adduser -D -g '' appuser
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./hello-world ./cmd/hello-world/main.go


FROM busybox
WORKDIR /home/tmp
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /etc/passwd /etc/passwd
COPY --from=BUILDER /home/tmp/hello-world ./hello-world
USER appuser
EXPOSE 8000
ENTRYPOINT ["./hello-world"]
