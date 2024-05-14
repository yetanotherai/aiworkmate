# aiworkmate

```
├── cmd/
│   └── main.go
├── internal/
│   ├── api/
│   │   └── handlers/
│   │       └── user_handler.go
│   ├── domain/
│   │   ├── model/
│   │   │   └── user.go
│   │   ├── repository/
│   │   │   └── user_repository.go
│   │   ├── service/
│   │   │   └── user_service.go
│   │   └── valueobject/
│   │       └── email.go
│   └── infrastructure/
│       ├── config/
│       │   └── config.go
│       ├── persistence/
│       │   └── mysql/
│       │       └── user_repository.go
│       └── web/
│           └── server.go
├── pkg/
├── configs/
│   └── config.yaml
├── scripts/
├── test/
└── docs/
```