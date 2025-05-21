# People API

### Overview
The REST API is a service for storing and enriching information about people with the integration of external APIs (age, gender, and nationality determination by name

### Structure

```
ef-test/
├── cmd/
├── docs/
├── internal/
│   ├── config/
│   ├── dto/
│   ├── models/
│   ├── repository/
│   ├── server/
│   └──  service/
├── migrations/
├── pkg/
│   ├── logger/
│   └── postgres/
├── docker-compose.yml   
├── .env            
└── README.md      
```
### API Endpoints
#### Create
`POST /api/people`

Creates a new record with data enrichment from external APIs

*Request Body:*
``` json
{
   "name": "string",
   "surname": "string",
   "patronomic": "string"/null
}
```

*Response:*
``` json
{
    "id": int,
    "name": "string",
    "surname": "string",
    "patronymic": "string",
    "age": int,
    "gender": "string",
    "gender_probability": float,
    "nationality": [
        {
            "country_id": "string",
            "probability": float
        },
        ...
    ]
}
```

#### Update
`PUT /api/people/{id}`

Updates the record of an existing person

*Request Body:*
``` json
{
   "name": "string",
   "surname": "string",
   "patronomic": "string"/null
}
```

*Response:*
``` json
{
    "id": int,
    "name": "string",
    "surname": "string",
    "patronymic": "string",
    "age": int,
    "gender": "string",
    "gender_probability": float,
    "nationality": [
        {
            "country_id": "string",
            "probability": float
        },
        ...
    ]
}
```

#### Delete
`DELETE /api/people/{id}`

Delete the record of an existing person

#### Get by id
`GET /api/people/{id}`

Get the record of an existing person

*Response:*
``` json
{
    "id": int,
    "name": "string",
    "surname": "string",
    "patronymic": "string",
    "age": int,
    "gender": "string",
    "gender_probability": float,
    "nationality": [
        {
            "country_id": "string",
            "probability": float
        },
        ...
    ]
}
```

#### Get 
`GET /api/people?name=&surname=&gender=&age_min=&age_max=&page=&per_page=`

Returns a list of people with the ability to filter and paginate

*Response:*
``` json
{
    "data": [
        {
            "id": int,
            "name": "string",
            "surname": "string",
            "patronymic": "string",
            "age": int,
            "gender": "string",
            "gender_probability": float,
            "nationality": [
                {
                    "country_id": "string",
                    "probability": float
                },
                ...
            ]
        },
       ...
    ],
    "pagination": {
        "total": int,
        "current_page": int,
        "per_page": int
    }
}
```


### Technologies
- Language: Go 1.23.3
- Framework: Gin
- Database: PostgreSQL
- Documentation: Swagger (OpenAPI 3.0)
- Logger: Zap
- Configuration: .env
