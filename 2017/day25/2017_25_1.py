'''When I first started this project I envisioned only 2 states (like the example)
in reality there are many states. So we need to make adjustments.'''

class Turing(object):

    STATE_A = 0
    STATE_B = 1
    STATE_C = 2
    STATE_D = 3
    STATE_E = 4
    STATE_F = 5

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

    def __setState(self, state):

        self.state = state

    def rep(self):

        states = ["A", "B", "C", "D", "E", "F"]
        print(self.tape)
        print("Checksum: ", sum(self.tape))
        print("Current Index: ", self.cursor, " Current State: ", states[self.state])
    
    def program(self, run_until):

        for step in range(run_until):

            #state A
            if self.state == Turing.STATE_A:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveRight()
                    self.__setState(Turing.STATE_B)
                else:
                    self.tape[self.cursor] = 0
                    self._moveLeft()
                    self.__setState(Turing.STATE_D)
                    
            elif self.state == Turing.STATE_B:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveRight()
                    self.__setState(Turing.STATE_C)
                else:
                    self.tape[self.cursor] = 0
                    self._moveRight()
                    self.__setState(Turing.STATE_F)

            elif self.state == Turing.STATE_C:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveLeft()
                    self.__setState(Turing.STATE_C)
                else:
                    self.tape[self.cursor] = 1
                    self._moveLeft()
                    self.__setState(Turing.STATE_A)

            elif self.state == Turing.STATE_D:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 0
                    self._moveLeft()
                    self.__setState(Turing.STATE_E)
                else:
                    self.tape[self.cursor] = 1
                    self._moveRight()
                    self.__setState(Turing.STATE_A)

            elif self.state == Turing.STATE_E:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 1
                    self._moveLeft()
                    self.__setState(Turing.STATE_A)
                else:
                    self.tape[self.cursor] = 0
                    self._moveRight()
                    self.__setState(Turing.STATE_B)

            elif self.state == Turing.STATE_F:
                current_value = self.getValue()
                if current_value == 0:
                    self.tape[self.cursor] = 0
                    self._moveRight()
                    self.__setState(Turing.STATE_C)
                else:
                    self.tape[self.cursor] = 0
                    self._moveRight()
                    self.__setState(Turing.STATE_E)

            else:

                raise ValueError("Internal state error")
                    
            #self.rep()
        self.rep()

        
tur = Turing()
tur.program(12317297)

            
