#from pdb import set_trace as bp #use pb() #as breakpoint

'''major error, was checking for grid size % 3 before 2. Worked correctly other way around.
secondary error, was missing a mirror rule'''

start_pattern = '''.#.
..#
###'''


class RuleBook(dict):

    def __init__(self, rule_string):

        for rule in rule_string.split("\n"):
            
            if rule:

                map_in, map_out = rule.split(" => ")
                map_out = self._string_to_grid(map_out)

                variants = self._permutation_of_rules(map_in)
                
                for variant in variants:
                    self[str(variant)] = map_out

    def _string_to_grid(self, rule_string):

        return [[char for char in row] for row in rule_string.split("/")]

    def _mirror_grid(self, grid):

        size = len(grid)
        if size == 2:

            #transpose = [[0, 1], [1, 0]]
            
            grid = [[grid[0][1], grid[0][0]]
                ,[grid[1][1], grid[1][0]]
                    ]

            grid = [['.' if sq == '' else sq for sq in row] for row in grid]
            
            return grid
                     
        elif size == 3:

            #transpose = [[0, 0, 1], [0, 1, 0], [1, 0, 0]]

            grid = [[grid[0][2], grid[0][1], grid[0][0]]
                    ,[grid[1][2], grid[1][1], grid[1][0]]
                    ,[grid[2][2], grid[2][1], grid[2][0]]
                    ]

            grid = [['.' if sq == '' else sq for sq in row] for row in grid]
            
            return grid
        else:
            raise ValueError("Grid size invalid: ", size)

    def _rotate_grid_clockwise(self, grid):

        size = len(grid)
        if size == 2:
                    
            #clockwise = [[0, -1], [1, 0]]

            grid = [[grid[1][0], grid[0][0]],
                     [grid[1][1], grid[0][1]]]
            #for i in range(len(grid)):
            #    for j in range(len(grid[i])):
            #        if grid[i][j] == '':
            #            grid[i][j] = '.'

            grid = [['.' if sq == '' else sq for sq in row] for row in grid]
            
            return grid
                     
        elif size == 3:

            #clockwise = [[0, -1, 0], [1, 0, 0], [0, 0, 1]]

            grid = [[grid[2][0],grid[1][0],grid[0][0]]
                    ,[grid[2][1],grid[1][1],grid[0][1]]
                    ,[grid[2][2],grid[1][2],grid[0][2]]
                    ]
            
            grid = [['.' if sq == '' else sq for sq in row] for row in grid]
            
            return grid
        else:
            raise ValueError("Grid size invalid: ", size)

    def _permutation_of_rules(self, rule):

        output = []

        base_rule = self._string_to_grid(rule)
        output = [base_rule, self._mirror_grid(base_rule)]
        
        base_rule = self._rotate_grid_clockwise(base_rule)
        output.append(base_rule)
        output.append(self._mirror_grid(base_rule))
        
        base_rule = self._rotate_grid_clockwise(base_rule)
        output.append(base_rule)
        output.append(self._mirror_grid(base_rule))
        
        base_rule = self._rotate_grid_clockwise(base_rule)
        output.append(base_rule)
        output.append(self._mirror_grid(base_rule))
        
        return output
        

class PixelGrid(object):

    def __init__(self, rules):

        self.grid = [['.','#','.'],
                     ['.','.','#'],
                     ['#','#','#']]
        self.rules = rules

    def merge(self, grids):

        grid_size = len(grids[0])
        root_size = int(len(grids) ** 0.5)

        output = [[] for i in range(root_size * grid_size)]
        
        for out_row in range(root_size):

            for grid in grids[out_row * root_size:out_row * root_size + root_size]:

                for row in range(grid_size):

                    [output[out_row * grid_size + row].append(item) for item in grid[row]]

        return output

    def get_part(self, grid, index, size):

        output = []
        for i in range(size):

            output.append(grid[index[0] + i][index[1]:index[1]+size])

        return output

    def partition(self, grid):

        size = len(grid)
        if size % 2 == 0:
            split_size = 2
        elif size % 3 == 0:
            split_size = 3
        else:
            raise ValueError("Grid size invalid.")

        '''grids_per_row = int(size / split_size)    
        output = [[[] for row in range(split_size)] for i in range(grids_per_row ** 2)]
        
        out_row = 0
        for row in grid:
            out_col = 0
            for item in row:
                
                out_index = int(out_row / split_size)*grids_per_row + int(out_col / split_size)
                print(item)
                output[out_index][out_row % split_size].append(item)
                
                out_col += 1
            out_row += 1'''
        output = []
        for i in range(0, size, split_size):

            for j in range(0, size, split_size):

                output.append(self.get_part(grid, (i, j), split_size))

        return output

    def apply_rules(self, grids):

        output = []
        for grid in grids:

            output.append(self.rules[str(grid)])

        return output

    def iter(self):

        components = self.partition(self.grid)
        self.grid = self.merge(self.apply_rules(components))

    def get_active_pixels(self):

        return sum([1 for row in self.grid for item in row if item == '#'])
        
        

rules = ''
#number 1


f = open("Inputs/2017_21.txt", "r")
#f = open("Inputs/2017_21_test.txt", "r")
for line in f:

    rules += line 
    
f.close()

rb = RuleBook(rules)
pg = PixelGrid(rb)

for i in range(5):
    pg.iter()

print(pg.get_active_pixels())



#number 2

rb = RuleBook(rules)
pg = PixelGrid(rb)

for i in range(18):
    pg.iter()

print(pg.get_active_pixels())


