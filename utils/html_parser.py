#!/usr/bin/env python3

from bs4 import BeautifulSoup
import requests

keycloak_doc = requests.get("http://www.keycloak.org/docs-api/3.2/rest-api/index.html").text

soup = BeautifulSoup(keycloak_doc, "lxml")

for div in soup.body.find_all('div'):
    if div.h2:
        if div.h2.string == "Definitions":
            definitions=div.div
            break

for div in definitions.find_all('div'):
    name = div.h3.string
    print("\n",name,"\n")
    #print(div)
    with open("./resources/{}".format(name), "w") as f :
        try:
            for tr in div.table.tbody.find_all('tr'):
                tds = tr.find_all('td')
                field_name = tds[0].p.strong.string
                field_req = tds[0].p.em.string
                if tds[1].p.string:
                    field_type = tds[1].p.string
                else:
                    field_type = tds[1].p.a.string
                if field_req not in set(['optional','required']):
                    field_req = 'optional'
                print(field_name + ' ' + field_req + ' ' + field_type)
                f.write(field_name + ' ' + field_req + ' ' + field_type + '\n')

        except :
            print("WARNING!!!", name)

#print(definitions.find_all('div')[-1])
