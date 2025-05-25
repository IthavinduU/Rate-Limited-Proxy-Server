# Rate-Limited Proxy Server

## Overview

This project implements an HTTP proxy server in Go that forwards requests to a target service while enforcing user-specific rate limits using a token bucket algorithm. The server includes:

- Rate limiting middleware
- Logging middleware
- Prometheus metrics for monitoring
- Integration with Grafana for visualization

---

## Features

- HTTP reverse proxy forwarding to a configurable target  
- Token bucket rate limiter to control request rates per user/IP  
- Structured logging of incoming requests and response statuses  
- Metrics endpoint (`/metrics`) compatible with Prometheus  
- Dockerized setup with `docker-compose` including Prometheus and Grafana  

---

