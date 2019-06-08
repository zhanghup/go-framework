#!/usr/bin/env bash

git add .
git commit -m remote
git push


version=v0.0.7
git tag $version
git push origin $version
