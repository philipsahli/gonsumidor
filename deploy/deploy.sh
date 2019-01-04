#!/bin/bash -xe

# Deploy app
oc create -f template.yaml
oc new-app --template=gonsumidor-template  --name gonsumidor
oc apply -f build.yaml
oc start-build gonsumidor-build -n gonsumidor