#!/usr/bin/env bash
  
echo The box of which user do you want?

read var1

echo The url of which box of this user do you want?

read var2

curl -sL \
  https://app.vagrantup.com/api/v1/box/$var1/$var2 | jq '.current_version.providers[0] .download_url'

echo $?
