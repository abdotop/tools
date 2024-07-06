# Package Kryptonite

Le package `kryptonite` offre des fonctionnalités de sécurité pour le hachage et la vérification des mots de passe en utilisant des techniques cryptographiques avancées.

## Fonctionnalités

- **Génération de hachage sécurisée** : Utilise PBKDF2 avec HMAC-SHA-256 pour créer un hachage sécurisé des mots de passe.
- **Vérification de mot de passe** : Permet de comparer un mot de passe fourni avec un hachage pour vérifier l'authenticité du mot de passe.
- **Gestion des erreurs** : Utilise un canal pour gérer les erreurs de manière asynchrone.

## Installation

Pour utiliser ce package, importez-le dans votre projet Go :
```go
import "github.com/abdotop/tools/kryptonite"
```

## Utilisation

### Création d'une instance

Créez une nouvelle instance de `kryptonite` en fournissant une clé secrète et une fonction de hachage :

```go 
k, err := kryptonite.New("votre_clé_secrète", sha256.New)
if err != nil {
// Gérer l'erreur
}
```

### Génération de hachage

Générez un hachage sécurisé pour un mot de passe :

```go
hash, err := k.GenerateHash("votre_mot_de_passe", []byte("votre_sel"))
if err != nil {
// Gérer l'erreur
}
```

### Vérification de mot de passe

Vérifiez si un mot de passe correspond au hachage :

```go
err = k.CompareHashAndPassword("hachage_enregistré", "mot_de_passe_à_vérifier", []byte("votre_sel"))
if err != nil {
// Gérer l'erreur
}
```


### Gestion des erreurs

Gérez les erreurs de manière asynchrone en utilisant un callback :

```go
k.OnError(func(e error) {
 fmt.Println("Erreur détectée :", e)
})
```


## Licence

Ce package est distribué sous la licence MIT. Veuillez consulter le fichier `LICENSE` pour plus de détails.