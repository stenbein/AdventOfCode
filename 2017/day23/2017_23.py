#from pdb import set_trace as bp #use pb() #as breakpoint

'''
set X Y sets register X to the value of Y.
sub X Y decreases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in
    register X by the value of Y.
jnz X Y jumps with an offset of the value of Y, but only if the value of X
    is not zero. (An offset of 2 skips the next instruction, an offset of -1
    jumps to the previous instruction, and so on.)
'''

def no_op(): pass

registers = {'counter':0,
             'a':1,
             'b':0,
             'c':0,
             'd':0,
             'e':0,
             'f':0,
             'g':0,
             'h':0} #address:value

instructions = {

    'set': lambda x, y: registers.update({x: registers.get(y,y)}),
    'add': lambda x, y: registers.update({x: registers.get(x,0) + registers.get(y,y)}),
    'sub': lambda x, y: registers.update({x: registers.get(x,0) - registers.get(y,y)}),
    'mul': lambda x, y: registers.update({x: registers.get(x,0) * registers.get(y,y)}),    
    'jnz': lambda x, y: instructions['add']('counter', y-1) if registers.get(x,x) != 0 else no_op(),
    'sta': lambda x, y: print(registers)

}

def eval(pro):

    degbug_stop = 400000000
    
    count_mult = 0
    program_length = len(pro)
    #registers['counter'] = 0
    while registers['counter'] < program_length:
        
        #print("Counter: ", registers.get('counter', 0))
        #print("State: ", registers)
        step = pro[registers.get('counter', 0)]
        #print("Next instruction: ", step)
        if step[0] == "mul": count_mult += 1
        #if len(step) == 2:
        #    instructions[step[0]](step[1])
        #else:
        instructions[step[0]](step[1], step[2])
        #
        registers['counter'] += 1

        if degbug_stop == 0:
            break
        degbug_stop -= 1
        
    print("Done")
    print("mul count: ", count_mult)
    print(registers)

#number 1

program = []
#f = open("Inputs/2017_23_test.txt", "r")
f = open("Inputs/2017_23_test.txt", "r")
for line in f:
    line = line.rstrip().split()
    for index in range(len(line)):
        if line[index].isnumeric() or line[index][0] == '-':
            line[index] = int(line[index])
    program.append(line)

eval(program)




#number 2







