# API Documentation

This is the API documentation for the lightweight Go server.

## Table of Contents
- [Models](#models)
- [Endpoints](#endpoints)
  - [Add Person](#add-person)
  - [Search Person](#search-person)
  - [Get Person](#get-person)
  - [Count People](#count-people)

## Models
The following models are used in this API:

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
Adds a new person to the database.

**Request:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "ID": "string",
  "Name": "string",
  "Nickname": "string",
  "Birth": "string",
  "CreatedAt": "time.Time",
  "UpdateAt": "time.Time"
}' http://localhost:8080/pessoas
```

**Response:**
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

### Search Person
Searches for a person in the database based on a search term.

**Request:**
```bash
curl -X GET 'http://localhost:8080/pessoas?t=searchTerm'
```

**Response:**
```json
[
  {
    "ID": "string",
    "Name": "string",
    "Nickname": "string",
    "Birth": "string",
    "CreatedAt": "time.Time",
    "UpdateAt": "time.Time"
  }
]
```

### Get Person
Retrieves information about a specific person from the database.

**Request:**
```bash
curl -X GET http://localhost:8080/pessoas/:id
```

**Response:**
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

### Count People
Counts the number of people in the database.

**Request:**
```bash
curl -X GET http://localhost:8080/contagem-pessoas
```

**Response:**
```json
{
  "count": 10
}
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*