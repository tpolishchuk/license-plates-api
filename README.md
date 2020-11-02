# License plates API

API, that allows you to receive data about license plate region, city and state by its abbreviation.

## Requirements

The latest version of Docker platform 


## Installation

- Clone the repository
- Navigate into the project root folder
- Run

```
docker-compose up
```

- Make GET request to [http://localhost:8088/plates](http://localhost:8088/plates) after you see the following output

```
license-plates-api    | [GIN-debug] Listening and serving HTTP on :8088
```

## Endpoints

- GET /plates --> lists all available plates for all countries
- GET /plates/:id --> shows plate information by its id
- POST /plates --> is used for a new plate creation. Schema:
```
{
    "abbreviation": "string",
    "city_or_region": "string",
    "state": "string",
    "country": "string"
}
```
- PUT /plates/:id --> is used for an existing plate amendment. Schema is the same as for the plate creation call
- DELETE /plates/:id --> deleted plate by a given id

## Project state

Work in progress. The project is used mostly to get the fist experience with Golang

## TODOs

- Add filtering by country
- Add search by abbreviation with all possible countries in the listing
- Add migrations with predefined DB seeds
- Add end-to-end tests for Travis CI run
