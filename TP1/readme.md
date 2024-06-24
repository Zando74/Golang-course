# TP d'initiation au langage GO

Ce TP vise à vous familiariser avec les différents aspects du langage GO. Vous serez amené à résoudre une série d'exercices de difficulté croissante, explorant les concepts de variables, fonctions, tableaux et slices, pointeurs, goroutines, maps, etc.

## Exercice 1: Variables et fonctions

1. Déclarez une variable `age` de type `int` et initialisez-la avec votre âge.
2. Écrivez une fonction `double` qui prend un entier en paramètre et retourne le double de sa valeur.
3. Appelez la fonction `double` en passant la variable `age` en argument et affichez le résultat.

## Exercice 2: Tableaux et slices

1. Déclarez un tableau `nombres` contenant les entiers de 1 à 5.
2. Utilisez la fonction `len` pour afficher la taille du tableau `nombres`.
3. Utilisez la fonction `cap` pour afficher la capacité du tableau `nombres`.
4. Utilisez la fonction `append` pour ajouter le nombre 6 à la fin du tableau `nombres`.
5. Utilisez la fonction `append` pour ajouter les nombres 7, 8 et 9 à la fin du tableau `nombres`.
6. Écrivez une fonction `somme` qui prend un tableau d'entiers en paramètre et retourne la somme de ses éléments.
7. Appelez la fonction `somme` en passant le tableau `nombres` en argument et affichez le résultat.

## Exercice 3: Pointeurs

1. Déclarez une variable `x` de type `int` et initialisez-la avec une valeur.
2. Déclarez un pointeur `ptr` vers `x`.
3. Modifiez la valeur de `x` en utilisant le pointeur `ptr`.
4. Affichez la nouvelle valeur de `x`.

## Exercice 4: Goroutines

1. Écrivez une fonction `afficher` qui affiche les nombres de 1 à 10.
2. Utilisez une goroutine pour exécuter la fonction `afficher` en parallèle.
3. Attendez la fin de l'exécution de la goroutine avant de terminer le programme.

## Exercice 5: Channels

1. Déclarez un channel `ch` de type `int`.
2. Écrivez une fonction `envoyer` qui envoie les nombres de 1 à 5 dans le channel `ch`.
3. Écrivez une fonction `recevoir` qui reçoit les nombres du channel `ch` et les affiche.
4. Utilisez une goroutine pour exécuter la fonction `envoyer` en parallèle.
5. Utilisez une autre goroutine pour exécuter la fonction `recevoir` en parallèle.
6. Attendez la fin de l'exécution des goroutines avant de terminer le programme.

## Exercice 6: Maps

1. Déclarez une map `personnes` qui associe des noms de personnes à leur âge.
2. Ajoutez quelques entrées à la map.
3. Parcourez la map et affichez les noms et les âges des personnes.

## Exercice 7: Structures de contrôle

1. Écrivez une boucle `for` qui affiche les nombres de 1 à 10.
2. Utilisez une structure `if` pour vérifier si un nombre est pair ou impair.
3. Utilisez une structure `switch` pour afficher un message différent en fonction de la valeur d'une variable.

## Exercice 8: Fonctions avancées

1. Écrivez une fonction `factorielle` qui calcule la factorielle d'un nombre.
2. Utilisez une fonction récursive pour implémenter la fonction `factorielle`.
3. Appelez la fonction `factorielle` avec différents nombres et affichez les résultats.

## Exercice 8Bis: Closures

1. Écrivez une fonction `compteur` qui retourne une fonction interne.
2. La fonction interne doit avoir une variable locale `count` qui est initialisée à 0.
3. À chaque appel de la fonction interne, la variable `count` doit être incrémentée de 1.
4. Appelez la fonction `compteur` pour obtenir une instance de la fonction interne.
5. Appelez la fonction interne plusieurs fois et affichez la valeur de `count` à chaque appel.

## Exercice 9: Interfaces

1. Déclarez une interface `Forme` avec une méthode `aire` qui retourne l'aire de la forme.
2. Implémentez cette interface pour les types `Cercle` et `Rectangle`.
3. Créez des instances de `Cercle` et `Rectangle`, appelez la méthode `aire` et affichez les résultats.

## Exercice 10: Gestion des erreurs

