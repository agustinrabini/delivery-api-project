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

# Domain models:
  
