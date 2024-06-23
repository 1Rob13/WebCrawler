learnings:

there is no waiting on threads even with channels

176.3µs record this means 176.3 / 5 = 35 per call

this could be ideally 100 when 2 extra rounds are not done with parallel


it seems like all the global variables in the scope fuck up parallel

i need one function to parallelizes that does not depend on any outside input

todo make it work without global variables

todo make it work without global variables, while not using parral 
todo add parallel

functions that are parallelized never contain any return values.
they also never use values that are not in their channel scope, they do not use globals

i need sth like:

for _,url := range urls {

    go CrawlWithoutDuplicate(channels..,waitgroup)

}

todo test this concept and ensure 2 functions running at the same time

fmt.Println() is slow as fuck

fmt.Printf is somehow slipery when used in parallel

S: [https://golang.org/pkg/fmt/ https://golang.org/pkg/os/ https://golang.org/ https://golang.org/pkg/ https://golang.org/ https://golang.org/pkg/]

fetched URLS: [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/]
51.93µs finally

path that worked was that the function encapsulates the parallel part