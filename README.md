# delivery-api-project

# Pre requirements:

You need to have Docker and docker-compose installed. See https://docs.docker.com/engine/install/ and https://docs.docker.com/compose/install/ for more information.

# How to run it:
On the root of the project run `docker-compose up` and thats it, the api will run in the port 8080 of your local host. 

# Endpoints:

You can find the postman collection on the root project. However this are the avaible endpoints:

- `http://localhost:8080/api/v1/order/ping`
- `http://localhost:8080/api/v1/order/create`
- `http://localhost:8080/api/v1/order/update/:id/:status`
- `http://localhost:8080/api/v1/order/get/:id`

# Request and responses examples:
  
- `http://localhost:8080/api/v1/order/ping`: `pong`

- `http://localhost:8080/api/v1/order/create`:
  Request:

  ```{
   "id_receiver":1,
   "id_remitter":20,
   "packages":[
      {
         "weight":10,
         "quantity_items":2
      },
      {
         "weight":20,
         "quantity_items":1
      }
   ],
   "delivery":{
      "origin_location":{
         "type":"remittent",
         "province":"caba",
         "city":"caba",
         "commune":"caba",
         "full_address":"caba",
         "lat":-34.6596342,
         "lng":-58.5050333
      },
      "destiny_location":{
         "type":"receiver",
         "province":"bsas",
         "city":"bsas",
         "commune":"bsas",
         "full_address":"bsas",
         "lat":-34.6088606,
         "lng":-58.5162579
      }
   }
}```
