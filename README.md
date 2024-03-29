# hedgeapp backend :hedgehog:

## dev quickstart :construction:

```bash
$ docker-compose up --build -d
```

---

## docs

### User

#### register

```http
POST /register
```

```js
{
  email: string,
  password: string
}
```

#### login

```http
POST /login
```

```js
{
  email: string,
  password: string
}
```

#### delete user

```http
DELETE /delete/users
```

### Status

#### create status

```http
POST /statuses
```

```js
{
  name: string,
}
```

#### get statuses

```http
GET /statuses
```

### Developer

#### create developer

```http
POST /developers
```

```js
{
  name: string,
}
```

#### get developers

```http
GET /developers
```

### Location

#### create location

```http
POST /locations
```

```js
{
  status_id: int,
  developer_id: int,
  name: string,
  issue_date: string, // ISO 8601
  inspection_date: string, // ISO 8601
  deforestation_date: string, // ISO 8601
  planting_date: string, // ISO 8601
  deforestation_done: boolean,
  planting_done: boolean
}
```

#### update location

```http
PUT /locations/:id
```

```js
{
  status_id: int,
  developer_id: int,
  name: string,
  issue_date: string, // ISO 8601
  inspection_date: string, // ISO 8601
  deforestation_date: string, // ISO 8601
  planting_date: string, // ISO 8601
  deforestation_done: boolean,
  planting_done: boolean
}
```

#### get locations

```http
GET /locations
```

#### get single location

```http
GET /locations/:id
```

#### delete location

```http
DELETE /locations/:id
```
