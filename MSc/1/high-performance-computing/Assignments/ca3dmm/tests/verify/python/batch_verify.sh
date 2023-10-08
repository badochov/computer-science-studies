#!/bin/bash -l

# This script runs the verification python script on Prometheus.
# Prometheus does not allow us to run MPI in an interactive session.
# So we need to run everything in a big batch job.

# Run instructions:
# (1) Move your .zip to ../solutions/, e.g., ../solutions/ab123456.zip
# (2) mkdir ../logs
# (3) Put ab123456.zip in the last line in this script
# (4) sbatch batch_verify.sh
# 
# After the batch job completes,
# stderr (to debug script problems) is in ../logs/batch-verify-*.err
# python logs (to debug the output of your program) is in ../logs/verifyprograms-run*.log
# result summary is in ../logs/results-run*.log

#SBATCH -J ca3dmm-batch-verify
#SBATCH -N 2
#SBATCH --ntasks-per-node=24
#SBATCH --mem-per-cpu=5GB
#SBATCH --time=00:10:00 
#SBATCH -A plgmimuwhpc23
#SBATCH -p plgrid-short
#SBATCH --output="../logs/batch-verify-%j.out"
#SBATCH --error="../logs/batch-verify-%j.err"

module add plgrid/tools/impi
module add plgrid/tools/cmake
module add plgrid/tools/python/3.9

cd $SLURM_SUBMIT_DIR

python3 verify.py hb417666