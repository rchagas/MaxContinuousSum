# Max Continuous Sum

## Introdução
Este software foi programado em go e implementa uma API GraphQl para gerar o valor e a sublista da maior soma em sequência de uma dada lista.

## Execução

Para rodar MaxContinuousSum use o comando:
```
go run server.go
```
Dessa forma um server de teste irá rodar na porta local 8080 onde poderá ser feito testes na API.

Foram feitos testes automaticos para as funções utilizadas dentro do software, os testes se encontram no arquivo maxContinuousSum_test.go dentro do pacote maxSum, para executar a bateria de testes automáticos entre dentro da pasta do pacote 
```
cd maxSum
```
maxSum o comando:
```
go test
```

A requisição pode ser feita com o server rodando utilizando: 
```
mutation{
  maxsum(list: [1, 2, 3, -10, 10,-2,5,-10]){
    sum
    positions
  }
}
```

## Algoritmo para encontrar a maior soma em série na lista

Foi criada uma função principal cujo nome é MaxContinuousSum que tem como objetivo retornar o valor mais alto de soma de números em sequência de uma lista de inteiros, além de uma função auxiliar GenerateSubList que gera uma lista apartir do índice inicial e final da sublista com os numeros dessa soma.
Sabendo que para uma lista com apenas números negativos a maior soma é o menor valor da lista sozinho, o algoritmo consiste em pecorrer a lista armazenando a maior soma atual enquanto busca outra maior.

É utilizada uma variável auxiliar que guarda uma nova soma (possivelmente maior que a atual), que quando fica abaixo de 0 pode ser descartada pois a soma atual é necessariamente maior ou igual a esse numero negativo e o index de início da variável auxiliar é reiniciado com o próximo index do laço
enquanto quando o valor da variável auxiliar ultrapassa o valor já armazenado na soma atual significa que uma soma mais alta foi encontrada e os valores atualizados.

Entrada:  - lista de inteiros de qualquer tamanho
Saída: 	  - Valor da maior soma de elementos em sequência na lista
			    - Sublista com as posições dos elementos da soma
```
func MaxContinuousSum(listNumbers []int) (int, []int) {
...
}
```
A função armazenará a maior soma contínua da lista, inicializado com o primeiro valor junto com o index inicial e final (como se trata de valores em sequência apenas precisamos do começo e o fim da série).
```
maxContinuousSum := listNumbers[0]
startIndex := 0
endIndex := 0
```
A variável nextSum é usada como auxiliar para buscar uma soma mais alta enquanto pecorre a lista nextSumStartIndex é usado como index auxiliar quando nextSum reinicia sua contagem.
```
nextSum := 0
nextSumStartIndex := 0
```
O for pecorre a lista de inteiros:
```
for  i := 0; i < len(listNumbers); i++ {
...
}
```
nextSum soma os elementos em série:
```
nextSum += listNumbers[i]
```
Se o valor da variável auxiliar é maior que o atual atualiza os valores:
```
if maxContinuousSum < nextSum {
  maxContinuousSum = nextSum
  startIndex = nextSumStartIndex
  endIndex = i
}
```
Se o valor da variável auxiliar é menor que 0 reinicia a contagem:
```
if nextSum < 0 {
  nextSum = 0
  nextSumStartIndex = i + 1
}
```

A função auxiliar GenerateSubList gera uma lista com a sequência de números de um valor até o outro

Entrada:  - Valor inicial
          - Valor Final
Saída: 	  - Lista com os valores em sequencia do Valor inicial até o final
```
func GenerateSubList(startIndex int, endIndex int) []int {
	list := make([]int, (endIndex - startIndex + 1))
	j := startIndex
	for i := 0; i < (endIndex - startIndex + 1); i++ {
		list[i] = j
		j++
	}
	return list
}
```