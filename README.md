### Description

This is my solution to the vorto challenge solving VRP. The solution uses the Nearest Neighbor heuristic + Two-Opt optimization to plan efficient routes for drivers, while making sure they are under their time constraint

### Installation

    git clone https://github.com/bnkc/vorto-challenge.git
    cd vorto-challenge


### Install Dependencies:


    go mod tidy

### Running the Project

    go build -o vorto cmd/main.go
    ./vorto problemDir/data/problem1.txt

### Run the Evaluation Script
Ensure you have Python installed. You can use the following command to run the evaluation script:

    python problemDir/evaluateShared.py --cmd ./vorto --problemDir problemDir/data

### Results

The current implementation yields the following results:

    Mean Cost: 51966.92022606448
    Mean Run Time: 4.548835754394531ms

### Other Considerations

While the current solution uses  Nearest Neighbor + Two-Opt for its simplicity and performance, I tried several other approaches:

- **Clarke-Wright Savings Heuristic**: A more sophisticated approach but I didn't see great improvements in cost vs performance.
- **Nearest Neighbor with Lookahead**: Actually a pretty good approach. improves cost greatly at the expense of run time. (look ahead by 3)
- **Greedy Insertion Heuristic:** Struggled to get this one working.
- **Simulated Annealing**: Also marginal improvements to cost `Mean Cost: 50722`
- **Multi-Start Local Search**: No clear improvements seen in cost or run time

Ultimately, I chose simplicity, performance, and straightforward implementation, making it a suitable choice for this challenge.

### References

- [Solving Vehicle Routing Problems with Python: Heuristics Algorithm](https://medium.com/@writingforara/solving-vehicle-routing-problems-with-python-heuristics-algorithm-2cc57fe7079c)
- [What is the Vehicle Routing Problem?](https://www.routific.com/blog/what-is-the-vehicle-routing-problem)
- [The Vehicle Routing Problem: Exact and Heuristic Solutions](https://towardsdatascience.com/the-vehicle-routing-problem-exact-and-heuristic-solutions-c411c0f4d734)

