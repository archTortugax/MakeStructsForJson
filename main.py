import jsonparser

def main():
    text = jsonparser.textFromJsonFile("testfile.json")
    data = jsonparser.dataFromJson(text)
    print(data)

if __name__ == "__main__":
    main()
