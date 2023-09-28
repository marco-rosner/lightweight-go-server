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
```json
{
  "id": "string",
  "name": "string",
  "nickname": "string",
  "birth": "string",
  "createdAt": "string",
  "updateAt": "string"
}
```

## Endpoints

### Add Person
- **URL:** `/pessoas`
- **Method:** `POST`
- **Request Body:**
```json
{
  "id": "string",
  "name": "string",
  "nickname": "string",
  "birth": "string"
}
```
- **Response:**
```json
{
  "id": "string",
  "name": "string",
  "nickname": "string",
  "birth": "string",
  "createdAt": "string",
  "updateAt": "string"
}
```
- **Example:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "id": "1",
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01"
}' http://localhost:8080/pessoas
```

### Search Person
- **URL:** `/pessoas`
- **Method:** `GET`
- **Query Parameters:**
  - `t` (required): search term
- **Response:**
```json
[
  {
    "id": "string",
    "name": "string",
    "nickname": "string",
    "birth": "string",
    "createdAt": "string",
    "updateAt": "string"
  }
]
```
- **Example:**
```bash
curl -X GET 'http://localhost:8080/pessoas?t=john'
```

### Get Person
- **URL:** `/pessoas/:id`
- **Method:** `GET`
- **Path Parameters:**
  - `id` (required): person ID
- **Response:**
```json
{
  "id": "string",
  "name": "string",
  "nickname": "string",
  "birth": "string",
  "createdAt": "string",
  "updateAt": "string"
}
```
- **Example:**
```bash
curl -X GET 'http://localhost:8080/pessoas/1'
```

### Count People
- **URL:** `/contagem-pessoas`
- **Method:** `GET`
- **Response:**
```json
{
  "count": 10
}
```
- **Example:**
```bash
curl -X GET 'http://localhost:8080/contagem-pessoas'
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*