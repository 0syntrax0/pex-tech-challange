# Pex Assignment

## How to run

While in the same directory run `docker-compose up` to build the `go` server up and running.

## Endpoints

The following endpoints are working on `localhost:8080`.

- `GET("/previous")` returns the **previous** number in the current sequence
- `GET("/current")` returns the **current** number in the current sequence and moves up to the next set
- `GET("/next")` returns the **next** number in the current sequence

## List of requirements

Requirements:

- [x] The API must be able to handle high throughput (~1k requests per second)
- [x] The API should also be able to recover and restart if it unexpectedly crashes
- [x] Assume that the API will be running on a small machine with 1 CPU and 512MB of RAM
- [x] You may use any programming language/framework of your choice
