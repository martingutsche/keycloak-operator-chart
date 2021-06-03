# keycloak-operator-chart
A helm chart for installing the keycloak operator into your kubernetes cluster.

## Updating keycloak-operator version
```
git subtree pull --prefix=charts/keycloak-operator/vendor/keycloak-operator/ keycloak-operator 13.0.1 --squash
```


## Updating the helm repo

```
helm package charts/keycloak-operator/
```

```
helm repo index --url https://martingutsche.github.io/keycloak-operator-chart/ .
```
