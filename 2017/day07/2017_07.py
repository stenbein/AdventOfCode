
class Tree(object):

    def __init__(self, name, weight, branches):
        
        self.name = name
        self.weight = weight
        self.root = None
        self.branches = {branch.replace(",",""):None for branch in branches}


    def add_branch(self, branch):

        if branch.name in self.branches:
            raise KeyError("Branch name already assigned")
        else:
            branch.add_root(self)
            self.branches[branch.name] = branch


    def add_root(self, root):

        if not self.root is None:
            raise ValueError("Branch already has root")
        else:
            self.root = root

    def assemble_branches(self, nodes):

        for branch in self.branches:
            
            if branch in nodes:
                self.branches[branch] = nodes[branch]
                del(nodes[branch])
            else:
                raise KeyError("Shouldn't happen.")

    def rep_weights(self):

        print(self.name, " weight is: ", self.weight)
        for branch in self.branches:
            print("Branch: ", branch, " weight: ", self.branches[branch].weight)
            print(self.branches[branch].branch_weights())

    def branch_weights(self):

        return self.weight + sum([self.branches[branch].branch_weights() for branch in self.branches])

    def print_branch_weights(self):

        print("Branch: ", self.name, " weight: ", self.weight)

        for branch in self.branches:
            if self.branches[branch] is None:
                break
            self.branches[branch].print_branch_weights()

    def solve(self):
        
        if self.branches is {}:
            return
        repeat = False
        #check weights of children
        weights = [(name, branch.branch_weights()) for name, branch in self.branches.items()]
        if weights[0][1] != weights[1][1] and weights[0][1] != weights[2][1]:
            off_balance = weights[0]
            repeat = True
        else:
            for weight in weights:
                if weight[1] != weights[0][1]:
                    off_balance = weight
                    repeat = True
                    break
        if repeat:
            self.branches[off_balance[0]].solve()
            self.branches[off_balance[0]].rep_weights()
        else:
            return
        

def process_line(line, nodes):

    if not "->" in line:
        
        parts = line.split("->")
        name = parts[0].split()[0]
        weight = int(parts[0].split()[1][1:-1])
        nodes[name] = Tree(name, weight, [])

    else:
        
        parts = line.split("->")
        
        name = parts[0].split()[0]
        weight = int(parts[0].split()[1][1:-1])
        branches = parts[1].split()
        nodes[name] = Tree(name, weight, branches)        


nodes = {}
f = open("Inputs/2017_07_test.txt","r")
#f = open("Inputs/2017_07_test.txt","r")

for line in f:

    process_line(line, nodes)

f.close()

for tree in [nodes[node] for node in nodes]:
    tree.assemble_branches(nodes)

    


for node in nodes.keys():
    #nodes[node].rep_weights()
    #nodes[node].print_branch_weights()
    #nodes[node].branch_weights()
    if nodes[node].root is None:
        nodes[node].solve()









        
