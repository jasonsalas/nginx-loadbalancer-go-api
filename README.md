# Using NGINX as a load balancer for a Go-based API in Docker containers

## h/t to [@zerefwayne](https://github.com/zerefwayne) for the [idea](https://github.com/zerefwayne/load-balancing-go-api-nginx)🔥🔥🔥!

***

This repo demonstrates how to use [NGINX as a reverse proxy/load balancer](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/nginx/nginx.conf) serving a web application based on [an API](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/main.go) written in Go, and [running in Docker containers](https://github.com/jasonsalas/nginx-loadbalancer-go-api/blob/main/docker-compose.yml) 🔀.

(It's also a staging area to demo how to [use GitHub Actions for scalable CI/CD](https://github.com/jasonsalas/nginx-loadbalancer-go-api/actions/runs/1096460347)) :chart_with_upwards_trend:.