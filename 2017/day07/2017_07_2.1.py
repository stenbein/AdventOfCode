
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

    #find the branch that is off balance
    def solve(self):
        
        if self.branches is {}:
            return

        branches = [self.branches[key] for key in self.branches.keys()]

        #check first branch
        if branches[0].branch_weights() != branches[1].branch_weights() and branches[0].branch_weights() != branches[2].branch_weights():
            unbalanced_branch = branches[0]
            difference = unbalanced_branch.branch_weights() - branches[1].branch_weights()
        else:
            #use first branch, which is ok, to compare rest
            for branch in branches:
                if branch.branch_weights() != branches[0].branch_weights():
                    unbalanced_branch = branch
                    difference = unbalanced_branch.branch_weights() - branches[0].branch_weights()
                    break
            else:
                return False, 0 #all branches balanced return nothing
        
        recurse_branch, discard = unbalanced_branch.solve()
        if not recurse_branch:
            return recurse_branch, difference
        else:
            return unbalanced_branch, difference
        
        

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
f = open("Inputs/2017_07.txt","r")
#f = open("Inputs/2017_07_test.txt","r")

for line in f:

    process_line(line, nodes)

f.close()

for tree in [nodes[node] for node in nodes]:
    tree.assemble_branches(nodes)

    

#start with root and recurse, root has no root below it
unbalance_branch = None
difference = 0
for node in nodes.keys():
    
    if nodes[node].root is None:
        unbalance_branch, difference = nodes[node].solve()
        difference = nodes[node].difference

print(unbalance_branch.name, unbalance_branch.weight, unbalance_branch.branch_weights())
print("Difference at this level: ", difference)






        
