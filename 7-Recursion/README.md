<h1 align="center">
  Estudo pessoal: Entendendo Algoritmos
</h1>

</br>

## Recursão
Aprender a separar um problema em um caso-base e um caso recursivo fazendo uso da estratégia dividir para conquistar quebra grandes problemas em pequenas soluções.
Recursão é quando uma função chama a si mesma, utilizado para tornar a resposta mais clara e não há nenhum benefício ao utilizar a recursão em termos de desempenho e os loops podem apresentar melhora para o desempenho do programa. Pelo fato da função recursiva chamar a si mesma, é mais fácil errar e acabar em um loop infinito.
Outro ponto importante é determinar a condição de parada. Sendo assim, o caso-base é quando a função chama a si mesma, já o caso-base é quando a função não chama a si mesma novamente, de forma que o programa não esteja num loop infinito. Vamos explorar um tipo de estrutura de dados que nos ajudará a entender o funcionamento da recursão.

## Pilha de Chamada
O computador usa uma pilha interna denominada pilha de chamada. Cada vez que você chama uma função, é alocada uma caixa de memória para a chamada. A variável passada por parâmetro da função é setada para o valor desejado e é salvo, isso é feito para todas as chamadas de função, o computador deve salvar na memória os valores para todas as variáveis. 
Para o espaço reservado na memória, é usada a estrutura de dados de pilha, sendo a primeira “caixa” de memória contida na base de pilhas quando a função é chamada dentro da outra, a chamada de função fica pausada em um estado parcialmente completo.
Quando utilizamos uma pilha para guardar variáveis de múltiplas funções é denominado pilha de chamada. 

### Exercício:
  3.1 Suponha que eu forneça uma pilha de chamada como esta:
Quais informações você pode retirar baseando-se apenas nessa pilha de chamada? Agora, vamos ver esta pilha de chamadas sendo executada com uma função recursiva. 
</br>

### Pilha de chamada com reversão
As funções recursivas também utilizam pilha de chamadas. Utilizar pilha é conveniente, mas pode custar muito espaço de memória. Cada uma das funções de chamada, ocupa um pouco de memória e quando a pilha está muito cheia é sinal que seu computador está salvando informações para muitas chamadas de funções, sendo possível: 
</br>
• Reescrever o código com loops;
</br>
• ⁠Utilizar tail recursion (recurso de cauda).
</br>

### Exercício:
  3.2 Suponha que você acidentalmente escreva uma função recursiva que fique executando infinitamente. Como você viu, seu computador aloca memória na pilha para cada chamada de função. O que acontece com a pilha quando a função recursiva executa infinitamente?
</br>


