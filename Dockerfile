FROM ubuntu

COPY go-simple-api-gateway-linux-amd64 /gateway

EXPOSE 1329

CMD ["/gateway"]
