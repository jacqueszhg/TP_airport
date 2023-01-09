# Projet Golang

### Build du projet 
Pour build le projet, exécuter le fichier ```script\build.sh``` ou ```script\build.bat```en fonction de l'OS

### Lancer tous les services
Pour lancer le broquer MQTT Mosquitto, la connexion à la base de données, les logs, et l'API exécuter le fichier ```runservices.bat``` présent dans le dossier du projet.

### Lancer tous les sensors 
Pour lancer tous les sensors d'un aéroport exécuter le fichier ```runsensors.bat``` présent dans le dossier de l'aéroport nommé avec son code IATA 

### IHM
Pour lancer l'IHM `npm install` ```npm run dev```

### Lancer les servies indépendamment

## _Mosquitto_
Lancer le broker MQTT Mosquitto ```mosquitto -v``` dans un terminal.

## _Lancer les pubs, les subs de log et de la BD et l'API_
Executer le fichier ```script\run.sh``` ou ```script\run.bat``` en fonction de l'OS

## _API_
La documentation swagger de l'API est à l'adresse : ```http://localhost:8080/swagger/index.html ```

Pour mettre à jour la doc swagger : ```swag init```
