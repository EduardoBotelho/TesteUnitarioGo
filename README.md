# Projeto de Testes em Go

Este projeto demonstra a implementação de testes unitários na linguagem Go, focando na função `soma`.

## Descrição

O objetivo deste projeto é ilustrar como escrever testes em Go utilizando o pacote `testing`. A função `soma` recebe três inteiros como parâmetros e retorna a soma deles. Dois testes foram implementados para verificar o comportamento da função:

- `TestSoma`: Verifica se a função retorna o resultado correto quando a soma dos números é igual ao esperado.
- `TestSoma2`: Verifica o comportamento da função quando o resultado esperado é diferente da soma real.

## Estrutura do Projeto

O projeto contém os seguintes arquivos:

- `main.go`: Contém a implementação da função `soma`.
- `main_test.go`: Contém os testes unitários para a função `soma`.

## Pré-requisitos

- [Go](https://golang.org/doc/install) instalado na máquina (versão 1.16 ou superior).

## Como Executar os Testes

Para rodar os testes, siga os passos abaixo:

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
   cd nome-do-repositorio
