# レイヤードアーキテクチャ + DI（Dependency Injection）実装パターン

![img.png](img.png)

https://mintaku-blog.net/go-ddd/ より画像引用

## ディレクトリ構成

```shell
├── domain # domain層
│   ├── repository.go
│   └── todo.go
├── handler # interface層
│   └── todo.go
├── infra # infrastructure層
│   ├── db.go
│   └── todo.go
├── main.go
└── usecase # usecase層
    └── todo.go
```

## 依存性逆転の原則（Dependency Inversion Principle）実装パターン

```go
// infrastructure層
// domain層のRepository（インターフェース）に依存させるようにしている
func NewTodoRepository(db *sql.DB) domain.TodoRepo {
	return &TodoInfra{
		db: db,
	}
}

// main.go
func main() {
    todoRepo := infra.NewTodoRepository(infra.ConnectDB())
    todoD := domain.NewTodoDomain(todoRepo)
    todoU := usecase.NewTodoUseCase(todoD)
    todoH := handler.NewTodoHandler(todoU)
}
```
