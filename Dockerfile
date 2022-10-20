## specify the base image for go app
FROM golang:1.19.2-alpine3.16

## create /app directory within the image that will hold app source files
RUN mkdir /app

## copy all files in the root directory into /app directory
ADD . /app

## set /app as working directory for any further commands
WORKDIR /app

## run go build to compile the binary executable of go app
RUN go build -o main .

## command to start the newly created binary executable
CMD ["/app/main"]