<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Algoritmos Gulosos
Os algoritmos gulosos (algoritmos gananciosos) permitem definir a cada etapa, o movimento ideal, selecionando a solução ideal e até obter a solução ideal global. Os algoritmos gulosos são simples de escrever e chegam perto da solução ótima, exceto os casos que atendem plenamente. O tipo de problema que os algoritmos gananciosos podem se encaixar bem são os problemas que envolvem conjuntos na tentativa de cortar custos, percorrendo os passos abaixo:

1. Listar subconjuntos chamados de conjuntos de partes (ou conjunto de potência).
2. Entre os subconjuntos, escolher o conjunto com o menor custo que abrange todo o problema. 

Para resolver o problema de maneira rápida, temos os algoritmos de aproximação.

### Algoritmos de aproximação
Quando é necessário tempo para calcular a solução exata, os algoritmos de aproximação destaca-se por sua rapidez bem como a capacidade de chegar na solução ideal. Considerando o algoritmo guloso, não é de fácil compreensão e simples, por ter tempo de execução de O (n^2), onde n compreende o número de itens a serem percorridos.
O que normalmente ocorre na prática é que é recebido um array inicial que é convertido em conjuntos (lista de itens que não são repetidos). Para cada item percorrido terá suas possibilidades de escolhas representadas em uma tabela hash para cada estação. Lembrando que os valores representados na tabela hash também são conjuntos que estão associados à uma chave (nome do vértice) e ajudam a compor o cálculo do conjunto final. Para chegarmos na resposta, deve-se escolher o conjunto que mais cobre os pontos, aplicado à teoria dos conjuntos. 
Em velocidade de execução, o algoritmo guloso tem tempo de execução de O (n^2) é mais rápido que o algoritmo exato de tempo de execução O (2^n).

### Exercícios:
</br>
  8. Para cada um desses algoritmos, diga se ele é um algoritmo guloso ou não.
  </br>
  8.3 Quicksort.
  </br>
  8.4 Pesquisa em largura.
  </br>
  8.5 Algoritmo de Dijkstra.
</br>

### Problemas NP-completos
É preciso calcular cada conjunto possível e isso deverá te fazer lembrar do problema do caixeiro-viajante que prevê visitar o menor caminho passando por todos os pontos somente uma vez de tempo de execução Big O de O(n!) ou tempo fatorial. Para um conjunto de pontos, defina:

1. Trace os diferentes caminhos entre dois pontos, resultando em possibilidades de caminhos em 2!.
2. Ao adicionar o terceiro ponto, simule os trajetos possíveis totalizando 6!.
3. Quando há um quarto ponto, os três primeiros pontos somam as possiblidades e para calcular o total seria: 4 (pontos de partida possíveis) * 6 (rotas possíveis para cada ponto de partida) = 24 rotas possíveis.

É por isso que é impossível calcular a solução correta para o problema do caixeiro-viajante caso o número de pontos seja elevado. Alguns problemas são difíceis de resolver como o do caixeiro-viajante e a cobertura dos conjuntos que não possuem uma resolução rápida aparente. 
Pensando em montar um time de futebol americano com requisitos bem definidos e com tamanho limitado, é possível trabalhar com a teoria dos conjuntos ou melhor, cobertura dos conjuntos.

1. Considerar o jogador que possui maior número de habilidades ainda não preenchidas e;
2. Repetir até o time já ser preenchido com todas as habilidades (ou até que fique sem espaço no time).

É sempre bom saber se o problema a ser solucionado é de natureza NP-completo e a partir daí, fazermos o uso correto do algoritmo de aproximação ao invés de tentar resolver perfeitamente. Porém, a diferença entre um problema de NP-completo e um problema que é fácil de resolver é mínima e pode se tornar difícil de identificar. 
Falando de cálculo do caminho mínimo, se quiser encontrar um caminho que conectaria com vários pontos, cairá no problema do caixeiro-viajante, que é um problema de NP-completo. As motivações para considerarmos que um problema é do tipo NP-completo, são:

- O algoritmo com alguns itens roda bem e conforme for aumentando, a performance vai sendo comprometida;
- Todas as combinações de "x" geralmente indicam um problema de NP-completo;
- Calcular cada "versão" de "x" porque não pode dividir em subproblemas menores;
- Se envolver sequência e é difícil de resolver, pode ser um NP-completo;
- Se envolve conjunto e é difícil de resolver, pode ser um NP-completo;
- Se for possível reescrever o problema somente com o problema de cobertura mínima ou problema do caixeiro-viajante.

### Exercícios:
</br>
  8.6 Um carteiro deve entregar correspondências para vinte casas. Ele deve encontrar a rota mais curta que passe por todas as vinte casas. Esse é um problema de NP-completo?
  </br>
  8.7 Encontrar o maior clique (em grafo não orientado é um subconjunto de seus vértices tais que cada dois vértices do subconjunto são conectados por uma aresta) em um conjunto de pessoas (conjunto de pessoas que todos se conhecem). Isso é um problema de NP-completo?
  </br>
  8.8 Você está fazendo o mapa dos Estados Unidos e precisa colorir os estados adjacentes com cores diferentes. Para isso, deve encontrar o número mínimo de cores para que não existam dois estados adjacentes com a mesma cor. Isso é um problema NP-completo?
</br>