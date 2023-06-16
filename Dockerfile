FROM golang:1.20-bullseye as build

ARG FDB_VERSION='7.1.33'
ARG FDB_CLIENTS_DEB_SHA256_SUM='d73002bc796de7ce0158f1eb64fab1f8e90bada2edd90599d9b40754220150af'
RUN wget "https://github.com/apple/foundationdb/releases/download/${FDB_VERSION}/foundationdb-clients_${FDB_VERSION}-1_amd64.deb" && \
    echo "${FDB_CLIENTS_DEB_SHA256_SUM}  foundationdb-clients_${FDB_VERSION}-1_amd64.deb" | sha256sum --check && \
    dpkg --force-all -i foundationdb-clients_${FDB_VERSION}-1_amd64.deb

WORKDIR /go/src/fdbmeter
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/fdbmeter ./cmd/fdbmeter

FROM gcr.io/distroless/base-debian11

COPY --from=build /go/bin/fdbmeter /usr/bin/
COPY --from=build /usr/lib/libfdb_c.so /usr/lib/libfdb_c.so
COPY --from=build /usr/include/foundationdb /usr/include/foundationdb

ENTRYPOINT ["/usr/bin/fdbmeter"]