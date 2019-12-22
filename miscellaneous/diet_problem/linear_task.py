from __future__ import print_function
from ortools.linear_solver import pywraplp


def main():

    solver = pywraplp.Solver(
                            'LinearExample',
                            pywraplp.Solver.GLOP_LINEAR_PROGRAMMING
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
    solver.Solve()

    print("x1: {}".format(x1.solution_value()))
    print("x2: {}".format(x2.solution_value()))    

    print("Cost: {}".format( 9 * x1.solution_value() + 7 * x2.solution_value()   ))

if __name__ == "__main__":
    main()