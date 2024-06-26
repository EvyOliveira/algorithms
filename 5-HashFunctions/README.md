<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Tabelas Hash
Ao comprar um produto no mercado no caixa, precisaria consultar um caderno não ordenado para obter o preço do produto. Pensando em uma pesquisa não ordenada, a pesquisa simples possuiria tempo de execução O(n) e para uma pesquisa ordenada, a pesquisa binária possui tempo de execução O(log n).
Para caderno de preços, uma forma seria implementar um array contendo o tipo de produto e o preço ordenado pelo nome, utilizando a pesquisa binária para buscar o preço de um item com tempo de execução O (log n).

### Funções Hash
Em poucas palavras, mapeia strings, relaciona com números e serve para colocar a função hash em um conjunto com array. Ao inserir qualquer tipo de dado, a função retorna um número baseado em alguns requisitos como retornar um número consistente para uma quantidade de strings e idealmente mapear palavras para diferentes números.
- Iniciando como um array vazio para armazenar os preços dos produtos;
- Adicionando o nome dos produtos na função hash, um número será retornado e;
- O valor devolvido pela função hash será o índice do array onde o preço do produto ficará armazenado.
Ao final, o array estará cheio e ao buscar pelo preço do produto, é necessário somente colocar o nome do produto, para a função hash retornar o índice onde o preço está, tornando o tempo de execução 0(1). As motivações para encontrá-lo são justificados por: 
- Mapeamento consistente de um nome para o índice correspondente. A primeira execução da função hash serve para verificar disponibilidade de armazenamento;
- A função hash mapeia strings (sequencia de bytes) para diferentes índices e;
- A função hash tem conhecimento sobre o tamanho do array e, retornará somente índices válidos.
A tabela hash é a primeira estrutura de dados com uma lógica adicional, pois enquanto arrays e listas mapeiam para a memória, a tabela hash usa a função hash para identificar onde armazenar elementos conhecidos. São conhecidos como hash maps, maps, dicionários e tabelas de dispersão. Constituídos de chave-valor e uma tabela hash mapeia chaves e valores. 

### Exercícios:
É importante que funções hash retornem o mesmo valor de saída quando o mesmo valor de entrada for inserido. Caso contrário, não será possível encontrar o item que você deseja na tabela hash. 
1. Quais destas funções hash são consistentes?
    1.1 f(x) = 1 (retorna "1" para qualquer entrada);
    1.2 f(x) = rand() (retorna um número aleatório a cada execução);
    1.3 f(x) = proximo_espaco_vazio() (retorna o índice do próximo espaço livre da tabela hash);
    1.4 f(x) = len(x) (usa o comprimento da string como índice).
Resposta: 1 e 4.

## Usando tabelas hash par pesquisas
Quando há uma agend telefônica em que o nome é associado ao número, as funcionalidades previstas devem ser:
- Adicionar o nome e o número de telefone associado ao telefone;
- Inserir o nome de uma pessoa e receber o número associado à ela.
Sendo assim, tabelas hash é ideal quando você precisar mapear algum item com relação a outro ou quando pesquisar algo. Outro exemplo é a aplicação para pesquisas muito maiores como um campo dedicado para url e a tradução para um endereço de IP, chamada de resolução de DNS.

### Evitando entradas duplicadas 
Numa votação feita manualmente, foi registrado os nomes dos votantes por uma lista. A ideia é criar uma tabela hash para quem já votou e quando alguém chegar para votar, seria necessário acessar um GET na tabela hash e caso não retornar, é porque ainda o cidadão ainda não teria votado.
Falando de tabela de pessoas que já tinham votado, se não tivéssemos uma tabela hash, teríamos que fazer uma lista simples percorrendo toda a lista com tempo de execução O(n) e checagem de duplicatas é feito de maneira mais veloz por tabela hash.

## Usando tabelas hash como cache
É visto como uma boa prática quando falamos de um site com um servidor que possui delay para processar dados para processar os dados e mostrar ao cliente. Para tornar o site mais rápido e fazer com que os servidores trabalhem menos, o cache memoriza os dados ao invés de recarregá-los à cada solicitação. O cache ajuda na resposta instantânea ou ao menos bem mais rápida e o site trabalharia menos.
Isto quer dizer que todos os cachings são armazenados em uma tabela hash. Dessa forma, o servidor trabalhará somente se a url não estiver armazenada no cache. Antes de retornar os dados, serão salvos os dados no cache. Sendo assim, a próxima vez que for solicitada a url, poderá enviar os dados do cache ao invés de fazer o servidor executar a tarefa. 

### Colisões
No cenário ideal, a função hash deve mapear diferentes chaves para diferentes espaços, mas nem sempre é feito na prática por conta de um cenário de colisão. Sabendo que o array tem sido ordenado pelos preços dos produtos em ordem alfabética e que cada espaço esteja destinado para armazenar produtos iniciadas com letras do alfabeto, eventualmente pode ocorrer de duas chaves estarem indicadas para um mesmo espaço. E para resolver esse problema, é recomendável que seja implementada um lista encadeada no espaço identificado como em colisão.
Dessa forma, quando quisermos buscar pela lista encadeada dentro de um array, a informação será retornada de maneira mais lenta, mas se a lista encadeada for pequena, não comprometerá o desempenho do algoritmo. Quando há mais elementos armazenados na lista encadeada ao invés da tabela hash, a lista encadeada diminuirá o tempo de execução da tabela hash, exceto se escolher uma boa função hash que evite colisões.

### Desempenho
Até o momento, foi abordado sempre o pior cenário e falando de cenário médio também, falaremos de desempenho das funções hashs quando comparados com outros algoritmos:

| Operação |                      Caso médio                         |                   Pior caso                    |
| -------- | ------------------------------------------------------- | -----------------------------------------------|
| Processa |                         O(1)                            |                      O(n)                      |
| Inserção |                         O(1)                            |                      O(n)                      |
| Remoção  |                         O(1)                            |                      O(n)                      |
           | Tempo constante (tempo estável independente do tamanho) | Tempo Linear (aumentando no decorrer do tempo) |

Lembrando que retornar um item do array também possui tempo constante.Fazendo um comparativo das tabelas hashs com arrays e listas, temos:

| Operação | Tabela Hash: Caso médio | Tabela Hash: Pior caso | Arrays | Listas encadeadas |
| -------- | ----------------------- | ---------------------- | ------ | ----------------- |
| Busca    |          O(1)           |         O(n)           |  O(1)  |       O(n)        |
| Inserção |          O(1)           |         O(n)           |  O(n)  |       O(1)        |
| Remoção  |          O(1)           |         O(n)           |  O(n)  |       O(1)        |