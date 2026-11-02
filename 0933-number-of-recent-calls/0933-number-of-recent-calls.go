type RecentCounter struct {
    counter []int
    count int
}


func Constructor() RecentCounter {
    return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
    //push to the queue
    this.counter = append(this.counter, t)
    for len(this.counter) > 0 && t-3000 > this.counter[this.count] {
        this.count++
    }

    return len(this.counter) - this.count
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */