package main

import (
    "fmt"
)

type Income interface {
    calculate() string
}

type google struct {
    workingYear, baseSalary, performance int
}

type yahoo struct {
    performance, workingYear, baseSalary int
}

func (g google) calculate() int {
    return g.baseSalary + (1000 * g.performance)
}

func (y yahoo) calculate() int {
    return y.baseSalary + (20000 * y.workingYear)
}

func main() {

    hanako := google{
        workingYear: 5,
        baseSalary: 100000,
        performance: 190,
    }
    choko := yahoo{
        baseSalary: 60000,
        workingYear: 50,
    }
    motoko := yahoo{
        baseSalary: 40000,
        workingYear: 25,
    }

    workers := []Income{hanako, choko, motoko}
    fmt.Printf("Total income: %d\n", calculateIncome(workers))
}

func calculateIncome(ic []Income) int {
    sum := 0
    for _, worker := range ic{
        sum += worker.calculate()
    }

    return sum
}