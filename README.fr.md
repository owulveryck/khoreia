khoreia est un project dont le but est de déployer une application à partir de son modèle topologique.
L'ordonnancement du déploiement se fait de manière décentralisée est sans plan d'orchestration.

Chaque noeud (au sens host) est indépendant et réagit et agit en fonction d'évènements.

# node

Chaque node représente un élément à déployer.

## cycle de vie

Chaque node implémente le cycle de vie dont les états sont représentés de la manière suivante:

* initial
* create
* created
* configure
* configured
* start
* started
* stop
* stopped
* delete
* deleted

## startup et shutdown

Il est possible de demander à chaque noeud d'atteindre un état particulier. Pour arriver à cet état le noeud va dérouler les éléments du cycle de vie.
Ainsi donc, pour passer à l'état started, le noeud va implémenter les méthodes create() configure() et start()

## Check

Chaque action (create, start, stop, configure, delete) dispose est implémenté via une interface qui dispose de deux méthodes `Do()` et `Check()`.

`Check()` retourne un booleen qui indique si l'état est atteint ou pas.
Si la méthode Check() de l'action retourne false, alors la méthode Do() est appelée.


# etcd

Le bus qui véhicule les évènements se base sur le daemon etcd.


