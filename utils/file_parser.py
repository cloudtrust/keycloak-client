#!/usr/bin/env python3

import sys
import re

#Parameter required or optional
OPT_TOKEN = 'optional'
REQ_TOKEN = 'required'
REQ_TOKENS = set([OPT_TOKEN,REQ_TOKEN])

#Direct Type tokens (directly mapped to golang)
DIRECT_INT_TYPE_TOKENS = set(['int16','int32','int64'])

#Typed Arrays
STL_OPEN_TOKEN = '<'
STL_CLOSE_TOKEN = '>'
ARRAY_TOKEN = 'array'

#Type tokens
MAP_TYPE_TOKEN = 'Map'
INT_TYPE_TOKEN = 'integer'
BOOL_TYPE_TOKEN = 'boolean'
STRING_TYPE_TOKEN = 'string'
ENUM_TYPE_TOKEN = 'enum'
TYPE_TOKENS = set([MAP_TYPE_TOKEN,INT_TYPE_TOKEN,BOOL_TYPE_TOKEN,STRING_TYPE_TOKEN, ENUM_TYPE_TOKEN])

MAP_TOKEN_TO_GO = {
    MAP_TYPE_TOKEN:"map[string]interface{}",
    BOOL_TYPE_TOKEN:"bool",
    STRING_TYPE_TOKEN:"string",
    ENUM_TYPE_TOKEN:"string",
}

def panic(*args):
    eprint(*args)
    sys.exit(1)

def eprint(*args):
    print(*args, file=sys.stderr)

def to_go_array(go_type):
    return "[]%s" % go_type

def is_req(split_line):
    req = split_line.pop(0)
    if req not in REQ_TOKENS:
        panic("Token %s does not match requirement tokens" % req)
    if req == REQ_TOKEN:
        return True
    return False

def get_type(split_line):
    token = split_line.pop(0)
    if token in TYPE_TOKENS:
        if token == INT_TYPE_TOKEN:
            token = split_line.pop()
            if token not in DIRECT_INT_TYPE_TOKENS:
                panic("%s is wrong int type!" % token)
            return token 
        return MAP_TOKEN_TO_GO[token]
    if token == STL_OPEN_TOKEN:
        return to_go_array(get_type(split_line))
    return token
            
def parse_line(split_line):
    field_name = split_line.pop(0)
    if field_name == type:
        field_name = "typ"
    exported_field_name = field_name[0].upper() + field_name[1:]
    required = is_req(split_line)
    go_type = get_type(split_line)
    if required:
        res = """\t%s *%s `json:"%s"`\n""" % (exported_field_name,go_type,field_name)
    else:
        res = """\t%s *%s `json:"%s,omitempty"`\n""" % (exported_field_name,go_type,field_name)
    return res


def main():
    if len(sys.argv) < 2:
        eprint("Usage : %s file" % sys.argv[0])
        return
    res = "type %s struct { \n" % sys.argv[1]
    with open(sys.argv[1], 'r') as table:
        for line in table.readlines():
            res += parse_line([x for x in re.split('\s+|\(|\)',line) if x])
    res += "}\n"
    print(res)
           
if __name__ == "__main__":
    main()
