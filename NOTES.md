Reading Notes
=============

The following is a collection of summaries, excerpts, and notes I took from the book.

**NOTE: This is NOT a complete summary of the book's content. Please read the book first to get a thorough understanding of all the topics covered in the book.**

Chapter 1. What Are Algorithms and Why Should You Care
------------------------------------------------------

- What is an algorithm?
  - A set of steps defined precisely enough for computers to execute, in order to complete a task.
- What we expect from proper algorithms
  - Correctness
    - We focus on the algorithms with knowable solutions.
      - But we sometimes have to accept algorithms that might produce incorrect answers, given that we have some control over how often such cases could occur.
      - Approximation Algorithms: "Almost Optimal" - its optimality within some known factor of the optimal solution's
  - Efficient Use of Computational Resources: Time, Memory Footprint
    - Order of growth of the running time

Chapter 2. How to Describe and Evaluate Computer Algorithms
-----------------------------------------------------------

- How to describe computer algorithms: Pseudocode? -> He will just use English..
  - Procedures, Parameters, Array, Loop, Iteration, Loop Variable
  - A sentinel: Empty box to prevent checking loop variables
- Linear Search: [See the code.](https://link.iamblogger.net/hiu3p)
- Asymptotic Notation
  - **Big-Theta**: Running time upper- and lower-bounded by the given function.
  - **Big-O**: Upper-bounded
  - **Big-Omega**: Lower-bounded
- Use a loop invariant to prove correctness
  - **Loop invariant**: A statement or property that holds at the start of each loop iteration
- **Recursion**
  - Base cases
  - Factorial

Chapter 3. Algorithms for Sorting and Searching
-----------------------------------------------

- If we don't have any information about how the elements of the array are ordered, we cannot do better than linear search.
- If the array is sorted into nondecreasing order: binary search in only `O(lg n)` time
- What do you mean when one element is less than another?
  - Numbers
  - Lexicographic ordering

- Key: The data that we do the sorting over to get an order
  - Satellite data: The information associated with the key but not the one we base our sorting

- Then how do we get the array to be sorted in the first place?
  - Let's use sorting algorithms: `Theta(n^2)` or `Theta(n * lg n)` in worst case.
    - Selection Sort
    - Insertion Sort
    - Merge Sort
    - Quick Sort

- Binary Search: [See the code.](https://link.iamblogger.net/mjw2d)
  - The idea: Given an array/list of elements, check the index right in the middle to see if the key is bigger or smaller than the one you are looking for.
    - If it's larger, then you can safely ignore the second half and only consider the first.
    - If smaller, you can discard the first half instead.
    - Continue checking the middle index to discard one of the two halves, until you get exactly the key you are looking for.
  - Time complexity: `O(lg n)`, since we are halving the array at each loop iteration, so there will be approximately `lg n` iterations.
  - Use loop variant to prove its correctness
    - Loop invariant: At the start of each iteration of the loop, if `x` is anywhere in the array `A`, then it is somewhere in the subarray `A[p:r]`.
      - By the contrapositive of the loop variant above, we can see that if `x` is not in `A[p:r]`, then it's not in `A`.
  - Recursive version: [See the code.](https://link.iamblogger.net/1orvi)

- Selection Sort: [See the code.](https://link.iamblogger.net/8zrxt)
  - The idea: Find the smallest among `A[i]` to `A[n]`, and put it in `A[i]`. Then increment `i` and repeat.
  - Time complexity: `Theta(n^2)`, since we run the outer loop `n` times, and execute the inner loop `n-i` times each.
  - However, actual moving of array elements occur only `n-1` times. Therefore, it might be reasonable to use selection sort when I/O costs are high.

- Insertion Sort: [See the code.](https://link.iamblogger.net/jsfq7)
  - The idea: Starting from the second element, sort elements among the first `(i+1)` items by swapping items in reverse order until reaching `A[0]`, or an element with sort key smaller than `A[i]`, since that would indicate that items located at that index and earlier are already sorted AND they are all smaller than `A[i]`, so they don't have to be shifted. Note that those first `(i+1)` items would include items that are already sorted in the previous iterations.
  - Time complexity
    - Best Case: `Theta(n)` time when the elements are already sorted. each iteration will take constant time for `n` iterations.
    - Worst Case: `Theta(n^2)` when items are in reversed sorted order. For each iteration of the outer loop, the inner loop iterates `i-1` times. Those `i-1`s will form an arithmetic series, and its sum will be `Theta(n^2)`.
      - In the worst case, therefore, selection sort and insertion sort have running times that are asymptotically the same.
    - Average Case: Still `Theta(n^2)`. If elements were to be initially given in a truly random order, then each elements would be expected to be larger than roughly half the elements preceding it and smaller than the another half. Then the inner loop would run `(i-1)/2` iterations per each outer loop run.

    - Insertion Sort is an excellent choice when the array starts out as "almost sorted":
      - If we can be certain about the expected number of movements each element would need, then the total running time would be `k*n` `~ Theta(n)`.
    - Selection Sort vs. Insertion Sort: While they are both `O(n^2)`, selection sort move elements `Theta(n)` times no matter what, while insertion sort could move elements up to `Theta(n^2)` times.

- Merge Sort: [See the code.](https://link.iamblogger.net/-ppw4)
  - The idea: Divide the array into two. Divide each two pieces again into two. Repeat the process until there's at most one element so it can't be divided into two any more. Then we start 'merging' all the pieces back into the full array, by comparing the each element from the last two pieces and put them in sorted order into new array, and so on until we get the fully sorted array.

  - Merge Sort employs **'divide-and-conquer'** paradigm.
    - We break the problem into subproblems that are similar to the original, solve the subproblems recursively, and then combine the solutions to the subproblems to solve the original problem.
      - DIVIDE by finding the index `q` roughly in the middle of `p` and `r`: add `p` and `r`, divide by `2`, and take the floor.
      - CONQUER by recursively sorting the elements in each of the two subproblems created by the divide step: recursively sort the elements that are in indexes `p` through `q`, and recursively sort the elements that are in indexes `q + 1` through `r`.
      - COMBINE by merging the sorted elements that are in indexes `p` through `q` and indexes `q+1` through `r`, so that all the elements in slots `p` through `r` are sorted.
    - The base case occurs when fewer than two elements need to be sorted (that is, when `p >= r`), since a set of elements with no elements or one element is already trivially sorted.

  - Time complexity: `Theta(n*lg n)` in ALL cases.
    - Since merging the subarrays divided at the same stage involves copying `n` elements, it takes `Theta(n)` time.
    - Plus, since we've been constantly halving the array into two, we get to do the merging over `n` elements `lg n` times, hence leading to `Theta(n*lg n)` running time.
  
  - Disadvantages:
    - The constant factor is higher than selection sort or insertion sort.
    - We have to make complete copies of the entire input array.
      - If space is at a premium, you might not want to use merge sort.

- Quick Sort: [See the code.](https://link.iamblogger.net/1x51x)
  - The idea: also employs divide-and-conquer. Pick the last point in the array and call it a 'pivot.' Organize them within the array in groups where one group `L` has all the elements smaller than the pivot, another called `R` with all bigger than the pivot, and the pivot `P` itself. Swap the pivot with the leftmost element of `R` to make the array to be in the order of `L`, `P`, and `R`. Then apply the steps again to `L` and `R` themselves.

  - Works in place
  - Better constant factors than merge sort's
  - Divide-and-conquer
    - DIVIDE: First choose any one element in indexes `p` through `r`. Call this index the *pivot*. Rearrange elements in the array so that all other elements with keys that come before the pivot's or have identical keys are to the left of the pivot, and all others with keys that come after the pivot's are to the right of the pivot.
    - CONQUER: Recursively sort the elements to the left of the pivot and to the right of the pivot. That is, if the divide step moves the pivot to slot `q`, then recursively sort the books in slots `p` through `q-1` and recursively sort the books in slots `q+1` through `r`.
    - COMBINE: Once the conquer step recursively sorts, everything is already combined. All the elements to the left of the pivot (in slots `p` through `q-1`) come before the pivot or have the same key as the pivot and ARE SORTED, and all the elements to the right of the pivot (in slots `q+1` through `r`) come after the pivot and are sorted. The elements in the indexes `p` through `r` can't help but be sorted!

    - Like merge sort, the base case occurs when the subarray to be sorted has fewer than two elements.

  - Time complexity
    - Each `partition()` would take `Theta(n)` time, since we have to iterate over all the elements in the given (sub)array.

    - Worst Case: `Theta(n^2)` - When everything is already sorted and every elements except the pivot gets into either `L` or `R`. `partition()` won't be able to divide elements into `L` and `R` evenly since they would all fall into just one of them. It would take `n` steps to partition everything, hence the total running time would be `Theta(n^2)`.
      - `T(n) = T(n-1) + Theta(n)`

    - Best Case: `Theta(n * lg n)` - Same as merge sort, when the pivot falls roughly in the middle every time.
      - `T(n) = T(n/2) + Theta(n)`

    - With some technical analysis, we can see that if the elements of the input array come in a *random* order, then on average we get splits that are close enough to even that quick sort takes `Theta(n * lg n)` time.

    - To prevent picking bad pivot points, we can choose the pivot randomly among the elements and swap with `A[r]` before proceeding to partition().
      - Even better, we can have alternative pivot-choosing strategies where
        - Choose three elements at random and swap the median of the three with `A[r]`.

    - Try avoid attempting 'swap' on elements to move to the exactly same positions.

- In empirical testing, randomized quicksort beats other algorithms examined in this chapter. But we can do even better by
  - Employing hybrid methods: Perform insertion sort, which is known to be `Theta(n)` in best cases, when the subarrays became small enough during the quicksort process.
  
  - But is it possible to beat `Theta(n * lg n)` time for sorting? It depends...
    - If we don't know about the elements beforehand, then **no**.

Chapter 4. A Lower Bound for Sorting and How to Beat It
-------------------------------------------------------

- Really simple sort: [See the code.](https://link.iamblogger.net/dalbf)
  - The idea: If we know that there's only `1` or `2` in the array, then we can just count the # of `1`s in the array to be `k`, set the first `k` indexes of an array to be `1`, and set the rest to be `2`. This algorithm involved no direct comparison between elements, and achieves `Theta(n)` time.

- The lower bound on "comparison sorting"
  - Comparison sort: Any sorting algorithm that outputs the sorted order by comparing the given elements. Note that really simple sort described above is NOT a comparison sort algorithm.
  - In the worst case, any comparison sorting algorithm for `n` elements requires `Omega(n * lg n)` comparisons between pairs of elements.
    - Note that it's saying something only about the WORST case. You might be able to do less comparisons if you assume other best cases.
      - We call this lower bound an **existential lower bound** because it says that there exists an input that requires `Omega(n * lg n)` comparisons.
      - Another type of lower bound is a **universal lower bound**, which applies to all inputs.
        - For sorting, the only universal lower bound we have is `Omega(n)`, since we at least have to look at each element once.
        - This lower bound does not depend on the particular algorithm.

- Counting Sort: [See the code.](https://link.iamblogger.net/x64ca)
  - The idea: Generalization of really simple sort - If we know that `k` elements have sort keys equal to `x` and that `l` elements have sort keys less than `x`, then we know that the elements with sort keys equal to `x` should occupy positions `l+1` through `l+k` in the sorted array.

  - `Theta(m + n)` time
    - `m` loop iterations for intializing the array containing the # of each `m` elements, and the `less` array.
    - `n` loop iterations for initial counting of the elements, and setting the output array by each values

  - Counting sort was possible because all the sort keys were integers; It wouldn't have been possible if they were like real numbers or strings.

  - Counting sort is a **stable sort**: In a stable sort, elements with the same sort key appear in the output array in the same order as they do in the input array.
    - In other words, a stable sort breaks ties between two elements equal sort keys by placing first in the output array whichever element appears first in the input array.

- Radix Sort
  - Given strings of characters of some fixed length, sequentially sort based on *each characters*
    - We can't do counting sort directly on those strings, because there will be infinite combinations, or still too many even if we limit the possible # of characters.

  - To generalize, in the radix sort algorithm, we assume that we can think of each sort key as a `d`-digit number, where each digit is in the range `0` to `m-1`.

  - We run a stable sort on each digit, going from *RIGHT* to *LEFT*. If we use counting sort as the stable sort, then the time to sort on one digit is `Theta(m + n)`, and the time to sort all `d` digits is `Theta(d*(m+n))`.
    - If `m` is a constant, then the time for radix sort is `Theta(d * n)`. If `d` is also a constant, then the time for radix sort is only `Theta(n)`.
    - It is critical that we use *stable sort* algorithm on individual digits, so that we retain sorted orders from previous digits.
    - Right to left, otherwise the element that is supposed to appear later could come earlier because it has lower sort key in later digits.

Chapter 5. Directed Acyclic Graphs
----------------------------------

- Directed Graphs
  - Vertices
  - Directed edges `(u, v)`
    - `v` is "adjacent" to `u`
    - "leaves" `u` and "enters" `v`

- Directed Acyclic Graphs (DAG)
  - There is no way to get from a vertex back to itself by following a sequence of one or more edges.
  - DAGs are great for modeling dependencies.

- Topological Sort: [See the code.](https://link.iamblogger.net/oy0zh)
  - A topological sort of a DAG produces a *linear order* such that if `(u, v)` is an edge in the DAG, then `u` appears before `v` in the linear order.
  
  - The linear order produced by a topological sort is not necessarily unique.

  - How to choose a new vertex to be added to the linear order:
    - Any vertex with no entering edges.
      - In-degree: The # of edges entering a vertex
      - Any vertex with `in-degree=0`
        - Every DAG must have at least one vertex with `in-degree=0` and at least one vertex with `out-degree=0`, for otherwise there would be a cycle.

  - Any vertex `v` that is adjacent to `u` must appear somewhere after `u` in the linear order.
    - Therefore, we can safely remove `u` and the edges leaving from `u` after adding it to the linear order.
    - What's left immediately after is *another DAG*.
      - we simply repeat the process until there's no remaining vertex.

    - How to represent a directed graph
      - `n * n` adjacency matrix: each row and column correspodns to one vertex. the entry is `1` if the edge `(u, v)` is present or ``0` if there's no such edge.
      - Adjacency-list representation: Lists of all the vertices adjacent to each vertex.
        - How should we represent such lists
          - Array: If we know/can fix the # of adjacent vertices
          - **Linked list**: If we need to insert into the middle of the list
            - Singly linked list: Only successor links
            - Doubly linked list: Also with predecessor links

    - Running Time:
      - `Theta(n + m)` time overall - assuming that adding and removing items to adjacency list takes constant time.

- Critical path in a PERT chart
  - PERT chart: "Program Evaluation and Review Technique"
    - The time to complete the entire job, even with as many tasks performed simultaneously possible, is given by the "critical path" in the PERT chart.

    - **Path**: A sequence of vertices and edges that allow you to get from one vertex to another (or back to itself)

    - A **critical path** in a PERT chart is a path for which the sum of the task times is *MAXIMUM* over *ALL PATHS*.
      - The sum of the task times along a critical path gives the minimum possible time for the entire job, no matter how many tasks are performed simultaneously.

    - Rather than checking paths between all pairs of vertices in which one has `in-degree=0` and one has `out-degree=0`, we can just add two *dummy* vertices, *start* and *finish*.

    - Now, solve the problem by *negating* task times and finding the *shortest path* from start to finish. (Negate so that we can in effect find a *maximum* path.)

    - Weighted directed graph
      - A shorted path from vertex `u` to vertex `v` is path whose sum of *edge* weights is minimum over all paths from `u` to `v`.
      - Shortest paths are not necessarily unique, as a directed graph from `u` to `v` could contain multiple paths whose weights achieve the minimum.

- *Single-source* shortest path in a DAG: [See the code.](https://link.iamblogger.net/q60eo)
  - The idea: Get a topologically sorted linear order first, and then relax every vertex and their edges according to that sorted order.
  - `Relax()`: Determine whether passing through `u` to reach `v` gives us a shorter route than the previously known shortest path to `v`.
  - Notice that we will end up relaxing **every single edges** in the given DAG.
    - The edges of each shortest path will be interspersed, in order, as we go through all the edges and relax each one.
    - Why does this work?
      - Suppose that a shortest path from `s` to `v` visits the vertices `s`,`v_1`, `v_2`, `v_3`, ... , `v_k`, `v`, in that order.
      - After edge `(s, v_1)` has been relaxed, `shortest[v_1]` must have the shortest-path weight for `v_1`, and `pred[v_1]` must be `s`. After `(v_1, v_2)` has been relaxed, `shortest[v_2]` and `pred[v_2]` must be correct. And so on, up through relaxing `(v_k, v)`, after which `shortest[v]` and `pred[v]` have their correct values.
  - `Theta(n + m)` running time.

Chapter 6. Shortest Paths
-------------------------

- Chapter 5 was about finding shortest paths in a directed *acyclic* graph - it had to be acyclic so that we can do topological sort.
  - But cycles are everywhere in real-life situation graphs
    - Road network: Must have cycles or you won't be able to return to where you came from

  - Single-pair shortest path problem: GPS finding the fastest route from one place to another
    - It would use an algorithm that finds *all* shortest paths from the source, but pays attention only to the path that leads to the relevant destination.
    - The thing is, you won't have to worry much about cycles as long as the weight of edges are *nonnegative*. Shortest paths are well defined for such graphs.

  - Negative-weight cycles: Are they relevant in the real world?
    - Yes. Arbitrage opportunity in currency trading.

- Dijkstra Algorithm: [See the code.](https://link.iamblogger.net/f8t23)
  - The idea: So this is similar to the shortest path algorithm for DAGs introduced in Chapter 5. However, instead of using topological sort to determine certain order of visits to all the vertices, we make use of priority queues to visit a vertex that have lowest `shortest[v]` each time.

  - Why does visiting the vertex with the lowest `shortest` value makes sense?
    - Loop invariant: At the start of each iteration of the loop, `shortest[v]=sp(s, v)` for each vertex `v` not in the queue. That is, for each vertex `v` not in the queue, the value of `shortest[v]` is *already* the weight of a shortest path from `s` to `v`.
      - So those `shortest[v]` values can never decrease after they leave the queue.
        - Consider the vertex `u` in the queue with the lowest `shortest` value.
        - Because the only edges remaining to be relaxed are edges leaving vertices in the queue, and every vertex in the queue has a `shortest[v]` value at least as large as `shortest[u]`.
        - Since all edge weights are nonnegative, we must have `shortest[u] <= shortest[v] + weight(v, u)` for every vertex `v` in the queue, and so no future relaxation step will decrease `shortest[u]`.

  - Priority Queue: It's an ADT (Abstract Data Type).
    - `Insert(Q, v)`
    - `Extract-Min(Q)`: Find the vertex with the smallest `shortest[v]`
    - `Decrease-Key(Q, v)`

    - Simple array implementation: If we use a simple array (for priority queue), it will make Dijkstra do `O(n^2)` time since it will take `O(n)` time to find the smallest `shortest[v]` for each of `n` calls. Calls to `Relax()` will take `O(m)` altogether, but `m` is magnitude smaller than `n^2` so it can be ignored.

    - Binary Heap Implementation: [See the code for binary heap priority queue here.](https://link.iamblogger.net/i2-r4)
      - Binary Tree: vertices are called *nodes*, and the edges are undirected. each node has 2 nodes directly below them at maximum, which are called *children*. Nodes with no children are called *leaves*.

      - Binary Heap: 3 properties
        1. The tree is COMPLETELY FILLED on all levels, except possibly the lowest, which is filled from the left up to a point.
        2. Each node contains a key, shown inside each node in the figure.
        3. The keys obey the **heap property**: The key of each node is less than or equal to the keys of its children.

        - Corollary:
          - We can store a binary heap in an array.
          - The node with the minimum key must always be at position `1`.
          - The children of the node at position `i` are at positions `2i` and `2i + 1`.
          - The parent will be at `floor(i/2)`.
          - The height of the tree will be `floor(lg n)`.
          - We can traverse a path from the root down to a leaf in `O(lg n)` time.

      - Because binary heaps have height `floor(lg n)`, **we can perform the three priority queue operations in `O(lg n)` time each.**

      - With a binary heap priority queue,
        - Inserting vertices would take `O(n * lg n)` time (`Theta(n)` in reality, since all of vertices will initially have infinity weights except for the source vertex.)
        - `ExtractMin(Q)` would take `O(n * lg n)`
        - `Decrease-Key(Q, v)` to take `O(m * lg n)`
        - Total time complexity:
          - `O((n+m) * lg n)` time
          - However, we could say `O(m * lg n)` instead as it is usually reasonable to assume that `m > n`.

      - When the graph is sparse - the number `m` of *edges* is much less than `n^2` (`n^2` is required by `ExtractMin` Operations in a simple array implementation) - implementing the priority queue with a binary heap is more efficient.
        - Example: graphs that model road networks `m` `~` `4n`.

      - On the other hand, when the graph is dense, `m` is close to `n^2`, so that the graph contains many edges, `O(m * lg n)` time that Dijkstra's algorithm spends in `Decrease-Key` calls can make it slower than using a simple array.

      - By the way, we can do heap sort using a binary heap. ([See the code for heap sort from here.](https://link.iamblogger.net/i9f-t))
        - `O(n * lg n)` time when inserting elements to a heap.
        - `O(n)` time if we build a heap directly within the array.

    - Fibonacci heap ('F-heap') implementation
      - `n` `Insert`/`ExtractMin` calls: `O(n * lg n)`
      - `m` `DecreaseKey` calls: `Theta(m)`
      - In total, `O(n * lg n + m)` time.
      - However, people don't use F-heap too much because
        - An individual operation might take much longer than average.
        - F-heaps are complicated, so the constant factors hidden are not as good as for binary heaps.

- The Bellman-Ford Algorithm: [See the code.](https://link.iamblogger.net/tqqyj)
  - The idea: With the initial shortest and pred values, it just relaxes all `m` edges `n-1` times.
    - Every *acyclic path* must contain *at most* `n-1` edges. (since there are `n` vertices and more than `n-1` edges implies that there's a cycle.)
    - Each time the loop runs through all the edges, each edge in the shortest path would get relaxed one by one.
    - After the `(n-1)`st time, all edges on the shortest path have been relaxed, in order, and therefore `shortest[v]` and `pred[v]` are correct.

  - Unlike Dijkstra, the Bellman-Ford algorithm works with negative edge weights, and we can use its output to detect a negative weight cycle.

  - Running time: `Theta(n * m)` since we go over exactly `m` edges `n-1` times.

  - Find a Negative Weight Cycle Procedure: [See the code.](https://link.iamblogger.net/uwj94)
    - Now, every edges should have been relaxed with `n-1` iterations. If we run the relaxation once more and find that there is a node that `shortest[v]` would decrease once more, it means there's at least one edge `(u, v)` on the cycle with a negative weight.

    - Running Time: To find whether a negative-weight cycle exists, relax each edge once more until either relaxing changes a shortest value or all edges have been relaxed, taking `O(m)` time. If there is a negative-weight cycle, it can contain at most `n` edges, and so the time to trace it out is `O(n)`.

    - Arbitrage Opportunity
      - Do the logarithm of exchange rates and negate them
      - See if there's any negative-weight cycles
      - The total # of edges `m` is `n + n * (n-1) = n^2`, so the Bellman-Ford takes `O(n^3)` time, along with another `O(n^2)` to determine whether there's a negative-weight cycle, and another `O(n)` to trace it out if it exists.

- The Floyd-Warshall Algorithm: [See the code.](https://link.iamblogger.net/afky5)
  - The All-Pairs Shortest-Paths Problem
    - Find a shortest path from every vertex to every vertex
      - If the graph is sparse, then we could just run Dijkstra (without negative-weight edges) or Bellman-Ford for each of the `n` vertices.
      - However, if the graph is dense, the running time would become unbearable.

    - We can solve the all-pairs problem in `Theta(n^3)` time with the Floyd-Warshall.
      - Regardless of the graph's sparsity
      - Works with negative-weight edges but not negative-weight cycles
      - The constant factor hidden in the Theta-notation is small.
      - Dynamic Programming

  - The idea: The algorithm relies on an obvious property of shortest paths:
    - The portion of the shortest route would be the shortest route for that particular source to destination as well.

  - The Floyd-Warshall keeps track of path weights and vertex predecessors in arrays indexed in not just one dimension, but in three dimensions.
    - More like a one-dimensional array of two-dimensional arrays.

  - We assume that all the vertices are numbered from `1` to `n`.
    - We define `shortest[u, v, x]` to be the weight of a shortest path from vertex `u` to vertex `v` in which each intermediate vertex other than `u` and `v` is numbered **`x`** or lower.
    - `shortest[u, v, n] ==  sp(u, v)` (the weight of a shortest path from `u` to `v`.)
  
  - Suppose `p` is the one with minimum weight, among all the paths from `u` to `v` that have no vertex numbered higher than `x`.
    - Then `p` would fall into one of two categories:
      1. `x` is not included in `p`. Then `shortest[u, v, x] == shortest[u, v, x-1]`.
      2. `x` is included in `p`. Then `shortest[u, v, x] == shortest[u, x, x-1] + shortest[x, v, x-1]`.
    - Because either `x` is an intermediate vertex in a shortest path from `u` to `v` or it's not, we can say that `shortest[u, v, x]` is the smaller of the two mentioned above.

  - `Theta(n^3)` time and space
    - But we can do `Theta(n^2)` if we don't need to keep track of changes between `x` ()

  - Dynamic Programming
    - *"Optimal substructure"* conditions for applying dynamic programming:
      1. We are trying to find an optimal solution to a problem.
      2. We can break an instance of the problem into instances of one or more subproblems.
      3. We use solutions to the subproblem(s) to solve the original problem, and
      4. if we use a solution to a subproblem within an optimal solution to the original problem, then the subproblem solution we use must be optimal for the subproblem.

    - In dynamic programming, we have some notion of *"size"* of a subproblem, and we often solve the subproblems in increasing order of size, so that we solve the smallest subproblems first, and then once we have optimal solutions to smaller subproblems, we can try to solve larger subproblems optimally using optimal solutions to the smaller subprolems.

    - Bottom Up: A common practice in DP is to store optimal solutions to subproblems in a table and then look them up as we compute an optimal solution to the original problem.

    - Top Down: Working from larger subproblems to smaller ones, storing the result of each subproblem in a table.

Chapter 7. Algorithms on Strings
--------------------------------

- Focusing on 3 tasks on strings:
  1. Find a longest common subsequence of two strings
  2. Given a set of operations that can transform one string to another, and the cost of each operation, find a lowest-cost way to transform one string to another.
  3. Find all occurrences of a pattern string within a string of text.

- Longest common sequence (LCS) problem
  - Subsequence isn't necsessarily a substring
  - We could check every single possible subsequence one by one: but there would be `2^m` subsequences possible.
  - Instead, we use **dynamic programming** methods: We can see that an LCS of 2 strings contains within it an LCS of the *prefixes* of the two strings.
  - We approach the problem of finding an LCS of strings `X` and `Y` in 2 steps:
    1. We'll find the length of an LCS `X` and `Y`, as well as the lengths of the longest common subsequences of all prefixes of `X` and `Y`.
    2. After computing the LCS lengths, we will 'reverse engineer' how we computed these lengths to find an actual LCS of `X` and `Y`.

    - Compute LCS table: [See the code.](https://link.iamblogger.net/vq3hp)
      - `Theta(m*n)` time.
    - Assemble LCS: [See the code.](https://link.iamblogger.net/hf8lb)
      - `O(m+n)` time. Why?
        - Observe that in each recursive call, either `i` decreases, `j` decreases, or both decreases.
        - After `m+n` recursive calls, therefore, we are guaranteed that one or the other of these indices hit 0 and the recursion bottoms out.

- Transforming one string to another
    - To find a sequence of operations that transform the string X into another string Y and has a minimum total cost.
    - Available Operations:
        - Copy: When we can directly make z_j = x_i. Then do i++, j++
        - Replace: When we have to put a new character a into z_j & ignore x_i. i++, j++
        - Delete: Ignore x_i and move on. i++
        - Insert: Insert a into z_j. j++

    - Delete and insert can be applied at every stage i,j > 0.
    - Only one of copy and replace can be applied.
    - Then cost[i, j] = Cost of transforming X_i into Y_j is the smallest of the following four:
        - cost[i-1, j-1] + c_C, but only if x_i and y_j are the same character.
        - cost[i-1, j-1] + c_R, but only if x_i and y_j differ.
        - cost[i-1, j] + c_D
        - cost[i, j-1] + c_I

    - The ComputeTransformTables procedure fills in each entry of the tables in constant time, just as the ComputeLcsTable procedure does. Because each of the tables contains (m+1) * (n+1) entries, ComputeTransformTables run in Theta(m*n) time.

    - AssembleTransformation: runs in O(m+n) time, with reasoning similar to AssembleLcs.

- String matching
    - Given a text string T and a pattern string P, we want to find all occurrences of P in T. 
    - Because we want to find all occurrences of the pattern P in the text T, a solution will be all the amounts that we can shift P by to find it in T.
        - Put another way, we say that pattern P occurs with shift s in text T if the substring of T that starts at t_(s+1) is the same as the pattern P.

        - Minimum possible shift = 0, Maximum = n-m.

    - Naive approach here whould be checking for m characters for every possible shifts from 0 to n-m.
        - O((n-m) * m).

    - Finite Automaton (FA)
        - A set of states and a way to go from state to state based on a sequence of input characters.
        - Based on the state it's in and the character it has just consumed, it moves to a new state.
        - In our string-matching application, the input sequence will be the characters of the text T, and the FA will have m+1 states, one more than the number of characters in the pattern P, numbered from 0 to m.
        - The FA starts in state 0. When it's in state k, the k most recent text characters it has consumed match the first k characters of the pattern. Whenever the FA gets to state m, therefore, it has just seen the entire pattern in the text.

        - The FA internally stores a table next-state, which is indexed by all the states and all possible input characters. The value in next-state[s, a] is the number of the state to move to if the FA is currently in state s and it has just consumed character a from the next. 

        - FA-String-Matcher running time: Theta(n)

        - How do I get next-state?
            - The goal: Given the states of prefixes P_0 through P_n, we need to determine what the next states would be when we add a new character a to these states.
            - Since we are at the current state with some prefix P_k, we've already seen the first k characters of P. If we add a new character to this P_k, then the new string could be P_(k+1), but the whole new string could also be non-prefix of P.
                - So we need to determine what is the longest prefix of P that could serve as a 'suffix' of P_k + a, so that even if the new string isn't P_(k+1), we could salvage the suffix as much as possible to continue looking for P without checking from the scratch.

        - Running time for building next-state: O(m^3 * q), 
            - since we are checking for m+1 possible states,
            - with suffix-checking comparing at most m suffix candidates, 
            - checking at most m characters each,
            - and lastly with q new character choices.

    - So total running time: Theta(n) + O(m^3 * q).

    - KMP (Knuth, Morris, and Pratt) algorithm: Avoids creating next-state and uses move-to array -> Theta(m) time. Actual pattern matching still takes Theta(n) time though

Chapter 8. Foundations of Cryptography
--------------------------------------
- Simple Substitution Ciphers
    - Shift cipher
    - Permutation
        - n! permutations
        - You can use letter frequencies and letter combinations to narrow down the choices.
    - Both the sender and receiver have to agree on the key

- Symmetric-key cryptography
    - When the sender and receiver use the same key
    - One-time pads
        - Works on bits
        - One-time pads apply XOR to bits.
            - X XOR 0 = X, X XOR 1 = OPP(X).
            - (X XOR Y) XOR Y = X

        - The key is called 'pad'
        - As the name implies, you should use a one-time pad just one time. If you use the same key k for plaintexts t_1 and t_2, then (t_1 XOR k) XOR (t_2 XOR k) = t_1 XOR t_2, which can reveal where the two two plaintexts have the same bits. 

    - Block ciphers and chaining
        - One-time pad need a key of equal lengths as the plaintext, which can be rather unwieldy.
        - Instead of doing that, they use a shorter key, and they chop up the plaintext into several blocks, applying the key to each block in turn.

        - AES: Uses elaborate methods to slice and dice a plaintext block to produce ciphertext.

        - Cipher block chaining: If the same block appears twice in the plaintext, then the same encrypted block will appear twice in the ciphertext.

            - First, create the first block of ciphertext => c_1 = E(t_1)
            - But before encrypting the second block, you XOR it, bit by bit, with c_1, so that
                - c_2 = E(c_1 XOR c_2)
            - For the 3rd, c_3 = E(c_2 XOR t_3)
            - And so on. In general, you compute the i-th block of ciphertext based on the (i-1)st block of ciphertext and the i-th block of PLAINtext => c_i = E(c_i-1 XOR t_i).

            - To decrypt, I first compute t_1 = D(c_1). From c_1 to c_2, I can compute t_2 by first computing D(c_2), which equals c_1 XOR t_2, and then XORing the result with c_1.

            - Still with block chaining, sending identical messages twice will generate the same sequence of ciphertext blocks each time.
                - One solution for this would be not starting c_0 being all 0s. Instead, you randomly generate c_0, you use that when encrypting the first block of plaintext, and so on.
                    - This c_0 is called "Initialization Vector".

    - Agreeing on common information
        - In order for symmetric-key cryptography to work, both the sender and receiver need to agree on the key. In addition, if they're using a block cipher with cipher block chaining, they might also need to agree on the initialization vector.
        - Hybrid cryptosystem

- Public-key cryptography
    - t = F_S(F_P(t))
        - In some cases, t = F_P(F_S(t)) might also be required, so that I encrypt plaintext with my secret key, anyone can decrypt the ciphertext.

    - The time required to successfully guess my F_S without knowing my secret key should be prohibitively large for anyone else.

    - Although we would want F_P to produce a different ciphertext for each possible plaintext, instead we could allow, or even prefer the element of randomization where F_P(t_1) == F_P(t_2).
        - The RSA is much more secure when the plaintext is only a small portion of what is encrypted, the bulk of the encrypted information being random "padding".

    - Another issue is that the plaintext t could take on an arbitrary number of possible values - in fact, it could be arbitrarily long - and the number of ciphertext values that F_P could convert t has to be at least as many as the number of values that t could take on.
        - Solution: block cipher

- The RSA cryptosystem
    - Modular arithmetic: We always divide by n and take the remainder.
    - The first property: If you have a number that is the product of two large secret prime numbers, then nobody else can determine these factors in any reasonable amount of time.
    - The second property: Even though factoring a large prime is hard, it's not hard to determine whether a large number is prime.
        - AKS primality test: the first algorithm to determine whether an n-bit number is prime in time O(n^c) for some constant c. Theoretically efficient, but	not yet practical for large numbers.
        - Miller-Rabin primality test
            - It can make errors, declaring a number that is actually composite to be prime. If it declares a number to be composite, however, then the number is definitely composite.)
            - The error rate is 1 in 2^s, where we can pick any positive value of s we want.

- Hybrid cryptosystems
    - You select a key k for AES and encrypt it with a RSA public key.

- Computing random numbers
    - We need random positive integers for some cryptosystems
    - A program running on a computer usually cannot be a random process, since it is built from well defined, deterministic instructions that produces the same result given the same data.
        - Some processors provide an instruction that generates random bits based on a random process, such as thermal noise within circuits.
    - Pseudorandom Number Generator (PRNG): A Deterministic program that produces 
        - a sequence of values, based on an initial value, or seed
        - a deterministic rule embodied in the program that says how to generate the next value in the sequence from the current value.

    - If you are using PRNG, you need to start with a random seed each time. 
        - In particular, the seed should be based on bits that are unbiased (not favoring either 0 or 1),
        - independent (no matter what you know about the previous bits generated, anyone has a 50% chance of correctly guessing the next bit),
        - and unpredictable to an adversary who is trying to break your cryptosystem.

Chapter 9. Data Compression
---------------------------
- Three questions
    1. Why would be want to compress information?
        - time
        - space
    2. What is the quality of the compressed information?
        - Lossy compression
        - Lossless compression
    3. Why is it possible to compress information?
        - Lossy: Tolerate how the precision decreases
        - Lossless: Remove redundant or useless bits.

- We focus on loseless compression

- Huffman codes
    - DNA strings: represented by n characters - 45% of A, 5% of C, 5% of G, and 45% of T, but the characters appear in the strand in no particular order.
    - If we used the ASCII, it would take 8n bits to represent the entire strand.
    - But since we only rely on 4 alphabets to represent the string, we really need only two bits to represent each character (00, 01, 10, 11) -> 2n bits.
    
    - But we can do even better by taking advantage of the relative frequencies of the characters.
        - Let's set A = 0, C = 100, G = 101, T = 11.
        - The more frequent characters get the shorter bit sequences.
        - Given the frequencies of the four characters, to encode the n-character strand, we need only
            - 0.45 * n * 1 + 0.05 * n * 3 + 0.05 * n * 3 + 0.45 * n * 2 = 1.65n bits.

    - Prefix-Free Code: In the encoding we used, we can see that no code is a prefix of any other code.
        - We can unambiguously match the compressed bits with their original characters.

    - If we measure the efficiency of compression methods according to the average length of the compressed information, then of the prefix-free codes, Huffman codes are the best.
        - One disadvantage though is that it requires the frequencies of all the characters to be known in advance, and therefore compression often requires 2-passes over the uncompressed text. 

        - Binary tree
            - The leaves of the tree, represent the characters, with the frequency of each.
            - The non-leaves, or internal nodes, contains the sum of the frequencies in the leaves below it.

        - We build the binary tree from bottom to top. We start with each of the n leaf nodes, corresponding to the uncompressed characters, as its own individual tree, so that initially, each leaf is also a root.
            - We then repeatedly find the two root nodes with the lowest frequencies, create a new root with these nodes as its children, and give this new root the sum of its children's frequencies.

            - The process continues until all leaves are under one root.
            - As we progress, we label each edge to a left child by 0 and each edge to a right child by 1, though once we select the two roots with the lowest frequencies, it doesn't matter which we make the left child and which we make the right child of the new root.
        
        - Running time: O(n * lg n) time, as each insert / extract-min operations of binary heap takes O(lg n) time and there are 2n - 1 of them.

        - BuildHuffmanTree procedure serve as an example of a greedy algorithm, wherein we make the decision that seems best at the moment.
            - Because we want the least-frequently appearing characters far from the root of the binary tree, the greedy approach always selects the two roots with the lowest frequency to place under a new node, which can later become a child of some other node.
            - Dijkstra's algorithm is another greedy algorithm, because it always relaxes edges from the vertex with the lowest shortest value of those remaining in its priority queue.

    - Adaptive Huffman codes
        - Instead of doing two-pass, let's update character frequencies and the binary tree as they compress or decompress in just one pass.
        - How can the decompression program determine whether the bits it's looking at represent an encoded or unencoded character? ==> Escape code.
        - Escape codes will usually appear infrequently, and so we don't want to assign them short bit sequences at the expense of a more frequently appearing character. 
        - A good way to ensure that escape codes are not short is to include an escape code character in the binary tree, but nail down its frequency to be 0 always. This will make the leaf for the escape code character to be the farthest from the root at all times.

    - Fax machines
        - Run-length encoding: Indicating the colors and lengths of runs of identical pels in the rows of the image being transmitted.
            - To determine which codes to use for which runs, a standards committee took a set of 8 representatitve documents and counted how often each run appeared.
                - Then they constructed Huffman codes for these runs.
                - The most frequent runs, and hence the shortest codes, were for runs of 2, 3, and 4 black pels, with code 11, 10, and 011, respectively.

        - In addition to compressing information only within each row of the image, some fax machines compress in both dimensions of the image.
            - A row would get encoded according to where it differs from the preceding row.
            - For most rows, the difference from the previous row is just a few pels.
            - However, this scheme entails the risk that errors propagate: an encoding or transmission error causes several consecutive rows to be incorrect.
                - For this reason, fax machines would limit the number of consecutive rows that can use this.

- LZW compression
    - Especially for text
    - Takes advantage of information that recurs in the text, though not necessarily in consecutive locations.
        - We could assign indexes to each of the words occurring in the text, but such approach would be unrealistic
        - But any recurring sequence of characters could help.

    - LZW makes a single pass over its input for compression and for decompression.
        - In both, it builds a dictionary of character sequences that it has seen, and it uses indices into this dictionary to represent character sequences.
        - Towards the beginning of the input, representing the sequences by indices could result in expansion, rather than compression, as the original sequences tend to be short. 
        - But as LZW progresses through its input, the sequences in the dictionary become longer, and representing them by an index can save quite a bit of space.

    - In LZW, the dictionary does not have to be stored with the compressed information.
        - LZW decompression rebuilds the dictionary directly from the compressed information.
        - By mirroring how the compressor would have added elements to the dictionary.
        - There is one edge case, where the index that got added to the dictionary when the previous element was added to the output, appears right after that previous element in the text. This would only occur when this element has a form of dictionary[previous] + string(previousEntry[0]).
    
    - How to find strings from the dictionary
        - Linear Search: O(n)
        - Special data structures: trie; it's like a binary tree we built for Huffman coding, except that each node can have many children, not just two, and each edge is labeled with an ASCII character.
        - The other data structure is Hash table.

    - LZW Improvements
        - Part of the problem stems from the large dictionary.
            - Moby Dick: With an output of 229,753 indices, the compressed conversion would require four times that, or 919,012 bytes.
            - Many of them are low numbers.
            - Some of the indices are occurring much more frequently than others.

        - When both of these properties hold, Huffman coding is likely to yield good results.
            - About 60% of original for Huffman, ~75% for LZW.
            - When we Huffman-encode the output of LZW, the size further reduces to ~ 40%
            - However, Huffman requires two-pass, and additional space for Huffman coding

        - Other approaches to LZW improvement:
            - Reducing the number of bits necessary to hold the indices that the compressor outputs.
                - Because many of the indices are small numbers, one approach is to use fewer bits for smaller numbers, but reserve, say, the first two bits to indicate how many bits the number requires.

            - In other approaches, we limit the size of the dictionary.
                - In one approach, once the dictionary reaches a maximum size, no other entries are ever inserted.
                - In another approach, once the dictionary reaches a maximum size, it is cleared out (except for the first 256 entries) and the process of filling out the dictionary restarts from the point in the text where the dictionary filled.

Chapter 10. Hard? Problems
--------------------------
- NP-Complete Problems
    - Traveling Salesman Problem
        - The # of possible routes that visit n locations is enormous: n!
        - No algorithm that runs in time O(n^c), for any constant c, has ever been found for the TSP.

    - Many problems are like TSP: For an input size of n, we know of no algorithm that runs in time O(n^c) for any constant c, yet nobody has proven that no such algorithm could exist.
        - If there were an algorithm that ran in O(n^c) time for any of these problems, where c is a constant, then there would be an algorithm that ran in O(n^c) time for all of these problems.
            - We call these problems NP-complete.

    - An algorithm that runs in time O(n^c) on an input of size n, where c is a constant, is a polynomial-time algorithm, so called because n^c with some coefficient would be the most significant term in the running time.

    - We know of no polynomial-time algorithm for any NP-complete problem, but nobody has proven that it's impossible to solve some NP-complete problem in polynomial time.

    - Many NP-Complete problems are almost the same as problems that we know how to solve in polynomial time. 
        - Bellman-Ford Algorithm, which finds shortest paths from a single source in a graph G, is polynomial.
            - However, finding a longest acyclic path between two vertices is NP-Complete.
            - In fact, merely determining whether a graph contains a path without cycles with at least a given # of edges is NP-complete.

        - Euler tour vs. Hamiltonian cycle
            - Both have to with finding paths in a connected, undirected graph.
                - Undirected: edges have no direction, so that (u, v) == (v, u). We say that (u, v) is incident on vertices u and v.
                - Connected: has a path between Every pair of vertices.

            - Euler tour: Starts and ends at the same vertex and visits each Edge exactly once, though it may visit each vertex more than once.

            - Hamiltonian cycle: Start and ends at the same vertex and visits each vertex exactly once (except for the vertex at which it starts and ends)

            - Deciding whether the graph has an Euler tour is easy: Determine the degree of each vertex, that is, how many edges are incident on it. The graph has an Euler tour iff the degree of each vertex is even.

            - However, asking if the graph has a Hamiltonian cycle is NP-complete.

            - Notice that the question is about whether the tour or cycle exists, not the actual cycle

- The classes P and NP and NP-completeness
    - Computer scientists generally regard problems solvable by polynomial-time algorithms as "tractable," meaning "easy to deal with."
    - If a polynomial-time algorithm exists for a problem, then we say that this problem is in the class P.

    - Now suppose that you're given a proposed solution to a problem, and you want to verify that the solution is correct.
        - Hamiltonian Cycle: A proposed solution would be a sequence of vertices. You could easily verify that this solution is correct in polynomial time.
        - If it is possible to verify a proposed solution to a problem in time polynomial in the size of the input to the problem, then we say that this problem is in the class NP.
        - We call the proposed solution a certificate, and in order for the problem to be in NP, the time to verify the certificate needs to be polynomial in the size of the input to the problem and the size of the certificate.

    - If you can solve a problem in polynomial time, then you can certainly verify a certificate for that problem in polynomial time. In other words, every problem in P is automatically in NP.
        - The reverse - is every problem in NP also in P? - is the question that has perplexed computer scientists for all these years. We often call it the "P = NP? problem."

    - The NP-complete problems are the "hardest" in NP. Informally, a problem is NP-complete if it satisfies two conditions:
        (1) It's in NP
        (2) If a polynomial-time algorithm exists for the problem, then there is a way to convert every problem in NP into this problem in such a way as to solve them all in polynomial time.

        - If a polynomial-time algorithm exists for any NP-complete problem - that is, any NP-complete problem is in P - then P = NP.

        - Because NP-complete problems are the hardest in NP, if it turns out that any problem in NP is not polynomial-time solvable, then none of the NP-complete problems are.

        - A problem is NP-hard if it satisfies the second condition for NP-completeness but may or may not be in NP.

- Decision problems and reductions
    - Decision problems: Their output is a single bit, indicating "yes" or "no".
    - Optimization problems: We want to find the best possible solutions, rather than decision problems.
        - Recast optimization problems into decision problems
                ex) Does the graph contain a path between two specific vertices whose path weight is at most a given value k?
                    - Keep doubling the value of k until the answer is yes.
                    - If that last value of k was k', then the answer is somewhere between k'/2 and k'.
                    - Then find the true answer by using binary search with an initial interval of k'/2 to k'.

    - Reduction: A way to convert every problem in NP into this problem that have a polynomial-time algorithm, so that we can solve every NP problem in polynomial time.
        - We're given some input x of size n to problem X.
        - We transform this input into an input y to problem Y, and we do so in time polynomial in n, say O(n^c) for some constant c.

        - The way we transform input x into input y has to obey an important property:
            - If algorithm Y decides "yes" on input y, then algorithm X should also decide "yes" on input x, and if Y decides "no" on y, then X should decide "no" on x.

        - We call this transformation a "polynomial-time reduction algorithm."

        - The reduction algorithm takes O(n^c) time, and its output cannot be longer than the time it took, and so the size of the reduction algorithm's output is O(n^c).
            - But this output is the input y to the algorithm for problem Y.
        - Since the algorithm for Y is a polynomial-time algorithm, on an input size of size m, it runs in time O(m^d) for some constant d.
            - Here, m is O(n^c), and so the algorithm for Y takes time O(n^(c*d)).
            - Because both c and d are constants, so is c*d, and we see that the algorithm for Y is a polynomial-time algorithm.
            - The total time for the algorithm for problem X is O(n^c + n^(c*d)), which makes it, too, a polynomial-time algorithm.

        - This approach shows that if problem Y is "easy" (solvable in polynomial time), then so is problem X.

        - But we'll use polynomial-time reductions to show not that problems are easy, but that they are hard:
            - If problem X is NP-hard and we can reduce it to problem Y in polynomial time, then problem Y is NP-hard as well.

                - Let's suppose that problem X is NP-hard and that there is a polynomial-time reduction algorithm to convert inputs to X into inputs to Y.
                - Because X is NP-hard, there is a way to convert any problem, say Z, in NP into X such that if X has a polynomial-time algorithm, so does Z.
                - Now we note that if we immediately follow the polynomial-time reduction for Z to X by the polynomial-time reduction from X to Y, we have a polynomial-time reduction from Z to Y.
                - The two polynomial-time reductions in sequence together constitute a single polynomial-time reduction.

        - Because X being NP-hard means that every problem in NP reduces to it in polynomial time, we picked any problem Z in NP that reduces to X in polynomial time and showed that it also reduces to Y in polynomial time.

        - Our ultimate goal is to show that problems are NP-complete.
            - Show that it's in NP by showing a way to verify a certificate in polynomial time.
            - Take some other problem X that we know to be NP-hard and give a polynomial-time reduction.

        - We need to start with some NP-complete problem M (the Mother Problem) that every problem in NP reduces to in polynomial time. Then we can reduce M to some other problem in polynomial time to show that the other problem is NP-hard, reduce the other problem to yet some other problem to show that the latter is NP-hard, and so on.

- A Mother Problem
    - Boolean Formula Satisfiability Problem: Decide whether any given boolean formula has a satisfying assignment to its variables
    - 3-CNF satisfiability
        - Instead of directly reducing from the boolean formula satisfiability, let's require that
            - the formula be ANDs of clauses
                - where each clause is an OR of three terms
            - each term is a literal
                - either a variable or the negation of a variable
            - A boolean formula in this form is in 3-conjunctive normal form or 3-CNF.

        - 3-CNF is NP-complete
            - NP: We can plugin the certificate to check if the formula holds
            - NP-hard: We reduce from (unrestricted) boolean formula satisfiability.

        - 2-CNF can be solved in polynomial time though.

    - Clique
        - A clique in an undirected graph G is a subset S of vertices such that the graph has an edge between every pair of vertices in S.
            - The size of a clique is the # of vertices it contains.

        - The clique problem takes 2 inputs, a graph G and a positive integer k,
            - and asks whether G has a clique of size k.

            - NP: Verifying a certificate is easy. We just have to check that each of the k vertices has an edge to the other k-1.

            - NP-hard: 
                - We start with a formula C_1 AND C_2 AND ... AND C_k, where each C_r is one of k clauses.
                - From this formula, we will construct a graph in polynomial time, and this graph will have a k-clique iff the 3-CNF is satisfiable.

                    - 3 things to consider:
                        - the construction
                        - An argument that the construction runs in time polynomial in the size of the 3-CNF formula
                        - A proof that the graph has a k-clique iff there is some way to assign to the variables of the 3-CNF so that it evaluates to 1.

                - The construction
                    - On C_r, it has 3 literals: C_r = l_1_r OR l_2_r OR l_3_r.
                    - Each literal is either a variable or the negation of a variable.
                    - We create one vertex for each literal, so that for clause C_r, we create a triple of vertices, v_1_r, v_2_r, and v_3_r. 
                    - We add an edge between v_i_r and v_j_s if 2 conditions hold:
                        - v_i_r and v_j_s are in different triples: that is, r and s are different clause numbers
                        - Their corresponding literals are not negations of each other.
                        - Note that we will do this for every vertices as long as the conditions hold.

                - Polynomial time of the construction
                    - If the 3-CNF formula has k clauses, then it has 3k literals, and so the graph has 3k vertices.
                    - At most, each vertex has an edge to all the other 3k - 1 vertices, and so the number of edges is at most 3k * (3k - 1) = 9 * k^2 - 3 * k.
                    - The size of the graph constructed is polynomial in the size of the 3-CNF input, and it's easy to determine which edges go into the graph.

                - The Proof
                    - We start by assuming that the formula is satisfiable, and we'll show that the graph has a k-clique.
                        - If there exists a satisfying assignment, each clause C_r contains at least one literal l_i_r that evaluates to 1, and each such literal corresponds to a vertex v_i_r in the graph.
                        - If we select one such literal from each of the k clauses, we get a corresponding set S of k vertices.
                        - I claim that S is a k-clique.
                            - Consider any 2 vertices in S.
                            - They correspond to literals in different clauses that evaluate to 1 in the satisfying assignment.
                            - They cannot be negations of each other, because it they were, then one of them would evaluate to 1 but the other would evaluate to 0.
                            - Since these literals are not negations of each other, we created an edge between the two vertices when we constructed the graph.
                            - Because we can pick any two vertices in S as this pair, we see that there are edges between all pairs of vertices in S.
                            - Hence, S, a set of k vertices, is a k-clique.

                    - The other direction: If the graph has a k-clique S, then the 3-CNF formula is satisfiable.
                        - No edges in the graph connect vertices in the same triple, and so S contains exactly one vertex per triple.
                        - For each vertex v_r_i in S, assign 1 to its corresponding literal l_r_i in the 3-CNF formula.
                            - We don't have to worry about assigning a 1 to both a literal and its negation, since the k-clique cannot contain vertices corresponding to a literal and its negation.
                        - Since each clause has a literal that evaluates to 1, each clause is satisfied, and so the entire 3-CNF formula is satisfied.

                        - If any variables don't correspond to vertices in the clique, assign values to them arbitrarily; they won't affect whether the formula is satisfied.

    - Vertex Cover
        - A vertex cover in an undirected graph G is a subset S of the vertices such that every edge in G is incident on at least one vertex in S. We say that each vertex in S "covers" its incident edges.
        - The size of a vertex cover is the number of vertices it contains.

        - As in the clique problem, the vertex-cover problem takes as input an undirected graph G and a positive integer m. It asks whether G has a vertex cover of size m.
        
        - Example: You have a building with hallways and cameras that can scan up to 360 degrees located at the intersections of hallways, and you want to know whether m cameras will allow you to see all the hallways.

        - It's easy to verify in time polynomial in the size of the graph that the proposed vertex cover has size m and really does cover all the edges, and so we see that this problem is in NP.

    - Hamiltonian cycle and Hamiltonian path
    - Traveling salesman
    - Longest acyclic path
    - Subset sum
    - Partition
    - Knapsack

- General strategies
    - There is no one-size-fits-all way to reduce one problem to another in order to prove NP-hardness.
    - Go from general to specific
        - When reducing problem X to problem Y, you always have to start with an arbitrary input to problem X. But you are allowed to restrict the input to problem Y as much as you like.
    - Take advantage of restrictions in the problem you're reducing from
    - Look for special cases
    - Select an appropriate problem to reduce from
        - Sometimes, however, it's best to leap from one domain to another. 3-CNF satisfiability often turns out to be a good choice to reduce from when crossing domains.
    - Make big rewards and big penalties
    - Design widgets

- Perspective
    - When a problem is NP-complete, it means that some inputs are troublesome, but not necessarily that all inputs are bad.
        - Finding a longest 'acyclic' path in a directed graph is NP-complete, but if you do know that the graph is acyclic, then you can find one in O(n+m) time.

    - Moreover, we can adopt some methods that would allow you to bring often very good results.
        - Branch and Bound
        - Neighborhood Search
        - Approximation Algorithms
            - Strangely enough, if two NP-complete problems are closely related, the solution produced by a good approximation algorithm for one might produce a poor solution for the other.

- Undecidable problems
    - For some problems, no algorithms is possible.
        - Halting problem
            - The input is a computer program A and the input x to A.
            - The goal is to determine whether program A, running on input x, ever halts. That is, does A with input x run to completion?
                - How do we know when to declare that A will never halt?
        - Post's Correspondence Problem (PCP)
            - At least 2 characters
            - 2 lists of n strings, A and B
            - The problem is to determine whether there exists a sequence of indices i_1, i_2, i_3, ..., i_m such that A_i_1 A_i_2 A_i_3 ... A_i_m concatenated together gives the same string as B_i_1 B_i_2 B_i_3 ... B_i_m.

            - If there's one such solution, then there are an infinite number of solutions, since you can just keep repeating the index sequence of a solution.

            - For PCP to be undecidable, we have to allow the strings in A and B to be used more than once, since otherwise you could just list all the possible combination of strings.

            - Although PCP might not seem particularly interesting on its own, we can reduce it to other problems to show that they too are undecidable.
            - Context-Free Grammars (CFGs): A set of rules for generating a formal language. 
                - By reducing from PCP, we can prove that it's undecidable whether two CFGs generate the same formal language, or the two generate any strings in common.