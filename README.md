# Fiber_CRUD_with_mongodb
CRUD Application in Golang with Fiber

Features:
- Import data from a CSV file into MongoDB.
- Concurrent goroutines for faster data import.
- CRUD operations (Create, Read, Update, Delete) for MongoDB.


## How to setup
1-Clone the project
<br/>
2-Open the project folder
<br/>
3-In your command line, run `docker-compose up --build` for building the project
<br/>
## How to use
### Full API documentation is available through Swagger.
You can access it by visiting the following link, after the build progress finished:
<br/>
#### [localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)
#### Important Note: if you want to test endpoints by swagger, be inform that the "series endpoint with GET method" might display response with delay (only if you generate large amount of json data, like using csv-to-mongo endpoint), for better experience, i also added postman collection so there be no time delay for test progress
### Postman collection is also available in project files (`Fiber_CRUD.postman_collection.json`)