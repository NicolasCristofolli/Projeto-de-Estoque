package main

import (
	"fmt"
	"os"
	/*Em resumo, a biblioteca bufio é útil quando você precisa otimizar a leitura ou gravação de dados em blocos,
	        em vez de caracter por caracter. Isso pode ser especialmente útil ao lidar com entrada/saída de dados em lotes
			como em operações de leitura ou gravação de arquivos.*/)

func exibeMenu() {
	for {
		fmt.Println("=== Vita's Estoques ===")
		fmt.Println("1- Cadastrar/Remover Produtos")
		fmt.Println("2- Adicionar/Remover Produtos")
		fmt.Println("3- Renomer Produto")
		fmt.Println("4- Mostrar Estoque")
		fmt.Println("0- Sair")

		comando := leComando()

		switch comando {
		case 1:
			cadastrarProdutos()
		case 2:
			adicionarProdutos()
		case 3:
			renomearProduto()
		case 4:
			mostrarEstoque()
		case 0:
			fmt.Println("saindo")
			os.Exit(0) //precisa importar o "os". Serve pra sair do progr
		default:
			fmt.Println("Nao esta certo")
			os.Exit(-1) //precisa importar o "os". Serve pra avisar q deu um erro e sai do prog

		}
	}
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) //comando pra pessoa digitar (& serve para falar onde vai a info)

	return comandoLido
}

func mostrarEstoque() {
	fmt.Println("=== Estoque Atual ===")
	conteudo, err := lerArquivo("estoque.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	for _, linha := range conteudo {
		fmt.Println("Linha do arquivo:", linha)

	}
}
