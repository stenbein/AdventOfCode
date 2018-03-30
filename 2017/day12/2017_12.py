#from pdb import set_trace as bp #use pb() #as breakpoint

class SimpleGraph(dict):

    def __init__(self):

        self.total = 0

    def add(self, edge):

        if not edge[0] in self:
            self[edge[0]] = []
            self[edge[0]].append(edge[1])
        else:
            self[edge[0]].append(edge[1])
            

    def edges(self, node):

        return self[node]

    def find_connected(self, node):

        all_connections = {node:True}
        to_check = [node]
        
        while to_check:

            node = to_check.pop()

            for edge in self.edges(node):
                if not edge in all_connections:
                    all_connections[edge] = True
                    to_check.append(edge)

        return all_connections
        
    def count_groups(self):
        
        output = {}
        count = 0
        for node in self:
            if node not in output:
                
                [output.update({edge:True}) for edge in self.find_connected(node)]
                count += 1

        return count
                    


#number 1



sg = SimpleGraph()
lines = ""
#f = open("Inputs/2017_12_test.txt", "r")
f = open("Inputs/2017_12.txt", "r")
for line in f:
    parts = line.split('<->')
    node = parts[0].strip()
    edges = parts[1].split(',')
    for edge in edges:
        sg.add((node, edge.strip()))
    
f.close()

print(len(sg.find_connected("0")))



#number 2
print(sg.count_groups())







