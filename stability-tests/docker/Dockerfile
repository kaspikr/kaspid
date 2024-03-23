ARG KASPID_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPID_IMAGE} as kaspid
FROM ${KASPAMINER_IMAGE} as kaspiminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/kaspikr/kaspid

WORKDIR /go/src/github.com/kaspikr/kaspid

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=kaspid /app/ /app/
COPY --from=kaspiminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/kaspikr/kaspid/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]