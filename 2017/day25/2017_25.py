

class Turing(object):

    def __init__(self):

        self.tape = [0]
        self.tape_length = 1
        self.state = 0 #0 will be state A, 1 will be state B
        self.cursor = 0 #index of the cursor position

    def getValue(self):

        if self.cursor >= self.tape_length:

            self.tape += [0 for i in range(self.tape_length)]
            self.tape_length *= 2

        elif self.cursor < 0:

            self.tape = [0 for i in range(self.tape_length)] + self.tape
            self.cursor = self.tape_length - 1 #new index of middle of tape
            self.tape_length *= 2

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
        print(self.tape)
        print("Checksum: ", sum(self.tape))
        print("Current Index: ", self.cursor, " Current State: ", states[self.state])

    def program(self, run_until):

        for step in range(run_until):

            if self.state == 0:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveRight()
                    self.__setStateB()
                else:
                    self.tape[self.cursor] = 0
                    self._moveLeft()
                    self.__setStateB()
            else:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveLeft()
                    self.__setStateA()
                else:
                    self.tape[self.cursor] = 0
                    self._moveRight()
                    self.__setStateA()
                    
            #self.rep()
        self.rep()
            
tur = Turing()
tur.program(12317297)

            
