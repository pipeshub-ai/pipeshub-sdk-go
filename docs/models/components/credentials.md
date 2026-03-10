# Credentials

Credentials based on the authentication method


## Supported Types

### PasswordCredentials

```go
credentials := components.CreateCredentialsPasswordCredentials(components.PasswordCredentials{/* values here */})
```

### OtpCredentials

```go
credentials := components.CreateCredentialsOtpCredentials(components.OtpCredentials{/* values here */})
```

### OAuthCredentials

```go
credentials := components.CreateCredentialsOAuthCredentials(components.OAuthCredentials{/* values here */})
```

### 

```go
credentials := components.CreateCredentialsStr(string{/* values here */})
```

