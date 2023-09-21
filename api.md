# API Documentation

## Table of Contents
- [Models](#models)
  - [Person](#person)
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
**POST /pessoas**

Adds a new person.

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
  "Name": "John",
  "Nickname": "Johnny",
  "Birth": "1990-01-01",
  "CreatedAt": "2022-01-01T12:00:00Z",
  "UpdateAt": "2022-01-01T12:00:00Z"
}' http://localhost:8080/pessoas
```

### Search Person
**GET /pessoas**

Searches for a person.

#### Query Parameters
- t: The search term.

#### Example
```bash
curl -X GET 'http://localhost:8080/pessoas?t=john'
```

### Get Person
**GET /pessoas/:id**

Gets a person by ID.

#### Path Parameters
- id: The ID of the person.

#### Example
```bash
curl -X GET 'http://localhost:8080/pessoas/1'
```

### Count People
**GET /contagem-pessoas**

Counts the number of people.

#### Example
```bash
curl -X GET 'http://localhost:8080/contagem-pessoas'
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*