

class Turing(object):

    def __init__(self):

        self.tape = {0:0}
        self.tape_length = 1
        self.state = 0 #0 will be state A, 1 will be state B
        self.cursor = 0 #index of the cursor position

    def getValue(self):

        if self.cursor not in self.tape:
            self.tape[self.cursor] = 0 # default value on the tape is 0

        return self.tape[self.cursor]
    
    def _moveLeft(self):

        self.cursor -= 1

    def _moveRight(self):

        self.cursor += 1

    def __setStateA(self):

        self.state = 0

    def __setStateB(self):

        self.state = 1

    def rep(self):

        states = ["A", "B"]
        #print(self.tape)
        print("Current Index: ", self.cursor, " Current State: ", states[self.state])

        out = []
        for key in sorted(self.tape.keys()):
            out.append("%s: %s" % (key, self.tape[key]))

        print(out)

    def program(self, run_until):

        for step in range(run_until):

            current_value = self.getValue()

            if self.state == 0:
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveRight()
                else:
                    self.tape[self.cursor] = 0
                    self._moveLeft()
                self.__setStateB()
            else:
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveLeft()
                else:
                    self.tape[self.cursor] = 1
                    self._moveRight()
                self.__setStateA()
                    
            #self.rep()
        self.rep()
            
tur = Turing()
tur.program(20)

            
