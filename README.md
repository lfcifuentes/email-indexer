# Email Indexer

[TOC]


## Tech Stack
- [Go 1.23.2](https://tip.golang.org/)
    - [Chi Router](https://github.com/go-chi/chi)
- [Search Engine](https://github.com/zincsearch/zincsearch)
- [Docker](https://www.docker.com)
    - [Docker Compose](https://docs.docker.com/compose/)


### ENVS
`ZINC_FIRST_ADMIN_USER` Usuario par el administrador de Zinc Search \
`ZINC_FIRST_ADMIN_PASSWORD` Clave par el administrador de Zinc Search \


### Build docker image

```shell
docker-compose up -d
```
Down docker 
```shell
docker-compose down

```
