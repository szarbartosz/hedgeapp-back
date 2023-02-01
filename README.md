# hedgeapp backend :hedgehog:

## REST api

### auth

```http
POST /register
```

```js
body: {
  email: string,
  password: string
}
response: {
  message: string
  error: string
}
```
