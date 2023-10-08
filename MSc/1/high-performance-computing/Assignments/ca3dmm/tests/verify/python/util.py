import numpy as np
import logging


log = logging.getLogger("util")


def read_denses(filename, num_matrices):
  f = open(filename, "r")
  matrices = []
  while len(matrices) < num_matrices:
    shape_txt = f.readline()
    shape = shape_txt.split()
    try:
        num_rows = int(shape[0])
        num_cols = int(shape[1])
    except (ValueError, IndexError):
        raise ValueError(f"shape string malformed. line \n{shape_txt}")
    rows = []
    for row_txt in f:
        row = [float(value) for value in row_txt.split()]
        if len(row) != num_cols:
            raise ValueError(
              f"expecting {num_cols} cols got {len(row)}; line\n{row_txt}")
        rows.append(row)
        if len(rows) == num_rows:
            break
    if len(rows) != num_rows:
        raise ValueError()
    matrices.append(np.matrix(rows))
  return matrices


def isclose(a, b, rel_tol=1e-04, abs_tol=1e-04):
  return abs(a-b) <= max(rel_tol * max(abs(a), abs(b)), abs_tol)


def compare_denses(mat_a, mat_b):
  if (mat_a.shape != mat_b.shape):
    return (
      False, 
      f"different shapes: expected {mat_a.shape} seen {mat_b.shape}")
  for row in range(0, mat_a.shape[0]):
    for col in range(0, mat_a.shape[1]):
      a = mat_a[row,col]
      b = mat_b[row,col]
      if not isclose(a,b):
        return (
          False, 
          f"different values at row {row} col {col} expecting {a} got {b}")
  return (True, "")