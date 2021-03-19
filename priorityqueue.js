class PriorityQueue {
    constructor() {
        this.end = 0;
        this.collection = [];
    }

    length() {
        return this.end;
    }

    // enqueue() takes an array of [value, priority]
    // adds to queue based on priority
    enqueue(element) {
        // If collection is empty, OR, element has lower priority / higher value
        // than last element in collection, add new element to end of queue.
        if (this.end == 0 || element[1] > this.collection[this.end - 1][1]) {
            this.collection[this.end] = element;
            this.end++;
        } else {
            // Find position nearest to front of queue based on priority level
            for (let i = 0; i < this.end; i++) {
                if (element[1] < this.collection[i][1]) {
                   this.collection.splice(i, 0, element);
                    this.end++;
                    return
                }
            }
        }
    }

    // dequeue() Removes and returns element at front of queue
    dequeue() {
        if (this.end == 0) { return null }
        
        const element = this.collection.shift();
        this.end--;
        return element;
    }

    // clear() Resets collection to empty []
    clear() {
        this.collection = [];
        this.end = 0;
    }
}


const kpqueue = new PriorityQueue();
console.log(`Priority Queue length: ${kpqueue.length()}`);
kpqueue.enqueue(['Pineapple', 3]);
kpqueue.enqueue(['Kev', 1]);
kpqueue.enqueue(['Mango', 2]);
console.log(`Priority Queue length: ${kpqueue.length()}`);
console.log(kpqueue.collection);
console.log(kpqueue.dequeue());
console.log(kpqueue.dequeue());
console.log(kpqueue.dequeue());
kpqueue.enqueue(['Pineapple', 3]);
kpqueue.enqueue(['Kev', 1]);
kpqueue.enqueue(['Mango', 2]);
kpqueue.clear();
console.log(`Cleared Priority Queue length: ${kpqueue.length()}`)
