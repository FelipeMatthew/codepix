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