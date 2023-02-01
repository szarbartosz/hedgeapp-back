# hedgeapp backend :hedgehog:

## REST api

### auth

```http
request: POST /register
body:
{
  email: string,
  password: string
}
response:
{
  message: string
  error: string
}
```
