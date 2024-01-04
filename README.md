# KEKS-EVENTS

## Usage

Using `go run`:
```shell
> go run server.go
```
Using `go build`:
```shell
> go build -o keks-events .
> ./keks-events
```
You can access the GraphQL playground at: [http://127.0.0.1:8080/](http://127.0.0.1:8080/)

## Development
After making changes to `graph` directory make sure to run the generate command:
```
go generate ./...
```
Create event:
```
mutation createEvent {
  createEvent(input: { 
    start: "2024-01-07T00:00:00", 
    end: "2024-01-07T23:59:59", 
    title: "Рождество", 
    description: "Отпраздновать рождество" 
    }) 
  {
    start
    end
    title
    description
  }
}
```
Query all events:
```
query findEvents {
  events {
    id
    title
    description
    start
    end
  }
}

curl -s 'http://localhost:8080/query' \
  -X POST \
  -H 'content-type: application/json' \
  --data '{
    "query":"query findEvents { events { id title description start end } }", "operationName":"findEvents"
  }' | jq
```
Delete event:
```
mutation deleteEvent {
  deleteEvent(id: "5256355a-71cc-4859-aa7e-b97b8600ed7c") 
}
```

## Built with

- [gqlgen](https://github.com/99designs/gqlgen)
- [sqlite](https://gitlab.com/cznic/sqlite)