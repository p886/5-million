# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o goapp

# final stage
FROM alpine:3.11.2
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --chown=appuser:appgroup --from=build-env /src/goapp /app/
USER appuser
ENTRYPOINT [ "./goapp" ]
CMD ["server"]
