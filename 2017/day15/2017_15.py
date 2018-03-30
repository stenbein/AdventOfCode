

class Judge(object):

    def __init__(self, gen1, gen2):

        self.generators = []
        self.generators.append(gen1)
        self.generators.append(gen2)

        self.mask = (1 << 16)-1

    def begin(self, steps):

        counts = 0
        for step in range(steps):

            if (self.generators[0].next() & self.mask) == (self.generators[1].next() & self.mask):
                counts += 1

        print(counts)


class Generator(object):

    DIV_CONSTANT = 2147483647

    def __init__(self, factor, start_value):
        
        self.factor = factor
        self.prior = 0
        self.current = start_value

    def next(self):

        self.prior = self.current
        self.current = (self.prior * self.factor) % 2147483647
        #print(self.current)
        return self.current


class ExtendedGenerator(object):

    def __init__(self, generator, mults):

        self.generator = generator
        self.mults = mults

    def next(self):

        possible = self.mults + 1
        #input(str(possible) + "...")
        while (possible % self.mults) != 0:

            possible = self.generator.next()
            #print(possible)
            

        #print("exiting with...", possible)
        return possible




genA = Generator(16807, 703) #703
genB = Generator(48271, 516) #516
xGenA = ExtendedGenerator(genA, 4)
xGenB = ExtendedGenerator(genB, 8)

judge = Judge(xGenA, xGenB)
#judge.begin(40000000)
judge.begin(5000000)
