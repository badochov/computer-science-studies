#!/bin/bash -l

#SBATCH --job-name mim-hello          # this will be shown in the queueing system
#SBATCH --output "mim-hello-%j.out"   # stdout redirection
#SBATCH --error "mim-hello-%j.err"    # stderr redirection
#SBATCH -p plgrid-short               # partition: plgrid-short fits jobs up to 1 hour
#SBATCH --account "plgmimuwhpc23"     # the number of our grant
#SBATCH --nodes 2                     # how many nodes we want
#SBATCH --tasks-per-node 24           # each node is 2 socket, 12 core, so we want 24 tasks on each node
#SBATCH --time 00:05:00               # if the job runs longer than this, it'll be killed

module add plgrid/tools/impi  # load Intel MPI module
module add plgrid/tools/cmake # load cmake module (if you need to compile)

# optional: compile before you run
# (Prometheus does not compile on the head node)
mkdir build
cd build
srun -N 1 -n 1 cmake ..
srun -N 1 -n 1 make
# end of compilation
