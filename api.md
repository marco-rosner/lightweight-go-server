# API Documentation

## Table of Contents
- [Models](#models)
- [Add Person](#add-person)
- [Search Person](#search-person)
- [Get Person](#get-person)
- [Count People](#count-people)

## Models<a name="models"></a>
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

## Add Person<a name="add-person"></a>
### Request
```bash
curl -X POST \
  http://localhost:8080/pessoas \
  -H 'Content-Type: application/json' \
  -d '{
    "ID": "1",
    "Name": "John Doe",
    "Nickname": "johnd",
    "Birth": "1990-01-01",
    "CreatedAt": "2022-01-01T12:00:00Z",
    "UpdateAt": "2022-01-01T12:00:00Z"
}'
```

### Response
```json
{
  "ID": "1",
  "Name": "John Doe",
  "Nickname": "johnd",
  "Birth": "1990-01-01",
  "CreatedAt": "2022-01-01T12:00:00Z",
  "UpdateAt": "2022-01-01T12:00:00Z"
}
```

## Search Person<a name="search-person"></a>
### Request
```bash
curl -X GET \
  'http://localhost:8080/pessoas?t=john' \
  -H 'Content-Type: application/json'
```

### Response
```json
[
  {
    "ID": "1",
    "Name": "John Doe",
    "Nickname": "johnd",
    "Birth": "1990-01-01",
    "CreatedAt": "2022-01-01T12:00:00Z",
    "UpdateAt": "2022-01-01T12:00:00Z"
  },
  {
    "ID": "2",
    "Name": "John Smith",
    "Nickname": "johns",
    "Birth": "1995-02-02",
    "CreatedAt": "2022-01-01T12:00:00Z",
    "UpdateAt": "2022-01-01T12:00:00Z"
  }
]
```

## Get Person<a name="get-person"></a>
### Request
```bash
curl -X GET \
  http://localhost:8080/pessoas/1 \
  -H 'Content-Type: application/json'
```

### Response
```json
{
  "ID": "1",
  "Name": "John Doe",
  "Nickname": "johnd",
  "Birth": "1990-01-01",
  "CreatedAt": "2022-01-01T12:00:00Z",
  "UpdateAt": "2022-01-01T12:00:00Z"
}
```

## Count People<a name="count-people"></a>
### Request
```bash
curl -X GET \
  http://localhost:8080/contagem-pessoas \
  -H 'Content-Type: application/json'
```

### Response
```json
10
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*