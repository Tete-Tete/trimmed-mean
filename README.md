# Trimmed Mean Usage

This repository contains an example Go program that imports the `trimmedmean` package and uses its `TrimmedMean` function to compute trimmed means of sample data.


## Requirements
- Go programming language installed (version 1.16 or higher recommended).

## Installation From Git and Set up for your own computer
### Step 1: Clone the Repository
Clone this repository to your local machine:
```sh
git clone <https://github.com/Tete-Tete/trimmed-mean.git>
```

### Step 2: Get the trimmed mean package from GitHub:

```sh
go get github.com/Tete-Tete/trimmedmean
```

### Step 2: Run the Application
To run the Go application, run the following command in your terminal:
```sh
go build -o trimmed.exe
```
This will create a file named `trimmed` in your current directory.


## Summary of Results
The program calculates the trimmed mean using a trimming proportion of 0.05 from each end of the data. The values were compared to results from R's `mean` function with trimming, and were found to be in agreement within acceptable limits of numerical precision.

## Example Comparison with R

Using the same dataset in R, the trimmed mean can be computed with:
```sh
mean(x, trim = 0.05)
```
The Go results and R results should closely match, providing validation for the Go package implementation.


