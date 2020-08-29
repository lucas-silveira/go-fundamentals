# Comandos básicos

Acessar as variáveis de ambiente:
```
go env
```

Compilar e executar um arquivo:
```
go run arquivo.go
```

Gerar um arquivo compilado:
```
go build arquivo.go
```

Baixar e instalar um pacote do Github:
```
go get -u github.com/caminho-do-pacote
```

Validar arquivo e expor erros:
```
go vet arquivo.go
```


## Comandos de teste

Efetuar todos os testes dentro de um diretório
```
go test
```

Efetuar os testes no modo "verboso" (com detalhes)
```
go test -v
```

Efetuar os testes e obter o coverage
```
go test -cover
```

Efetuar os testes e obter o coverage file
```
go test -coverprofile=coverage.out
```

Efetuar os testes e obter o coverage por função e o coverage file
```
go tool cover -func=coverage.out
```

Efetuar os testes e obter um arquivo html do coverage
```
go tool cover -html=coverage.out
```

Efetuar todos os testes dentro de sub diretórios
```
go test ./...
```