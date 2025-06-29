PE(0,0) => Once {
    START:{
        MOV, ^[South, R] -> [$0]
        ADD, [$0], IMM[1] -> [$1]
        MOV, [$1] -> ^[North, R]
    }
}

PE(1,0):
{
    Entry => Loop {
        {
            MOV, ^[South, R] -> [$0]
            ADD, [$0], IMM[1] -> [$1]
            MOV, [$1] -> ^[North, R]
        }
        {
            MOV, ^[South, R] -> [$0]
            MUL, [$0], IMM[1] -> [$1]
            MOV, [$1] -> ^[North, R]
        }
        {
            MOV, ^[South, R] -> [$0]
            DIV, [$0], IMM[1] -> [$1]
            MOV, [$1] -> ^[North, R]
        }
    } 
     Entry => Loop {
        {
            MOV, ^[South, R] -> [$0]
            ADD, [$0], IMM[1] -> [$1]
            MOV, [$1] -> ^[North, R]
        }
    } 
}