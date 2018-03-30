
'''
snd X plays a sound with a frequency equal to the value of X.
set X Y sets register X to the value of Y.
add X Y increases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y
    (that is, it sets X to the result of X modulo Y).
rcv X recovers the frequency of the last sound played, but only when the value of X is not zero.
    (If it is zero, the command does nothing.)
jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero.
    (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
'''
def no_op(): pass

registers = {} #address:value

instructions = {

    'snd': lambda x: registers.update({'played': registers.get(x,0)}),
    'set': lambda x, y: registers.setdefault(x, y),
    'add': lambda x, y: registers.update({x: registers.get(x,0) + y}),
    'mul': lambda x, y: registers.update({x: registers.get(x,0) * registers.get(y,y)}),
    'mod': lambda x, y: registers.update({x: registers.get(x,0) % registers.get(y,y)}),
    'rcv': lambda x: registers.update({'??': registers.get('played',0)}) if x != 0 else no_op(),
    'jgz': lambda x, y: instructions['add']('counter', y) if registers.get(x,0) > 0 else no_op()

}

def eval(pro):

    program_length = len(pro)
    registers['counter'] = 0
    while registers['counter'] < program_length:
        print(registers)
        step = pro[registers.get('counter', 0)]
        print(step)
        if len(step) == 2:
            instructions[step[0]](step[1])
        else:
            instructions[step[0]](step[1], step[2])
        registers['counter'] += 1
    print("Done")

program = []
f = open("Inputs/2017_18_test.txt", "r")
for line in f:
    line = line.rstrip().split()
    for index in range(len(line)):
        if line[index].isnumeric() or line[index][0] == '-':
            line[index] = int(line[index])
    program.append(line)

eval(program)
