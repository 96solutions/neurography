FROM docker.io/library/golang:1.22.1 as builder

WORKDIR /app/src

RUN groupadd -g 1000 builder && useradd -u 1000 -ms /bin/bash -g builder builder
RUN chown -R builder:builder /app/src

ADD ./ /app/src/

RUN mkdir /var/www
RUN chown -R builder:builder /var/www

USER builder

RUN go install go.uber.org/mock/mockgen@v0.4.0
RUN go generate ./...
RUN ls -lha &&  go mod vendor && \
    CGO_ENABLED=0 GOOS=linux GOFLAGS=-buildvcs=false go build -a -o /app/src/bin/neurography .


FROM scratch

COPY --from=builder /app/src/bin/neurography /app/

EXPOSE 8080

CMD ["/app/neurography"]
