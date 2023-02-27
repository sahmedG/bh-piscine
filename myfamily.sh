export HERO_ID=$HERO_ID
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq --arg id "$HERO_ID" ".[] | select(.id==$HERO_ID) | .connections.relatives" | tr -d '"'