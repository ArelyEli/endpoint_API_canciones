# Requerimientos
- Docker
- Docker Compose

# Instalación:
```
apt update
apt install docker.io docker-compose
git clone https://github.com/ArelyEli/endpoint_API_canciones.git
cd endpoint_API_canciones
docker-compose up -d
```
# Funcionamiento

## Creación de token valido
La creación de un nuevo token valido se puede hacer directamente en el navegador o en programas de prueba de API (Postman)
```
localhost/singup?username=<NOMBRE_USUARIO>
```
EL token solo sera valido durante 5 horas a partir de su creación, un ejemplo de una consulta valida es la siguiente.

```
localhost/singup?username=Arely
```

## Cómo evaluar la operación:
La evaluación se puede hacer directamente en el navegador o en programas de prueba de API (Postman) y mandando los paramentros artist, song y token
```sh
localhost/canciones?artist=<NOMBRE_ARTISTA>&song=<NOMBRE_CACIÓN>&token=<TOKEN_VALIDO>
```
Un ejemplo de una consulta valida es la siguiente

```
localhost/canciones?artist=Arjona&song=hongos&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFyZWx5IiwiZXhwIjoxNjY3OTkzMzI0fQ.H9aPgIA2waXftkxYqIa7GOyUuKDJU-YbGIWLkHPMV8U
```

