# kcliques

My solution uses presented in the paper approach of orienting graph the graph and the traversing it in a mix of DFS and BFS.

## Graph preprocessing
Edges are oriented from vertex with smaller degree to the vertex with higher degree. Then the graph is converted to the CSR representation.

This part is done on the CPU.

## K-Clique counting
As described in the paper each block counts k-cliques of a separate vertex and after finishing processing vertex takes next vertex using work stealing implemented via atomic counter.
Then each block computes induced subgraph of the vertex.
Then each block divides into smaller WARP size groups and each warp processes in parallel different neighbour of root vertex.
Each thread in a block holds its own counter of k-cliques found, at the end those local counters are reduces to block counter and block leader adds cliques to the global counter.
I put emphasis to increase locality and minimise number of global memory accesses.

## Optimisations
- Graph orientation
- Induced subgraph
- Binary representation of the binary subgraph
- Mixed BFS and DFS approach
- 3 level parallelization

## Notes
Had I started the task earlier I would like to implement one more optimization of using smaller `__shared__` adjacency list for small subgraph, as most of the subgraphs are small.
