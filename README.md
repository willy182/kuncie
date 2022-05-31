# Kuncie

## Requirement
- golang version 1.16

## How to run

## Running docker postgres and jaeger

```
$ docker-compose up -d
```

UP migration:
```
$ make migration
```

## Build and run service
```
$ make run
```

## Open Browser
```
http://localhost:8000/graphql/playground
```

## Schema Mutation
```
mutation{
  orders{
    createOrders(
      data:{
        cart: [
          {
            sku: "A304SD"
            qty: 3
          }
        ]
      }
    )
  }
}
```
