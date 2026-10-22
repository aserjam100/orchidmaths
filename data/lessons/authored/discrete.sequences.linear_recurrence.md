# Solving linear recurrences

A **linear homogeneous recurrence** expresses each term as a fixed linear combination of earlier terms, such as \(a_n = 5a_{n-1} - 6a_{n-2}\). Its solutions are built from the roots of a polynomial, and the initial conditions fix the coefficients.

## The characteristic equation

Substituting the guess \(a_n = r^n\) gives \(r^n = 5r^{n-1} - 6r^{n-2}\), and dividing by \(r^{n-2}\) leaves the **characteristic equation** \(r^2 - 5r + 6 = 0\). Its roots are 2 and 3, so both \(2^n\) and \(3^n\) solve the recurrence.

Because the recurrence is linear and homogeneous, any combination \(A \cdot 2^n + B \cdot 3^n\) also solves it. That is the general solution when the roots are distinct; a repeated root introduces a factor of \(n\).

## Initial conditions

The coefficients come from the first two terms. With \(a_0 = 2\) and \(a_1 = 5\), the equations are \(A + B = 2\) and \(2A + 3B = 5\). Subtracting twice the first from the second gives \(B = 1\), so \(A = 1\), and \(a_n = 2^n + 3^n\).

The check is a substitution rather than a rearrangement: \(a_0 = 1 + 1 = 2\), \(a_1 = 2 + 3 = 5\), and \(5a_{n-1} - 6a_{n-2} = 5(2^{n-1} + 3^{n-1}) - 6(2^{n-2} + 3^{n-2})\), which simplifies to \(2^n + 3^n\).

## Checking the closed form

A closed form can be written in more than one way, and equivalent expressions are equally correct: \(3^n + 2^n\) is the same function as \(2^n + 3^n\). What is not equivalent is a wrong combination, such as \(2 \cdot 3^n\), which satisfies the recurrence but not the initial conditions.

The general solution always has as many free coefficients as the order of the recurrence, and the initial conditions are what pin them down. A formula with the right roots and the wrong coefficients is the most common error, and substituting \(n = 0\) and \(n = 1\) catches it.
