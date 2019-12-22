from ortools.sat.python import cp_model

def main():

    model = cp_model.CpModel()
    maxT = 8    

    # initial set of predicates
    WolfOnThisSide     = [None] * maxT
    CabbageOnThisSide  = [None] * maxT
    SheepOnThisSide    = [None] * maxT
    FerrymanOnThisSide = [None] * maxT
    WolfOnTheOtherSide = [None] * maxT
    CabbageOnTheOtherSide  = [None] * maxT
    SheepOnTheOtherSide    = [None] * maxT
    FerrymanOnTheOtherSide = [None] * maxT

    # initial set of operators
    MoveWolfAcross    = [None] * maxT
    MoveCabbageAcross = [None] * maxT
    MoveSheepAcross   = [None] * maxT
    MoveWolfBack      = [None] * maxT
    MoveCabbageBack   = [None] * maxT
    MoveSheepBack     = [None] * maxT
    MoveBack          = [None] * maxT
    MoveAcross        = [None] * maxT


    for t in range(maxT):
        WolfOnThisSide[t]         = model.NewBoolVar("WolfOnThisSide"+str(t))
        CabbageOnThisSide[t]      = model.NewBoolVar("CabbageOnThisSide"+str(t))
        SheepOnThisSide[t]        = model.NewBoolVar("SheepOnThisSide"+str(t))
        FerrymanOnThisSide[t]     = model.NewBoolVar("FerrymanOnThisSide"+str(t))
        WolfOnTheOtherSide[t]     = model.NewBoolVar("WolfOnTheOtherSide"+str(t)) 
        CabbageOnTheOtherSide[t]  = model.NewBoolVar("CabbageOnTheOtherSide"+str(t))
        SheepOnTheOtherSide[t]    = model.NewBoolVar("SheepOnTheOtherSide"+str(t))
        FerrymanOnTheOtherSide[t] = model.NewBoolVar("FerrymanOnTheOtherSide"+str(t))

    # initial constraints
    # predicates are on this side at initial step
    model.AddBoolAnd ([
        WolfOnThisSide[0],
        CabbageOnThisSide[0],
        SheepOnThisSide[0],
        FerrymanOnThisSide[0]
    ])

    # predicates are not on the other side on initial state
    model.AddBoolAnd ([
        WolfOnTheOtherSide[0].Not(),
        CabbageOnTheOtherSide[0].Not(),
        SheepOnTheOtherSide[0].Not(),
        FerrymanOnTheOtherSide[0].Not()
    ])    

    # at final step, goal is to predicates to be on the other side
    model.AddBoolAnd([
        WolfOnTheOtherSide[maxT-1],
        CabbageOnTheOtherSide[maxT-1],
        SheepOnTheOtherSide[maxT-1]
    ])

    for t in range(maxT-1):
        MoveWolfAcross[t] = model.NewBoolVar("MoveWolfAcross"+str(t))
        MoveCabbageAcross[t] = model.NewBoolVar("MoveCabbageAcross"+str(t))
        MoveSheepAcross[t] = model.NewBoolVar("MoveSheepAcross"+str(t))
        MoveAcross[t] = model.NewBoolVar("MoveAcross"+str(t))
        MoveWolfBack[t] = model.NewBoolVar("MoveWorldBack"+str(t))
        MoveCabbageBack[t] = model.NewBoolVar("MoveCabbageBack"+str(t))
        MoveSheepBack[t] = model.NewBoolVar("MoveSheepBack"+str(t))
        MoveBack[t] = model.NewBoolVar("MoveBack"+str(t))

    # pre and post conditions for all predicates
    for t in range(maxT-1):
        
        # wolf
        model.AddBoolAnd([
            WolfOnThisSide[t],
            FerrymanOnThisSide[t],
            WolfOnTheOtherSide[t+1],
            FerrymanOnTheOtherSide[t+1],
            WolfOnThisSide[t+1].Not(),
            FerrymanOnThisSide[t+1].Not()
        ]).OnlyEnforceIf(MoveWolfAcross[t])

        # cabbage
        model.AddBoolAnd([
            CabbageOnThisSide[t],
            FerrymanOnThisSide[t],
            CabbageOnTheOtherSide[t+1],
            FerrymanOnTheOtherSide[t+1],
            CabbageOnThisSide[t+1].Not(),
            FerrymanOnThisSide[t+1].Not()
        ]).OnlyEnforceIf(MoveCabbageAcross[t])        

        # cabbage
        model.AddBoolAnd([
            SheepOnThisSide[t],
            FerrymanOnThisSide[t],
            SheepOnTheOtherSide[t+1],
            FerrymanOnTheOtherSide[t+1],
            SheepOnThisSide[t+1].Not(),
            FerrymanOnThisSide[t+1].Not()
        ]).OnlyEnforceIf(MoveSheepAcross[t]) 

        # wolf opposite side - back
        model.AddBoolAnd([
            WolfOnTheOtherSide[t],
            FerrymanOnTheOtherSide[t],
            WolfOnThisSide[t+1],
            FerrymanOnThisSide[t+1],
            WolfOnTheOtherSide[t+1].Not(),
            FerrymanOnTheOtherSide[t+1].Not()
        ]).OnlyEnforceIf(MoveWolfBack[t])   

        # cabbage opposite side - back
        model.AddBoolAnd([
            CabbageOnTheOtherSide[t],
            FerrymanOnTheOtherSide[t],
            CabbageOnThisSide[t+1],
            FerrymanOnThisSide[t+1],
            CabbageOnTheOtherSide[t+1].Not(),
            FerrymanOnTheOtherSide[t+1].Not()
        ]).OnlyEnforceIf(MoveCabbageBack[t])   

        # sheep opposite side - back
        model.AddBoolAnd([
            SheepOnTheOtherSide[t],
            FerrymanOnTheOtherSide[t],
            SheepOnThisSide[t+1],
            FerrymanOnThisSide[t+1],
            SheepOnTheOtherSide[t+1].Not(),
            FerrymanOnTheOtherSide[t+1].Not()
        ]).OnlyEnforceIf(MoveSheepBack[t])    

        # ferryman across
        model.AddBoolAnd([
            FerrymanOnThisSide[t],
            FerrymanOnTheOtherSide[t+1],
            FerrymanOnThisSide[t+1].Not()
        ]).OnlyEnforceIf(MoveAcross[t])                           

        # ferryman back
        model.AddBoolAnd([
            FerrymanOnTheOtherSide[t],
            FerrymanOnThisSide[t+1],
            FerrymanOnTheOtherSide[t+1].Not()
        ]).OnlyEnforceIf(MoveBack[t])

    
    # definition of the axioms

    for t in range(maxT-1):

        # Move Across Axioms

        # wolf is not on this side of the river
        # unless it has been there before
        # or has moved back
        model.AddBoolOr([
            WolfOnThisSide[t+1].Not(),
            WolfOnThisSide[t],
            MoveWolfBack[t]
        ])

        # cabbage axiom
        model.AddBoolOr([
            CabbageOnThisSide[t+1].Not(),
            CabbageOnThisSide[t],
            MoveCabbageBack[t]
        ])

        # sheep axiom
        model.AddBoolOr([
            SheepOnThisSide[t+1].Not(),
            SheepOnThisSide[t],
            MoveSheepBack[t]
        ])        


    # move back axioms    
    for t in range(maxT-1):
            
        # wolf
        model.AddBoolOr([
            WolfOnTheOtherSide[t+1].Not(),
            WolfOnTheOtherSide[t],
            MoveWolfAcross[t]
        ])

        # cabbage
        model.AddBoolOr([
            CabbageOnTheOtherSide[t+1].Not(),
            CabbageOnTheOtherSide[t],            
            MoveCabbageAcross[t]
        ])       

        # sheep 
        model.AddBoolOr([
            SheepOnTheOtherSide[t+1].Not(),
            SheepOnTheOtherSide[t],            
            MoveSheepAcross[t]
        ])

    # ferryman axioms
    for t in range(maxT-1):

        model.AddBoolOr([
            FerrymanOnThisSide[t+1].Not(),
            FerrymanOnThisSide[t],
            MoveWolfBack[t],
            MoveCabbageBack[t],
            MoveSheepBack[t],
            MoveBack[t]
        ])

        model.AddBoolOr([
            FerrymanOnTheOtherSide[t+1].Not(),
            FerrymanOnTheOtherSide[t],
            MoveWolfAcross[t],
            MoveWolfAcross[t],
            MoveSheepAcross[t],
            MoveAcross[t]
        ])        

    # exclusion axioms
    
    # 8 predicates
    # Formula: 8! / 2! 6!

    for t in range(maxT-1):

        # wolf and cabbage cant move simmountaneously
        model.AddBoolOr([
            MoveWolfAcross[t].Not(),
            MoveCabbageAcross[t].Not()
        ])      

        model.AddBoolOr([
            MoveWolfAcross[t].Not(),
            MoveSheepAcross[t].Not()
        ])               

        model.AddBoolOr([
            MoveWolfAcross[t].Not(),
            MoveAcross[t].Not()
        ])    

        model.AddBoolOr([
            MoveWolfAcross[t].Not(),
            MoveWolfBack[t].Not()
        ])                

        model.AddBoolOr([
            MoveWolfAcross[t].Not(),
            MoveCabbageBack[t].Not()
        ])              

        model.AddBoolOr([
            MoveWolfAcross[t].Not(),
            MoveSheepBack[t].Not()
        ])

        model.AddBoolOr([
            MoveWolfAcross[t].Not(),
            MoveBack[t].Not()
        ])
        
        # Cabbage Across Exclution Axios. 7 in total
        model.AddBoolOr([
            MoveCabbageAcross[t].Not(),
            MoveWolfAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveCabbageAcross[t].Not(),
            MoveSheepAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveCabbageAcross[t].Not(),
            MoveAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveCabbageAcross[t].Not(),
            MoveWolfBack[t].Not()
        ])

        model.AddBoolOr([
            MoveCabbageAcross[t].Not(),
            MoveCabbageBack[t].Not()
        ])

        model.AddBoolOr([
            MoveCabbageAcross[t].Not(),
            MoveSheepBack[t].Not()
        ])

        model.AddBoolOr([
            MoveCabbageAcross[t].Not(),
            MoveBack[t].Not()
        ])

        # exclusion sheep across. 7 in total
        model.AddBoolOr([
            MoveSheepAcross[t].Not(),
            MoveWolfAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveSheepAcross[t].Not(),
            MoveCabbageAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveSheepAcross[t].Not(),
            MoveCabbageAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveSheepAcross[t].Not(),
            MoveWolfBack[t].Not()
        ])

        model.AddBoolOr([
            MoveSheepAcross[t].Not(),
            MoveCabbageBack[t].Not()
        ])

        model.AddBoolOr([
            MoveSheepAcross[t].Not(),
            MoveSheepBack[t].Not()
        ])
        
        model.AddBoolOr([
            MoveSheepAcross[t].Not(),
            MoveBack[t].Not()
        ])

        # move across exclusion. 7 in total
        model.AddBoolOr([
            MoveAcross[t].Not(),
            MoveWolfAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveAcross[t].Not(),
            MoveCabbageAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveAcross[t].Not(),
            MoveSheepAcross[t].Not()
        ])

        model.AddBoolOr([
            MoveAcross[t].Not(),
            MoveWolfBack[t].Not()
        ])

        model.AddBoolOr([
            MoveAcross[t].Not(),
            MoveCabbageBack[t].Not()
        ])

        model.AddBoolOr([
            MoveAcross[t].Not(),
            MoveSheepBack[t].Not()
        ])  

        model.AddBoolOr([
            MoveAcross[t].Not(),
            MoveBack[t].Not()
        ])                    

    # final constraints
    # 1. wolf and sheep cant be on this side alone, unless ferry is on this side
    # 2. sheep and cabbage cant be on this side alone, unless ferry is on this side
    # 3. wolf and sheep cant be in opposite side alone, unless ferry is in opposite side
    # 4. sheep and cabbage cant be in opposite side alone, unless ferry is in opposite side    
    for t in range(maxT):
        
        model.AddBoolOr([
            WolfOnThisSide[t].Not(),SheepOnThisSide[t].Not()
        ]).OnlyEnforceIf(FerrymanOnThisSide[t].Not())                                                                                                                                                                    

        model.AddBoolOr([
            WolfOnTheOtherSide[t].Not(),SheepOnTheOtherSide[t].Not()
        ]).OnlyEnforceIf(FerrymanOnTheOtherSide[t].Not())

        model.AddBoolOr([
            SheepOnThisSide[t].Not(),CabbageOnThisSide[t].Not()
        ]).OnlyEnforceIf(FerrymanOnThisSide[t].Not())

        model.AddBoolOr([
            SheepOnTheOtherSide[t].Not(),CabbageOnTheOtherSide[t].Not()
        ]).OnlyEnforceIf(FerrymanOnTheOtherSide[t].Not())                


    solver = cp_model.CpSolver()

    variables = []
    for t in range(maxT-1):
        variables.append(
            {
                "move accross":MoveAcross[t],
                "move back":MoveBack[t],
                "move wolf accross":MoveWolfAcross[t],
                "move wolf back":MoveWolfBack[t],
                "move sheep accross":MoveSheepAcross[t],
                "move sheep back": MoveSheepBack[t],
                "move cabbage accross": MoveCabbageAcross[t],
                "move cabbage back": MoveCabbageBack[t]    
            })

    solver.SearchForAllSolutions(model, SolutionPrinter(variables))
    
    print(WolfOnThisSide)
    for t in range(maxT-1):

        if solver.Value(MoveWolfAcross[t]): print("move wolf across")
        if solver.Value(MoveCabbageAcross[t]): print("move cabbage across")
        if solver.Value(MoveSheepAcross[t]): print("move sheep across")
        if solver.Value(MoveAcross[t]): print("move across")
        if solver.Value(MoveWolfBack[t]): print("move wolf back")
        if solver.Value(MoveCabbageBack[t]): print("move cabbage back")
        if solver.Value(MoveSheepBack[t]): print("move sheep back")
        if solver.Value(MoveBack[t]): print("move back")                                                                                    

   

class SolutionPrinter(cp_model.CpSolverSolutionCallback):
    def __init__(self, variables):
        cp_model.CpSolverSolutionCallback.__init__(self)        
        self.variables_ = variables        
        self.solutions_ = 0            
    
    def OnSolutionCallback(self):        
        self.solutions_ = self.solutions_ + 1        
        print("solution", self.solutions_ )        
        i=0        
        for vars_in_timestep in self.variables_:            
            i=i+1
            #            print(" - Timestep: ", i)            
            for op in vars_in_timestep:                
                if self.Value(vars_in_timestep[op]):                    
                    print("   ", op)        
        print()



if __name__ == "__main__":
    main()
