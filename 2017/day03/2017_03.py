
testNum = 289326

'''
#number 1
def ring_size(index):

    return (1 + 2*(index - 1)) ** 2

def ring_length(index):

    return 1 + 2*(index - 1)

target_ring = 1
while ring_size(target_ring) < testNum:

    target_ring += 1

size_length = ring_length(target_ring)

distance_to_mid = (testNum - ring_size(target_ring - 1)) % size_length

#we want distance to middle index in square
if distance_to_mid > (size_length - 1) / 2:
    distance_to_mid -= int((size_length - 1) / 2)

print(distance_to_mid + target_ring)

'''


#number 2

def ring_size(index):

    return (1 + 2*(index - 1)) ** 2

def ring_length(index):

    return 1 + 2*(index - 1)

def get_adjacent(index):

    ring_number = 1
    while ring_size(ring_number) < index:
        ring_number += 1

    first_corner = ring_size(ring_number)
    length = ring_length(ring_number)
    corners = [first_corner - (length - 1) * i for i in range(4)]

    #can't be bigger than corner[0]
    if index > corners[1]:
        
        return index - (length-1)*3 - 1
    elif index > corners[2]:
        
        return index - (length-1)*2 - 1
    else:
        
        return ring_size(ring_number - 1) - (index - ring_size(ring_number - 1)-1)

def get_inner(index):

    ring_number = 1
    while ring_size(ring_number) < index:
        ring_number += 1

    first_corner = ring_size(ring_number)
    length = ring_length(ring_number)
    corners = [first_corner - (length - 1) * i for i in range(4)]

    #can't be bigger than corner[0]
    if index == corners[0]:
        return index - (length-1)*3 - (length-2) - 1
    elif index == corners[1]:
        return index - (length-1)*2 - (length-2)*2
    elif index == corners[2]:
        return index - (length-1) - (length-2)*3 + 1
    elif index == corners[3]:
        return index - (length-2)*4 + 2
    else:
        raise ValueError("Corner wrong")


'''neighbours will be one prior, adjacent and diagionals
    diagionals are just adjacent +1 and -1. If at a corner
    you have one prior and one diagional'''
def sum_of_neighbours(index):

    total = get_value(index - 1) #the value of the preceeding index
    
    if is_corner(index):
        total += get_value(get_inner(index))
    elif is_corner(index-1):
        total += get_value(index-2)
        adjacent_index = get_adjacent(index)
        total += get_value(adjacent_index + 1)
        total += get_value(adjacent_index)
    else:
        
        adjacent_index = get_adjacent(index)
        total += get_value(adjacent_index + 1)
        total += get_value(adjacent_index)
        total += get_value(adjacent_index - 1)
        
    return total

def get_value(index):

    if index < 1:
        return 0
    elif index == 1:
        return 1
    else:
        return sum_of_neighbours(index)


def memoize(f):
    
    memo = {1:1, 2:1, 3:2, 4:4, 5:5, 6:10}
    def helper(x):
        if x not in memo:
            print("Didn't find in memo, calculating.")
            memo[x] = f(x)
        return memo[x]
    return helper

get_value = memoize(get_value)
    

def is_corner(index):

    ring_number = 1
    while ring_size(ring_number) < index:
        ring_number += 1

    first_corner = ring_size(ring_number)
    length = ring_length(ring_number)
    corners = [first_corner - (length - 1) * i for i in range(4)]

    if index in corners:
        return True
    else:
        return False


#we need the number which is bigger then test number
index_of_number = 1
current_num = 1 #first number
while current_num < testNum:
    
    current_num = get_value(index_of_number)

    print("Value is: ", current_num)
    print("Found at: ", index_of_number)
    
    input("Enter")

    index_of_number += 1



print(current_num)
