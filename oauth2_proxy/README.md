# OAuth2 Proxy Demo

OAuth2 demo using [bitly/oauth2_proxy](https://github.com/bitly/oauth2_proxy).

## Using Github

Register a new github oauth application from [here](https://github.com/settings/applications/new).

### Install and configure nginx

```console
$ sudo apt install nginx
```

```console
$ subl /etc/nginx/sites-enabled/default 

server {
	listen      80;

	location / {
	  proxy_pass http://127.0.0.1:4180;
	  proxy_set_header Host $host;
	  proxy_set_header X-Real-IP $remote_addr;
	  proxy_set_header X-Scheme $scheme;
	  proxy_connect_timeout 1;
	  proxy_send_timeout 30;
	  proxy_read_timeout 30;
	}
}

server {
	listen      8080;
	location / {
	  add_header Content-Type text/plain;
	  return 200 "HELLO WORLD";
	}
}
```

```console
$ sudo service nginx restart 
```

### Configure OAuth2 proxy

Download and extract binary from [here](https://github.com/bitly/oauth2_proxy/releases).

```console
./oauth2_proxy \
-provider=github \
-client-id=<your_client_id> \
-client-secret=<your_client_secret> \
-login-url=https://github.com/login/oauth/authorize \
-email-domain=gmail.com \
-redirect-url=http://127.0.0.1:4180 \
-upstream=http://127.0.0.1:8080 \
-cookie-secret=secretsecret \
-cookie-secure=false \
--cookie-expire=5s
```
