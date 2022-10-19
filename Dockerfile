FROM node:16.18.0-alpine3.15 as reactBuild
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
# ENV PATH /app/node_modules/.bin:$PATH

COPY frontend/package.json ./
COPY frontend/package-lock.json ./
RUN npm install

COPY frontend/ .
RUN npm run build 

FROM golang:1.17-alpine as golangBuild
WORKDIR /app
COPY backend ./
RUN go mod tidy && \
    go mod download && \
    go build -o main

FROM alpine:3.15 as final
WORKDIR /app
ENV PROD=true
COPY --from=golangBuild /app/main ./main
COPY --from=reactBuild /app/build/ ./build
EXPOSE 8080

CMD ["/app/main"]

# build command: docker build -t test:0.0.1 .
# run command: docker run --network="host" test:0.0.1
# remove dangling images: docker rmi $(docker images --filter "dangling=true" -q --no-trunc)