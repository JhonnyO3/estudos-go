# Estudos Go

Repositório de estudos da linguagem Go, organizado em duas trilhas:

## Estrutura

```
.
├── go-from-zero/        # Fundamentos da linguagem
│   ├── 1-pacotes/       # Pacotes e módulos
│   ├── 2-variaveis/     # Variáveis, constantes e declarações
│   ├── 3-tipos-de-dados/# Tipos primitivos e conversões
│   ├── 4-funcoes/       # Funções, retornos múltiplos
│   ├── 5-operadores/    # Operadores aritméticos e lógicos
│   ├── 6-structs/       # Structs e composição
│   └── 7-heranca-go/    # Herança via embedding
│
└── go-advanced/         # Tópicos avançados
    ├── 1-hello-world/
    └── 2-slices-arrays-ponteiros/
        ├── 2.1-introducao/          # Introdução a slices e arrays
        ├── 2.2-copia-simples/       # Cópia superficial
        ├── 2.3-deep-copy/           # Cópia profunda
        ├── 2.4-ponteiros/           # Ponteiros
        ├── 2.5-deep-copy-ponteiros/ # Deep copy com ponteiros
        └── 2.6-desafio/             # Desafio: agendamento de salas (WIP)
```

## Como executar

Cada pasta é um módulo independente. Para rodar um exemplo:

```bash
cd go-from-zero/2-variaveis
go run variaveis.go
```

Para os módulos em `go-advanced`, que usam Go Workspaces:

```bash
cd go-advanced/1-hello-world
go run main.go
```

### Sincronizar o go.work

Para atualizar o workspace após adicionar novos módulos em `go-advanced`:

```bash
cd go-advanced
go run sync-go-work.go
```

## Requisitos

- Go 1.21+
