[![genDocsGPT](https://img.shields.io/badge/Doc%20generated%20by-genDocsGPT-blue)](https://github.com/marco-rosner/genDocsGPT)

# API Documentation

## Table of Contents
- [Models](#models)
- [Endpoints](#endpoints)
  - [Add Person](#add-person)
  - [Search Person](#search-person)
  - [Get Person](#get-person)
  - [Count People](#count-people)

## Models<a name="models"></a>

### Person
```json
{
  "id": string,
  "name": string,
  "nickname": string,
  "birth": string,
  "createdAt": string,
  "updatedAt": string
}
```

## Endpoints<a name="endpoints"></a>

### Add Person<a name="add-person"></a>
Adds a new person to the database.

**URL:** `/pessoas`

**Method:** `POST`

**Request Body:**
```json
{
  "id": string,
  "name": string,
  "nickname": string,
  "birth": string
}
```

**Example:**
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "name": "John Doe",
    "nickname": "johnd",
    "birth": "1990-01-01"
  }' \
  http://localhost:8080/pessoas
```

### Search Person<a name="search-person"></a>
Searches for a person in the database based on a term.

**URL:** `/pessoas`

**Method:** `GET`

**Query Parameters:**
- `t` (required): The search term.

**Example:**
```bash
curl -X GET \
  "http://localhost:8080/pessoas?t=johnd"
```

### Get Person<a name="get-person"></a>
Retrieves a specific person from the database based on their ID.

**URL:** `/pessoas/:id`

**Method:** `GET`

**Path Parameters:**
- `id` (required): The ID of the person.

**Example:**
```bash
curl -X GET \
  http://localhost:8080/pessoas/1
```

### Count People<a name="count-people"></a>
Counts the number of people in the database.

**URL:** `/contagem-pessoas`

**Method:** `GET`

**Example:**
```bash
curl -X GET \
  http://localhost:8080/contagem-pessoas
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*