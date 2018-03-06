$ sudo apt install nginx

$ subl /etc/nginx/sites-enabled/default 

server {
  listen      80;
  server_name your.company.com;

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
  listen      8090;
  root        /home/dipta/Downloads;

  location / {
    try_files $uri $uri/ index.html index.php =404;
  }
}

$ sudo service nginx restart 

./oauth2_proxy -client-id=... \
              -client-secret=... \
              -provider=github \
              -email-domain=gmail.com \
              -upstream=http://127.0.0.1:3000/upstream \
              -cookie-secret=secretsecret \
              -login-url=https://github.com/login/oauth/authorize \
              -cookie-secure=false \
              -redirect-url=http://127.0.0.1:4180



