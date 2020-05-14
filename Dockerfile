FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go mod init demo
RUN go build -o main . 
CMD ["/app/main"]