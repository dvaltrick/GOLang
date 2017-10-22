# Teste de aptidão com GO

# Instruções
* Clonar o repositório para o diretório C:/Go/src ou para dentro da instalação padrão do Go
* O arquivo novo.csv contém os dados para inicializar a base de dados e deve estar no diretório raiz, quando executada pela primeira vez a aplicação verifica se a base já possui dados caso contrário importa as informações do arquivo novo.csv 
* Para o correto funcionamento do postgres deve ser criado um banco de dados e executado o script contido no arquivo dml.sql
* No arquivo github.com/teste/services/database.go devem ser alteradas as constantes de acesso ao postgre

const (
	dbuser = "usuário do banco de dados"
	dbpwd = "senha do banco de dados"
	dbname = "nome do banco de dados" 
)

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

