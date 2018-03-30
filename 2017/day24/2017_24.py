


ports_input = '''0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10'''


ports_input = '''14/42
2/3
6/44
4/10
23/49
35/39
46/46
5/29
13/20
33/9
24/50
0/30
9/10
41/44
35/50
44/50
5/11
21/24
7/39
46/31
38/38
22/26
8/9
16/4
23/39
26/5
40/40
29/29
5/20
3/32
42/11
16/14
27/49
36/20
18/39
49/41
16/6
24/46
44/48
36/4
6/6
13/6
42/12
29/41
39/39
9/3
30/2
25/20
15/6
15/23
28/40
8/7
26/23
48/10
28/28
2/13
48/14'''


'''
f = open("13_input.txt", "r")
for line in f:
    wall = process_line(line)
    while len(walls) < wall[0]:
        walls.append(0)
    walls.append(wall[1])

f.close()
'''

def get_strength(start, rest):
	
    max_strength = 0
	
    for i in range(len(rest)):
                    
        cur_strength = 0
        if rest[i][0] == start:
            cur_strength = rest[i][0] + get_strength(rest[i][1], rest[:i] + rest[i+1:])
        elif rest[i][1] == start:
            cur_strength = rest[i][1] + get_strength(rest[i][0], rest[:i] + rest[i+1:])
        else:
            pass
        
        if cur_strength > max_strength:
            max_strength = cur_strength
                            
    return start + max_strength


def get_length(start, length, rest):

    max_strength = 0
    max_length = 0
	
    for i in range(len(rest)):
                    
        cur_strength = 0
        cur_length = 0
        if rest[i][0] == start:
            cur_strength, cur_length = get_length(rest[i][1], 1, rest[:i] + rest[i+1:])
            cur_strength += rest[i][0]
        elif rest[i][1] == start:
            cur_strength, cur_length = get_length(rest[i][0], 1, rest[:i] + rest[i+1:])
            cur_strength += rest[i][1]
        else:
            pass
        
        if cur_length > max_length:
            max_strength = cur_strength
            max_length = cur_length
        elif cur_length == max_length:
            if cur_strength > max_strength:
                max_strength = cur_strength
                            
    return start + max_strength, length + max_length



def solve_for_strength(ports):

    max_strength = 0
	
    for i in range(len(ports)):

        if ports[i][0] == 0 or ports[i][1] == 0:
		
            strength = get_strength(ports[i][0] + ports[i][1], ports[:i] + ports[i+1:])
            if strength > max_strength:
                max_strength = strength
    return max_strength

def solve_for_length(ports):

    max_strength = 0
    max_length = 0
	
    for i in range(len(ports)):

        if ports[i][0] == 0 or ports[i][1] == 0:
		
            strength, length = get_length(ports[i][0] + ports[i][1], 1, ports[:i] + ports[i+1:])
            if length >= max_length:
                if strength > max_strength:
                    max_strength = strength
                max_length = length
                print(max_length)
    return max_strength


ports = []
for port in ports_input.split("\n"):
    ports.append([int(size) for size in port.split("/")])

#print(solve_for_strength(ports))
print(solve_for_length(ports))
