# Manually port the ngnix helm chart to an Ansible Operator

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

# To Run
```shell script
/bin/zsh ap.sh
```

or

```shell script
ansible-playbook deploy.yml
```

# Current State
WIP

This offering is limited and is just a scratch space.  In fact, the current example will fail the last Ansible task,
which is to actually deploy to k8s.  Since this is the whole point of this task, this work is unfinished in its current
form.
