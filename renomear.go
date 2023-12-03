package main

import (
	"fmt"
	"os"
	"strings"
	/*Em resumo, a biblioteca bufio é útil quando você precisa otimizar a leitura ou gravação de dados em blocos,
	        em vez de caracter por caracter. Isso pode ser especialmente útil ao lidar com entrada/saída de dados em lotes
			como em operações de leitura ou gravação de arquivos.*/)

func renomearProduto() {
	var produtoAntigo, novoNome string

	fmt.Println("Qual produto você deseja renomear?")
	fmt.Scan(&produtoAntigo)

	if !produtoExiste(produtoAntigo) {
		fmt.Println("Produto não encontrado no estoque.")
		return
	}

	fmt.Println("Digite o novo nome para o produto:")
	fmt.Scan(&novoNome)

	if novoNome == produtoAntigo {
		fmt.Println("O novo nome deve ser diferente do nome atual.")
		return
	}

	if produtoExiste(novoNome) {
		fmt.Println("Já existe um produto com o novo nome no estoque.")
		return
	}

	renomearProdutoNoEstoque(produtoAntigo, novoNome)

	fmt.Printf("Produto %s renomeado para %s com sucesso.\n", produtoAntigo, novoNome)
}

func renomearProdutoNoEstoque(produtoAntigo, novoNome string) {
	conteudo, err := lerArquivo("estoque.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Encontrar a linha correspondente ao produto antigo
	for i, p := range conteudo {
		if strings.TrimSpace(p) == produtoAntigo {
			// Atualizar o nome na linha do slice
			conteudo[i] = novoNome
			break
		}
	}

	// Escrever o novo conteúdo de volta no arquivo
	arquivo, err := os.Create("estoque.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer arquivo.Close()

	for _, p := range conteudo {
		arquivo.WriteString(p + "\n")
	}
}
