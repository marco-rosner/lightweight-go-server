[![genDocsGPT](https://img.shields.io/badge/Doc%20generated%20by-genDocsGPT-blue)](https://github.com/marco-rosner/genDocsGPT)

# API Documentation

## Table of Contents
- [Models](#models)
- [Endpoints](#endpoints)
  - [Add Person](#add-person)
  - [Search Person](#search-person)
  - [Get Person](#get-person)
  - [Count People](#count-people)

## Models
### Person
| Field      | Type      | Description                                |
|------------|-----------|--------------------------------------------|
| ID         | string    | Unique identifier for the person            |
| Name       | string    | Name of the person                          |
| Nickname   | string    | Nickname of the person                      |
| Birth      | string    | Date of birth of the person (required)      |
| CreatedAt  | time.Time | Timestamp of when the person was created    |
| UpdateAt   | time.Time | Timestamp of when the person was last updated |

## Endpoints

### Add Person
- Method: POST
- URL: /pessoas

**Request Body:**
```json
{
  "id": "string",
  "name": "string",
  "nickname": "string",
  "birth": "string"
}
```

**Example:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "id": "123",
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01"
}' http://localhost:8080/pessoas
```

### Search Person
- Method: GET
- URL: /pessoas

**Query Parameters:**
- t: Search term (required)

**Example:**
```bash
curl -X GET 'http://localhost:8080/pessoas?t=john'
```

### Get Person
- Method: GET
- URL: /pessoas/:id

**URL Parameters:**
- id: Person ID (required)

**Example:**
```bash
curl -X GET 'http://localhost:8080/pessoas/123'
```

### Count People
- Method: GET
- URL: /contagem-pessoas

**Example:**
```bash
curl -X GET 'http://localhost:8080/contagem-pessoas'
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*