# gitlab-issue-cli

## Pré requisitos

- Crie um arquivo `.env` contendo a variável `GITLAB_TOKEN`

```
GITLAB_TOKEN={token}
```

## Objetivo

Consultar o nome e a descrição de issues no Gitlab a partir do terminal,
da maneira mais rápida possível.

### Input esperado

```sh
issue 770
```

### Output esperado

```
MSA-770: {Título da issue}

{Descrição da issue}
```

## Desenvolvimento

### Testando

```sh
go run issue.go {int}
```

### Build

```sh
go build issue.go
```

