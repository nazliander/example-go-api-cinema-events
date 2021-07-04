# Cinema Events

This is a simple Go web application that has two methods:

1. `GET` for retrieving all cinema events from `/events` endpoint.
2. `PUT` for inserting a new cinema event to the events list at `/event` endpoint.

## Run the application

This application listens on port `6111`

To run the application, use the following command:

```bash
docker build -t cinema-events . && docker run -p 6661:6661 cinema-events
```

Access the application on: http://127.0.0.1:6111/

## Available at Docker Hub

Link: https://hub.docker.com/r/nandercc/cinema-events

```bash
docker pull docker.io/nandercc/cinema-events
```
