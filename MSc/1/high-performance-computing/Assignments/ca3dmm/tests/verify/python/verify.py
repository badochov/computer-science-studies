# This script is configured for Prometheus and needs to be run in a Slurm batch job.
# See batch_verify.sh for instructions


import logging
import os
import subprocess
import signal
import shutil
import zipfile
import itertools
import sys

log = logging.getLogger('verifyprograms')

from util import read_denses, compare_denses

class Alarm(Exception):
  pass


def alarm_handler(signum, frame):
  raise Alarm


def get_run_number(filebase='results'):
  run_number = 1
  while os.path.isfile(f'{filebase}-run-{run_number:02}.csv'):
    run_number += 1
  return f'{run_number:02}'


if __name__ == '__main__':
  logging.basicConfig(level=logging.DEBUG)
  solutions = sys.argv[1].split(' ')

  try:
    os.mkdir('../logs/')
  except FileExistsError:
    pass
  try:
    os.mkdir('../tests/')
  except FileExistsError:
    pass

  run_number = get_run_number('../logs/results')
  log.info('run number %s', run_number)

  fh = logging.FileHandler('../logs/verifyprograms-run'+run_number+'.log')
  formatter = logging.Formatter('%(asctime)s : %(message)s')
  fh.setFormatter(formatter)
  log.addHandler(fh)

  timelimit = 10  # in seconds, should be way less

  basedir = os.getcwd()
  ref_results_dir = basedir + "/../verimat/"

  testdir = '../tests/tests-'+run_number+'/'
  os.makedirs(testdir)

  with open('../logs/results-run'+run_number+'.csv', 'a') as results_stat:
    for solution in solutions:
      results_stat.flush()
      results_stat.write('\n'+solution+',')
      os.chdir(basedir)
      log.info('solution: %s' % solution)

      zip_location = "../solutions/" + solution + '.zip'
      log.info('zip file location: %s', zip_location)
      if not os.path.exists(zip_location):
        log.error('zip file missing')
        results_stat.write('-1, missing')
        continue
      shutil.copy(zip_location, testdir)
      os.mkdir(testdir + solution)
      zipf = zipfile.ZipFile(testdir + solution + '.zip')
      try:
        zipf.extractall(testdir + solution)
      except zipfile.BadZipFile:
        log.error('zip file misformatted')
        results_stat.write('-1, misformatted')
        continue

      build_dir = testdir + solution + '/build'
      try:
        os.rmdir(build_dir)
      except FileNotFoundError:
        pass
      os.mkdir(build_dir)
      os.chdir(build_dir)
      os.mkdir('outputs')

      retcode = subprocess.call(
        ('srun', '-N', '1', '-n', '1', 'cmake', '..'), 
        stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
      if retcode != 0:
        log.error('cmake non-0 retcode')
        results_stat.write('-1, cmake')
        continue
      retcode = subprocess.call(
        ('srun', '-N', '1', '-n', '1', 'make'),
        stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
      if retcode != 0:
        log.error('make non-0 retcode')
        results_stat.write('-1, make')
        continue

      mat_sizes = [(10, 6, 4), (30, 30, 20), (6, 6, 80), (500, 150, 50)]
      ge_nums = [1, 5, 20, 12]
      seeds = [42, 442, 13, 17, 25, 41]
      mpi_sizes = [4, 8, ]
      with open('errors.txt', 'w') as error_file:
        for ((mat_size, ge_num), mpi_size, out_flag) in itertools.product(
          zip(mat_sizes, ge_nums), mpi_sizes, ('-v', '-g')):

          str_mat_size = "_".join(str(size) for size in mat_size)
          results_filename = f'outputs/mpiresult_{mpi_size}_{str_mat_size}_{out_flag[1]}.txt'
          with open(results_filename, 'w') as results_file:
            params = ['mpiexec', '-n', 
                      str(mpi_size),
                      './ca3dmm',
                      str(mat_size[0]), str(mat_size[1]), str(mat_size[2]),
                      '-s', ",".join([str(seed) for seed in seeds]),
                      out_flag]
            if out_flag == '-g':
              params.append(str(ge_num))
            compact_params = " ".join(params)
            log.info("dir: %s params: %s", solution, compact_params)
            signal.signal(signal.SIGALRM, alarm_handler)
            signal.alarm(timelimit)
            child = None
            try:
              child = subprocess.Popen(
                params, stdout=results_file, stderr=subprocess.DEVNULL)
              outcode = child.wait()
              signal.alarm(0)
              log.info("subprocess finished")
              if outcode != 0:
                log.info("non-zero exit code %d" % (outcode))
                results_stat.write('-1, outcode, ')
                error_file.write(compact_params+"\n")
                continue
            except FileNotFoundError:
              signal.alarm(0)
              log.info("exec not found")
              results_stat.write('-1, execfile, ')
              error_file.write(compact_params+"\n")
              continue
            except Alarm:
              log.info("timeout!")
              if child is not None:
                log.info("killing the child")
                child.kill()
                log.info("child killed")
                results_stat.write('-1, timeout, ')
                error_file.write(compact_params+"\n")
              continue

          correct_count = 0
          ref_result = None

          if out_flag == '-g':
            reference_filename = f'{ref_results_dir}ge_{str_mat_size}_{ge_num}.txt'
            ref_result = []
            with open(reference_filename) as ref_file:
              ref_result = ref_file.readlines(len(seeds)//2)
            ref_result = [int(result) for result in ref_result]

            with open(results_filename) as res_file:
              prog_result = res_file.readlines(len(seeds)//2)
            
            try:
              prog_result = [int(result) for result in prog_result]
            except ValueError:
              log.info('result is malformed')
              results_stat.write('-1, formaterror, ')
              error_file.write(compact_params+"\n")
              continue
            
            if len(prog_result) != len(ref_result):
              log.info('insufficient number of results')
              results_stat.write('-1, numresult, ')
              error_file.write(compact_params+"\n")
              continue
            
            for (ref, res) in zip(ref_result, prog_result):
              if ref != res:
                log.info('incorrect result: expected %d got %d', ref, res)
                error_file.write(compact_params+"\n")
              else:
                correct_count += 1
            
          else: # out_flag == '-v'
            reference_filename = f'{ref_results_dir}mat_{str_mat_size}.txt'
            ref_result = read_denses(reference_filename, len(seeds)/2)
            try:
              results = read_denses(results_filename, len(seeds)/2)
            except ValueError:
              log.info("result is malformed")
              results_stat.write('-1, formaterror, ')
              error_file.write(compact_params+"\n")
              continue

            correct_count = 0      
            for (ref, res) in zip(ref_result, results):
              (is_equal, message) = compare_denses(ref, res)
              if not is_equal:
                log.info('incorrect result: %s', message)
                error_file.write(compact_params+"\n")
              else:
                correct_count += 1

          status = 'OK' if correct_count == len(ref_result) else 'err'
          results_stat.write(f'{correct_count}, {status}, ')
          log.info("num result %d status %s", correct_count, status)

    log.info('Correctly quitting')