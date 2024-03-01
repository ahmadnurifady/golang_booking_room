# ENIGMA CAMP 2.0 [BOOKING ROOM]

### Deskripsi

Projek ini adalah Final Project Golang Batch 12, di projek ini kita membuat aplikasi Booking Room untuk menggantikan sistem pemesanan ruangan yang saat ini masih dilakukan secara lisan dengan aplikasi berbasis digital.

---

## API Spec

### Employe

#### Create Employe

Request :

- Method : `POST`
- Endpoint : `/users`
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

- Method : `GET`
- Endpoint : `/users/:id`
- Header :
  - Accept : application/json
  - Authorization : bearer-token

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

- Method : `PUT`
- Endpoint : `/users/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
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

- Method : `DELETE`
- Endpoint : `/users/:id`
- Header :
  - Accept : application/json
  - Authorization : bearer-token
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

- Method : `GET`
- Endpoint : `/users`
- Header :
  - Accept : application/json
  - Authorization : bearer-token
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
```


### ROOM

### Create Room

Request : 

- Method : `POST`
- Endpoint : `/rooms/create`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
- Body :
```json

{
    "id": "String",
    "roomtype": "String",
    "maxcapacity": "Number",
    "facility": {
        "description": "String",
        "wifi": "String",
        "projector": "String",
        "screenProjector": "String",
        "chairs": "String",
        "tables": "String",
        "soundProof": "String",
        "smokingArea": "String",
        "television": "String",
        "ac": "String",
        "bathroom": "String",
        "coffeMaker": "String"
    },
    "status": "available"
}
```
Response :

```json
{
    "status": {
        "code": 200,
        "description": "Ok"
    },
    "data": {
    "id": "String",
    "roomtype": "String",
    "maxcapacity": "Number",
    "facility": {
        "description": "String",
        "wifi": "String",
        "projector": "String",
        "screenProjector": "String",
        "chairs": "String",
        "tables": "String",
        "soundProof": "String",
        "smokingArea": "String",
        "television": "String",
        "ac": "String",
        "bathroom": "String",
        "coffeMaker": "String"
    },
        
        "status": "available",
        "createdAt": "2023-11-15T14:15:08.960824Z",
        "updatedAt": "2023-11-15T14:15:08.959496Z"
    }
}

```

### GetByRoomType Room

Request : 

- Method : `GET`
- Endpoint : `/rooms`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
- Params :
  - Key : roomType
- Body : -


Response :

```json
{
    "status": {
        "code": 200,
        "description": "Ok"
    },
    "data": {
    "id": "String",
    "roomtype": "String",
    "maxcapacity": "Number",
    "facility": {
        "description": "String",
        "wifi": "String",
        "projector": "String",
        "screenProjector": "String",
        "chairs": "String",
        "tables": "String",
        "soundProof": "String",
        "smokingArea": "String",
        "television": "String",
        "ac": "String",
        "bathroom": "String",
        "coffeMaker": "String"
    },
        
        "status": "available",
        "createdAt": "2023-11-15T14:15:08.960824Z",
        "updatedAt": "2023-11-15T14:15:08.959496Z"
    }
}

```

### GetById Room

Request : 

- Method : `GET`
- Endpoint : `/rooms/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
- Body : -


Response :

```json
{
    "status": {
        "code": 200,
        "description": "Ok"
    },
    "data": {
    "id": "String",
    "roomtype": "String",
    "maxcapacity": "Number",
    "facility": {
        "description": "String",
        "wifi": "String",
        "projector": "String",
        "screenProjector": "String",
        "chairs": "String",
        "tables": "String",
        "soundProof": "String",
        "smokingArea": "String",
        "television": "String",
        "ac": "String",
        "bathroom": "String",
        "coffeMaker": "String"
    },
        
        "status": "available",
        "createdAt": "2023-11-15T14:15:08.960824Z",
        "updatedAt": "2023-11-15T14:15:08.959496Z"
    }
}

