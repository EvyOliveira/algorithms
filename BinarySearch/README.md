<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Pesquisa ou Busca Binária
Consiste em repartir uma lista ordenada até chegar no conteúdo buscado, uma solução que possibilita ser realizada com o menos número de tentativas possíveis. A estratégia da busca binária é chutar um número intermediário e eliminara metade dos números a cada vez que a estrutura for repartida.
A pesquisa simples é quando você precisa percorrer uma lista inteira para identificar o conteúdo buscado, minando a performance do programa. De maneira geral, uma lista de n números, a pesquisa simples necessitará verificar em n etapas, enquanto a pesquisa binária necessita de log² n. 

### Exercícios:
1.1 Suponha que você tenha uma lista com 128 nomes e esteja fazendo uma pesquisa binária. Qual seria o número máximo de etapas que você levaria para encontrar o nome desejado?
Cálculo dos 128 nomes: 128 -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1, totalizando 8 etapas até encontrar o número desejado.

1.2 Suponha que você duplique o tamanho da lista. Qual seria o número máximo de etapas agora?
Cálculo dos 256 nomes: 256 -> 128 -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1, totalizando 9 etapas até encontrar o número desejado.

## Tempo de execução
Quando o número máximo de tentativas é igual ao tamanho da lista, é chamado de tempo linear. A pesquisa binária é executada de acordo com o tempo logarítimo.
</br>