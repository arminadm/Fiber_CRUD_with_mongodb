FROM golang:latest

WORKDIR /app

# Required packages my not be avaiable in Iran, changing the Goprivate can handle the problem
ENV GOPRIVATE "github.com"

COPY go.mod go.sum ./

RUN go mod tidy 

# Copy the source code to the working directory
COPY . .

WORKDIR /app/src

RUN go build -o Fiber_CRUD

# Run the Go binary
CMD ["./Fiber_CRUD"]
