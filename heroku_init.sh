#/bin/sh

echo ${PWD##*src/} > .godir
echo "web: $(basename `pwd`)" > Procfile
#heroku create -b https://github.com/kr/heroku-buildpack-go.git $(basename `pwd`)

