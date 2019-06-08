git add .
git commit -m remote
git push


set version=v0.0.7
git tag %version%
git push origin %version%
