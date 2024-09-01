<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Algoritmo de Dijkstra
Determina o caminho mínimo até X para grafos ponderados (atribuir pesos nas arestas). Quando há ciclos em grafos, o algoritmo de Djikstra não funciona. 
Retomando o conceito de pesquisa em largura, em grafos de um ponto para o outro, pode ser adicionado um tempo de deslocamento entre os nós. Sendo assim, para calcular o caminho com menor número de etapas (arestas), poderá ser atendido com a pesquisa em largura, já o caminho mais rápido pode ser obtido por meio do algoritmo de Djikstra. As etapas que o algoritmo de Djikstra percorre estão listados abaixo:
</br>
  1. Encontre o vértice mais "barato", pois será o nó que possibilitará chegar no menor tempo possível.
</br>
  2. Verifique se há um caminho mais "barato" para os vizinhos desse vértice e atualize os custos dos vizinhos do vértice atual. 
</br>
  3. Repita até que tenha feito isso para cada vértice do grafo. 
</br>
  4. Calcule o caminho final.
</br>
Detalhando melhor o que ocorre após escolher o ponto mais próximo do ínicio e que tenha menos tempo de deslocamento, é realizada uma atualização de cálculo de custo para chegar em todos os vértices próximos ou vizinhos de qual estamos em menos tempo e reduzir o custo total. As etapas devem ser executadas recursivamente até chegarmos ao final do percurso.

### Terminologia
Cada aresta do grafo tem um número associado chamado de peso, sendo um grafo com pesos denominado de grafo ponderado (grafo valorado). Por outro lado, um grafo sem peso é chamado de grafo não ponderado (grafo não valorado). Em outras palavras, calcular o caminho mínimo em um grafo não ponderado deve-se utilizar a pesquisa em largura e para calcular o caminho mínimo de um grafo ponderado, utilize o algoritmo de Djikstra. 
Grafos pode conter ciclos, indicando que é possível começar em um vértice, viajar nele e terminar no mesmo vértice ou até mesmo evitar o caminho do ciclo, avaliando o que terá mais ou menos peso. Sendo assim, se seguir o ciclo adicionará mais peso no peso total, seguir o ciclo jamais fornecerá o caminho mínimo. 
Um grafo não direcionado é quando cada vértice adiciona um novo ciclo, pondendo apontar para seu respectivo vértice se tratando de dois vértices. O algoritmo de Djikstra só funciona em grafos sem cilos ou em grafos com um ciclo de peso positivo, uma vez que seguir o ciclo nunca indicará um ciclo de peso positivo.

### Na prática
Depois de compreender o custo entre os vértices, é importante fazer uma tabela para registrar a custo para chegar em cada um dos vértices. Quando falamos de vértices diretos, sabemos o quanto de custo teremos, mas para vértices indiretos, consideraremos como custo infinito.
- Analisaremos qual é o vértice com menos peso para seguirmos, além de identificarmos qual é o nó pai.
- Passando pelo vértice da troca mais barata, precisamos recalcular o custo dos vértices diretos do segundo vértice, que antes tinha custo infinito.
- O vértice pai agora é o segundo vértice. Recalcular para saber se existe um vértice pai que resulte num custo total inferior ao que temos.
- Devemos repetir o processo até chegarmos no menor caminho dos grafos ponderados, compreendendo todas as trocas possíveis.
O caminho mínimo não precisa ser somente a menor distância física, mas pode ser algo como reduzir o custo de troca.

## Arestas com pesos negativos
Não é possível usar o algoritmo de Djikstra se tiver arestas com pesos negativos, pois o algoritmo acaba seguindo o segundo caminho de menos custo por ignorar esses pesos negativos. Caso encontra-se nesse cenário, recomenda-se:
- Criar uma tabela de custos com valores iniciais.
- Encontrar o vértice de menor valor e atualizar o preço dos seus vizinhos, desconsiderando os caminhos de pesos negativos.
- Ao processar o vértice presente na tabela de custos, mas não tendo atualizado o preço dele, não é uma prática recomendável. O motivo se deve ao vértice não ter sido processado, é porque não existe uma forma mais barata de se chegar até ele, mesmo que tenha encontrado um caminho mais viável.
Para os casos de arestas de pesos negativos, o algoritmo de Bellman-Ford podem ajudar a resolver.

### Implementação
Para programar o grafo, precisaremos de tabelas hashs, uma destinada para os pontos (arestas) do grafo, uma segunda para os custos e a último para os pais, sendo as duas últimas tabelas executadas em tempo de execução. 
- Implementar grafo armazenando os vizinhos e o custo para chegar ao vizinho, é resolvido ao criar uma tabela hash com mais tabelas hashs para representar pesos das arestas. 
- Adicionar o restante dos vértices e seus vizinhos no grafo nas tabelas.
Para a tabela de custos, temos como definição o custo de cada vértice como a quantia necessária para chegar, a partir do ínicio, no vértice em questão.
- Analisar como representar o infinito na linguagem de programação a ser escolhida.
- Criar um array adicional para mantermos o registro dos vértices processados, pois não precisam ser processados mais de uma vez. 
Pensando em implementação, para representar os vizinhos do grafo, será composto por uma tabela hash.
- Caso achar um caminho mais barato para algum vértice, deve ser realizada a atualização de custos na tabela.
- Para um novo caminho, considere o novo vértice como pai.
- Para custo final infinito, qualquer valor de custo total será menor que infinito, isto é, considerar o menor custo.
- Ao final, marcar o vértice e para seus vizinhos como processados depois que calcular os custos e continuar.
Encerrar o algoritmo depois de estar completamente processados os vértices. 

### Caso médio e pior caso
No pior caso, a complexidade do algoritmo Big O é de O(V² (número de vértices do grafo) + E (número de arestas no grafo)), já no médio caso é mais complexa a análise com exatidão, pois depende da distribuição dos pesos das arestas e estrutura do grafo. Geralmente, o desempenho no médio caso é próximo ao pior caso, especialmente em grafos mais densos.

### Exercícios:
</br>
  7.1 Em cada um desses grafos, qual é o peso do caminho mínimo do ínicio ao fim?
  </br>
  A. Início -> 5 -> 2 -> 1 -> Fim. Total: 8.
  </br>
  B. Início -> 10 -> 20 -> 30 -> Fim. Total: 60.
  </br>
  C. Ciclo de peso negativo.
</br>