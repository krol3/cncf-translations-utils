# cncf-translations
CNCF translations utils for the Spanish group team

### Requisites

You can test out a GitHub Personal Access Token using this simple example.
// You can generate them here: https://github.com/settings/tokens

### Run

````
go mod tidy
go run main.go
````

Validate [this issue sample here](https://github.com/kubernetes/website/issues/42736)

```
go run ./main.go
Enter yaml file: k8s-es.yml
GitHub Token: 
--- Create issue in kubernetes/website 
fileTarget: docs/concepts/security/security-checklist/
Issue Title: [es] Translate - `docs/concepts/security/security-checklist/` into Spanish
Issue Labels: [kind/feature language/es]
Issue Body:
**This is a Feature Request**
Translate this page [docs/concepts/security/security-checklist/](https://kubernetes.io/docs/concepts/security/security-checklist/) into Spanish.

**What would you like to be added**
The translation will be added to `content/es/`docs/concepts/security/security-checklist/

**Why is this needed**
There is no Spanish localization for this file. It would be great if we could read this in Spanish.

**Comments**
The Spanish translation team would like to facilitate your contribution's journey,  any doubt participate in the weekly meetings in the kubernetes slack channel `kubernetes-docs-es`

We also encourage new contributors to participate in their native speaking language 💥.

/triage accepted
/kind feature
/language es
/sig docs
----------------
```


## Github API Docs

- [Create an issue](https://docs.github.com/en/free-pro-team@latest/rest/issues/issues?apiVersion=2022-11-28#create-an-issue)

## Golang libraries

- https://github.com/google/go-github
- https://github.com/google/go-github/tree/master/example