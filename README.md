# MT2022_PROJ03
*Florian Leroy, Steven Nativel, Filoupatir Hanna*

### Table de contenu

* [1. Description](#description)
* [2. Lancer le projet](#lancer-projet)
* [3. Tester le projet](#tester-projet)
* [3. Requêtes](#requete)



## 1 - Description du projet
<a name="description"/>
Une API de recherche de livres.

## 2 - Lancer le projet
<a name="lancer-projet"/>

docker-compose build  
docker-compose up

## 3 - Tester le projet
<a name="tester-projet"/>

Pour aller sur Elasticsearch : http://localhost:9200    
Pour aller sur Kibana : http://localhost:5601    
Pour aller sur l'API GO : http://localhost:8080

## 4 - Requêtes

| Méthodes | Endpoint | Action |
|---|---|---|
| POST | /books | Création d'un livre  |
| POST | /books/:id | Modification d'un article  |
| GET | /books | Récupération des livres |
| GET | /books:id | Récupération d'un livre |
| GET | /search/:value | Recherche d'un livre |
| DELETE | /books/:id | Suppression d'un livre |
