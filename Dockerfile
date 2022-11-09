FROM golang:1.18.8-alpine AS build

RUN apk update && \
    apk upgrade && \
    apk add build-base

WORKDIR /tem/app

COPY API .

RUN GOOS=linux go build -o ./out/api .

#################################################3

FROM  alpine:latest

RUN apk add ca-certificates 

COPY --from=build /tem/app/out/api /app/api

WORKDIR /app

EXPOSE 6767 

CMD ["./api"]


