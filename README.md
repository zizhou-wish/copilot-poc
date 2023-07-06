# copilot-poc

## Getting started

> Version: go 1.19

```sh
# Install dependencies
go mod download

# Start GraphQL server
go run cmd/main.go
```

## Usage

Visite http://localhost:8080

### Sample mutation

Add new user

```
mutation createUser {
  createUser(input: {name: "zi"}) {
    id
    name
  }
}
```

Add new todo

```
mutation createTodo {
  createTodo(input: { text: "todo", userId: "<userID>" }) {
    user {
      id
    }
    text
    done
  }
}
```

### Sample query

```
query todos {
  todos(userID: "<useID>") {
    id
    text
  }
}
```