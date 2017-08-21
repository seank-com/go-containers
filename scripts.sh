
start-go()
{
  if [[ $(docker ps -aqf name=go-interactive) =~ ^[:space:]*$ ]]; then
    docker run -it --name go-interactive -v $PWD/go:/go golang bash
  else
    docker start -ai go-interactive
  fi
}
