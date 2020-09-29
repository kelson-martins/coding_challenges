from __future__ import print_function
from ortools.linear_solver import pywraplp


def main():

    solver = pywraplp.Solver(
                            'mixedinteger',
                            pywraplp.Solver.CBC_MIXED_INTEGER_PROGRAMMING
                            )    


    # x1 and x2 representing the 2 mixes
    x1 = solver.NumVar(0, solver.infinity(), 'x1')
    x2 = solver.NumVar(0, solver.infinity(), 'x2')

    # adding the constraints
    c1 = solver.Constraint(12, solver.infinity())
    c1.SetCoefficient(x1, 2)
    c1.SetCoefficient(x2, 4)    

    # adding the constraints
    c2 = solver.Constraint(15, solver.infinity())
    c2.SetCoefficient(x1, 5)
    c2.SetCoefficient(x2, 3)    

    # adding the constraints
    c3 = solver.Constraint(8, solver.infinity())
    c3.SetCoefficient(x1, 4)
    c3.SetCoefficient(x2, 1)            

    objective = solver.Objective()
    objective.SetCoefficient(x1, 9)
    objective.SetCoefficient(x2,7)
    objective.SetMinimization()
    status = solver.Solve()

    if status == pywraplp.Solver.OPTIMAL:
        print('Solution:')
        print('Objective value =', solver.Objective().Value())
        print('x =', x1.solution_value())
        print('y =', x2.solution_value())
    else:
        print('The problem does not have an optimal solution.')
if __name__ == "__main__":
    main()