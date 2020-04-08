#Porting Ngnix helm chart to Ansible Operator

This repository is a scratch-pad for converting the
[Bitnami ngnix helm chart](https://github.com/bitnami/charts/tree/master/bitnami/nginx) to an Ansible Operator.  There
are two main endeavours:

1) Manually port the ngnix helm chart to an Ansible Operator
2) Automatically port the ngnix helm chart to Ansible Operator



## Manually port the ngnix helm chart to an Ansible Operator

There are several advantages to using the operator-framework rooted in advanced OLM capabilities.  This is a first stab
at converting the ngnix helm chart to an ngnix ansible operator.

Steps:
1) Create an empty ansible operator using operator-sdk.
2) Copy Values.yml into roles/ngnix/defaults/main.yml
3) Copy the helm charts into roles/ngnix/templates.  Rename each file with a ".j2" extension to indicate that it now
stores jinja2 markup.
4) Manually convert each template in step #3 to something that Ansible can understand.  Simply put, helm charts offer a
whole lot of really cool functionality that Ansible doesn't natively understand.  This process includes several steps
that are explained a bit more in depth in [./conversion_steps.txt].

Finally, a deploy.yml file was created to debug the conversion of the helm templates into proper jinja2 templates that
ansible-playbook can understand.

### To Run
```shell script
/bin/zsh ap.sh
```

or

```shell script
ansible-playbook deploy.yml
```

### Current State
WIP

This offering is limited and is just a scratch space.  In fact, the current example will fail the last Ansible task,
which is to actually deploy to k8s.  Since this is the whole point of this task, this work is unfinished in its current
form.

## Automatically port the ngnix helm chart to Ansible Operator
WIP

This is also a work in progress effort.  The code for this is located in the "go" subdirectory, and right now just
exemplifies how to use the Go "text/template" and "text/parse" modules to look at the different NodeType(s) in a given
template.  For now, I just print out the different top level NodeIf sections.  Note:  this code does not yet handle
nested if statements.

The ability to take a look at NodeType is of great importance to us, since it implies the actions necessary to convert
to valid jinja2 template syntax, which Ansible Playbook can understand.  There will likely not be a 1:1 mapping, at
least to start, but that is okay.  Eventually, we may be able to become smarter about our translation technique.

In the case of NodeIf, we will be able to appropriately:
1) Change from Go Template tags like:
```go
{{ if serverBlock }}
```
to the appropriate Jinja2 substitute:
```yaml
{% if serverBlock is defined %}
```

It is important to reiterate that this is just a scratchpad, and will likely never be perfect, but it is a start in the
right direction.

### Preflighting the helm charts

Since Helm charts require some additional translation, right now I have removed "include" and sprig function references
completely.  This exercise doesn't consider those functions, but their translation should be considered at a later date.

### Running the NodeIf example
```bash
> cd go
> go run main.go
```
