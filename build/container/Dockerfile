FROM alpine:3.17 AS build

RUN apk add --no-cache go
WORKDIR /mnt
COPY . .
RUN go build -ldflags '-s -w' -o skeletor cmd/main.go

FROM alpine:3.17

COPY --from=build /mnt/skeletor /usr/local/bin/skeletor
USER nobody:nobody
ENTRYPOINT ["skeletor"]
