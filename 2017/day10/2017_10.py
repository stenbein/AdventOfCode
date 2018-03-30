#from pdb import set_trace as bp #use pb() #as breakpoint


#number 1
class TwistList(object):

    def __init__(self, size):

        self.size = size
        self.list = [x for x in range(size)]
        self.current_index = 0
        self.skip_size = 0

    def twist(self, length):
        
        list_slice = []

        for i in range(length):

            mod_index = self.current_index + i
            list_slice.append(self.list[mod_index % self.size])

        list_slice = list(reversed(list_slice))

        for i in range(length):

            mod_index = self.current_index + i
            self.list[mod_index % self.size] = list_slice[i]


        self.current_index += length + self.skip_size
        self.skip_size += 1


#number 2




def hash_reduce(ascii_array):

    if len(ascii_array) % 16 != 0:
        raise ValueError("Array size not equally divisible by 16.")
    
    output = []
    for i in range(0,int(len(ascii_array) / 16)):

        val = 0
        print(ascii_array[i*16:i*16+16])
        for numb in ascii_array[i*16:i*16+16]:
            val ^= numb
        output.append(val)

    return output

def to_hex_str(dec_array):

    output = ""
    for dec in dec_array:
        val = str(hex(dec)[2:])
        if len(val) == 1:
            val = '0' + val
        output += val
    
    return output







lines = ""
PADDING = [17, 31, 73, 47, 23]

f = open("Inputs/2017_10.txt", "r")
for line in f:
    lines += line.rstrip() #hidden newline in file input
f.close()

tl = TwistList(256)



'''
f = open("Inputs/2017_10_test.txt", "r")
for line in f:
    lines += line
f.close()

tl = TwistList(5)
'''
line = [ord(char) for char in lines]
[line.append(numb) for numb in PADDING]

for itter in range(64):
    #for length in line.split(","):
    for length in line:
        tl.twist(int(length))




print(to_hex_str(hash_reduce(tl.list)))

