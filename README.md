# Projet Golang

### Build du projet 
Pour build le projet, exécuter le fichier ```script\build.sh``` ou ```script\build.bat```en fonction de l'OS

### _Mosquitto_
Lancer le broker MQTT Mosquitto ```mosquitto -v``` dans un terminal.

### _Lancer les pubs, les subs de log et de la BD et l'API_
Executer le fichier ```script\run.sh``` ou ```script\run.bat``` en fonction de l'OS

### _API_
La documentation swagger de l'API est à l'adresse : ```http://localhost:8080/swagger/index.html ```

Pour mettre à jour la doc swagger : ```swag init```
### IHM
Pour lancer l'IHM ```npm run dev```