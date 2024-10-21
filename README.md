# CloudRun Weather

CloudRun Weather is an API that provides weather updates and forecasts based on CEP (Brazil postal code). It uses Google Cloud Run for deployment.

## Features

- **Weather Updates**: Get weather information for any location in Brazil based on CEP.
- **Scalability**: Uses Google Cloud Run for automatic scaling.

## Usage

1. Set your `.env` vars
2. Start the development server

```
docker-compose up -d
```

3. Open your browser and navigate to `http://localhost:8080?cep=41210100`.
4. This will return the temperatures in Celsius, Kelvin and Farenheit.

PS. You need to provide a `weatherapi.com` API key in the env vars

## Access

[You can access the deployed API here](https://cloundrun-weather-578459701422.us-central1.run.app?cep=41210100)

#### PS. This project is a quick implementation for the Go-Expert postgraduate diploma I'm taking.
