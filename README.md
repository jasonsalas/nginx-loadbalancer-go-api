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

You can spin-up a pre-built image [on Docker Hub](https://hub.docker.com/repository/docker/jasonsalas/nginx-loadbalancer-go-api) without needing to build from source using this command from a terminal: 
- `docker-compose -f docker-compose-for-docker-hub.yml up -d` 

_(Notice that the [modified docker-compose file](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/docker-compose-for-docker-hub.yml#L5) uses the image from the remote repo instead of building it locally.)_

Alternately, you can always pull the remote image to your server: 
- `docker pull jasonsalas/nginx-loadbalancer-go-api:1.0` 