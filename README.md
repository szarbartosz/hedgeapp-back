# hedgeapp backend :hedgehog:

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

#### response

```js
{
  message: string;
  error: string;
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

#### response

```js
{
  token: string;
  error: string;
}
```

### delete user

```http
DELETE /delete/user
```

#### response

```js
{
  message: string;
  error: string;
}
```
