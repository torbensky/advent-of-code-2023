#!/bin/sh

runall() {
    echo "Typescript"
    npx ts-node solution.ts ./input.txt
    echo "Go"
    go run main.go input.txt
    echo "Python"
    python3 ./solution.py ./input.txt
}

runall
runall
runall