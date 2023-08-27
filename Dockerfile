FROM golang:1.18 AS build

RUN mkdir -p /opt/build

WORKDIR /opt/build

# copy only necessary files
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the files
COPY . .

# do the build
RUN CGO_ENABLED=0 GOOS=LINUX GOARCH=amd64 go build -o main main.go

FROM gcr.io/distroless/static
USER nobody:nobody
WORKDIR /
COPY --from=build /opt/build/main .
ENV GIN_MODE=release
CMD ./main
