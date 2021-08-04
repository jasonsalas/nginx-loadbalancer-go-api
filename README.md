# Using NGINX as a load balancer for a Go-based API in Docker containers

## h/t to [@zerefwayne](https://github.com/zerefwayne) for the [idea](https://github.com/zerefwayne/load-balancing-go-api-nginx)🔥🔥🔥!

***
**PURPOSE**

This repo demonstrates how to use [NGINX as a reverse proxy/load balancer](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/nginx/nginx.conf) serving a web application based on [an API](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/main.go) written in Go, and [running in Docker containers](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/docker-compose.yml) 🔀.

(It also demos how to [use GitHub Actions for scalable CI/CD](https://github.com/jasonsalas/nginx-loadbalancer-go-api/actions/runs/1096460347)) :chart_with_upwards_trend:.

***
**DOCUMENTATION**

You can find the full explanation for the original repo [on codeburst](https://codeburst.io/load-balancing-go-api-with-docker-nginx-digital-ocean-d7f05f7c9b31).


***
**USAGE**

Spin-up the containers hosting the API server - there are [5 by default](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/docker-compose.yml#L11), but can be scaled up/down: 
- `docker-compose up -d --build`

...then make an HTTP request against the load balancer: 
- `curl -X GET localhost/hi`


***
**DOCKER HUB IMAGE**

You can spin-up a pre-built image [on Docker Hub](https://hub.docker.com/repository/docker/jasonsalas/nginx-loadbalancer-go-api) by pulling the image to your server:
- `docker pull jasonsalas/nginx-loadbalancer-go-api:latest`

...or, you can also run this load balancer application on your server as a pre-built disk image by running the following commands in a terminal:
- `docker-compose -f docker-compose-for-docker-hub.yml up -d`

Then, you'll need to verify which randomized port the API server your server is listening on (it'll be bound to port 5000 on the image):
- `docker-compose ps -a`

For example, let's assume the port is `12345`. Use this port to make a HTTP request to the load balancer's endpoint:

`curl -X GET localhost:12345/hi`