```


### Update Room

Request : 

- Method : `PUT`
- Endpoint : `/rooms/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
- Body : 
```json
{
    "status": {
        "code": 200,
        "description": "Ok"
    },
    "data": {
    "id": "String",
    "roomtype": "String",
    "maxcapacity": "Number",
    "facility": {
        "description": "String",
        "wifi": "String",
        "projector": "String",
        "screenProjector": "String",
        "chairs": "String",
        "tables": "String",
        "soundProof": "String",
        "smokingArea": "String",
        "television": "String",
        "ac": "String",
        "bathroom": "String",
        "coffeMaker": "String"
    },
        
        "status": "available",
        "createdAt": "2023-11-15T14:15:08.960824Z",
        "updatedAt": "2023-11-15T14:15:08.959496Z"
    }
}

```

Response :

```json
{
    "status": {
        "code": 200,
        "description": "Ok"
    },
    "data": {
    "id": "String",
    "roomtype": "String",
    "maxcapacity": "Number",
    "facility": {
        "description": "String",
        "wifi": "String",
        "projector": "String",
        "screenProjector": "String",
        "chairs": "String",
        "tables": "String",
        "soundProof": "String",
        "smokingArea": "String",
        "television": "String",
        "ac": "String",
        "bathroom": "String",
        "coffeMaker": "String"
    },
        
        "status": "available",
        "createdAt": "2023-11-15T14:15:08.960824Z",
        "updatedAt": "2023-11-15T14:15:08.959496Z"
    }
}

```

### Delete Room

Request : 

- Method : `DEL`
- Endpoint : `/rooms/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
- Body : -


Response :
```json
{
    "status": {
        "code": 200,
        "description": "ok"
    },
    "data": null
}
```


### GetAll Room

Request : 

- Method : `GET`
- Endpoint : `/rooms/get`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
- Body : -


Response :

```json
{
    "status": {
        "code": 200,
        "description": "Ok"
    },
    "data": []
}

```


### GetByStatus Room

Request : 

- Method : `GET`
- Endpoint : `/rooms/status`
- Header :
  - Content-Type : application/json
  - Accept : application/json
  - Authorization : bearer-token
- Params :
  - Key : status
- Body : -


Response :

```json
{
    "status": {
        "code": 200,
        "description": "Ok"
    },
    "data": {
    "id": "String",
    "roomtype": "String",
    "maxcapacity": "Number",
    "facility": {
        "description": "String",
        "wifi": "String",
        "projector": "String",
        "screenProjector": "String",
        "chairs": "String",
        "tables": "String",
        "soundProof": "String",
        "smokingArea": "String",
        "television": "String",
        "ac": "String",
        "bathroom": "String",
        "coffeMaker": "String"
    },
        
        "status": "available",
        "createdAt": "2023-11-15T14:15:08.960824Z",
        "updatedAt": "2023-11-15T14:15:08.959496Z"
    }
}

```

### Booking

#### Create Booking

Request :

- Method : `POST`
- Endpoint : `/booking`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "bookingDetails": [
    {
      "description": "ini desciption",
      "rooms": {
        "id": "String"
      }
    }
  ]
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "status": {
    "code": 201,
    "description": "Ok"
  },
  "data": {
    "bookingId": "String",
    "employe": {
      "id": "String",
      "name": "String",
      "divisi": "String",
      "jabatan": "String",
      "email": "String",
      "role": "String",
    },
    "bookingDetails": [
      {
        "id": "String",
        "bookingId": "String",
        "rooms": {
          "id": "String",
          "roomType": "String",
          "maxcapacity": "Number",
          "facility": {
            "id": "String",
            "description": "String",
            "wifi": "String",
            "soundSystem": "String",
            "projector": "String",
            "screenProjector": "String",
            "chairs": "String",
            "tables": "String",
            "soundProof": "String",
            "smokingArea": "String",
            "television": "String",
            "ac": "String",
            "bathroom": "String",
            "coffeMaker": "String",
          },
          "status": "available",
        },
        "description": "String",
        "status": "pending",
        "bookingDate": "2023-11-21T08:45:00.954394Z",
        "bookingDateEnd": "2023-11-21T11:45:00.954394Z",
        "createdAt": "2023-11-21T08:45:00.928153Z",
        "updatedAt": "2023-11-21T08:45:00.954394Z"
      }
    ],
    "createdAt": "2023-11-21T08:45:00.928153Z",
    "updatedAt": "2023-11-21T08:45:00.929702Z"
  }
}
```

