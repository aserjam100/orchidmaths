# Solving Systems with Matrices

**Augmented matrix:** a rectangular array holding the coefficients of a linear system with the constants in an extra column, reduced by row operations to solve the system.

## Building the augmented matrix

Each equation contributes a row, with the coefficient of each variable in its own column and the constant in the rightmost augmented column. Missing variables get a zero, which keeps the columns aligned, and the vertical bar separates the coefficients from the constants. Once the system is in this form the variables themselves are absent from the work, and only the numbers are manipulated.

## Row operations

Three operations preserve the solution set: swap two rows, multiply a row by a nonzero constant, and add a multiple of one row to another. These correspond to reordering equations, scaling an equation, and adding a multiple of one equation to another. Reducing to row-echelon form with these operations turns the system into one that is solved by back-substitution, and reducing further to reduced row-echelon form reads the solution directly.

## Reading the outcome

The reduced matrix states plainly whether the system has one solution, none, or infinitely many. A pivot in every variable column gives a single solution; a row of zeros in the coefficient part with a nonzero constant states an impossibility and gives none; and a missing pivot gives a free variable and hence infinitely many solutions. The matrix reports the geometry of the system in a single glance.
