<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Ordenação por seleção
Arrays e lista encadeada serão abordados nesse capítulo, onde aprenderemos quando a lista encadeada é a melhor opção do que o array.
A maioria das linguagens de programação possuem nativamente os algoritmos de seleção, então raramente terá de escrever o algoritmo do zero.
O computador é composto por um conjunto de gavetas e em cada gaveta possui um endereço de um slot de memória, mas quando queremos armazenar diversos valores, podemos fazer isso de duas formas, arrays e listas. Lembrando que não existe uma forma correta para armazenar itens para cada caso, mas é importante conhecer as diferenças. 

### Arrays e listas encadeadas
Ao necessitar armazenar uma lista de elementos na memória, optaremos primeiro pelo array devido fácil compreensão, armazenando de maneira contígua (uma ao lado da outra). Uma forma de fazer uso desse mecanismo é para armazenar uma lista de afazeres. Ao optar por um array nem sempre o conteúdo será de armazenamento sequencial, pelo fato da memória estar em uso.
Quando precisamos adicionar conteúdo aos slots de memória já ocupados, é necessário buscar um novo espaço de memória que comporte todo o conteúdo e, por isso, adicionar itens a um array torna-se mais lento. Uma maneira de contornar o problema é reservando um número superior de espaços desejados a fim de evitar deslocamento. Vale ressaltar que existem desvantagens como desperdício de memória que por sua vez não ficará disponível para outros processos e possivelmente necessitará mover os itens de qualquer forma pois os "espaços reservados" serão limitados.

## Listas encadeadas
Permite manter os itens em qualquer lugar da memória, em que cada item armazena o endereço do próximo item da lista e endereços aleatórios de memória se mantém interligados, não sendo mais necessário deslocar os itens da lista para novos slots de memória quando escalado o volume dos dados. O tempo de execução para as operações em listas encadeadas dependem de fatores como a operação a ser realizada, tipo de lista encadeada, implementação da estrutura de dados na linguagem de programação escolhida e tamanho da lista. Geralmente, temos o tempo de execução médio para algumas operações em listas encadeadas conforme abaixo:

| Operação                          | Lista Encadeada: Caso médio | Lista Encadeada: Pior caso | Explicação                            |
| --------------------------------- | --------------------------- | -------------------------- |
| Acesso aleatório ao elemento      |              O(1)           |             O(n)           | Em listas encadeadas, para acesso aleatório, o elemento desejado é acessado diretamente no tempo médio, já no pior caso, se falamos de uma lista longa e o elemento estiver no final, será percorrida toda a lista.                                       |
| Busca                             |              O(n)           |             O(n)           | A busca precisa percorrer a lista sequencialmente até encontrar o elemento buscado.                                                                                                                               | 
| Inserção no início                |              O(1)           |             O(1)           | Inserção no ínicio consiste em somente atualizar o ponteiro inicial da lista para apontar para o novo elemento.                                                                                                                              |
| Inserção no final                 |              O(1)           |             O(n)           | Percorrer a lista até o último elemento e atualizar o ponteiro "next" do penúltimo elemento para apontar para um novo elemento e no pior caso, se a lista for muito longa, será necessário percorrer toda a lista para encontrar o penúltimo elemento.       |
| Remoção no ínicio                 |              O(1)           |             O(1)           | Somente é necessário atualizar o ponteiro inicial da lista para apontar para o próximo elemento.                                                                                                                              |
| Remoção no final                  |              O(n)           |             O(n)           | É necessário percorrer a lista até o penúltimo elemento para atualizar o ponteiro "next" para apontar para null.                                                                                                                                  |
| Remoção de um elemento arbitrário |              O(n)           |             O(n)           | Precisa percorrer a lista até encontrar o elemento a ser removido, e os ponteiros "next" e "prev" (se for uma lista duplamente encadeada) precisam ser atualizados.                                                                                        |