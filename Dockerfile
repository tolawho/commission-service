# FROM 461429446948.dkr.ecr.ap-southeast-1.amazonaws.com/medici-base:medici-base-go-1-19 as base
# WORKDIR /app/go
# COPY . .
# RUN go mod download

# FROM base as image-dev
# RUN go get github.com/cosmtrek/air
# EXPOSE 3000
# CMD $(go env GOPATH)/bin/air

# FROM base as builder
# RUN mkdir dist
# RUN go build -o dist/web .

# FROM 461429446948.dkr.ecr.ap-southeast-1.amazonaws.com/medici-base:medici-base-go-1-19 as image-prod
# WORKDIR /app
# COPY --from=builder ./app/go/dist/ ./
# EXPOSE 3000
# CMD  ./web

FROM 461429446948.dkr.ecr.ap-southeast-1.amazonaws.com/medici-base:medici-base-go-1-19 as build

WORKDIR /go/src/app
COPY go.* ./
RUN go mod download
COPY . .

RUN go get -u
RUN GOOS=linux GOARCH=amd64 go build -a -v -tags musl

FROM 461429446948.dkr.ecr.ap-southeast-1.amazonaws.com/medici-base:medici-base-go-1-19
RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/* /tmp/*
# RUN apk add nodejs-current
# RUN apk add nodejs-npm
# RUN npm install pm2 -g

WORKDIR /
COPY --from=build /go/src/app/commission-serivce /commission-serivce

EXPOSE 8010

# CMD [ "pm2-runtime", "start", "go run main.go"]

CMD [ "./commission-serivce"]

# ENTRYPOINT [ "/commission-serivce" ]
