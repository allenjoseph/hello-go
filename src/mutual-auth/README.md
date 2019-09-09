# Mutual Authentication

## Creating certificates
Private key and Certificate Signing Request
``` bash
openssl req -new -nodes -keyout client.key -out client.csr -days 3650
openssl req -new -nodes -keyout server.key -out server.csr -days 3650
```

## Signing certificates
``` bash
openssl x509 -req -days 3650 -in client.csr -signkey client.key -extfile domains.ext -out client.crt
openssl x509 -req -days 3650 -in server.csr -signkey server.key -extfile domains.ext -out server.crt
```