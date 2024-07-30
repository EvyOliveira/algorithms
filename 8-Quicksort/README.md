<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Quicksort
Nem sempre haverá um algoritmo de seu conhecimento para solucionar e para tentar resolver esse problema com a técnica de dividir para conquistar (DC) que trabalha com recursividade. A técnica de dividir para conquistar oferece uma nova maneira de pensar sobre resolução de problemas. 
Lembrando que o algoritmo de quicksort é um algoritmo de ordenação muito mais rápido que a ordenação por seleção (arrays e listas encadeadas). 
O algoritmo de ordenação mais rápido que a ordenação por seleção, além de ser amplamente usado e que também segue a estratégia de dividir para conquistar.
Para ordenarmos um array com quicksort, precisamos que seja um array que contenha mais de um elemento. Quando falamos de estruturas maiores, precisamos escolher um elemento do array para ser chamado de pivô e a partir de um particionamento, é possível encontrar os valores abaixo e acima do pivô. Após esse particionamento, teremos: 
</br>
• o pivô;
</br>
• um subarray contendo os números menores que o pivô e;
</br>
• um subarray contendo os números maiores que o pivô.
</br>
Caso os subarrays não estiverem ordenados, seria simples realizar a ordenação. Ao contrário, caso os subarrays estiverem ordenados, poderá combiná-los com o array esquerdo + pivô + array direito, como ocorre no exemplo abaixo:
[10, 15] + [33] + [] = [10, 15, 33]

Para o caso-base do quicksort, é possível ordenar os arrays de dois elementos (o subarray esquerdo) e também os arrays vazios (o subarray direito). Assim, se utilizar o quicksort em ambos os subarrays e então combinar os resultados, obtemos um array ordenado. 
Quando escolhemos um pivô e possuímos um elemento no subarray com um valor inferior e superior ao pivô, precisamos rodar recursivamente para ambos os subarrays. Sendo assim, independente de qual for o pivô e quantos elementos no subarrays haverão, é somente necessário rodar o quicksort recursivamente. 
O que acabamos de ver é um exemplo de provas por indução, quando é provado de diversas formas que um algoritmo está correto, sendo cada prova acompanhada de um caso-base e um caso indutivo. A mesma lógica é aplicada ao quicksort, ao funcionar para todos os tamanhos de array. 
Referente a notação Big O, a velocidade de execução ao quicksort depende do pivô escolhido, mas no caso Big O (pior caso) tem tempo de execução de O (n log n), tão lento como a ordenação por seleção. No caso médio, temos tempo de execução de O(n log n).

### Merge Sort vs. Quicksort
Quando falamos de execução de um algoritmo, possuímos uma forma de estimar a quantidade determinada de tempo que seu algoritmo gasta para ser executado, representado da seguinte forma: O(n) -> c (constante em unidade de medida de tempo) * n.
Importante dizer que quando temos dois algoritmos com tempos Big O diferentes, a constante é ignorada. A constante fará diferença quando há uma constante menor numa comparação, como acontece com o quicksort e merge sort, pois apesar de terem tempos de execução O(n log n), acaba sendo o quicksort mais rápido por desempenhar melhor em casos médios do que no pior caso.

### Caso médio e pior caso
Uma vez que o desempenho do quicksort depende da escolha do pivô e que o quicksort não verificará previamente se o array está ou não ordenado, a ordenação do array será sempre uma etapa a ser realizada em ambos os casos.
Quando escolhemos um pivô cada vez mais perto do subarray esquerdo na pilha é perdido, pois o subarray esquerdo será vazio e a pilha de chamada consequente é muito longa. Nesse caso, o tamanho da pilha é o O(n), sendo o pior dos casos possíveis.
Já quando o pivô fica mais ao centro, a pilha de chamada fica reduzida quando comparado com o exemplo anterior, justamente por dividirmos o array pela metade e evitar recursos desnecessários, sendo o tamanho da pilha O(log n), considerado o melhor caso. Lembrando que o melhor caso também é o caso médio O(n log n).

### Exercícios:
</br>
  Quanto tempo levaria, em notação Big O, para completar cada uma destas operações?
  </br>
  4.5 Imprimir o valor de cada elemento em um array.
  </br>
  4.6 Duplicar o valor de cada elemento em um array.
  </br>
  4.7 Duplicar o valor apenas do primeiro elemento do array.
  </br>
  4.8 Criar uma tabela de multiplicação com todos os elementos do array. Assim, caso o seu array seja [2, 3, 7, 8, 10], você primeiro multiplicará cada elemento por 2. Depois, multiplicará cada elemento por 3 e então por 7, e assim por diante.
</br>

## Dividir para Conquistar
Para resolver uma estratégia fazendo uso de Dividir para Conquistar, é importante saber que DC são recursivos e para realizar, precisamos:
</br>
• Descobrir o caso-base, que deve ser o caso mais simples possível.
</br>
• Divida ou diminua o seu problema até que ele se torne o caso-base. 
</br>
Para descobrir o caso-base, a ideia é primeiramente verificar se os lados são múltiplos e identificar o maior quadrado possível dentro de uma área. O caso recursivo deve reduzir o problema a cada vez que é chamado, onde o algoritmo deve ser aplicado à uma área para encontrar o maior quadrado possível até a área ser inteiramente calculada. Ao chegar ao final do cálculo da figura, teremos descoberto o caso-base, terá descoberto o último e não haverá segmentos sobrando na figura.
Bom lembrar que a recursão tem memória dos estados anteriores. Quando estiver escrevendo uma função de recursão que envolva um array, o caso-base será muitas vezes um array vazio ou um array de apenas um elemento. 
Quando falamos de programação funcional, existem linguagens de programação que não possuem loops e para isso é necessária que uma função seja chamada recursivamente. 

### Exercícios:
</br>
  4.1 Escreva o código para a função soma, vista anteriormente.
  </br>
  4.2 Escreva uma função recursiva que conte o número de itens em uma lista.
  </br>
  4.3 Encontre o valor mais alto em uma lista.
  </br>
  4.4 Você se lembra da pesquisa binária no Capítulo 1? Ela também é um algoritmo do tipo dividir para conquistar. Você consegue determinar o caso-base e o caso recursivo para a pesquisa binária?
</br>