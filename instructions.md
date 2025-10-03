# Exercice de Programmation : Cookify 🍳

## 1. Présentation
**Cookify** est une application qui permet aux cuisiniers, débutants ou expérimentés, d’accéder rapidement et simplement à leurs recettes via une interface intuitive.  
Ton objectif est de développer cette application en **Go (gRPC)** pour le backend et **React (JavaScript)** pour le frontend.

---

## 2. Objectifs pédagogiques
- Découvrir et pratiquer le développement backend en **Go** avec **gRPC**.  
- Développer une interface utilisateur moderne en **React**.  
- Manipuler une base de données (PostgreSQL recommandé).  
- Mettre en place une architecture claire et maintenable.  
- Gérer l’internationalisation (français et anglais).  
- Bonus : Approfondir avec un système d’utilisateurs et d’interactions sociales.  

---

## 3. Fonctionnalités attendues

### a) Recherche de recettes
- Rechercher une recette par **nom**, **ingrédient** ou **ustensile**.  

### b) Création de recettes
- Ingrédients  
- Ustensiles  
- Temps de préparation et de cuisson  
- Étapes de préparation  
- Images illustratives  
- Notes ou conseils optionnels  

### c) Bonus
- Système d’utilisateurs (authentification).  
- Notation des recettes.  
- Commentaires.  

---

## 4. Spécifications techniques

### Backend (Go + gRPC)
- Définir un fichier `.proto` décrivant les services et messages.  
- Implémenter un serveur gRPC en Go.  
- Connecter le serveur à une base PostgreSQL.  
- Fournir des services pour :  
  - Créer, lire, mettre à jour et supprimer des recettes.  
  - Rechercher des recettes par nom, ingrédient ou ustensile.  
  - (Bonus) Gérer utilisateurs, notes et commentaires.  

### Frontend (React)
- Interface responsive (desktop + mobile).  
- Page d’accueil avec barre de recherche.  
- Formulaire de création de recette.  
- Page de détail d’une recette.  
- (Bonus) Gestion utilisateur, notation et commentaires.  

### Internationalisation
- Support **français** et **anglais** dès le départ.  
- Prévoir une structure extensible pour d’autres langues.  

---

## 5. Checklist pas-à-pas

### Étape 1 : Mise en place du projet
- [ ] Créer un repo Git pour le projet.  
- [ ] Initialiser deux dossiers : `backend/` (Go) et `frontend/` (React).  
- [ ] Configurer un `docker-compose.yml` avec PostgreSQL.  

### Étape 2 : Backend gRPC
- [ ] Installer `protoc` et les plugins Go (`protoc-gen-go`, `protoc-gen-go-grpc`).  
- [ ] Créer un fichier `recipe.proto` décrivant :  
  - Le service `RecipeService` avec méthodes `CreateRecipe`, `GetRecipe`, `ListRecipes`, `SearchRecipes`.  
  - Les messages `Recipe`, `Ingredient`, `Utensil`, etc.  
- [ ] Générer le code Go à partir du `.proto`.  
- [ ] Implémenter le serveur gRPC en Go.  
- [ ] Connecter le serveur à PostgreSQL (via `gorm` ou `pgx`).  
- [ ] Implémenter les méthodes CRUD et recherche.  
- [ ] Ajouter des tests unitaires simples.  

### Étape 3 : Frontend React
- [ ] Initialiser un projet React (`create-react-app` ou Vite).  
- [ ] Installer un client gRPC-Web (ex. `grpc-web` + `envoy` comme proxy).  
- [ ] Créer une page d’accueil avec barre de recherche.  
- [ ] Créer un formulaire de création de recette.  
- [ ] Créer une page de détail de recette.  
- [ ] Ajouter la gestion de l’internationalisation (ex. `react-i18next`).  

### Étape 4 : Intégration
- [ ] Configurer Envoy ou un proxy pour exposer gRPC au frontend.  
- [ ] Vérifier la communication entre React et le backend gRPC.  
- [ ] Tester la recherche et la création de recettes depuis l’UI.  

### Étape 5 : Bonus (optionnel)
- [ ] Ajouter un service `UserService` dans le `.proto`.  
- [ ] Implémenter l’authentification (JWT).  
- [ ] Ajouter la notation et les commentaires.  
- [ ] Déployer l’application (Heroku, Render, Railway, etc.).  

---

## 6. Livrables attendus
- Code source du backend (Go + gRPC).  
- Code source du frontend (React).  
- Fichier `docker-compose.yml`.  
- Documentation minimale :  
  - Comment lancer le projet.  
  - Description des services gRPC.  
  - Instructions pour changer la langue.  

---

## 7. Pistes d’amélioration
- Filtres avancés (temps de cuisson, difficulté).  
- Stockage d’images (S3, Cloudinary, ou local).  
- Tests unitaires et d’intégration.  
- CI/CD avec GitHub Actions.  