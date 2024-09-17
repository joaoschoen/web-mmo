npx nodemon --watch './**/*' \
  --ext 'html,css,go,templ' \
  --signal SIGTERM \
  --exec "tailwind -i 'base.css' -o 'static/css/tailwind.css'"