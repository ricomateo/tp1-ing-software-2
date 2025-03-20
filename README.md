# tp1-ing-software-2

## Introducción

Para resolver este trabajo práctico se implementó una API REST-like en Go utilizando la librería gin-gonic. 
Se incluyen 2 Dockerfiles, uno para buildear la imagen que corre el servicio, y otro para buildear la imagen que ejecuta los tests.
Para ejecutar los tests se usa Docker Compose.
La solución no cuenta con base de datos.

## Lo más desafiante del proyecto

En mi caso lo que me resultó más desafiante fue Docker, ya que no había tenido suficiente experiencia previa con la herramienta.

## Prerrequisitos

## Tests

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

Este comando ejecuta internamente `docker compose up` lo cual levanta el servicio y luego ejecuta una serie de tests que envian requests al servicio y chequean las respuestas.
 