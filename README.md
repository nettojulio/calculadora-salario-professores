# Calculadora de Salário de Professores

## Programação de Soluções Computacionais

Desenvolva um programa que funcione como uma calculadora de valor a receber para um professor.
O programa deve solicitar ao usuário o nome do professor e seu regime de pagamento, que pode ser: CLT, Horista ou PJ.
Em seguida, o programa deve pedir informações adicionais de acordo com o regime selecionado.
Caso o regime de pagamento seja CLT, o programa deve solicitar o salário mensal do professor.
Se o regime for Horista, o programa deve pedir o número de horas trabalhadas pelo professor e o valor da hora de trabalho.
No caso de PJ, o programa deve pedir o valor do contrato do professor.

Após receber as informações necessárias, o programa deverá realizar os cálculos adequados para determinar o valor a receber pelo professor, de acordo com seu regime de pagamento. O cálculo será feito da seguinte maneira:

- Para regime CLT, o valor a receber será igual ao salário mensal informado.
- Para regime Horista, o valor a receber será o produto entre o número de horas trabalhadas e o valor da hora de trabalho.
- Para regime PJ, o valor a receber será igual ao valor do contrato.

Por fim, o programa deverá imprimir o resultado, exibindo o nome do professor e o valor a receber.

O código pode ser desenvolvido em qualquer linguagem de programação, utilizando classes ou estruturas para resolver o problema.

ATENÇÃO: O código deverá estar versionado no GitHub

### ⚙️ Cofigurações

- Para execução do projeto, verifique se o Docker está ativo e execute os comandos abaixo:

```bash
make docker-up
```

```bash
make run
```
