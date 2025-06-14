# Weather API Go net/http

This is my golang solution to https://roadmap.sh/projects/weather-api-wrapper-service, i tried to minimize dependencies using only the `net/http` standard library for both server and client code, i only use the redis package and godotenv to load .env files.

## Features

- *Weather Data Fetching*: Retrieves weather data for a specified location using the Visual Crossing Weather API.
- *Caching with Redis*: Caches the weather data in Redis to reduce API calls and improve performance.
- *Rate Limiting with custom middleware*: Limits the number of API requests a user can make to prevent abuse.


## Running application

### 1. Clone the repo

```bash
git clone https://github.com/umdalecs/weather-api
cd weather-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Environment variables

copy the .env example and populate with your data

```bash
cp .env.example .env
```

### 4. Build and run the application

```bash
go build -o out/
./out/weather-api
```

## Usage

I put the `requests.http` file with an example fetching the London weather data but you're free to use any other tool out there 