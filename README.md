# KEKS-EVENTS
---
## Usage

Using `go run`:
```shell
> go run main.go
```
Using `go build`:
```shell
> go build -o keks-events .
> ./keks-events
```
You can access the GraphQL playground at: [http://127.0.0.1:8080/](http://127.0.0.1:8080/)

## Development
После изменения `graph` обязательно запустить перегенерацию `graphql`
```
go generate ./...
```
- Создать событие:
```
mutation createEvent {
  createEvent(input: { 
    dateStart: "2023-01-07T00:00:00Z", 
    dateEnd: "2023-01-07T23:59:59Z", 
    eventName: "Рождество", 
    description: "Отпраздновать рождество" 
    }) 
  {
    dateStart
    dateEnd
    eventName
    description
  }
}

```
- Выборка всех событий:
```
query findEvents {
  events {
    id
    eventName
    description
    dateStart
    dateEnd
  }
}

```
## Built with

- [gqlgen](https://github.com/99designs/gqlgen)
- [sqlite](https://gitlab.com/cznic/sqlite)