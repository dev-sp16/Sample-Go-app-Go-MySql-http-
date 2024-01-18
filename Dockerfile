# docker file for building Go application
FROM ubuntu:latest

# Install dependencies
RUN sudo apt install -y git go wget

COPY . /app

WORKDIR /app

# Build the application
RUN go build -o main .

CMD [ "main" ]