# Simple Telegram SOCKS5 Proxy

Small and fast proxy. It written in Go and supports username/password authentication

Telegram was blocked in Russia and to bypass those blocks, you can use this free proxy server.

To compile proxy clone repo from github and run:

> cd src && go get && go build tgproxy && go build -o tgproxy

after successful build run it with:

> tgproxy --proxyname=<username> --proxypasswd=<password>

To build docker container, clone repo and run:

> docker build -t tgproxy

after successful build run it with:

> docker run -p 6785:6785 --name tgproxy --detach tgproxy --proxyname=<username> --proxypasswd=<password>

Also you can use prebuilded docker image:
https://hub.docker.com/r/ikozinov/tgproxy/

Just run it with the following command:

> docker run -p 6785:6785 --name tgproxy --detach ikozinov/tgproxy --proxyname=<username> --proxypasswd=<password>