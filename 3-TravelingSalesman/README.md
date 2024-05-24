<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Notação Big O
Determina a rapidez de um algoritmo cresce, comparando o número de operações e é importante compreendermos quando lidamos com código feito por outros.Tempos de execução mais comuns e entendimentos sobre Big O.

## Tempo de execução dos algoritmos cresce a taxas diferentes
Para um algoritmo que será utilizado quando um foguete da NASA estiver em vias de pousar na Lua, o algoritmo ajudará a calcular o melhor local de pouso. Nesse caso, para resolver a situação problema há uma necessidde em ser um algoritmo rápido e assertivo em sua execução.
A pesquisa binária atenderia em termos de rapidez, por outro lado a pesquisa simples é mais assertiva. Decidiu-se cronometrar os algoritmos para uma lista de 100 elementos, sabendo-se que para verificar um elemento será gasto 1 milissegundo (ms).
Seriam 100 milissegundos (ms) considerando o pior cenário para a buscaa simples, já que a pesquisa binária levaria 7 milissegundos (log2 100). Sendo assim, a rapidez de um algoritmo não é medida em segundos, mas pelo crescimento no número de operações.

### Ordenação de tempos de execução Big O
Ordenação de tempos de execução do mais rápido para o mais lento:
* Tempo logarítmico O(log n)
* Tempo linear O(n)
* Algoritmo de ordenação rápida 0(n* log n)
* Algoritmo lento de ordenação O(n²)
* Algoritmo mais lento O(n!)


### Exercícios:
1. Forneça o tempo de execução para um dos casos a seguir em termos da notação Big O.
  1.3 Você tem um nome e deseja encontrar o número de telefone para esse nome em uma agência telefônica.
  1.4 Você tem um número de telefone e deseja encontrar o dono dele em uma agenda telefônica.
  1.5 Você quer ler o número de cada pessoa na agenda telefônica.
  1.6 Você quer ler os números apenas dos nomes que começam com A.
</br>

## O Caixeiro-Viajante
Algoritmo de crescimento exorbitante e entendedores acreditam que esse algoritmo possa ser melhorado, chamado de o problema do caixeiro-viajante. A ideia é simular o menor trecho possível a ser percorrido, analisando as possibilidades que devem ser visitadas. O tempo de execução para o problema do caixeiro-viajante é de O(n!) ou tempo fatorial e performa melhor em menor números.
Quando temos um número maior a ser calculado pela proposta deste algoritmo, não temos uma opção mais veloz capaz de trazer resposta, isto é, o que se consegue fazer é chegar numa resposta aproximada.