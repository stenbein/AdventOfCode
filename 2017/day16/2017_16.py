

def parse(programs, move):

    instruction = move[0]

    if move[0] == "s":
        programs = spin(programs, int(move[1:]))
    elif move[0] == "x":
        programs = exchange(programs, int(move[1:].split("/")[0]), int(move[1:].split("/")[1]))
    elif move[0] == "p":
        programs = partner(programs, move[1:].split("/")[0], move[1:].split("/")[1])
    else:
        raise ValueError("Unmapped parse: ", move)

    return programs


def spin(programs, index):

    length = len(programs)
    
    #if len(programs[-index:] + programs[:length-1]) != length:
        #raise ValueError("Wrong!")

    #print("-------")
    #print(length)
    #print(programs)
    #print(programs[length - (index+1):] + programs[:index])
    #print(index)
    return programs[-index:] + programs[:-index]
    

def exchange(programs, index_one, index_two):

    try:
        temp_program = programs[index_one]
        programs[index_one] = programs[index_two]
        programs[index_two] = temp_program
    except:
        raise ValueError("Wrong")

    return programs

def partner(programs, named_one, named_two):

    try:
        index = 0
        while programs[index] != named_one:
            index += 1
        index_one = index

        index = 0
        while programs[index] != named_two:
            index += 1
        index_two = index
    except:
        print(named_one, named_two, index)
        raise ValueError

        
    return exchange(programs, index_one, index_two)



#programs = ["a", 'b', 'c', 'd', 'e']
programs = ["a", 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p']
#dance_moves = ["s2"]

completed = {}

f = open("Inputs/2017_16.txt", "r")
dance_moves = f.readline().rstrip().split(",")

setdance = ''.join(programs)

total = 28 #1000000000 % cycle length
count = 0
while total:
    
    '''if ''.join(programs) in completed:
        total = total % completed[''.join(programs)]
        break
    else:
        completed[''.join(programs)] = count'''

    for move in dance_moves:
        programs = parse(programs, move)

    total -= 1
    count += 1

    if ''.join(programs) == setdance:
        print(count)
        break
        
print(programs)
print(''.join(programs))
