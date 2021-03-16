package main

import (
    "fmt"
)

func partition(in []int, low, high int) int {
    pivot := in[high]
    for j := low; j < high; j++ {
     if in[j] < pivot {
         //Swap
        in[j],in[low] = in[low],in[j]
        low++
     }
    }
    in[low],in[high] = in[high],in[low]
    return low
}

func quickSort(in []int, low, high int){
    if low > high {
        return
    }

    part := partition(in, low,high)
    quickSort(in, low, part - 1)
    quickSort(in, part+1, high)
}

func main(){
    in := []int{12,44,13,1,99,22,0,200,8391,461}
    quickSort(in, 0, len(in) - 1)
    for _,v := range in {
        fmt.Printf("Value %d\n",v)
    }
}
