# log-geo-anomaly-detect

## Rodando a aplicação
Para compilar a aplicação, deve-se apenas compilar o código usando o comando `go build`, e então chamar o executavél passando por parametro o arquivo de configuração `./log-detection ./path/to/config/file.json`

## Configuração da aplicação:
A configuração deve estar em formato JSON, e contém os seguintes campos:
fieldUserId: Nome do campo que refere-se à identificação do sujeito do log. 
fieldLat: Nome do campo que refere-se à coordenada de latitude do registro. 
fieldLong: Nome do campo que refere-se à coordenada de longitude do registro. 
fieldTime: Nome do campo que refere-se ao tempo do registro. 
fieldTimeFormat: Formato da data do campo de tempo do registro.  
acceptableDisplacement: Deslocamento aceito em km/h do sujeito para detectar anomalia. 
filePath: Caminho para o arquivo de log a ser analisado. 

**Exemplo de arquivo de configuração:**
```json
{
  "fieldUserId":"uid",
  "filePath":"./log.log",
  "fieldLat":"lat",
  "fieldLong":"long",
  "fieldTime":"time",
  "fieldTimeFormat":"2006-01-02 15:04:05",
  "acceptableDisplacement":200
}
```
CONSIDERAÇÕES:
Este aplicação considera que os logs sejam escritos em um arquivo (centralizada), e que a inserção destes logs sejam temporalmente sequenciais.
