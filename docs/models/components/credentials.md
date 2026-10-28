# Credentials

Credentials based on the authentication method; the shape each method needs is
listed above. The validator only checks that this is a non-empty string or an
object whose `password`, `otp`, `token`, `code`, `accessToken` and `idToken`,
when present, are strings; other fields of the object are passed through. An
object without the field its method needs is refused later as a failed sign-in,
not as a validation error: `microsoft` without `idToken` answers `400`, and a Google
token that is not a JWT answers `500`.



## Supported Types

### PasswordCredentials

```go
credentials := components.CreateCredentialsPasswordCredentials(components.PasswordCredentials{/* values here */})
```

### OtpCredentials

```go
credentials := components.CreateCredentialsOtpCredentials(components.OtpCredentials{/* values here */})
```

### GoogleCredentials

```go
credentials := components.CreateCredentialsGoogleCredentials(components.GoogleCredentials{/* values here */})
```

### MicrosoftCredentials

```go
credentials := components.CreateCredentialsMicrosoftCredentials(components.MicrosoftCredentials{/* values here */})
```

### 

```go
credentials := components.CreateCredentialsStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch credentials.Type {
	case components.CredentialsTypePasswordCredentials:
		// credentials.PasswordCredentials is populated
	case components.CredentialsTypeOtpCredentials:
		// credentials.OtpCredentials is populated
	case components.CredentialsTypeGoogleCredentials:
		// credentials.GoogleCredentials is populated
	case components.CredentialsTypeMicrosoftCredentials:
		// credentials.MicrosoftCredentials is populated
	case components.CredentialsTypeStr:
		// credentials.Str is populated
}
```
