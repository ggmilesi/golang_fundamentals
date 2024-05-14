# golang_fundamentals

Aula 01 - Overview

Aula02 

Fundamentos de Multi Task - Multithread
Pre-multitask -> Sistema de Tempo Compartilhado -> OS'sMultitask -> HYper-threading->multicore
1940/1960->1960/1970->1980->1990->2000->2000->2010

Processo (Estudar sobre como funciona)
Uma instancia em execução
Componentes?

Processos usa HEAP(espansivo) ou STACK(fixo)

Criação de processo
-unix/linux
efetua fork e gera um filho do processo raiz

- Scheduler Decide qual processo é executado e alterna entre os processos
obs: POssui um algoritmo de maximizar o CPU
2 tipos de Scheduler
- Colaborativo & Cooperativo
Processos que estão sendo executados tem contro quando liberam a CPU para outros processos.
  Problema: Pode monopolizar a CPU
- Preemptivo
SO tem a capacidade de interromper um processo em execução e ceder o uso da CPU

Threads
Processo são instancia de programas em execução
Thread são unidades basicas de utilização de CPU que fazer parte do processos
São sequencias de execução dentro do mesmo processo, compartilhando o msm espaço de memória e recurso

Dentro de um processo, vaárias threads podem existir


Scheduler no GO
Runtime cria um várias threas de 2kb (em outros geralmente são 2MB)

Pesquisar -> Go routine
GO Routines tem um Global QUeue e Mutex e P (processador Logico)

Cada processo é criado um P para cada Go Routines
![image](https://github.com/ggmilesi/golang_fundamentals/assets/24300009/6a0f570c-84b2-40ec-8f2b-a36d99dfaaff)
![image](https://github.com/ggmilesi/golang_fundamentals/assets/24300009/531f5d4f-96ee-4920-9df2-49ce24879768)


