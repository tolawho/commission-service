FROM 461429446948.dkr.ecr.ap-southeast-1.amazonaws.com/medici-base:medici-base-go-1-19 as build

WORKDIR /go/src/app
COPY go.* ./
RUN go mod download

COPY . .
COPY scripts/env/.env.dev.5009 .env

RUN go mod tidy
RUN go build -o godocker

EXPOSE 8010

ENTRYPOINT [ "/go/src/app/godocker" ]
