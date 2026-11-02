type RecentCounter struct {
    counter []int
}


func Constructor() RecentCounter {
    return RecentCounter{[]int{}}
}


func (this *RecentCounter) Ping(t int) int {
    //push to the queue
    this.counter = append(this.counter, t)
    for len(this.counter) > 0 && t-3000 > this.counter[0] {
        this.counter = this.counter[1:]
    }

    return len(this.counter)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */