# tp1-ing-software-2

# Tabla de contenidos

1. [Introducción](#introducción)
2. [Lo más desafiante del proyecto](#lo-más-desafiante-del-proyecto)
3. [Prerrequisitos](#prerrequisitos)
4. [Tests](#tests)
5. [Comandos para buildear las imágenes](#comandos-para-buildear-las-imágenes)
6. [Cómo levantar el servicio](#cómo-levantar-el-servicio)
7. [Cómo correr los tests](#cómo-correr-los-tests)

## Introducción

Para resolver este trabajo práctico se implementó una API REST-like en Go utilizando la librería gin-gonic. 

Se incluyen 2 Dockerfiles, uno para buildear la imagen que corre el servicio, y otro para buildear la imagen que ejecuta los tests.

Para ejecutar los tests se usa Docker Compose.

La solución no cuenta con base de datos.

## Lo más desafiante del proyecto

En mi caso lo que me resultó más desafiante fue el uso de Docker Compose, ya que no había tenido suficiente experiencia previa con la herramienta.

## Prerrequisitos

Los requisitos para levantar el entorno de desarrollo son:

* Go (versión 1.23.2)
* Docker
* Docker Compose

## Tests

Se utilizó la librería [net/http](https://pkg.go.dev/net/http) para los tests E2E del servicio.

La librería permite instanciar un cliente HTTP y hacer requests, lo cual es justo lo que necesitamos para testear el servicio.

En los tests se realizan requests a cada uno de los endpoints, probando los distintos casos y chequeando tanto las respuestas como los status code.

Las instrucciones para ejecutar los tests se encuentran en la sección [Cómo correr los tests](#cómo-correr-los-tests).

## Comandos para buildear las imágenes

> [!IMPORTANT]
> 
> Para correr los Makefile targets es necesario tener un archivo `.env` con las variables de entorno.
> Se provee un archivo `.env.example` el cual se puede copiar a un `.env` con el siguiente comando
> 
> ```bash
> cp .env.example .env
> ```

Para buildear la imagen del servicio, simplemente ejecutar

```bash
make server_image
```

Para buildear la imagen de los tests, ejecutar

```bash
make test_image
```

## Cómo levantar el servicio

El servicio se levanta con el siguiente comando

```bash
make start_server
```

Esto levanta un container con la imagen `server`, el cual escucha por conexiones en el address definida por las variables de entorno `HOST` y `PORT`.

## Cómo correr los tests

Los tests se pueden ejecutar con el siguiente comando

```bash
make tests
```

**Nota:** el comando requiere que ambas imágenes (del servicio y de los tests) hayan sido buildeadas previamente.
Este comando ejecuta internamente `docker compose up` lo cual levanta el servicio y luego ejecuta una serie de tests que envian requests al servicio y chequean las respuestas.
 