# dbcrudops

Le package `dbcrudops` fournit une interface simplifiée pour effectuer des opérations CRUD (Créer, Lire, Mettre à jour, Supprimer) sur une base de données en utilisant GORM, un ORM populaire pour Go.

## Fonctionnalités

- **Gestion des erreurs** : Utilisation d'un canal pour gérer et propager les erreurs de manière asynchrone.
- **Migration** : Permet de migrer des modèles dans la base de données.
- **Opérations CRUD** : Fonctions simplifiées pour créer, lire, mettre à jour, et supprimer des données.
- **Exécution de SQL** : Exécute des requêtes SQL directement.
- **Recherche par clé** : Permet de trouver des données en fonction d'une clé spécifique.

## Utilisation

### Création d'une instance

Pour utiliser les opérations définies, commencez par créer une instance de `Operator` :

```go
db, := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
op := New(db)
```

### Gestion des erreurs

Vous pouvez gérer les erreurs de manière asynchrone en utilisant la méthode `OnError` :
```go
op.OnError(func(err error) {
    fmt.Println("Erreur rencontrée:", err)
})
```

### Migration

Pour migrer des modèles :
```go
err := op.Migrate(&UserModel{})
if err != nil {
// Gérer l'erreur
}
```

### Opérations CRUD

Exemple de création de données :
```go
user := User{Name: "John Doe"}
err := op.Create(&user)
if err != nil {
// Gérer l'erreur
}
```

## Licence

Ce package est distribué sous la licence MIT. Voir le fichier `LICENSE` pour plus d'informations.
