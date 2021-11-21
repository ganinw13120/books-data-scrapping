# Books Data API (Scrapping)
> Scrapping from <a href='https://naiin.com'>Naiin</a>

> Practice projects for implementing redis, with golang. Applying with hexagonal architecture.

## Installation
Runing following containers :
<li>Golang Container (Server) (running Dockerfile)</li>
<li>Redis Container (Caching) </li>
<li>K6 Container (Load Testing) </li>
<li>Influxdb Container (Storing load testing result) </li>
<li>Grafana Container (Dashboard for load testing result (include data from influxdb)) </li>

Docker compose file located at `./docker-compose.yml`

```sh
docker-compose up
```

## Load testing
Run load testing with K6 with test file in `test/test.js`
```sh
docker-compose run --rm k6 run /scripts/test.js
```

