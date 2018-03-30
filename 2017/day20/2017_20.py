

'''
It transmits to you a buffer (your puzzle input) listing each particle in order
(starting with particle 0, then particle 1, particle 2, and so on). For each
particle, it provides the X, Y, and Z coordinates for the particle's position
(p), velocity (v), and acceleration (a), each in the format <X,Y,Z>.

Each tick, all particles are updated simultaneously. A particle's properties
are updated in the following order:

Increase the X velocity by the X acceleration.
Increase the Y velocity by the Y acceleration.
Increase the Z velocity by the Z acceleration.
Increase the X position by the X velocity.
Increase the Y position by the Y velocity.
Increase the Z position by the Z velocity.
Because of seemingly tenuous rationale involving z-buffering, the GPU would
like to know which particle will stay closest to position <0,0,0> in the 
long term. Measure this using the Manhattan distance, which in this situation is simply the sum of the absolute values of a particle's X, Y, and Z position.

For example, suppose you are only given two particles, both of which stay
entirely on the X-axis (for simplicity). Drawing the current states of
particles 0 and 1 (in that order) with an adjacent a number line and diagram
of current X positions (marked in parenthesis), the following would take place:

p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>                         (0)(1)

p=< 4,0,0>, v=< 1,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=< 2,0,0>, v=<-2,0,0>, a=<-2,0,0>                      (1)   (0)

p=< 4,0,0>, v=< 0,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=<-2,0,0>, v=<-4,0,0>, a=<-2,0,0>          (1)               (0)

p=< 3,0,0>, v=<-1,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=<-8,0,0>, v=<-6,0,0>, a=<-2,0,0>                         (0)

At this point, particle 1 will never be closer to <0,0,0> than particle 0,
and so, in the long run, particle 0 will stay closest.

Which particle will stay closest to position <0,0,0> in the long term?
'''

class Particle(object):

    def __init__(self, count, pos, vel, acc):

        self.id = count
        self.pos = pos
        self.vel = vel
        self.acc = acc

    def step(self):

        self.mod_vel()
        self.mod_pos()

    def mod_vel(self):

        self.vel = (self.vel[0] + self.acc[0], self.vel[1] + self.acc[1], self.vel[2] + self.acc[2])

    def mod_pos(self):

        self.pos = (self.pos[0] + self.vel[0], self.pos[1] + self.vel[1], self.pos[2] + self.vel[2])

    def distance(self):

        return sum([abs(dis) for dis in self.pos])

    def velocity(self):

        return sum([abs(vel) for vel in self.vel])

    def __str__(self):

        #print(self.pos)
        return "id: {} pos: {} {} {} vel: {} {} {} acc {} {} {}".format(self.id,
                                                                        self.pos[0], self.pos[1], self.pos[2],
                                                                        self.vel[0], self.vel[1], self.vel[2],
                                                                        self.acc[0], self.acc[1], self.acc[2])


def parse_particle(particle_string, count):

    #p=< 4,0,0>, v=< 0,0,0>, a=<-1,0,0>
    particle = particle_string.rstrip()[:-1].replace(" ","").split(">,")

    pos = tuple([int(x) for x in particle[0][3:].split(",")])
    vel = tuple([int(x) for x in particle[1][3:].split(",")])
    acc = tuple([int(x) for x in particle[2][3:].split(",")])

    particle = Particle(count, pos, vel, acc)
    
    return particle



particles = []


input_str = ""
f = open("Inputs/2017_20.txt", "r")
for line in f:
    input_str += line

particles = []
for particle in input_str.rstrip().split("\n"):

    particles.append(parse_particle(particle, len(particles)))


#question 1
'''
particles_static = [s_p for s_p in particles if s_p.vel == (0,0,0) and s_p.acc == (0,0,0)]
particles_steady = [s_p for s_p in particles if s_p.acc == (0,0,0)]

if particles_static:
    particles_static.sort(key=lambda part: part.distance())
    print("Statics: ", particles_static[0])

else:
    particles_steady.sort(key=lambda part: part.distance())
    print("No statics - Closest Steady State: ", particles_steady[0])
'''

#question 2

for count in range(100):
    for particle in particles:
        particle.step()
    positions = [particle.pos for particle in particles]
    counts = {position:positions.count(position) for position in positions}

    i = 0
    while i < len(particles):
        if counts[particles[i].pos] > 1:
            del(particles[i])
            print("Bam!~")
        else:
            i += 1

    print(len(particles))