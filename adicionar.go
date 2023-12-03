package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	/*Em resumo, a biblioteca bufio é útil quando você precisa otimizar a leitura ou gravação de dados em blocos,
	        em vez de caracter por caracter. Isso pode ser especialmente útil ao lidar com entrada/saída de dados em lotes
			como em operações de leitura ou gravação de arquivos.*/)

func adicionarProdutos() {
	for {
		fmt.Println("Você deseja adicionar(1) ou remover(2) um produto do estoque? (Para voltao ao menu digite 0)")
		comando := leComando()

		switch comando {
		case 1:
			adicionarProduto()
		case 2:
			removerProduto()
		case 0:
			exibeMenu()
		default:
			fmt.Println("Comando inválido")
		}

	}

}

func adicionarProduto() {
	var produto string
	var quantidade int

	fmt.Println("De qual produto do estoque você quer adicionar?")
	fmt.Scan(&produto)

	if !produtoExiste(produto) {
		fmt.Println("Produto não encontrado no estoque.")
		return
	}

	fmt.Println("Quantidade a adicionar:")
	fmt.Scan(&quantidade)

	if quantidade < 0 {
		fmt.Println("Quantidade inválida. Deve ser maior ou igual a 0.")
		return
	}

	// Aqui você pode adicionar a quantidade ao estoque
	adicionarQuantidade(produto, quantidade)

	fmt.Printf("%d unidades de %s adicionadas com sucesso.\n", quantidade, produto)
}
func produtoExiste(produto string) bool {
	conteudo, err := lerArquivo("estoque.txt")
	if err != nil {
		return false
	}

	for _, p := range conteudo {
		// Verifica se a linha começa com o nome do produto
		if strings.HasPrefix(strings.TrimSpace(p), produto+":") {
			return true
		}
	}

	return false
}

func adicionarQuantidade(produto string, quantidade int) {
	conteudo, err := lerArquivo("estoque.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	encontrado := false

	for i, p := range conteudo {
		parts := strings.SplitN(p, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Formato inválido no arquivo:", p)
			continue
		}

		nomeProduto := strings.TrimSpace(parts[0])

		if nomeProduto == produto {
			// Produto encontrado
			quantidadeAtual, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				fmt.Println("Erro ao converter a quantidade:", err)
				return
			}

			novaQuantidade := quantidadeAtual + quantidade

			// Atualizar a linha no slice
			conteudo[i] = fmt.Sprintf("%s: %d", produto, novaQuantidade)
			encontrado = true
			break
		}
	}

	if !encontrado {
		// Se o produto não foi encontrado, adiciona uma nova linha
		novaLinha := fmt.Sprintf("%s: %d", produto, quantidade)
		conteudo = append(conteudo, novaLinha)
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

func removerProduto() {
	var produto string
	var quantidade int

	fmt.Println("De qual produto do estoque você quer remover?")
	fmt.Scan(&produto)

	if !produtoExiste(produto) {
		fmt.Println("Produto não encontrado no estoque.")
		return
	}

	fmt.Println("Quantidade a remover:")
	fmt.Scan(&quantidade)

	if quantidade < 0 {
		fmt.Println("Quantidade inválida. Deve ser maior ou igual a 0.")
		return
	}

	// Aqui você pode adicionar a quantidade ao estoque
	removerQuantidade(produto, quantidade)

	fmt.Printf("%d unidades de %s removidas com sucesso.\n", quantidade, produto)
}

func removerQuantidade(produto string, quantidade int) {
	conteudo, err := lerArquivo("estoque.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Encontrar a linha correspondente ao produto
	for i, p := range conteudo {
		if strings.TrimSpace(p) == produto {
			// Dividir a linha em produto e quantidade
			parts := strings.Fields(p)
			if len(parts) != 2 {
				fmt.Println("Formato inválido no arquivo.")
				return
			}

			// Obter a quantidade atual
			quantidadeAtual, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Erro ao converter a quantidade:", err)
				return
			}

			// Verificar se há quantidade suficiente para remover
			if quantidadeAtual < quantidade {
				fmt.Println("Quantidade insuficiente para remover.")
				return
			}

			// Calcular a nova quantidade
			novaQuantidade := quantidadeAtual - quantidade

			// Atualizar a linha no slice
			conteudo[i] = fmt.Sprintf("%s %d", produto, novaQuantidade)
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
