# SCIM 2.0 GO SDK (GO-lang)

System for Cross-domain Identity Management (**SCIM**) specification is designed to make managing 
user identities in cloud-based applications and services easier. The specification suite seeks to 
build upon experience with existing schemas and deployments, placing specific emphasis on 
simplicity of development and integration, while applying existing authentication, authorization, and 
privacy models. 

Its intent is to reduce the cost and complexity of user management operations by providing a common 
user schema and extension model, as well as binding documents to provide patterns for exchanging 
this schema using standard protocols. 
In essence: make it fast, cheap, and easy to move users in to, out of, and around the cloud.


**More Info on:**

http://www.simplecloud.info/

https://tools.ietf.org/html/rfc7644

https://tools.ietf.org/html/rfc7643


#### JAVA SDK

You can find Java SDK here

https://github.com/suvera/scim2-sdk


## Installation

This repo contains three components **client** and **server** that supports scim2 protocol.

These libraries can be included in any of your GO applications.

```
go get -u github.com/suvera/goScim2

```


## SCIM2 Client

This library contains SCIM 2.0 compatible Http client.

You can build a **client** object like below.

```
TODO: working on

```


## SCIM2 Server

scim can be enabled in your server like below

```
import (
    "http"
    "github.com/gorilla/mux"
    "github.com/suvera/goScim2/scim2"
)


r := mux.NewRouter()

d := scim2.DefaultHandler{}

scim2.Server("/scim2", r, d)

http.ListenAndServe(":8090", r)
``` 

- You need to implement interface defined in **scim2.RequestHandler** for full compatibility. **scim2.DefaultHandler** is default implementation.


## Is your server is compliant to SCIM 2.0?

Here is the tool to test the compliance level  https://github.com/suvera/scim2-compliance-test-utility
