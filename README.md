# Trainberry - Serveur

Ce dépôt contient les sources du serveur central utilisé pour contrôler les trains. Dans le détail, le serveur s'occupe
de :

- Conserver une liste des trains actuellement présents sur la table ;
- Communiquer les commandes envoyées par l'utilisateur au(x) train(s) concerné(s) ;
- Vérifier que l'ensemble des trains actuellement affichés sont bien disponibles.

# Démarrer le serveur

## Via Docker depuis DockerHub

1. Démarrez le container : `docker run -p 8080:8080 --name trainberry_server a1hex/trainberry_server:<version>`
2. Le serveur est désormais disponible sur `http://localhost:8080`. 🎉

## Via Docker depuis les sources (recommandé)

1. Clonez ce dépôt ;
2. Construisez l'image docker : `docker build -t trainberry_server:latest .`
3. Démarrez le container : `docker run -p 8080:8080 --name trainberry_server trainberry_server:latest`
4. Le serveur est désormais disponible sur `http://localhost:8080`. 🎉

## Depuis les sources

Pour démarrer ce serveur depuis les sources, vous devez disposer de Go 1.20 ou ultérieur sur votre machine.

Ensuite, il vous suffit de lancer `go run ./cmd`, le serveur sera disponible sur `http://localhost:8080`.

# Configuration

Aucune configuration n'est requise ! Le serveur démarre directement sur l'ensemble des IPs de votre machine sur le port
8080, et n'utilise aucune base de données. Si votre serveur fonctionne, alors vous êtes prêt.

# API

Le détail de l'API exposée est documentée dans un fichier Swagger dans le dossier `api/`.