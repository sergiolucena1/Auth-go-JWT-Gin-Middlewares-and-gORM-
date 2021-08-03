# Rest-API-challenge (JWT, Gin Middlewares e gORM)

Rest api em golang, criando endpoint de blogposts que cria, edita, lista pelo ID e deleta.

usando o Gin para criar o server, routes, instanciar controllers : $ go get -u github.com/gin-gonic/gin

Nessa aplicação uso o Gorm como ORM : $ go get gorm.io/gorm
criei no banco de dados a tablea de blogposts e de usuário  (para ser atutenticado)

Autentico usuário usando JWT
*Escondendo o Password para SHA256
*Gerar um JWT
*Criar Login model e controller 
*Criar Middleware no GIN 

Tudo foi testado e aplicado usando o Insomnia 

