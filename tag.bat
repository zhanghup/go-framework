git add .
git commit -m remote
git push


set version=v%1%
git tag %version%
git push origin %version%
