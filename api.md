# API Documentation

## Table of Contents
- [Models](#models)
- [Person](#person)
  - [Add Person](#add-person)
  - [Get Person](#get-person)
  - [Search Person](#search-person)
  - [Count People](#count-people)

## Models
### Person
| Field      | Type      | Description              |
|------------|-----------|--------------------------|
| ID         | string    | ID of the person         |
| Name       | string    | Name of the person       |
| Nickname   | string    | Nickname of the person   |
| Birth      | string    | Birth date of the person |
| CreatedAt  | time.Time | Creation timestamp       |
| UpdatedAt  | time.Time | Update timestamp         |

## Person
### Add Person
Add a new person to the database.

**Request**
```bash
curl -X POST \
  http://localhost:8080/pessoas \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "1",
    "name": "John Doe",
    "nickname": "johndoe",
    "birth": "1990-01-01"
}'
```

**Response**
```json
{
  "id": "1",
  "name": "John Doe",
  "nickname": "johndoe",
  "birth": "1990-01-01",
  "createdAt": "2021-01-01T00:00:00Z",
  "updatedAt": "2021-01-01T00:00:00Z"
}
```

### Get Person
Get a person by ID.

**Request**
```bash
curl -X GET http://localhost:8080/pessoas/1
```

**Response**
```json
{
  "id": "1",
  "name": "John Doe",
  "nickname": "johndoe",
  "birth": "1990-01-01",
  "createdAt": "2021-01-01T00:00:00Z",
  "updatedAt": "2021-01-01T00:00:00Z"
}
```

### Search Person
Search for a person by term.

**Request**
```bash
curl -X GET http://localhost:8080/pessoas?t=john
```

**Response**
```json
[
  {
    "id": "1",
    "name": "John Doe",
    "nickname": "johndoe",
    "birth": "1990-01-01",
    "createdAt": "2021-01-01T00:00:00Z",
    "updatedAt": "2021-01-01T00:00:00Z"
  }
]
```

### Count People
Get the total number of people in the database.

**Request**
```bash
curl -X GET http://localhost:8080/contagem-pessoas
```

**Response**
```json
10
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*