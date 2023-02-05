# hedgeapp backend :hedgehog:

## User

### register

```http
POST /register
```

```js
{
  email: string,
  password: string
}
```

### login

```http
POST /login
```

```js
{
  email: string,
  password: string
}
```

### delete user

```http
DELETE /delete/user
```

## Status

### create status

```http
POST /status
```

```js
{
  name: string,
}
```

### get statuses

```http
GET /status
```

## Developer

### create developer

```http
POST /developer
```

```js
{
  name: string,
}
```

### get developers

```http
GET /developer
```

## Locations

### create location

```http
POST /location
```

```js
{
  status_id: int,
  developer_id: int,
  name: string,
  issue_date: string, // ISO 8601
  inspection_date: string, // ISO 8601
  deforestation_date: string, // ISO 8601
  planting_date: string // ISO 8601
}
```

### get locations

```http
GET /locaiton
```
