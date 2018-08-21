#!/usr/bin/env python

import json
import urllib

user = raw_input("The box of which user do you want?:")

boxx = raw_input("The url of which box do you want?:")

url = "https://app.vagrantup.com/api/v1/box/" + user + "/" + boxx

data = urllib.urlopen(url).read()
output = json.loads(data)
print output["versions"][0]["providers"][0]["download_url"]

