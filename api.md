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
```json
{
  "ID": "string",
  "Name": "string",
  "Nickname": "string",
  "Birth": "string",
  "CreatedAt": "time.Time",
  "UpdateAt": "time.Time"
}
```

## Endpoints

### Add Person
- Method: POST
- Path: /pessoas

#### Request Body
```json
{
  "ID": "string",
  "Name": "string",
  "Nickname": "string",
  "Birth": "string",
  "CreatedAt": "time.Time",
  "UpdateAt": "time.Time"
}
```

#### Example
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "ID": "1",
  "Name": "John Doe",
  "Nickname": "JD",
  "Birth": "1990-01-01",
  "CreatedAt": "2022-01-01T00:00:00Z",
  "UpdateAt": "2022-01-01T00:00:00Z"
}' http://localhost:8080/pessoas
```

### Search Person
- Method: GET
- Path: /pessoas

#### Query Parameters
- t: string (required) - The search term

#### Example
```bash
curl -X GET 'http://localhost:8080/pessoas?t=John'
```

### Get Person
- Method: GET
- Path: /pessoas/{id}

#### Path Parameters
- id: string (required) - The ID of the person

#### Example
```bash
curl -X GET 'http://localhost:8080/pessoas/1'
```

### Count People
- Method: GET
- Path: /contagem-pessoas

#### Example
```bash
curl -X GET 'http://localhost:8080/contagem-pessoas'
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*