(
  npx nodemon --watch './**/*' \
  --ext 'go,html,css' \
  --signal SIGTERM \
  --exec 'go run main.go'
)&
(
  npx nodemon --watch './**/*' \
  --ext 'html,css,go' \
  --signal SIGTERM \
  --exec "tailwind -i 'base.css' -o 'static/css/tailwind.css'"
)& 
(
  npx browser-sync start --proxy 'localhost:3000' -e * --ignore=node_modules -w --reload-delay="1000" --files '**/*.html, **/*.css, **/*.go'
)
