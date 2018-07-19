## nginx to go_server

>About: This is a simple demo which can run in a stand-alone machine. Hope this demo is helpful to you.

### How to use

First,you should install [docker](www.docker.com) and `docker-compose`.

Clone this repository:

```bash
git clone git@github.com:7Ethan/nginx2go_server.git
```

Then you can modify the file accroding to your requirement,such as `nginx.conf`,webserver etc.

Build you http server container:
```bash
$ cd web
$ make build
```

Start you server:

```bash
$ docker-compose build --no-cache 
$ docker-compose up 
```

Finally check your http server or nginx in browser.

### License
MIT
