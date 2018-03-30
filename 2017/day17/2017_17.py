

class Ratchet(object):

    def __init__(self, value, next_ratchet):
        
        self.value = value
        if next_ratchet is None:
            self.next = self
        else:
            self.next = next_ratchet

    def next_ratchet(self):

        return self.next

class SpinLock(object):

    def __init__(self, skip_size, counts):

        self.skip_size = skip_size
        self.counts = counts

        self.first = Ratchet(0, None)
        self.current_pos = self.first

    def ratchet(self, steps, value):

        current_ratchet = self.current_pos
        
        for count in range(self.skip_size):

            current_ratchet = current_ratchet.next_ratchet()

        new_ratchet = Ratchet(value, current_ratchet.next_ratchet())
        current_ratchet.next = new_ratchet

        self.current_pos = new_ratchet

    def output(self):

        out = []
        current_ratchet = self.first
        for count in range(self.counts+1):
            
            out.append(current_ratchet.value)
            current_ratchet = current_ratchet.next_ratchet()

        return out

    def main(self):

        for count in range(self.counts):

            self.ratchet(self.skip_size, count+1)
            if count % 100000 == 0:
                print(count)


input_ratchet = 386
#input_ratchet = 3

sl = SpinLock(input_ratchet, 50000000)
sl.main()
#print(sl.output())

i = 0
outs = sl.output()
while outs[i] != 0:
    i += 1

print(outs[i])
print(outs[i+1])
