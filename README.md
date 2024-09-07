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

### Investor

#### create investor

```http
POST /investors
```

```js
{
  name: string,
}
```

#### get investors

```http
GET /investors
```

### Location

#### create location

```http
POST /locations
```

```js
{
  statusId: int,
  investorId: int,
  name: string,
  issueDate: string, // ISO 8601
  inspectionDate: string, // ISO 8601
  deforestationDate: string, // ISO 8601
  plantingDate: string, // ISO 8601
  deforestationDone: boolean,
  plantingDone: boolean
}
```

#### update location

```http
PUT /locations/:id
```

```js
{
  statusId: int,
  investorId: int,
  name: string,
  issueDate: string, // ISO 8601
  inspectionDate: string, // ISO 8601
  deforestationDate: string, // ISO 8601
  plantingDate: string, // ISO 8601
  deforestationDone: boolean,
  plantingDone: boolean
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
