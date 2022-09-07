FROM golang:1.19-buster AS builder

ARG SERVICE_NAME
ARG SERVICE_EXEC
ARG WDIR=/go/src/app

WORKDIR ${WDIR}
COPY proto proto
COPY shared_pkg shared_pkg

WORKDIR ${WDIR}/services/$SERVICE_NAME
COPY services/${SERVICE_NAME}/go.mod go.mod
COPY services/${SERVICE_NAME}/go.sum go.sum
RUN go mod download

COPY services/${SERVICE_NAME} .

RUN go build -o /go/bin/app cmd/${SERVICE_EXEC}/main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /app

COPY --from=builder /go/bin/app app
CMD [ "./app" ]