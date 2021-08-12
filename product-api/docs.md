### All the points that are learned while implementing microservices using go

- Handler in go is anything which implements Handler interface
  - Any type should which implements ServeHTTP(http.ResponseWriter,*http.Request) function qualifies as a handler
