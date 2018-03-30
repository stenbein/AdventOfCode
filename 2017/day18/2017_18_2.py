
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

class OneWayChannel(object):
    
    def __init__(self, ids):
        
        self.count = 0
        self.buff = []
        self.ids = ids

    def enqueue(self, msg):
        
        self.count += 1
        self.buff.append(msg)
        
    def dequeue(self):
        
        value = self.buff.pop(0) 
        #print("dequeue call: ", value)
        
        return value
        
    def ready(self):
        #print("ready check")
        
        return bool(self.buff)
        

class RegisterMachine(object):
    
    def __init__(self, ident, chan_in, chan_out, program):
        
        self.pro = program
        
        self.reg = {}
        self.inst = {}
        
        self.inst['set'] = lambda x, y: self.reg.update({x: self.reg.get(y,y)})
        self.inst['snd'] = lambda x: self.cout.enqueue(self.reg.get(x, x))
        self.inst['add'] = lambda x, y: self.reg.update({x: self.reg[x] + self.reg[y]}) if y in self.reg else self.reg.update({x: self.reg[x] + y})
        self.inst['mul'] = lambda x, y: self.reg.update({x: self.reg[x] * self.reg[y]}) if y in self.reg else self.reg.update({x: self.reg[x] * y})
        self.inst['mod'] = lambda x, y: self.reg.update({x: self.reg[x] % self.reg[y]}) if y in self.reg else self.reg.update({x: self.reg[x] % y})
        self.inst['rcv'] = lambda x: self.reg.update({x: self.cin.dequeue()}) if self.cin.ready() else self.listenMode()
        self.inst['jgz'] = lambda x, y: self.incrememt_counter(y) if self.reg.get(x,x) > 0 else no_op()
        
        
        self.reg['a'] = 0
        self.reg['b'] = 0
        self.reg['i'] = 0
        self.reg['f'] = 0
        self.reg['p'] = ident

        self._counter = 0
        self.listen = False

        self.cin = chan_in
        self.cout = chan_out
        
        self.complete = False
        
    def listenMode(self):
        
        self.listen = True
        
    def incrememt_counter(self, val):
        
        if val in self.reg:
            val = self.reg[val]
        self._counter += val - 1

    def eval(self, limit=None, debug=False):

        program_length = len(self.pro)
        self.listen = False
        
        if self._counter < program_length:
            
            step = self.pro[self._counter]
            
            if len(step) == 2:
                self.inst[step[0]](step[1])
            else:
                self.inst[step[0]](step[1], step[2])
            
            if debug:
                print(step)
                print(self.reg)
            
            
            if self.listen:
                self._counter -= 1
            self._counter += 1
            
        else:
            self.complete = True
        
    def step(self):
        
        self.eval(1, False)
    
 
    

program = []
f = open("Inputs/2017_18.txt", "r")
for line in f:
    line = line.rstrip().split()
    for index in range(len(line)):
        if line[index].isnumeric() or line[index][0] == '-':
            line[index] = int(line[index])
    program.append(line)

    
chan0to1 = OneWayChannel(0)
chan1to0 = OneWayChannel(1)

rm0 = RegisterMachine(0, chan1to0, chan0to1, program)
rm1 = RegisterMachine(1, chan0to1, chan1to0, program)

i = 0
while not rm0.complete or not rm1.complete:

    rm0.step()
    rm1.step()
    
    if rm0.listen and rm1.listen:
        print("deadlock")
        break
    
    if i > 10000000:
        break
    
    i += 1
    
print(chan1to0.count)







