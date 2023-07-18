# user service

## todo

- log centralize
- github actions
- image push
- readme update

## notes

```
openssl genrsa -out keys/private.pem 2048
openssl rsa -in keys/private.pem -pubout -out keys/public.pem
```

```
{
  "kty": "RSA",
  "n": "0rr7AsdQ3H5qACZkVZI_DRASwARrvDgAxiGZgOntdoXrAVNPow8gw_bfkSjloS_WV16toXjJEfZOCBGDNpW2OT9qAdhRQC0o8Nc90yoahBkhv9ZaaIwYBjHB_6Q7rNNLrkkjCm1pjEmMYy-gYcJVm_FgCR2-gIER6EInXAmH08qXGuj_0yHAsjfN79MzCc_hWObrYAQeSfYUT5wHeNgMclLFzlypSQoHF4jNAeZI6v3K08ORx7LOp-AutI1fGH8uTnEsCLg8f2B4xHsP9Vq8nbnfOAcqBVGEXQg4uCCbgqda5bQyZLBkfHTy6rhVGqzlFqWNwdbtXvMESmzWvhwBtw",
  "e": "AQAB",
  "alg": "RS256"
}
```
