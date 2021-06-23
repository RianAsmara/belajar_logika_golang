## import base image 
FROM golang

## Create /app directory that contain out application
RUN mkdir /app

# copy every file to /app folder
ADD . /app

# specify execute program in /app folder

WORKDIR /app

# run go build to execute our binary
RUN go build -o main .

# created executable
CMD ["/app/main"]