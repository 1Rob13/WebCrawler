learnings:

there is no waiting on threads even with channels

176.3µs record this means 176.3 / 5 = 35 per call

this could be ideally 100 when 2 extra rounds are not done with parallel


it seems like all the global variables in the scope fuck up parallel

i need one function to parallelizes that does not depend on any outside input

todo make it work without global variables