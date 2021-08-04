# Using NGINX as a load balancer for a Go-based API in Docker containers

## h/t to [@zerefwayne](https://github.com/zerefwayne) for the [idea](https://github.com/zerefwayne/load-balancing-go-api-nginx)🔥🔥🔥!

***
**PURPOSE**
This repo demonstrates how to use [NGINX as a reverse proxy/load balancer](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/nginx/nginx.conf) serving a web application based on [an API](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/main.go) written in Go, and [running in Docker containers](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/docker-compose.yml) 🔀.

(It's also a staging area to demo how to [use GitHub Actions for scalable CI/CD](https://github.com/jasonsalas/nginx-loadbalancer-go-api/actions/runs/1096460347)) :chart_with_upwards_trend:.

***
**DOCUMENTATION**
You can find the full explanation [on codeburst](https://codeburst.io/load-balancing-go-api-with-docker-nginx-digital-ocean-d7f05f7c9b31).


***
**USAGE**

- Spin-up the containers hosting the API server (there are 5 by default but [this is a setting that can be scaled up/down](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/docker-compose.yml#L11)): `docker-compose up -d --build`
- Make an HTTP request against the load balancer: `curl -X GET localhost/hi`


***
**DOCKER IMAGE**
You can spin-up a pre-built image [on Docker Hub](https://hub.docker.com/repository/docker/jasonsalas/nginx-loadbalancer-go-api) using this command from a terminal: 
- `docker pull jasonsalas/nginx-loadbalancer-go-api:1.0` 

...and then
- `docker run --name loadbalancer --rm` 