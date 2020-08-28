# starwars_planets

##Required:
### Docker: https://docs.docker.com/engine/install/

### Docker Compose: https://docs.docker.com/compose/install/


## Run
```bash
sudo docker-compose up -d --build
```

#### Routes:
##### GET
:3000/planets/all <br> <br>

:3000/planets <br>

{
    "name": "Tatooine",
}

OR


{
    "id": "Planet_ID",
}


##### POST
:3000/planets

###### Example
 {
    "name": "Tatooine",
    "climate": "arid",
    "terrain": "desert"
}
##### DELETE
:3000/planets

{
    "id": "Planet_ID"
}
