# MVC + S 実装パターン

MVC + S は、MVC パターンに Service レイヤーを追加した実装パターンです。

## ディレクトリ構成

```shell
├── controller # Controller
│   └── todo.go
├── handler # View
│   └── todo.go
├── main.go
├── model # Model
│   ├── db.go
│   └── todo.go
└── service # Service
    └── todo.go
```
