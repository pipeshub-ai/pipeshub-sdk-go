# Jwk

JSON Web Key (RFC 7517).
Public key for verifying JWT signatures.



## Fields

| Field                            | Type                             | Required                         | Description                      | Example                          |
| -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- |
| `Kty`                            | **string*                        | :heavy_minus_sign:               | Key type                         | RSA                              |
| `Use`                            | **string*                        | :heavy_minus_sign:               | Key use (sig = signature)        | sig                              |
| `Alg`                            | **string*                        | :heavy_minus_sign:               | Algorithm                        | RS256                            |
| `Kid`                            | **string*                        | :heavy_minus_sign:               | Key ID                           |                                  |
| `N`                              | **string*                        | :heavy_minus_sign:               | RSA modulus (base64url encoded)  |                                  |
| `E`                              | **string*                        | :heavy_minus_sign:               | RSA exponent (base64url encoded) | AQAB                             |