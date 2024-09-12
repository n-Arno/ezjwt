ezjwt
=====

Quickly generate JWT tokens for Supabase helm install.

Usage
-----

```
$ ezjwt --help
Usage of ezjwt:
  -issuer string
    	token issuer (mandatory)
  -kv value
    	key/value to add: ie foo=bar (can be used multiple times)
  -secret string
    	encoding secret string (mandatory)
  -years int
    	number of years before expiration (default 10)
```

Example:

```
$ export JWTSECRET="$(openssl rand 64 | base64)"
$ echo $JWTSECRET
P3Ozhwtqqqv0o8YI4HzBPbxIKpMrreepTkmC8FAz5p/pQPkYhB9ZJCSFae7/erD//IMtHWWAYhzEbGM1ApnHCA==

$ ezjwt -issuer my-supabase -years 10 -kv role=anon -secret "$JWTSECRET"
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjIwNDE1MDkzNjYsImlhdCI6MTcyNjE0OTM2NiwiaXNzIjoibXktc3VwYWJhc2UiLCJyb2xlIjoiYW5vbiJ9.X_ui1YZvNAlqNrojpkTG7ObdZe7M4Z0cafC8lUJDC1c

$ ezjwt -issuer my-supabase -years 10 -kv role=service_role -secret "$JWTSECRET"
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjIwNDE1MDk0OTUsImlhdCI6MTcyNjE0OTQ5NSwiaXNzIjoibXktc3VwYWJhc2UiLCJyb2xlIjoic2VydmljZV9yb2xlIn0.HFWzOdzxXKZzEObMtWIuKIROrfkN4mCMpplGIFIylxw
```
