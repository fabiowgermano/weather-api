# weather-api

Esta aplicação permite consultar a temperatura atual em graus Celsius, Fahrenheit, Kelvin, pesquisando por CEP em um endpoint **REST**.

## Executando a aplicação no CloudRun
* O serviço está disponível no CloudRun, no host `weather-api-557662975481.southamerica-east1.run.app` e endpoint `/temp`. Consulte pelo curl abaixo (ajustar o CEP):
    ```bash
    curl https://weather-api-557662975481.southamerica-east1.run.app/temp?CEP=13180000
    ```
    
* Se preferir, use o modelo disponível em `api/temp_cloudrun.http` para consumir o seviço.

## Executando a aplicação em ambiente local
1. Certifique-se de ter o Docker instalado.
2. Suba os containers necessários executando o comando:
    ```bash
    docker-compose up app --build
    ```
3. Aguarde até que a mensagem de que a aplicação está rodando na porta :8080 seja exibida nos logs.
4. O serviço esta disponível no ambiente local. Pode ser consumido usando o modelo disponível em `api/temp_local.http` ou pela curl abaixo (ajustar o CEP):
    ```bash
    curl http://localhost:8080/temp?CEP=13180000
    ```

## Testes automatizados
Durante o desenvolvimento, foram criadas classes de testes para garantir o funcionamento correto da aplicação. Os testes automatizados foram criados nas camadas de regras negócio (entity), regras de aplicação (usecase) e na intergração com serviços externos (adapters/api).

Passo a passo para execução dos testes automatizados:
1. Certifique-se de ter o Docker instalado.
2. Dispare os testes executando o comando:
    ```bash
    docker-compose run test
    ```