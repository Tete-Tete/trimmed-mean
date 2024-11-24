Trimmed Mean Usage

This repository contains an example Go program that imports the "trimmedmean" package and uses its TrimmedMean function to compute trimmed means of sample data.

Getting Started

Prerequisites

Make sure you have Go installed on your machine. You can download and install it from golang.org.

Installation

Clone this repository:

git clone https://github.com/yourusername/trimmedmean-usage.git
cd trimmedmean-usage

Get the trimmed mean package from GitHub:

go get github.com/yourusername/trimmedmean

Running the Program

To run the program, use the following command:

go run main.go

Building an Executable

To create an executable file for this application, use the following command:

go build -o trimmedmean-usage.exe

If you are on Mac or Linux and want to create a Windows .exe file, use the following command:

GOOS=windows GOARCH=amd64 go build -o trimmedmean-usage.exe

After building, you will have an executable file named trimmedmean-usage.exe that you can run:

./trimmedmean-usage.exe

The program will display the trimmed means of sample data and wait for you to press Enter before closing.

Summary of Results

The program calculates the trimmed mean using a trimming proportion of 0.05 from each end of the data. The values were compared to results from R's mean function with trimming, and were found to be in agreement within acceptable limits of numerical precision.

Example Comparison with R

Using the same dataset in R, the trimmed mean can be computed with:

mean(x, trim = 0.05)

The Go results and R results should closely match, providing validation for the Go package implementation.


