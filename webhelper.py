import requests

def loadtextfromweb(link):
    response = requests.get(link)
    return response.text
