# Atividade 4

Implementar em Go as Camadas de Infraestrutura, Distribuição e Serviços Comuns (Serviço de Nomes) de um middleware baseado em RPC utilizando os seguintes padrões de projeto:

- Client Request Handler
> TODO tratar timeouts e erros de invocação

O Client Request Handler é responsável por gerenciar conexões de rede, enviar/receber invocações, tratar timeouts e erros de invocação

- Server Request Handler<br>
> TODO aparentemente nada, provavelmente tratar os erros que aparecerem

O Server Request Handler é responsável por receber a invocação, combinar fragmentos de mensagens (quando for o caso), encaminhar a invocação para o Invoker correto

- Requestor<br>
> TODO tratar erros 

Considerando que: na rede só passa sequência de bytes, é preciso estabelecer uma conexão, a invocação precisa ser enviada ao objeto remoto, precisa receber o resultado da invocação, precisa tratar erros

- Invoker
> TODO tudo

- Marshaller
> ok

Dados dos requests: object ID, nome da operação, parâmetros e valores de retorno, [informações de contexto]<br>
Apenas streams de bytes são transportados pela rede

- Absolute Object Reference
> TODO tudo

- Lookup
> TODO tudo

Construa uma aplicação sobre o middleware implementado e sobre um middleware existente baseado em RPC (e.g., gRPC ou Go RPC). Utilize esta aplicação para realizar uma avaliação comparativa do desempenho dos dois sistemas de middleware.
