# Golang Weather API wrapper service

This is my golang solution to https://roadmap.sh/projects/weather-api-wrapper-service.

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
go build -o out/ .
./out/weather-api
```

## Usage

```bash
curl http://localhost:8080/api/v1/London
```