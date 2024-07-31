FROM golang:1.22-alpine as final

ENV GOPATH=/

COPY ./ ./

RUN go mod download

RUN apk add curl

RUN curl -sSf https://atlasgo.sh | sh

RUN go build -o /bin/server ./server/main.go

# Expose the port that the application listens on.
EXPOSE 8081

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]