#### Get Booking By Id

Request :

- Method : `GET`
- Endpoint : `/booking/:id`
- Header :
  - Accept : application/json

Response :


- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 201,
    "description": "Ok"
  },
  "data": {
    "bookingId": "String",
    "employe": {
      "id": "String",
      "name": "String",
      "divisi": "String",
      "jabatan": "String",
      "email": "String",
      "role": "String",
    },
    "bookingDetails": [
      {
        "id": "String",
        "bookingId": "String",
        "rooms": {
          "id": "String",
          "roomType": "String",
          "maxcapacity": "Number",
          "facility": {
            "id": "String",
            "description": "String",
            "wifi": "String",
            "soundSystem": "String",
            "projector": "String",
            "screenProjector": "String",
            "chairs": "String",
            "tables": "String",
            "soundProof": "String",
            "smokingArea": "String",
            "television": "String",
            "ac": "String",
            "bathroom": "String",
            "coffeMaker": "String",
          },
          "status": "available",
        },
        "description": "String",
        "status": "pending",
        "bookingDate": "2023-11-21T08:45:00.954394Z",
        "bookingDateEnd": "2023-11-21T11:45:00.954394Z",
        "createdAt": "2023-11-21T08:45:00.928153Z",
        "updatedAt": "2023-11-21T08:45:00.954394Z"
      }
    ],
    "createdAt": "2023-11-21T08:45:00.928153Z",
    "updatedAt": "2023-11-21T08:45:00.929702Z"
  }
}
```

#### Get All Booking

Request :

- Method : `GET`
- Endpoint : `/booking`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
 "status": {
    "code": 200,
    "description": "Ok"
  },
  "data": []booking

```

#### Get Booking By Status

Request :

- Method : `GET`
- Endpoint : `/booking/status/:status` ("pending", "accept", "decline")
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
 "status": {
    "code": 200,
    "description": "Ok"
  },
  "data": []booking

```

#### Change Status Booking By Booking Details ID

Request :

- Method : `GET`
- Endpoint : `/booking/approval`
- Header :
  - Accept : application/json
- Body :

```json
{
  "approval": "accept",  ("accept", "decline")
  "bookingDetailId": "a543db66-6df7-4cde-a14d-8bf2be6edd2f"
}
```

Response :

- Status : 201
- Body :

```json
{
  "status": {
    "code": 201,
    "description": "Ok"
  },
  "data": {
    "bookingId": "String",
    "employe": {
      "id": "String",
      "name": "String",
      "divisi": "String",
      "jabatan": "String",
      "email": "String",
      "role": "String",
    },
    "bookingDetails": [
      {
        "id": "String",
        "bookingId": "String",
        "rooms": {
          "id": "String",
          "roomType": "String",
          "maxcapacity": "Number",
          "facility": {
            "id": "String",
            "description": "String",
            "wifi": "String",
            "soundSystem": "String",
            "projector": "String",
            "screenProjector": "String",
            "chairs": "String",
            "tables": "String",
            "soundProof": "String",
            "smokingArea": "String",
            "television": "String",
            "ac": "String",
            "bathroom": "String",
            "coffeMaker": "String",
          },
          "status": "available",
        },
        "description": "String",
        "status": "decline",
        "bookingDate": "2023-11-21T08:45:00.954394Z",
        "bookingDateEnd": "2023-11-21T11:45:00.954394Z",
        "createdAt": "2023-11-21T08:45:00.928153Z",
        "updatedAt": "2023-11-21T08:45:00.954394Z"
      }
    ],
    "createdAt": "2023-11-21T08:45:00.928153Z",
    "updatedAt": "2023-11-21T08:45:00.929702Z"
  }
}
```
