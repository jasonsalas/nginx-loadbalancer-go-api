# build container
FROM golang:1.14-alpine AS builder

# update the repository and install Git
RUN apk update && apk upgrade && \
	apk add --no-cache git

WORKDIR /tmp/app
COPY . .
RUN GOOS=linux go build -o ./out/api .

# build container
FROM alpine:latest
RUN apk add ca-certificates
COPY --from=builder /tmp/app/out/api /app/api
WORKDIR /app
EXPOSE 5000

CMD [ "./api" ]
