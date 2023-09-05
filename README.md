# The problem

We have a job that calculates the total earnings per store from collection of orders. 			
This job needs to do the caluclations councurently assuning that it's a heavy task to process each order, we can change number of the go routines and gather the results at the end. 