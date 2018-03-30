
rawData = "0	5	10	0	11	14	13	4	11	8	8	7	1	4	12	11"
data = [int(numb) for numb in rawData.split()]

#testData = "0    2  7    0"
#data = [int(numb) for numb in testData.split()]


def redistribute(memory):

    memory_size = len(memory)

    max_index = 0
    max_value = memory[0] #always assumed to be memory
    #find max and index of max
    for i in range(len(memory)):
        if memory[i] > max_value:
            max_value = memory[i]
            max_index = i

    #reset memory at max
    memory[max_index] = 0
    next_block = 1
    while max_value:

        memory[(max_index + next_block) % memory_size] += 1

        next_block += 1
        max_value -= 1

    return memory

states = {}
count = 0
while True:
    state = str(data)
    #print(str(data))
    #input("Press Enter to continue...")
    if not state in states:
        states[state] = 1
        #count += 1
    else:
        if states[state] == 2:
            break
        else:
            states[state] += 1
            count += 1
    data = redistribute(data)

print(count)
