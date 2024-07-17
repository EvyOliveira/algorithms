<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Pesquisa em largura (BFS)
Permite encontrar o menor caminho entre objetos e pode ser aplicado em uma série de casos. Ajuda a responder se existe caminho entre dois pontos e qual é o caminho mínimo. 
Imagine uma rede de amigos para buscar quem é o vendedor de mangas, sendo possível a busca nos amigos de amigos, adicionando todos a lista até encontrar. Além de encontrar um vendedor de mangas, é importante identificar o vendedor de mangas mais próximo e a pesquisa em largura faz essa varredura do primeiro grau antes de seguir adiante, varrendo de acordo com a ordem que foram adicionados à lista.

### Encontrando o caminho mínimo
A estrutura de dados de fila é responsável por ajudar a pesquisar na ordem que foram adicionados.

### Grafos
Por meio do Problema do Caminho Mínimo, o menor número de etapas possíveis para chegar de um ponto para o outro. O algoritmo que resolve o Problema do Caminho Mínimo é o algoritmo de pesquisa em largura.
Um grafo é o conjunto de conexões, formado por vértices e arestas onde um vértice pode ser conectados com outros vértices chamados de vizinhos e os grafos servem para modelarmos eventos diferentes conectados entre si. A forma que devemos usar para expressar relações com o grafo é fazendo uso da tabela hash, uma vez que a mesma permite mapear uma chave a um valor, ou seja, mapear um vértice a todos os vizinhos. Pelo fato das tabelas hashs não possuirem ordenação, não importa a ordem que são adicionadas os pares chave-valor.
Quando temos um vértice (nó) que não possui vizinhos, é chamado de dígrafo (grafo não direcionado). Um grafo não direcionado (ou simplesmente grafo) não contém setas, e ambos os vértices são vizinhos um do outro. Sendo assim, podem existir casos que o grafo direcionale o não direcional poderão significar a mesma coisa. 
Um grande problema é quando um nó pode ser verificado mais de uma vez, ou seja, possuir relação com diferentes nós. Quando isso acontece, o mesmo será adicionado todas as vezes que estiver relacionado e por ter valores duplicados, verificar mais de uma vez será perda de tempo e dessa forma você poderá marcar o campo como verificado para não ser analisado novamente. Caso contrário, teremos um loop infinito e para resolver esse problema, é necessário criar uma lista de pessoas que já foram verificadas. 

#### Tempo de execução
Caso seja procurado um vendedor de mangas em toda a rede, cada aresta (seta ou conexão entre vértice e outro) será analisado. O tempo de execução é de O(número de arestas). Caso seja mantida a lista de pessoas já verificadas, possui tempo constante de O(1) e fazer a operação para cada pessoa terá tempo de O(número de pessoas) no total. A pesquisa em largura tem tempo de execução O(número de pessoas + número de arestas), escrito como O(V (número de vértices) + A (número de arestas)).

### Exercícios:
</br>
  Este é um pequeno grafo da minha rotina matinal. Ele mostra que não posso tomar café da manhã antes de escovar meus dentes. Então, "tomar café da manhã" depende de "escovar os dentes". Por outro lado, tomar banho não depende de escovar os dentes. A partir desse grafo você pode fazer uma lista relacionando a ordem das atividades da minha rotina matinal.
  </br>
  1. Acordar
  </br>
  2. Tomar banho
  </br>
  3. Escover os dentes
  </br>
  4. Tomar café da manhã
  </br>
  Note que "tomar banho" pode ser movido, logo essa lista também é válida:
  </br>
  1. Acordar
  </br>
  2. Escovar os dentes
  </br>
  3. Tomar banho
  </br>
  4. Tomar café da manhã
  </br>
  6.3 Quanto à essas três listas, marque se elas são válidas ou inválidas.
  </br>
   A. Inválido.
  </br>
   B. Válido.
  </br>
   c. Inválido.
  </br>
  6.4 Aqui temos um grafo maior. Faça uma lista válida para ele.
  </br>
  1. Acordar, 2. Praticar Exercício, 3. Tomar banho, 4. Trocar de roupa, 5. Escovar os dentes, 6. Tomar café da manhã, 7. Embrulhar o lanche.
  </br>
  Você poderia dizer que essa lista é, de certa forma, ordenada. Se a tarefa A depende da tarefa B, a tarefa A aparece depois na lista. Isso é chamado de ordenação topológica, e é uma maneira de criar uma lista ordenada a partir de um grafo. Imagine que você esteja planejando um casamento e tenha um grafo enorme de tarefas a serem realizadas. Porém, você não sabe nem por onde começar. Assim, uma ordenação topológica do grafo poderia ser feita e, dessa forma, uma lista de tarefas já em ordem seria elaborada. Suponha que você tenha uma árvore genealógica.
  </br>
  Essa árvore é um grafo, pois existem vértices (as pessoas) e arestas, e as arestas apontam para os pais dos vértices. Porém, todas as arestas apontam para baixo, pois não faria sentido uma árvore genealógica ter arestas apontando para cima! Seu pai não pode ser o pai do seu avó!
  </br>
  Isso é chamado de árvore. Uma árvore é um tipo especial de grafo em que nenhuma aresta jamais aponta de volta.
  6.5 Quais desses grafos também são árvores?
  </br>
  A. Árvore.
  </br>
  B. Não é uma árvore.
  </br>
  C. É uma árvore, mas na lateral. Árvores são um subconjunto de grafos. Assim, uma árvore sempre será um grafo, mas um grafo pode ou não ser uma árvore.
</br>

### Filas
Suponha que há uma fila em um ponto de ônibus e que você está numa posição à frente do seu amigo, quer dizer que entrará antes no ônibus do que seu amigo. As filas possuem funcionamento similar as pilhas e por isso, não é possível acessar elementos aleatórios na fila. As únicas operações possíveis são enqueue (enfileirar) e dequeue (desenfileirar) se referindo à atualização das filas, mas que podem ser chamadas de push (enqueue) e pop (dequeue). Quando dois itens são enfileirados na lista, o primeiro item será desenfileirado antes do segundo. Aplicando a ideia da lista de pesquisas, os itens adicionados primeiro na lista serão desenfileirados e verificados primeiro. A fila é uma estrutura de dados FIFO (First In, First Out), já a pilha é uma estrutura de dados LIFO (Last In, First Out).

### Exercícios:
</br>
  Execute o algoritmo de pesquisa em largura em cada um desses grafos para encontrar a solução.
  </br>
  6.1 Encontre o menor caminho do ínicio ao fim.
  Resposta: O menor caminho possui 2 etapas.
  </br>
  6.2 Encontre o menor caminho de "jato" até "gato".
  Resposta: Jato possui conexão direta com gato, sendo assim, somente uma etapa.
  </br>
</br>