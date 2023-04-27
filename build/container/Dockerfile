FROM alpine:3.17 AS build

RUN apk add --no-cache go
WORKDIR /mnt
COPY . .
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}"  -o skeletor cmd/main.go
#-ldflags '-s -w'

FROM alpine:3.17

COPY --from=build /mnt/skeletor /usr/local/bin/skeletor
EXPOSE 8080
USER nobody:nobody
WORKDIR /usr/local/bin
ENTRYPOINT ["skeletor"]
