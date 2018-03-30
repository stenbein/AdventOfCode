from pdb import set_trace as bp


#number 1
'''
def process_stream(stream, index, rank):

    total = 0
    
    lenght = len(stream)
    while index < lenght:

        if stream[index] == "<":
            index = process_garbage(stream, index + 1)
        elif stream[index] == "{":
            val, index = process_stream(stream, index + 1, rank + 1)
            total += val
        elif stream[index] == "}":
            break
        if stream[index] == "!":
            index += 2
        else:
            index += 1

    return rank + total, index

def process_garbage(stream, index):

    lenght = len(stream)
    while index < lenght:

        if stream[index] == "!":
            index += 2
        elif stream[index] == ">":
            break
        else:
            index += 1
            
    return index
'''

#number 2
def process_stream(stream, index, rank):

    total = 0
    
    lenght = len(stream)
    while index < lenght:

        if stream[index] == "<":
            val, index = process_garbage(stream, index + 1)
            total += val
        elif stream[index] == "{":
            val, index = process_stream(stream, index + 1, rank + 1)
            total += val
        elif stream[index] == "}":
            break
        if stream[index] == "!":
            index += 2
        else:
            index += 1

    return total, index

def process_garbage(stream, index):

    total = 0
    lenght = len(stream)
    while index < lenght:

        if stream[index] == "!":
            index += 2
        elif stream[index] == ">":
            break
        else:
            total += 1
            index += 1
            
    return total, index

lines = ""
f = open("Inputs/2017_09.txt", "r")
for line in f:
    lines += line
f.close()

print(process_stream(lines, 0, 0))
