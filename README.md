# Crud Golang gRPC RESTful redis 

  En este código se encontrará un CRUD en el lenguaje Go con almacenamiento en un base de datos postgreSQL utilizando gorm.

## Requisitos
- Docker
- go


## Ejecutar con docker
**Primero** clonar este repositorio.
`git clone https://github.com/pacho328/golangCRUD`

**Segundo** en la raíz del proyecto donde esta el archivo main.go

`docker-compose up` o `docker compose up`

  

## gRPC

  

importar en postman el archivo **medidorgRPC.proto** esta en el path gRPC/medidorRPC/medidorgRPC.prot

  

El servicio queda escuchando en el puerto 50001 como se ve en la imagen, así que seria localhost:50001

  

![enter image description here](https://i.ibb.co/LxR6kcf/Captura-de-pantalla-de-2022-12-25-18-58-08.png)


### Lista de los procedimientos


![enter image description here](https://i.ibb.co/Gfw2CZt/Captura-de-pantalla-de-2022-12-25-18-53-33.png)

  

## RESTful Api

  
  

El servicio queda escuchando en el puerto 5000 la url **localhost:5000**

  

![enter image description here](https://i.ibb.co/0DW9qrW/Captura-de-pantalla-de-2022-12-25-19-45-14.png)

  

Dentro de este repositorio hay un archivo **EnerBit.postman_collection.json** que contiene los request CRUD de medidores.