#!/usr/bin/env bash
  
# user or org in vagrant cloud
echo The box of which user do you want?
read var1

# box in user or org in vagrant cloud
echo The url of which box of this user do you want?
read var2

# grab user/box json
# filter current version
# grab first privider
# grab downloar_url

curl -sL \
  https://app.vagrantup.com/api/v1/box/${var1}/${var2} | jq '.current_version.providers[0] .download_url'

echo ${?}
