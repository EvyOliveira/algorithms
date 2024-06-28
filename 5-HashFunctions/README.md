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
</br>
  É importante que funções hash retornem o mesmo valor de saída quando o mesmo valor de entrada for inserido. Caso contrário, não será possível encontrar o item que você deseja na tabela hash.
  </br>
  1. Quais destas funções hash são consistentes?
  </br>
  1.1 f(x) = 1 (retorna "1" para qualquer entrada);
  </br>
  1.2 f(x) = rand() (retorna um número aleatório a cada execução);
  </br>
  1.3 f(x) = proximo_espaco_vazio() (retorna o índice do próximo espaço livre da tabela hash);
  </br>
  1.4 f(x) = len(x) (usa o comprimento da string como índice).
  </br>
  Resposta: 1 e 4.
</br>

## Usando tabelas hash para pesquisas
Quando há uma agenda telefônica em que o nome é associado ao número, as funcionalidades previstas devem ser:
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
|          | Tempo constante (tempo estável independente do tamanho) | Tempo Linear (aumentando no decorrer do tempo) |

Lembrando que retornar um item do array também possui tempo constante.Fazendo um comparativo das tabelas hashs com arrays e listas, temos:

| Operação | Tabela Hash: Caso médio | Tabela Hash: Pior caso | Arrays | Listas encadeadas |
| -------- | ----------------------- | ---------------------- | ------ | ----------------- |
| Busca    |          O(1)           |         O(n)           |  O(1)  |       O(n)        |
| Inserção |          O(1)           |         O(n)           |  O(n)  |       O(1)        |
| Remoção  |          O(1)           |         O(n)           |  O(n)  |       O(1)        |

Para um caso médio da tabela hash, são velozes quanto os arrays na busca (pegar valores dos índices) e velozes quanto as listas de inserção e remoção de itens. Mas no pior caso, as tabelas hashs são lentas para os casos anteriores e para não operar no pior caso, evite colisões através de um baixo fator de carga e uma boa função hash.

### Fator de carga
O cálculo é feito considerando o número de itens de uma tabela hash / número total de espaços. Para o cenário de armazenamento de cem produtos na tabela hash considerando 100 espaços, cada item terá seu espaço. Neste caso, terá fator de carga de 1. Quando temos um fator de carga superior a 1, indicará que será impossível que cada item terá seu próprio espaço e caso o fator de carga cresça demais, precisará de mais espaços na tabela hash, processo denominado de redimensionamento. Normalmente, o redimensionamento é aplicado quando o fator de carga estiver acima de 0.7 e muito próximo de 1, duplicando o tamanho da estrutura do array reinserindo os valores para a nova tabela hash, garantindo menos colisões e com melhor desempenho.

### Uma boa função hash
É considerado uma boa função hash quando os valores são distribuídos simetricamente, não ideal uma função hash que agrupa valores e produz diversas colisões.

### Exercícios:
É importante que funções hash tenham uma boa distribuição. Dessa forma, elas ficam com o mapeamento mais amplo possível. O pior caso é uma função hash que mapeia todos os itens para o mesmo espaço da tabela hash. Suponha que você tenha estas quatro funções hash que operam com strings:
</br>
  A. Retorne "1" para qualquer entrada.
  </br>
  B. Use o comprimento da string como o índice.
  </br>
  C. Use o primeiro caractere da string como índice. Assim, todas as strings que iniciam com a letra a são hasheadas juntas e assim por diante.
  </br>
  D. Mapeie cada letra para um número primo: a = 2, b = 3, c = 5, d = 7, e = 11, e assim por diante. Para uma string, a função hash é a soma de todos os caracteres-módulo conforme o tamanho da hash. Se o tamanho da sua hash for 10, por exemplo, e a string for "bag", o índice será (3 + 2 + 17) % 10 = 22 % 10 = 2.
  </br>
  Para cada um desses exemplos, qual função hash fornecerá uma boa distribuição? Considere que o tamanho da tabela hash tenha dez espaços.
  </br>
  1.1 Uma lista telefônica em que as chaves são os nomes e os valores são os números telefônicos. Os nomes são os seguintes: Esther, Ben, Bob e Dan.
  </br>
  1.2 Um mapeamento do tamanho de baterias e sua devida potência. Os tamanhos são A, AA, AAA e AAAA.
  </br>
  1.3 Um mapeamento de títulos de livros e autores. Os títulos são Maus, Fun Home e Watchmen.
  </br>
</br>