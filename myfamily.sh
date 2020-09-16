curl https://ytrack.learn.ynov.com/assets/superhero/all.json | jq ' .[]  | select( .id == '$HERO_ID' ) | .connections | .relatives ' | tr -d '"'
