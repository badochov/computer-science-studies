# CA3DMM

## Optimizations
- custom communicators for:
  - pk group
  - cannon group
  - reducing results
  - creating replicas of matrix fragment
- final results are reduces using tree instead of linear communication and summing
- matrix is pre-skewed at generation time reducing unnecessary additional communication round
- shift of matrix A and matrix B is done concurrently