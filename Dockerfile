FROM golang:alpine

WORKDIR /build

# Copy and install dependencies 
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the app into the container
COPY . .

# Build
RUN go build -o main .

# Navigate where the binary is built
WORKDIR /dist

RUN cp /build/main .

EXPOSE 8080

CMD ["dist/main"]