1. Écrivez une fonction `diviser` qui divise deux nombres et retourne le résultat.
2. Gérez l'erreur lorsque le diviseur est égal à zéro en retournant une erreur.
3. Appelez la fonction `diviser` avec différents nombres et affichez les résultats ou les erreurs.

## Exercice 11: Tests unitaires

1. Écrivez des tests unitaires pour la fonction `somme` de l'exercice 2.
2. Vérifiez que la fonction retourne la somme correcte pour différents tableaux d'entiers.
3. Exécutez les tests et vérifiez les résultats.

## Exercice 12: Mini-Projet

Dans ce dernier exercice, nous allons mettre en œuvre le maximum de concepts que nous avons appris jusqu'à présent. Nous allons créer un mini-projet qui consiste à construire une application de gestion de tâches.

Voici les fonctionnalités que nous souhaitons implémenter :

1. Création d'une tâche : L'utilisateur doit pouvoir créer une nouvelle tâche en spécifiant un titre et une description.

2. Affichage des tâches : L'utilisateur doit pouvoir afficher la liste des tâches existantes avec leur titre et leur description.

3. Marquage des tâches terminées : L'utilisateur doit pouvoir marquer une tâche comme terminée.

4. Suppression des tâches : L'utilisateur doit pouvoir supprimer une tâche de la liste.

5. Sauvegarde des tâches : Les tâches doivent être sauvegardées dans un fichier pour pouvoir être récupérées ultérieurement.

Pour implémenter cela, nous allons utiliser les concepts suivants :

- Structures : Nous allons créer une structure `Task` avec les champs `title`, `description` et `completed` pour représenter une tâche.

- Slices : Nous allons utiliser un slice de `Task` pour stocker toutes les tâches.

- Fonctions : Nous allons créer des fonctions pour chaque fonctionnalité, par exemple `createTask`, `displayTasks`, `markTaskAsCompleted`, `deleteTask`, etc.

- Boucles et conditions : Nous utiliserons des boucles et des conditions pour itérer sur les tâches, afficher les informations et effectuer les actions demandées par l'utilisateur.

- Fichiers : Nous utiliserons la lecture et l'écriture de fichiers pour sauvegarder et récupérer les tâches.

- Interfaces : Nous pourrions également utiliser une interface pour définir les méthodes communes à toutes les fonctionnalités, par exemple `create`, `display`, `markAsCompleted`, `delete`, etc.

Ce mini-projet vous permettra de mettre en pratique tous les concepts que vous avez appris jusqu'à présent et de consolider vos connaissances en programmation avec le langage GO.

## Recommandations pour l'exercice 12: Mini-Projet

Pour enrichir votre mini-projet de gestion de tâches, vous pouvez envisager l'utilisation de Design Patterns. Voici quelques recommandations :

1. **Pattern Observer** : Utilisez ce pattern pour permettre à d'autres parties de votre application d'être informées lorsque des tâches sont créées, modifiées ou supprimées. Cela peut être utile pour mettre à jour l'interface utilisateur ou pour effectuer d'autres actions en fonction des événements liés aux tâches.

2. **Pattern Factory** : Utilisez ce pattern pour créer des instances de tâches de manière flexible et centralisée. Vous pouvez avoir une fabrique de tâches qui prend en charge la création de différents types de tâches en fonction des paramètres fournis par l'utilisateur.

3. **Pattern Strategy** : Utilisez ce pattern pour permettre à l'utilisateur de choisir différentes stratégies de tri ou de filtrage des tâches. Par exemple, vous pouvez implémenter des stratégies de tri par date, par priorité, par état, etc.

4. **Pattern MVC** : Si vous souhaitez créer une application web pour piloter votre application de gestion de tâches, vous pouvez envisager d'utiliser le pattern MVC (Modèle-Vue-Contrôleur). Cela vous permettra de séparer la logique métier (Modèle), la présentation (Vue) et la gestion des interactions utilisateur (Contrôleur).

N'oubliez pas que la forme de votre application est libre. Vous pouvez choisir de créer une interface en ligne de commande (CLI) ou une application web avec une interface utilisateur graphique (GUI). Vous pouvez également explorer d'autres options telles que les applications mobiles ou les interfaces de programmation (API) pour interagir avec votre application de gestion de tâches.
