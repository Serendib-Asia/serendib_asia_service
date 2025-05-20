FROM golang:1.22-bullseye as build 

ARG SKIP_UNIT_TEST=false
ENV SKIP_UNIT_TEST=${SKIP_UNIT_TEST}

ENV GOLANG_VERSION="1.22.0"

RUN apt-get clean
RUN apt-get update && apt-get install -y ca-certificates openssl git tzdata
ARG cert_location=/usr/local/share/ca-certificates

RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.ctr
RUN openssl s_client -showcerts -connect gitlab.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/gitlab.ctr
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/proxy.golang.ctr
RUN openssl s_client -showcerts -connect gopkg.in:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/gopkg.ctr
RUN openssl s_client -showcerts -connect storage.googleapis.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/storage.googleapis.ctr
RUN openssl s_client -showcerts -connect sum.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/sum.golang.ctr

RUN update-ca-certificates
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/russross/blackfriday-tool@latest

ADD . /app
WORKDIR /app
RUN make build

FROM gcr.io/distroless/base-debian11
COPY --from=build /app/build /

ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD
ARG IS_CLOUD_SQL
ARG SRV_LISTEN_PORT
ARG AUTH_SERVER_BASE_URL
ARG TOKEN_VALIDATE_ENDPOINT
ARG SEC_SRV_GET_ALL_USER_DETAILS_ENDPOINT
ARG DEFAULT_CLIENT
ARG CLIENT_DB_HOST
ARG CLIENT_DB_PORT
ARG CLIENT_DB_NAME
ARG CLIENT_DB_USER
ARG CLIENT_DB_PASSWORD

ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV IS_CLOUD_SQL=${IS_CLOUD_SQL}
ENV SRV_LISTEN_PORT=${SRV_LISTEN_PORT}
ENV AUTH_SERVER_BASE_URL=${AUTH_SERVER_BASE_URL}
ENV TOKEN_VALIDATE_ENDPOINT=${TOKEN_VALIDATE_ENDPOINT}
ENV SEC_SRV_GET_ALL_USER_DETAILS_ENDPOINT=${SEC_SRV_GET_ALL_USER_DETAILS_ENDPOINT}
ENV DEFAULT_CLIENT=${DEFAULT_CLIENT}
ENV CLIENT_DB_HOST=${CLIENT_DB_HOST}
ENV CLIENT_DB_PORT=${CLIENT_DB_PORT}
ENV CLIENT_DB_NAME=${CLIENT_DB_NAME}
ENV CLIENT_DB_USER=${CLIENT_DB_USER}
ENV CLIENT_DB_PASSWORD=${CLIENT_DB_PASSWORD}

CMD ["/icx_dashboard"]
