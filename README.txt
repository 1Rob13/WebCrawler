TODO: 

Do tutorial

Change to generic? The function applied to a website should be interchangable 



i fucked up somewhere:


❯ go run cmd/main.go
found: [https://golang.org/pkg/ https://golang.org/cmd/] "The Go Programming Language"
found: [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/] "Packages"
found: [https://golang.org/pkg/ https://golang.org/cmd/] "The Go Programming Language"
found: [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/] "Packages"
not found: https://golang.org/pkg/
not found: https://golang.org/cmd/
not found: https://golang.org/
not found: https://golang.org/cmd/
found: [https://golang.org/ https://golang.org/pkg/] "Package fmt"
found: [https://golang.org/pkg/ https://golang.org/cmd/] "The Go Programming Language"
not found: https://golang.org/
found: [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/] "Packages"
not found: https://golang.org/pkg/
not found: https://golang.org/pkg/fmt/
found: [https://golang.org/ https://golang.org/pkg/] "Package os"
found: [https://golang.org/pkg/ https://golang.org/cmd/] "The Go Programming Language"
not found: https://golang.org/
found: [https://golang.org/ https://golang.org/cmd/ https://golang.org/pkg/fmt/ https://golang.org/pkg/os/] "Packages"
not found: https://golang.org/pkg/
not found: https://golang.org/pkg/os/
not found: https://golang.org/pkg/
not found: https://golang.org/cmd/
not found: https://golang.org/
time since start 516.712µs% 

❯ go run cmd/main.go
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
not found: https://golang.org/cmd/
found: https://golang.org/pkg/fmt/ "Package fmt"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
time since start 133.75µs%                                                                                                                                   ❯ go run cmd/main.go
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
not found: https://golang.org/cmd/
found: https://golang.org/pkg/fmt/ "Package fmt"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
time since start 204.4µs% 


15.73µs
13.2µs
3.93µs
3.61µs
2.7µs
4.27µs
5.76µs
33.03µs
15.36µs
13.99µs

the sum of these should be the max execution time 

goal 100 micro Seconds