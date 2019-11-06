Graph Note
=====
This is a note for the [Graph Theory Playlist](https://www.youtube.com/playlist?list=PLDV1Zeh2NRsDGO4--qE8yH72HFL1Km93P) by [WilliamFiset](https://www.youtube.com/channel/UCD8yeTczadqdARzQUp29PJw). 

- [Graph Note](#graph-note)
    - [Overview](#overview)
      - [Types of graphs](#types-of-graphs)
      - [Graph representations](#graph-representations)
        - [Adjacency matrix](#adjacency-matrix)
        - [Adjacency list](#adjacency-list)
        - [Edge list](#edge-list)
    - [Common Graph Theory Problems](#common-graph-theory-problems)
      - [Shortest Path Problem](#shortest-path-problem)
      - [Connectivity](#connectivity)
      - [Detecting Negative Cycles](#detecting-negative-cycles)
      - [Strongly Connected Components](#strongly-connected-components)
      - [Traveling Salesman Problem (NP-Hard)](#traveling-salesman-problem-np-hard)
      - [Bridges](#bridges)
      - [Articulation Points](#articulation-points)
      - [Minimum Spanning Tree (MST)](#minimum-spanning-tree-mst)
      - [Max Flow](#max-flow)
    - [Depth-First Search (DFS)](#depth-first-search-dfs)
      - [Connected Components](#connected-components)
    - [Breadth-First Search (BFS)](#breadth-first-search-bfs)
      - [Construct the Shortest Path](#construct-the-shortest-path)
      - [BFS Shortest Path on a Grid](#bfs-shortest-path-on-a-grid)

### Overview
#### Types of graphs
- Undirected graph
- Directed graph (Digraph)
- Weighted graph
- Tree 
  - Graph **without cycles**
- Directed acyclic graph (DAG)
- Bipartite graph
  - bipartite graph is **two colorable**
- Complete graph

#### Graph representations
##### Adjacency matrix 
A $V^2$ matrix `m` in which `m[i][j]` represents the weight from node `i` to `j`. 

- Pros
  - Space efficient for dense graphs
  - Constant edge weight lookup
  - Simple
- Cons
  - O($V^2$) space, not efficient for sparse graphs
  - O($V^2$) complexity for iterating over all edges, not friendly for sparse graphs

##### Adjacency list
A map to track the **outgoing edges** for each node.

- Pros
  - Space efficient for sparse graph
  - Efficient iteration
- Cons
  - Less space efficient for denser graphs
  - O(E) edge weight lookup (rarely needed)

##### Edge list
An unordered list of edges (seldomly used).

- Pros
  - Efficient iteration
  - Simple
- Cons
  - Less space efficient for denser graphs
  - O(E) edge weight lookup (rarely needed)


### Common Graph Theory Problems
#### Shortest Path Problem
- Breadth-first search
- Dijkstra's
- Bellman-Ford
- Flyod-Warshall
- A*

#### Connectivity
- Union-find
- Depth-first search

#### Detecting Negative Cycles
- Bellman-Ford
- Floyd-Warshall

#### Strongly Connected Components
- Tarjan's
- Kosaraju's 

#### Traveling Salesman Problem (NP-Hard)
- Held-Karp 
- Branch and bound

#### Bridges
A bridge (or cut edge) is any edge in a graph whose removal increases the number of connected components.

#### Articulation Points
An articular point (or cut vertex) is any node in a graph whose removal increases the number of connected components.

#### Minimum Spanning Tree (MST)
A MST is a subset of the edges of a connected, edge-weighted graph that connects all the vertices together, **without any cycles** and **with the minimum possible total edge weight**.

Algorithms:
- Kruskal's
- Prim's 
- Boruvka's

#### Max Flow
In graphs with edges whose weight represents capacity. 

Algorithms:
- Ford-Fulkerson
- Edmons-Karp
- Dinic's

### Depth-First Search (DFS)
- Complexity: O(V+E)

#### Connected Components 
- Perform DFS on every unvisited node
- Label a node as visited after the visit
- Increment the number of connected components by one once the DFS on a node finishes
- In the end, we can get the total number of connected components in the graph


### Breadth-First Search (BFS)
- Complexity: O(V+E)
- Application: finding the shortest path on a unweighted graph

#### Construct the Shortest Path
- Use an array to keep track of the parent of every visited node

#### BFS Shortest Path on a Grid