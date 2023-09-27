[![genDocsGPT](https://img.shields.io/badge/Doc%20generated%20by-genDocsGPT-blue)](https://github.com/marco-rosner/genDocsGPT)

# API Documentation

## Table of Contents
- [Models](#models)
- [Handlers](#handlers)

## Models
### Person
| Field      | Type     | Description     |
|------------|----------|-----------------|
| ID         | string   | Unique identifier of the person |
| Name       | string   | Name of the person (required) |
| Nickname   | string   | Nickname of the person (required, unique) |
| Birth      | string   | Date of birth of the person (required) |
| CreatedAt  | time.Time| Time when the person was created |
| UpdateAt   | time.Time| Time when the person was last updated |

## Handlers

### CountPeople
Count the number of people in the database.

**URL:** `/contagem-pessoas`

**Method:** `GET`

**Response:**
```json
200 OK
{
  "count": 10
}
```

### GetPerson
Get a person by their ID.

**URL:** `/pessoas/:id`

**Method:** `GET`

**Response:**
```json
200 OK
{
  "id": "1",
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01",
  "createdAt": "2021-01-01T00:00:00Z",
  "updatedAt": "2021-01-01T00:00:00Z"
}
```

### SearchPerson
Search for people by a given term.

**URL:** `/pessoas`

**Method:** `GET`

**Query Parameters:**
- `t` (required): The search term

**Response:**
```json
200 OK
[
  {
    "id": "1",
    "name": "John Doe",
    "nickname": "johnd",
    "birth": "1990-01-01",
    "createdAt": "2021-01-01T00:00:00Z",
    "updatedAt": "2021-01-01T00:00:00Z"
  },
  {
    "id": "2",
    "name": "Jane Smith",
    "nickname": "janes",
    "birth": "1995-02-01",
    "createdAt": "2021-02-01T00:00:00Z",
    "updatedAt": "2021-02-01T00:00:00Z"
  }
]
```

### AddPerson
Add a new person to the database.

**URL:** `/pessoas`

**Method:** `POST`

**Request Body:**
```json
{
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01"
}
```

**Response:**
```json
201 Created
{
  "id": "1",
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01",
  "createdAt": "2021-01-01T00:00:00Z",
  "updatedAt": "2021-01-01T00:00:00Z"
}
```

Note: All requests and responses are in JSON format.

## Examples

### CountPeople
**JSON Example:**
```json
GET /contagem-pessoas
```

**Curl Example:**
```bash
curl -X GET http://localhost:8080/contagem-pessoas
```

### GetPerson
**JSON Example:**
```json
GET /pessoas/1
```

**Curl Example:**
```bash
curl -X GET http://localhost:8080/pessoas/1
```

### SearchPerson
**JSON Example:**
```json
GET /pessoas?t=john
```

**Curl Example:**
```bash
curl -X GET http://localhost:8080/pessoas?t=john
```

### AddPerson
**JSON Example:**
```json
POST /pessoas
{
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01"
}
```

**Curl Example:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01"
}' http://localhost:8080/pessoas
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*