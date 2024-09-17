# Golang CRUD Application

## Description

Ce projet est une application CRUD (Create, Read, Update, Delete) simple développée en Go (Golang). Il permet de gérer des ressources via une API RESTful en utilisant les méthodes HTTP correspondantes : `POST`, `GET`, `PUT`, et `DELETE`.

L'objectif de ce projet est d'offrir un exemple basique mais fonctionnel d'une application backend en Go, qui peut être utilisée comme point de départ pour des projets plus complexes.

## Prérequis

Avant de pouvoir exécuter le projet, assurez-vous d'avoir les éléments suivants installés sur votre machine :

- [Golang](https://golang.org/dl/) version 1.16 ou supérieure
- Une base de données (si nécessaire, selon l'implémentation, comme MySQL, PostgreSQL, etc.)
- Un gestionnaire de dépendances Go (Go Modules est recommandé)

## Installation

1. Clonez ce dépôt sur votre machine locale :

   ```bash
   git clone https://github.com/R-Thibault/Golang---CRUD.git
   ```

2. Accédez au répertoire du projet :

   ```bash
   cd Golang---CRUD
   ```

3. Installez les dépendances du projet à l'aide de Go Modules (ou autre gestionnaire de dépendances) :

   ```bash
   go mod tidy
   ```

4. Configurez les variables d'environnement si nécessaire (par exemple pour la connexion à la base de données). Copiez le fichier `config.json.sample`, renommez le `config.json` et ajouter vos variable d'environnement, ou définissez les variables d'environnement :

   ```bash
   export DB_USER=your_database_user
   export DB_PASSWORD=your_database_password
   export DB_NAME=your_database_name
   export DB_HOST=your_database_host
   ```

## Utilisation

### Démarrer l'application

Pour démarrer l'application :

```bash
air

```
