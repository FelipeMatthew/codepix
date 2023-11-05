# CODEPIX - Full cycle application

## Basic annotation

  ### Microserviços - Pequenos serviços 
    - Empresas atuais utilizam microserviços para realizar arquitetura, desenvolver, testar, monitorar um microserviço. Então não é mais dividido em backend e front, uma equipe de desenvolvedores realiza todo esse procedimento.
    
  ### Full Cycle Developer

  #### Operate what u build 
  - Realizar a configuração e manter operação da aplicação criada.

  #### Ferramentas para escalar
  - Ferramentas que vai realizar a gestão, monitoramento e manutenção dinamica e ativa, geralmente se utiliza, docker, cicd, jeanks, etc.

  ### Full Cycle Developer != Full stack
  Principal diferença seria na parte de arquitetura, entregar o produto final já testado e deployado com suscesso, tento um monitoramento. Sendo assim, o full stack gera o codigo pronto e outros desenvolvedores vai fazer essa parte de configuração da aplicação. TOTALMENTE ESCALÁVEIS.

  <!-- PROJECT  -->

  ### Estudo de caso prático
    - sistema para criar o pix cadastrando o banco e a chave de cada banco 
    - detalhe importante foi a respeito que se caso algum banco ou aplicação esteja fora, não será perdida a chave pix, ela terá que ser mantida sem que haja perca do valor.

  ### Sobre os bancos
  - Banco vai ser um microserviço que vai fazer o cadastro da conta e a transferencia. 
  
  #### Tecnologias
  - Backend - nestjs
  - Fronted - nextjs
  - Docker como container

  ### Sobre o codepix
  - Micro serviço de transferencia bancária
  - Comunicação entre duas pontas
  - Envia requisição para o outro banco com status PENDENTE
  - assim que a req foi aprovada ela vai ser status CONFIRMED
  - assim que o destino confirmar, vai mandar uma comunicação para a origem. 
  - Somente quando a origem e destino estiverem confirmada, ai sim será será retornado como transação completa com sucesso. 

  #### Cadastro e consulta do codepix
  - Chaves vão ficar armazenada no banco do codepix
  - Banco vai verificar se a chave existe, e retorna valor
  - Caso não exista, a requisição de criar chave vai ser liberada. 

  ### Dinâmica do processo
  ![Process Dynamic](https://media.discordapp.net/attachments/901313407083495446/1164034634376544326/image.png?ex=6541bebb&is=652f49bb&hm=4bcbffd67db44b35cc5f5b4342778c6393cfc562ffcc622ccb280989e86800ec&=&width=1439&height=613)


  ### GRPC - KAFKA

  - GRPC - Framework que trabalha com http 2 e transita dados em binário, comunicação entre sistemas.

  - Kafka - Sistema que realiza processamento de dados, joga de um lado p outro ele consome e sistema de stream de dados

  ### Detalhes codepix

  - codepix vai atuar com servidor grpc 
  - vai trocar mensagens com apache kafka
  - ações sendo executadas simultaneamente em GO
  - Orientado a DDD - Domain Driven Design 

  - Camada de domino - camada de aplicação.


  ### Arquitetura Hexagonal / Ports and adapters 

  - Usar padrão existente para criar algo sustentável e sem disperdicio 
  - SUSTENTABILIDADE DE SOFTWARE

  ![Process Dynamic](https://media.discordapp.net/attachments/901313407083495446/1166125423105687614/image.png?ex=654959ee&is=6536e4ee&hm=a1e772d8bf3b67bda36d1cfed0abd6ece5f5cd266e4696b218fbbc91c2abcd77&=)

  ### Recursos a serem utilizados 

  - Docker 
  - GoLang 
  - Apache Kafka


  ### Desenvolvimento 

  zookiper - 

  #### Codigos no docker 

  - para iniciar a aplicação - docker-compose up -d
  - para verificar quais aplicações estão sendo rodadas - docker-compose ps
  - para acessar uma imagem você utiliza - docker exec -it [`nome da imagem`] bash


### Go mod

vai conmseguir baixar e configurar para toda aplicação

### Domain

- coração da aplicação - onde tem entidade e regras de negocio 
- go não é orientado a objeto, é orienteado a estrutura 

Serialização: converter automaticamente a classe para o tipo que você deseja - ex: JSON

Processos de validação e tudo mais tem que ser realizado no coração da aplicação, e não depois

### Realations 

BANK -> Accounts[] -> PixKeys[]

Pix conta e o banco 

### Contrato

Agregado principal - pixkeys

Vai ser definido um padrão - repostirio de interface, para quando for criado um banco de dados ele seja criado nesse padrão.



# Aula 02 - GRPC

GRPC X API

- framework da google que facilita comunicação entre sistemas. 
- ideal para quando for trabalhar em micro serviços (API REST)
- Gera codigo de forma automática com base na sua regra de negócio
- STREAMING BIDIRECIONAL UTILIZANDO HTTP 2 {
  consegue passar informações de modo binário, e algo mais consistente, dados de forma continua
}

Linguagens que tem suporte: 
  - GO
  - JAVA
  - C
    - C++
    - Python 
    - Ruby
    - Node
    - Dart
    - Kotlin
    - C#
    - PHP


### RPC 

Remote  procedure call - função que o servidor oferece para o usuário, e o usuário passa o parametro 

Clitent - sever.soma(A+B) = sever - func soma(a,b){return a+b}

#### Protocol Buffers - 

 É uma plataforma ao qual vc cria uma estrutura de dados, e faz serialização usando XML

 Protocol buffer x JSON
  - Protocol trabalha com arquivos binários, o json é plaintex 
  - é mais rapido e leve, alto desempenho
  - gasta menos recurso de redes

  ![Process Dynamic](https://cdn.discordapp.com/attachments/901313407083495446/1169430023858569237/image.png?ex=65555f94&is=6542ea94&hm=79a6df939d18b09490327d081c07b3a81789cff8caff1ed4ed95179cc84ad17c&)


 ### HTTP 2

- Alto ganho de performance 
- Utiliza a mesma conexão TCP tanto para requisição quanto para resposta. 
- Headers comprimidos para comunicação. 
- migrar tudo para http 2


### GRPC API

Retornar alguns formatos de comunicação 

"UNARY"
- formato unário 
![Process Dynamic](https://media.discordapp.net/attachments/901313407083495446/1169433739173048340/image.png?ex=6555630a&is=6542ee0a&hm=502ddbaf1db78d7064e1d4383c2a09528f7f3fd8c55552c486130cd68ae297d0&=)

// Vai continuar retornando response para o cliente, não vai ser uma requisição unica, ela vai ter retorno continuo chamado de STREAMING DE DADOS
"SERVER STREAMING"
![Process Dynamic](https://cdn.discordapp.com/attachments/901313407083495446/1169434115204984952/image.png?ex=65556364&is=6542ee64&hm=7e68f1912eb36ff6541fce371df0033e83955d83a157941dae14ac0c41f58d7e&)

"CLIENT STREAMING"
![Process Dynamic](https://media.discordapp.net/attachments/901313407083495446/1169434512720134154/image.png?ex=655563c2&is=6542eec2&hm=962973d4d35133cb261b8823110396fbb2e4f1ccdf1128391909fed1741351b1&=)


Model se autovalida e por isso na infraestrutura no repositorio não necessita dessa validação, como é uma regra de negócio é necessária que seja validada na raiz do documento. 


### Bisu

domain - Coração da aplicação
infrastructure - como faz as tratativas do coração da aplicação - Persistencia de dados
application - regra da aplicação - cria banco, conta, chave e depois a aplicação - Regras que vão acessar o dominio a todo o momento

Caso de uso - Passo a passo das ações que serão realizada. 
  - Ex - regista conta, mas ela só é registrada se tiver uma key. Passo a passo 


### GRPC -

 Vai ser responsável para coletar os dados atuais e vai colocar da parte interna da aplicação para parte externa para quando tenha uma api request ele já esteja configurado corretamente

 COMANDO PARA EXECUTAR NO TERMINAL 
 - 