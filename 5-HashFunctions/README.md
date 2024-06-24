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

