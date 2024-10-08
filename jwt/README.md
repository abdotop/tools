# Package JWT

Ce package fournit des fonctionnalités pour gérer JSON Web Tokens (JWT) en Go.

## Installation

Pour utiliser ce package, vous pouvez l'ajouter à votre projet Go en important le chemin du package dans votre code :

   ```go
   import "github.com/abdotop/tools/jwt"
   ```

## Génération des clés RSA

Avant d'utiliser ce package, vous devez générer des clés privées et publiques au format PEM pour la signature et la vérification des JWT. Voici comment vous pouvez générer ces clés :

1. **Génération de la clé privée RSA**:
   ```bash
   openssl genrsa -out private.pem 2048
   ```

2. **Génération de la clé publique RSA à partir de la clé privée**:
   ```bash
   openssl rsa -in private.pem -outform PEM -pubout -out public.pem
   ```

## Configuration des variables d'environnement

1. **Encodez vos clés en Base64**:
   ```bash
    base64 private.pem > private_key.base64
    base64 public.pem > public_key.base64

   ```

2. **Ajoutez les clés encodées à vos variables d'environnement**:
   Créez un fichier `.env` à la racine de votre projet et ajoutez-y les lignes suivantes:
   ```
    PRIVATE_KEY=$(cat  private_key.base64) <!-- content of the file -->
    PUBLIC_KEY=$(cat public_key.base64) <!-- content of the file -->
   ```

## Utilisation de jwt_tools

- **Initialisation**:
  Importez `jwt_tools` et créez une instance en spécifiant la durée de validité du token en heures.
  ```go
  import "github.com/abdotop/tools/jwt"

  func main() {
      jwtTool := jwt.New(24) // Token valide pour 24 heures
  }
  ```

- **Chargement des clés**:
  Avant de générer ou valider des tokens, chargez les clés RSA depuis les variables d'environnement.
  ```go
  err := jwtTool.LoadPrivateKeyFromEnv("PRIVATE_KEY")
  if err != nil {
      log.Fatal(err)
  }

  err = jwtTool.LoadPublicKeyFromEnv("PUBLIC_KEY")
  if err != nil {
      log.Fatal(err)
  }
  ```

- **Génération de token**:
  ```go
  token, err := jwtTool.GenerateToken("your_payload_here")
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Generated Token:", token)
  ```

- **Validation de token**:
  ```go
  claims, err := jwtTool.ValidateToken(token)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println("Claims:", claims)
  ```

### Gestion des erreurs

Gérez les erreurs de manière asynchrone en utilisant un callback :

```go
k.OnError(func(e error) {
 fmt.Println("Erreur détectée :", e)
})
```

## Conclusion
Suivez ces étapes pour configurer et utiliser `jwt_tools` pour la gestion sécurisée des tokens JWT dans vos applications Go.


## Licence

Ce package est distribué sous la licence MIT. Veuillez consulter le fichier `LICENSE` pour plus de détails.
