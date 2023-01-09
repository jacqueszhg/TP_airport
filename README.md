# Projet Golang

### Build du projet 
Pour build le projet, exécuter le fichier ```script\build.sh``` ou ```script\build.bat```en fonction de l'OS
#### Configuration du build
```
SET INSTALL_PATH="C:\Users\lucas\airport" // le dossier pour générer tout les éxécutables
SET AIRPORT="NTE" // Le code IATA de l'aéroport
```

### Lancer tous les services
Pour lancer le broquer MQTT Mosquitto, la connexion à la base de données, les logs, et l'API exécuter le fichier ```runservices.bat``` présent dans le dossier du projet.

### Lancer tous les sensors 
Pour lancer tous les sensors d'un aéroport exécuter le fichier ```runsensors.bat``` présent dans le dossier de l'aéroport nommé avec son code IATA
#### Configuration des sensors
Chaque sensors possède un fichier `config.yml` placer avec son exécutable
```
// Connexion au broker mosquitto
mqtt:
  protocol: 'tcp'
  url: 'localhost'
  port: '1883'

sensor:
  id: 1 // Doit-être différent pour chaque sensors
  airport: 'NTE' // Code IATA
  frequency: 10 // Fréquence d'envoie de donnée tout les N secondes
  QOSLevel: 1 // Qualité d'envoie de donnée du sensors
  altitudeAirport: 100 // Utile pour la simulation de nos données
```

### IHM dans `website`
Pour lancer l'IHM : 
-   `npm install` : pour obtenir les dépendances
- ```npm run dev``` : pour lancer le projet

### Lancer les servies indépendamment
## _Mosquitto_
Lancer le broker MQTT Mosquitto ```mosquitto -v``` dans un terminal.

## _Lancer les pubs, les subs de log et de la BD et l'API_
Lancer les éxécutables un par un situer dens les dossiers  :
-   sub
-   api 
-   IATA/sensors/pressure
- IATA/sensors/temperature
- IATA/sensors/wind


## _API_
La documentation swagger de l'API est à l'adresse : ```http://localhost:8080/swagger/index.html ```

Si vous ne possédez pas la commande `swag` exécuter la commande suivante : `go install github.com/swaggo/swag/cmd/swag@latest`
Pour mettre à jour la doc swagger : ```swag init```
