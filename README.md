# ENIGMA CAMP 2.0 [BOOKING ROOM]

### Deskripsi

Projek ini adalah Final Project Golang Batch 12, di projek ini kita membuat aplikasi Booking Room untuk menggantikan sistem pemesanan ruangan yang saat ini masih dilakukan secara lisan dengan aplikasi berbasis digital.

---

## API Spec

### Employe

#### Create Employe

Request :

- Method : `POST`
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "name": "string",
  "divisi": "string",
  "jabatan": "string",
  "email": "string",
  "password": "string",
  "role": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
```

#### Get Employe By Id

Request :

- Method : GET
- Endpoint : `/employees/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
```

#### Update Employee

Request :

- Method : PUT
- Endpoint : `/employees/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "name": "string",
  "divisi": "string",
  "jabatan": "string",
  "email": "string",
  "password": "string",
  "role": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
```

#### Delete Employee

Request :

- Method : DELETE
- Endpoint : `/employees/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": "OK"
}
```

#### Get List Employee

Request :

- Method : GET
- Endpoint : `/employees`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
`


``
