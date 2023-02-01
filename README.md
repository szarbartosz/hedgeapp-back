# hedgeapp backend :hedgehog:

## auth

### request

```http
POST /register
```

```js
{
  email: string,
  password: string
}
```

### response

```js
{
  message: string;
  error: string;
}
```
