# hedgeapp backend :hedgehog:

Backend of an application created to streamline and enhance the inventory processes for individuals and teams managing green areas.

## tech stack :wrench:

![Go](https://img.shields.io/badge/go-%23000000.svg?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/gin-%23000000.svg?style=for-the-badge&logo=gin&logoColor=white)
![Postgresql](https://img.shields.io/badge/postgresql-%23000000.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%23000000.svg?style=for-the-badge&logo=docker&logoColor=white)

## dev quickstart :construction:

```bash
$ docker compose up --build -d
```

---

## project structure :deciduous_tree:

```bash
.
├── Dockerfile
├── README.md
├── controllers
│   ├── investorsController.go
│   ├── locationsController.go
│   ├── officeController.go
│   ├── statusesController.go
│   └── usersController.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── initializers
│   ├── dbConnector.go
│   ├── dbSeeder.go
│   ├── dbSynchronizer.go
│   ├── envVariablesLoader.go
│   └── seeds
│       ├── offices.json
│       └── statuses.json
├── main.go
├── middleware
│   ├── handleCORS.go
│   └── requireAuth.go
├── models
│   ├── addressModel.go
│   ├── applicationModel.go
│   ├── authModel.go
│   ├── commonModel.go
│   ├── investorModel.go
│   ├── locationModel.go
│   ├── noteModel.go
│   ├── officeModel.go
│   ├── statusModel.go
│   └── userModel.go
├── test
│   └── example_test.go
├── tmp
└── validators
    └── usersValidator.go
```

## api endpoints :book:

<details>
  <summary>Auth</summary>

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

#### validate

```http
GET /validate
```

#### delete user

```http
DELETE /users
```

</details>

<details>
  <summary>Statuses</summary>

#### get statuses

```http
GET /statuses
```

#### create status

```http
POST /statuses
```

```js
{
  name: string;
}
```

#### update status

```http
PUT /statuses/:id
```

```js
{
  name: string;
}
```

#### delete status

```http
DELETE /statuses/:id
```

</details>

<details>
  <summary>Investor</summary>

#### get investors

```http
GET /investors
```

#### get single investor

```http
GET /investors/:id
```

#### create investor

```http
POST /investors
```

```js
{
  name: string,
  contactPerson: string,
  phone: string,
  email: string,
  nip: string,
  regon: string,
  address: {
    city: string,
    street: string,
    number: string,
    zipCode: string
  }
}
```

#### update investor

```http
PUT /investors/:id
```

```js
{
  name: string,
  contactPerson: string,
  phone: string,
  email: string,
  nip: string,
  regon: string,
  address: {
    city: string,
    street: string,
    number: string,
    zipCode: string
  }
}
```

#### delete investor

```http
DELETE /investors/:id
```

</details>

<details>
  <summary>Offices</summary>

#### get offices

```http
GET /offices
```

#### get single office

```http
GET /offices/:id
```

#### create office

```http
POST /offices
```

```js
{
  name: string,
    address: {
    city: string,
    street: string,
    number: string,
    zipCode: string
  }
}
```

#### update office

```http
PUT /offices/:id
```

```js
{
  name: string,
  address: {
    city: string,
    street: string,
    number: string,
    zipCode: string
  }
}
```

#### delete office

```http
DELETE /offices/:id
```

</details>

<details>
  <summary>Locations</summary>

#### get locations

```http
GET /locations
```

#### get single location

```http
GET /locations/:id
```

#### create location

```http
POST /locations
```

```js
{
  name: string,
  issueDate: string, // ISO 8601
  inspectionDate: string, // ISO 8601
  deforestationDate: string, // ISO 8601
  plantingDate: string, // ISO 8601
  deforestationDone: boolean,
  plantingDone: boolean,
  address: {
    city: string,
    street: string,
    number: string,
    zipCode: string
  },
  statusId: int,
  investorId: int,
}
```

#### update location

```http
PUT /locations/:id
```

```js
{
  name: string,
  issueDate: string, // ISO 8601
  inspectionDate: string, // ISO 8601
  deforestationDate: string, // ISO 8601
  plantingDate: string, // ISO 8601
  deforestationDone: boolean,
  plantingDone: boolean,
  address: {
    city: string,
    street: string,
    number: string,
    zipCode: string
  },
  statusId: int,
  investorId: int,
}
```

#### delete location

```http
DELETE /locations/:id
```

</details>
