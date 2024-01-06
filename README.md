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
Endpoint for query http://localhost:8080/query

Create event:
```
mutation createEvent {
  createEvent(input: { 
    start: "2024-01-07T00:00:00+03:00", 
    end: "2024-01-07T23:59:59+03:00", 
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

curl -s -X POST 'http://localhost:8080/query' \
  -H 'content-type: application/json' \
  -H 'Authorization: Bearer xxx' \
  --data '{
    "query":"mutation createEvent { createEvent(input: { start: \"2024-03-08T00:00:00+03:00\", end: \"2024-03-08T23:59:59+03:00\", title: \"Международный женский день\", description: \"День женщин\" }) { start end title description }}","operationName":"createEvent"
  }' | jq
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

curl -s -X POST 'http://localhost:8080/query' \
  -H 'content-type: application/json' \
  -H 'Authorization: Bearer xxx' \
  --data '{
    "query":"query findEvents { events { id title description start end } }", "operationName":"findEvents"
  }' | jq
```
Update event:
```
mutation updateEvent {
    updateEvent(input: { id: "0a9f34bb-6ddf-40a3-9464-a30ee6f8960e", start: "2024-01-09T10:00:00+03:00", end: "2024-01-09T11:00:00:+03:00", title: "Что-то очень важное", description: "Что-то очень очень важное"}) {
    id
    start
    end
    title
    description
    }
}
```
Delete event:
```
mutation deleteEvent {
  deleteEvent(id: "5256355a-71cc-4859-aa7e-b97b8600ed7c") 
}

curl -s -X POST 'http://localhost:8080/query' \
  -H 'content-type: application/json' \
  -H 'Authorization: Bearer xxx' \
  --data '{
    "query":"mutation deleteEvent { deleteEvent(id: \"5b76910d-0756-4f63-8c6e-aff89f7f8739\")}","operationName":"deleteEvent"
  }' | jq
```

## Built with

- [gqlgen](https://github.com/99designs/gqlgen)
- [sqlite](https://gitlab.com/cznic/sqlite)