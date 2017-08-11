# GO-SIMPLE-API-GATEWAY

![status](https://ci.fornever.org/job/go-simple-api-gateway/badge/icon)

A simple API gateway written by golang.

Support for authenticate and authorization, and web applications will be protected after the gateway.

in development now.

documents will be wrote later.

## ARCH

![arch](https://res.cloudinary.com/digf90pwi/image/upload/v1502367434/go-simple-api-gatway_1_grgl5o.png)

## CONFIGURATION

You could use **cli option** or **environment varibles** to config your api gateway

```bash
./go-simple-api-gateway --help
Options:

  -h, --help                                display help information
  -c, --*conn[=$GATEWAY_CONN_STR]          *mysql connection str
  -l, --*listen[=$GATEWAY_LS]              *gateway listen host and port
  -r, --*resource[=$GATEWAY_RESOURCE_URL]  *gateway resource url

```

* -c --conn **GATEWAY_CONN_STR**, mysql connection string, format is *user:pass@tcp(domain:port)/dbname*

* -l --listen **GATEWAY_LS**, gateway listen addr, format is *host:port*, example: *0.0.0.0:1329*

* -r --resource **GATEWAY_RESOURCE_URL**, gateway protect target, the resource server, could be a api server, format is *http://host:port*

## DOCKER

you could find docker image from [here](https://hub.docker.com/r/theosun/go-simple-api-gateway/)

example:

```bash
docker run -d --restart=always -p 11329:1329 -e GATEWAY_CONN_STR='user:pass@tcp(mysql:3306)/db_name' -e GATEWAY_LS=':1329' -e GATEWAY_RESOURCE_URL='http://api:1323' --link mariadb:mysql --link citi_api:api --name citi_gateway theosun/go-simple-api-gateway

```

## DOWNLOAD

You could download the latest build binaries from [here](https://download.fornever.org/go-simple-api-gateway/latest/)