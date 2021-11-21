# Books Data API (Scrapping)
> Scrapping from naiin.com


## Installation
Runing following containers :
<li>Golang Container</li>

```sh
docker-compose up
```

## Load testing
Run load testing with K6 with test file in `test/test.js`
```sh
docker-compose run --rm k6 run /scripts/test.js
```

