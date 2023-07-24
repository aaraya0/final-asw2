# Proyecto Final: Sistema de publicacion de clasificados

## Arquitectura de Computadoras II 

### Autoras:
* Aguilar Belen
* Araya Agustina
* Cersofios Sofia
* Morellato Agostina

## Busqueda
### GET 
+ Busca una consulta especifica: 
``` GET- /search=:searchQuery ``` 
+ Busca todo lo que coincida con un parametro de la consulta: 
``` GET - /searchAll=:searchQuery ```
+ Busca la publicacion cuyo id coincida con el proporcionado a traves de solr: 
``` GET - /items/:id ```
### DELETE
+ Elimina la publicacion cuyo id coincida con el id proporcionado a traves de solr:
``` DELETE - /item/:item_id ```

## Publicaciones
### GET
+ Busca el item que coincide con el id proporcionado: 
``` GET - /items/:item_id ```
+ Busca los items del usuario cuyo id coincide con el id proporcionado: 
``` GET - /users/:id/items ```
+ Busca las imagenes del item cuyo id coincida con el id proporcionado: 
``` GET - /users/:id/items ```
### POST
+ Crea una publicacion:  
``` POST - /item ```
```
Body:
{
"title": "",
"location": "",
"seller": "",
...
}
```
+ Crea varias publicaciones: 
``` POST - /items ```
``` 
Body:
[
{
"title": "",
"location": "",
"seller": "",
...
},
...
{
"title": "",
"location": "",
"seller": "",
...
}
]
```
### DELETE
+ Elimina la publicacion cuyo id coincida con el id proporcionado:
``` DELETE - /item/:item_id ```
+ Elimina las publicaciones del usuario cuyo id que coincida con el id proporcionado:
``` DELETE - /users/:id/items```

## Comentarios
### GET
+ Busca comentarios: 
``` GET- /messages ```
+ Busca el comentario cuyo id coincida con el id proporcionado: 
``` GET- /messages/:id ```
+ Busca comentarios del usuario cuyo id coincida con el id proporcionado: 
``` GET-/users/:id/messages ```
+ Busca comentarios en la publicacion  cuyo id coincida con el id proporcionado: 
``` GET-/items/:id/messages ```
### POST
+ Crea un comentario en una publicacion: 
``` POST-/message ```
```
Body:
{
    "message_id": ,
    "user_id": ,
    "body": "...",
    ...
}
```
### DELETE
+ Elimina el comentario que coincida con el id proporcionado: 
```DELETE-/messages/:id ```
+ Elimina comentarios de un usuario cuyo id coincida con el id proporcionado: 
```DELETE-/users/:id/messages ``` 

## Usuarios
### GET
+ Busca usuarios:
``` GET-/users ``` 
+ Busca el usuario que coincida con el id proporcionado:
``` GET-/users/:id ```

### POST
+ Crea un nuevo usuario:
``` POST-/user ```
```
Body:
{
	"username": "",
	"password": "",
  "first_name": "",
	...
}
```
+ Crea:
``` POST -/login```
### DELETE
+ Elimia el usuario que coincide con el id proporcionado:
``` DELETE-/user/:id ``` 

