class Queue {
    constructor() {
        this.end = 0;
        this.collection = [];
    }

    length() { 
        return this.end 
    }

    enqueue(element) {
        this.collection[this.end] = element;
        this.end++;
    }

    dequeue() {
        if (this.end == 0) { return null }
        const element = this.collection[0];
        this.collection = this.collection.slice(1, this.end);
        return element
    }

    clear() {
        this.collection = [];
    }
}

const kqueue = new Queue();
console.log(`kqueue length: ${kqueue.length()}`);
kqueue.enqueue('Kev');
kqueue.enqueue('Pineapple');
console.log(`kqueue length: ${kqueue.length()}`);
console.log(kqueue.dequeue());
console.log(kqueue.dequeue());
console.log(`kqueue length: ${kqueue.length()}`);
