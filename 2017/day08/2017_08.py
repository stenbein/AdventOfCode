


def eval_conditional(conditional):
    
    conditional = conditional.split()
    register = conditional[0]
    op = conditional[1]
    value = int(conditional[2])

    if not register in registers:
        registers[register] = 0

    if op == "<":
        return registers[register] < value
    elif op == ">":
        return registers[register] > value
    elif op == ">=":
        return registers[register] >= value
    elif op == "<=":
        return registers[register] <= value
    elif op == "==":
        return registers[register] == value
    elif op == "!=":
        return registers[register] != value
    else:
        raise ValueError("Op: ", op, " not mapped.")


def eval_ident(ident):

    global largest #for puzzle 2
    ident = ident.split()
    register = ident[0]
    direction = ident[1]
    value = int(ident[2])

    if not register in registers:
        registers[register] = 0
    
    if direction == "dec":
        registers[register] -= value
        if registers[register] > largest:
            largest = registers[register]
    elif direction == "inc":
        registers[register] += value
        if registers[register] > largest:
            largest = registers[register]
    else:
        raise ValueError("Direction: ", direction, " not mapped.")

    
def process_line(line):

    line = line.split(" if ")
    ident_part = line[0]
    conditional_part = line[1]

    if eval_conditional(conditional_part):
        eval_ident(ident_part)


largest = 0
registers = {}

f = open("Inputs/08_input.txt", "r")
for line in f:
    process_line(line)
f.close()

for key, value in sorted(registers.items(), key=lambda register: (register[1],register[0])):
    print(key, ":", value)
print(largest)
