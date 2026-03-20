## History

2007 first OAuth 1.0 released

## Use cases

One of them is single sign on for multiple apps

## Request Params for /oauth/authorize

### response_type 

response_type usually is `code` which is a code being returned in case of successful requests. Later that will be exchanged with a token.

## Authorization Server

    * Authenticates users(login)
    * Issues Authorization codes
    * Exchanges them for tokens
    * Validate clients
    * Returns access/id/refresh tokens

## Resource Server

Your APIs

## OAuth vs OpenID Connect

OAuth issues access tokens to your app, opendID Connect issues id tokens to your app. 

OpenID Connect is a extension to OAuth.
