 npx nodemon --watch './**/*' \
  --ext 'templ' \
  --signal SIGTERM \
  --exec "templ generate"