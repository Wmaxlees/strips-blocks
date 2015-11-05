# strips-blocks
STRIPS method for the block problem

## Installation Instructions (for Windows)

### Install Go
1. Visit https://golang.org/doc/install and click the 'Download Go' link
2. Open the MSI file and follow the propts to install the Go tools
3. Restart any running command prompts
4. Choose a location where all go projects will live (e.g. C:/Users/user/goworkspace/)
5. Go to Control Panel->System->Advanced->Environment Variables and set GOPATH to your selected directory

### Installing and Running Strips-Blocks
1. Type `go install github.com/Wmaxlees/strips-blocks` to download the source into your go workspace and compile the source into an executable in your go bin folder
2. Type `$GOPATH/bin/strips-blocks' to run the newly compiled executable

## Command Line Options

The program will run just fine without any command line arguments but for full functionality, use the following:

-verbose: Tells the program to show the Strips Stack as it determines which blocks to move

-problem <int>: Tells the program which blocks problem to attempt to solve.
    1 : Simple Blocks Problem
    2 : Sussman Anomaly
    3 : Extra Problem
