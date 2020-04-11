# SimpleRestApi A sample Rest Api 


### Build
./build

### RUN
./run

POST http://localhost:8080/contact
body: {
 "deviceid": "saurabhtest123",
 "name": "saurabh",
 "mobile": "9999999999"
}

GET http://localhost:8080/contact/saurabhtest123

PATCH http://localhost:8080/contact/9999999999
body
 {
 "affectedstatus": 1
 }

DELETE http://localhost:8080/contact/saurabhtest123


Similar datapoints for 
/registration
/affectedlist
/contact

