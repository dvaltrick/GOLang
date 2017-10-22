# Teste de aptidão com GO

# Instruções
* Clonar o repositório para o diretório C:/Go/src ou para dentro da instalação padrão do Go
* Da pasta src executar o comando *go run github.com/teste/main/main.go*; 
  Esse comando irá inicializar o serviço no endereço http://localhost:12345;
* Para testar a aplicação utilize o programa Postman ou similar para testar apis
* Envie uma requisição com o método *POST* para o endereço *http://localhost:12345/universities*, os parâmetros podem ser passados no formato json da seguinte forma, sendo que todos os parâmetros são opcionais:
{
  "unitid":"455354",
	"city":"San An",
	"instnm":"univer",
	"stabbr":"tx"
}

