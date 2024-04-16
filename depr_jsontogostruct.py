import json
import loadfilefromweb

go_types = {
    type(0) : "int",
    type(0.5) : "float32",
    type("") : "string",
    type({}) : "struct",
    type([]) : "[]",
    type(True) : "bool",
    type(None) : "nil",
}

def main():
    jsondata = json.loads(loadfilefromweb.loadtextfromweb(f"https://ddragon.leagueoflegends.com/cdn/{version}/data/{lang}/champion.json"))

def generate_champions_structs():
    version = "14.7.1"
    lang = "en_US"
    jsondata = json.loads(loadfilefromweb.loadtextfromweb(f"https://ddragon.leagueoflegends.com/cdn/{version}/data/{lang}/champion.json"))
    all_champions_ids = list(jsondata["data"].keys())

    for champion_id in all_champions_ids:
        champion_json = json.loads(loadfilefromweb.loadtextfromweb(f"https://ddragon.leagueoflegends.com/cdn/{version}/data/{lang}/champion/{champion_id}.json"))
        print(champion_json["data"].keys())
        write_go_struct_from_json(champion_json, f"LoLChampion{champion_id}", f"structs/struct_champion_{champion_id.lower()}.go", "structs")

def write_go_struct_from_json(json_data, struct_name, file_out = "outputstruct.go", package = "main"):
    s = generate_field_struct({struct_name : json_data}, 0)
    write_go_repr(s, file_out, package)

def generate_field_struct(field, depth):
    field_name = list(field.keys())[0]
    field_value = field[field_name]
    
    spacing = depth * "\t"
    json_field = "`json:\""+ field_name +"\"`"
    go_field = field_name[0].upper() + field_name[1:]

    if type(field_value) == type({}):
        return spacing + go_field + " struct{\n" \
        + "".join([generate_field_struct({k : v}, depth + 1) for k, v in field_value.items()]) \
        + spacing + "}" + (depth > 0) * json_field + "\n"
    elif type(field_value) == type([]):
        return spacing + go_field + " struct{\n" \
        + "".join([generate_field_struct({field_name + str(i) : v}, depth + 1) for i, v in enumerate(field_value)]) \
        + spacing + "}" + (depth > 0) * json_field + "\n"
    else:
        return spacing + go_field + " " + generate_go_type(field_value, depth) + (depth > 0) *  json_field + "\n"

def write_go_repr(s, out, package):
    struct_go_repr = "type " + s
    
    with open(out, "w", encoding="utf-8") as gofile:
        gofile.write("package " + package + "\n")
        gofile.write(struct_go_repr)

def generate_go_type(value, depth):
    if type(value) == type([]):
        if len(value) > 0:
            return go_types[type(value)] + go_types[type(value[0])]
        else:
            return go_types[type(value)] + go_types[type("")]
    return go_types[type(value)]

if __name__ == "__main__":
    main()
