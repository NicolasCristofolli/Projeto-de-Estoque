package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cadastrarProdutos() {

	for {
		fmt.Println("Você deseja cadastrar(1) ou remover(2) um produto do estoque? (Para voltao ao menu digite 0)")
		comando := leComando()

		switch comando {
		case 1:
			cadastrarProduto()
		case 2:
			removerProdutoCadastrado()
		case 0:
			exibeMenu()
		default:
			fmt.Println("Comando inválido")
		}

	}
}

func cadastrarProduto() {
	var produto string
	fmt.Println("Qual produto você deseja cadastrar?")
	fmt.Scan(&produto)

	if produtoExiste(produto) {
		fmt.Println("Produto já existe no estoque.")
		return
	}

	// Adiciona o nome do produto seguido por ": 0"
	linhaProduto := fmt.Sprintf("%s: 0", produto)

	arquivo, err := os.OpenFile("estoque.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("Erro ao adicionar produto!", err)
		return
	}

	defer arquivo.Close()
	if _, err := arquivo.WriteString(linhaProduto + "\n"); err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
	fmt.Println("Produto adicionado com sucesso")
}

func removerProdutoCadastrado() {
	var produto string
	fmt.Println("Qual produto você deseja remover?")
	fmt.Scan(&produto)

	if !produtoExiste(produto) {
		fmt.Println("Produto nao encontrado no estoque.")
		return
	}

	conteudo, err := lerArquivo("estoque.txt")
	if err != nil {
		return
	}

	novoConteudo := []string{}
	for _, p := range conteudo {
		if p != produto {
			novoConteudo = append(novoConteudo, p)
		}
	}

	arquivo, err := os.Create("estoque.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo", err)
		return
	}
	defer arquivo.Close()
	for _, p := range novoConteudo {
		arquivo.WriteString(p + "\n")
	}
	fmt.Println("Produto removido com sucesso")
}

func produtoExiste1(produto string) bool {
	conteudo, err := lerArquivo("estoque.txt")
	if err != nil {
		return false
	}

	for _, p := range conteudo {
		if strings.TrimSpace(p) == produto {
			return true
		}
	}

	return false
}

func lerArquivo(nomeArquivo string) ([]string, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	var linhas []string
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	return linhas, nil
}
