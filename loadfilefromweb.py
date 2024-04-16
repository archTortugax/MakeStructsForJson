import requests
import sys

def main():
    link = sys.argv[1]
    loadfilefromweb(link)

def loadfilefromweb(link):
    response = requests.get(link)
    with open("file.json", "w", encoding="utf-8") as file:
        file.write(response.text)

def loadtextfromweb(link):
    response = requests.get(link)
    return response.text

if __name__ == "__main__":
    main()