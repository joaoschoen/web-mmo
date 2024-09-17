 npx nodemon --watch './**/*' \
  --ext 'go,html,css,templ' \
  --signal SIGTERM \
  --exec 'go run main.go'