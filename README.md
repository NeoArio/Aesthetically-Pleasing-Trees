# Aesthetically-Pleasing-Trees

## Introduction
This piece of code is the solution for the aestheically pleassing garden problem.

## How to run
Please use the makefile to run and test the code:
* make build
* make run input="[3,4,5,3,7]" --> You can put your input
* make test

## Solutions

### First Solution(order n)
We can make the problem easier if we consider some patterns. The basic pattern 
 that makes our garden not pleasing is 3 ascending/descending trees. we don't want 
any of them. So:
* if we find 4 ascending or descending trees, there is no way to cut 1 tree out and 
make the garden pleasing. Because still we have at least 3 ascending or descending trees in 
the garden
* if we have more than two ascending/descending trees with the lenght of 3, still we can not make the
garden pleasing because one of them will be left in the place.
* So the only case left is that we have exactly 1 ascending/descending trees with
the length of 3.

This is what I did:
* Iterate through the input using a for loop.
* compare the value of each element with next 2 elements. if these 3 elements 
are making a 3 ascending/descending, save the index of the first element.
* if the numbers saved are more than 1, there is no way.
* if the numbers saved are 0, the garden is already pleasing.
* if only one number is saved, we can make garden pleasing by cutting trees
in 1 or 2 or 3 ways.
### Second Solution(order n^2)
At first, I try to find subset of (n-1) size. Then for each subset check if it is aesthetically pleasing
or not. Number of aesthetically pleasing subsets is the answer.