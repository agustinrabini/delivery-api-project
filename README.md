# delivery-api-project

# Pre requirements:

You need to have Docker and docker-compose installed. See https://docs.docker.com/engine/install/ and https://docs.docker.com/compose/install/ for more information.

# How to run it:
On the root of the project run `docker-compose up` and thats it, the api will run in the port 8080 of your local host. 

# Endpoints:

You can find the postman collection on the root project. However this are the avaible endpoints:

- `http://localhost:8080/api/v1/order/ping` `[GET]`
- `http://localhost:8080/api/v1/order/create` `[POST]`
- `http://localhost:8080/api/v1/order/update/:id/:status` `[UPDATE]`
- `http://localhost:8080/api/v1/order/get/:id` `[GET]`

# Request example:
  
- `http://localhost:8080/api/v1/order/create`: 

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
  }

# Response example:

- `http://localhost:8080/api/v1/order/get/4`:

```
{
    "id": 4,
    "id_receiver": 1,
    "id_remitter": 20,
    "package": [],
    "delivery": {
        "id": 4,
        "origin_location": {
            "id": 8,
            "type": "remitter",
            "province": "caba",
            "city": "caba",
            "commune": "caba",
            "full_address": "caba",
            "lat": -34.659634,
            "lng": -58.50503
        },
        "destiny_address": {
            "id": 7,
            "type": "receiver",
            "province": "bsas",
            "city": "bsas",
            "commune": "bsas",
            "full_address": "bsas",
            "lat": -34.60886,
            "lng": -58.51626
        },
        "pick_up_date": "2022-11-15 15:13:17",
        "delivery_date": "2022-11-16 15:13:17"
    },
    "status": "creado",
    "creation_date": "2022-11-14 15:13:17"
}
