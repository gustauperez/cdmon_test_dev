# CDmon test

This repo contains the code for the CDmon test. It's been split in two different directories, each one contains the code and stuff needed.

## Exercise 1

This one a simple test in which we have to code a simple program that has to follow this specification:

- Print all the integer numbers from 1 to 100
- If the number can be divided by 3, modulo 0, instead of the number the program has to print 'CD'
- On the other hand, if the number can be divided by 5, modulo 0, instead of the number the program has to print 'mon'
- If the number can be divided by both, the string 'CDmon' has to be printed

Three different scripts are provided. You can find them under the *ex1* directory.

- The bash script can be run with *./prova1.sh*. It uses parameter substitution to print the default value (when the integer can't be divided neither by 3 nor by 5)
- The second one uses awk. It can be run with *awk -f prova1.awk*
- The third one uses golang. It has to be compiled and run: *go build prova.go && ./prova*

## Exercise 2

Here you can find a brief specification of the API:

- *GET* */hostings/*: This will list all the hostings in the 'database'.
- *GET* */hosting/{id}*: This will fetch a single hosting. If not found a 404 will be returned.
- *PUT* */hosting/{id}*: This will put a new hosting or replace an existing one. The new hosting and an HTTP 200 will be returned in both cases.
- *DELETE* */hoting/{id}*: This will delete an already existing hosting. If not found it will do nothing. In either case a 204 will be returned (which is, no content returned because a hosting has been removed).

To solve this one, golang has been used. To handle the API, the gorilla mux golang package has been used. The API needs an HTTP header token (X-Session-Token) to validate the API call. 

### Deployment

For this one you'd need a personal Digital Ocean key. Remember that you need to install docker in your machine (your mileage may vary, for Linux check your package manager, for Windows and MacOs, go to the official site and download them). After that, open a command line and run the following commands (I’m assuming a UNIX machine, for Windows probably you’d need to use back slashes instead):

```
# docker-machine create --driver digitalocean \ --digitalocean-access-token ${DIGITAL_OCEAN_TOKEN} ${DROPLET_NAME} [1]
# eval $(docker-machine env ${DROPLET_NAME})
# cd ${GUS_GITHUB_REPO}/ex2
# docker-compose build --up -d
```
Please replace **${DIGITAL_OCEAN_TOKEN}** with your personal token generated in your Digital Ocean account and **${DROPLET_NAME}** with the name of your droplet. **${GUS_GITHUB_REPO}** is the directory where you cloned the repo.

This will connect to your droplet, build the latest image of the official docker container from Apache (which uses the latest version of 2.4, tagged as latest) and run it in the droplet.

If you want to save yourself from the hassle of deploying it by yourself, you can check it in:

	http://174.138.53.103:8000/

### Examples

- Fetching all the hostings:

```
MacBook-Air-de-Gustau:cdmon_test_dev Gus$ curl -XGET -H "X-Session-Token:23994ff1197350ee94e0052d21bff2a21154846" -vv http://localhost:8000/hostings
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> GET /hostings HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Mon, 12 Feb 2018 19:58:13 GMT
< Content-Length: 144
< Content-Type: text/plain; charset=utf-8
<
[{"id":"1","name":"Hosting1","cores":"2","memory":"4096","disc":"1TB"},{"id":"2","name":"Hosting2","cores":"4","memory":"8192","disc":"500MB"}]
* Connection #0 to host localhost left intact
```

- Fetching a hosting:

```
MacBook-Air-de-Gustau:cdmon_test_dev Gus$ curl -XGET -H "X-Session-Token:23994ff1197350ee94e0052d21bff2a21154846" -vv http://localhost:8000/hosting/2
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> GET /hosting/2 HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Mon, 12 Feb 2018 19:57:00 GMT
< Content-Length: 72
< Content-Type: text/plain; charset=utf-8
<
{"id":"2","name":"Hosting2","cores":"4","memory":"8192","disc":"500MB"}
* Connection #0 to host localhost left intact
```

- Creating a hosting:

```
MacBook-Air-de-Gustau:cdmon_test_dev Gus$ curl -XPUT -vv -H "X-Session-Token:23994ff1197350ee94e0052d21bff2a21154846" -d '{"name":"Hosting4","cores":"3","memory":"4096","Disc":"1TB"}' http://localhost:8000/hosting/3
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> PUT /hosting/3 HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Length: 60
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 60 out of 60 bytes
< HTTP/1.1 200 OK
< Date: Mon, 12 Feb 2018 19:59:26 GMT
< Content-Length: 70
< Content-Type: text/plain; charset=utf-8
<
{"id":"3","name":"Hosting4","cores":"3","memory":"4096","disc":"1TB"}
* Connection #0 to host localhost left intact
```

- Updating a hosting:

```
MacBook-Air-de-Gustau:cdmon_test_dev Gus$ curl -XPUT -vv -H "X-Session-Token:23994ff1197350ee94e0052d21bff2a21154846" -d '{"name":"Hosting5","cores":"4","memory":"16384","Disc":"1TB"}' http://localhost:8000/hosting/3
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> PUT /hosting/3 HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Length: 61
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 61 out of 61 bytes
< HTTP/1.1 200 OK
< Date: Mon, 12 Feb 2018 20:00:56 GMT
< Content-Length: 70
< Content-Type: text/plain; charset=utf-8
<
{"id":"3","name":"Hosting4","cores":"3","memory":"4096","disc":"1TB"}
* Connection #0 to host localhost left intact
```

- Removing a hosting:

```
MacBook-Air-de-Gustau:cdmon_test_dev Gus$ curl -XDELETE -vv -H "X-Session-Token:23994ff1197350ee94e0052d21bff2a21154846" http://localhost:8000/hosting/3
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> DELETE /hosting/3 HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 204 No Content
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Mon, 12 Feb 2018 20:01:40 GMT
<
* Connection #0 to host localhost left intact
```

[1] eval $(docker-machine env ${DROPLET_NAME)}
