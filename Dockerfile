#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------
FROM golang:1.15.4-alpine AS development

# Force the go compiler to use modules
ENV GO111MODULE=on
#ENV GOPROXY=http://172.17.0.1:3333

# Update OS package and install Git
RUN apk update && \
    apk add git && \
    apk add build-base && \
    apk add openssh && \
    apk add --no-cache bash && \
    apk add git mercurial && \
    apk add linux-headers

# Set working directory
WORKDIR /go/src/github.com/NoobMM/golang

# Install Fresh for local development
RUN go get github.com/pilu/fresh

# Install go tool for convert go test output to junit xml
RUN go get -u github.com/jstemmer/go-junit-report &&\
    go get github.com/axw/gocov &&\
    go get github.com/AlekSi/gocov-xml

# Install wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /usr/local/bin/wait-for &&\
    chmod +x /usr/local/bin/wait-for

# Copy Go dependency file
COPY go.mod go.mod
COPY go.sum go.sum

# Download dependency
RUN go mod download

# Copy src
COPY app app

# Use CMD instead of RUN to allow command overwritability
CMD cd app && fresh


# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------
FROM golang:1.15.4-alpine AS build

ENV GO111MODULE=on

# Set working directory
WORKDIR /go/src/github.com/NoobMM/golang

# Copy stuff from development stage
COPY --from=development /go/src/github.com/NoobMM/golang .

# Update OS package and install Git
RUN apk update && apk add git mercurial bzr && apk add build-base

# Build the binary
RUN cd app && go build -o /go/bin/server.bin

# ------------------------------------------------------------------------------
# Application image
# ------------------------------------------------------------------------------
FROM golang:1.15.4-alpine

RUN apk add --no-cache tini tzdata
RUN addgroup -g 211000 -S appgroup && adduser -u 211000 -S appuser -G appgroup

# Set working directory
WORKDIR /app

#Get artifact from builder stage
COPY --from=build /go/bin/server.bin /app/server.bin

# Copy SQL stuff
COPY --from=development /go/src/github.com/NoobMM/golang/app/migrations /app/migrations

# Set Docker's entry point commands
RUN chown -R appuser:appgroup /app
USER appuser
ENTRYPOINT ["/sbin/tini","-sg","--","/app/server.bin"]
