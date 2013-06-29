// Implement a genetic algorithm with D genes that range over [0, 1).
//
// Mutation is a clipped Gaussian that decreases in magnitude when the
// population starts to stagnate.
//
// The population always includes N% of the best-ever individuals and
// M% new individuals.

package optbench

import (
	"math/rand"
	"math"
	"sort"
	"fmt"
)

type Individual struct {
	Fitness float64
	Genes []float64
}

func NewIndividual(d int) *Individual {
	ind := new(Individual)
	ind.Genes = make([]float64, d)
	for i := range ind.Genes {
		ind.Genes[i] = rand.Float64()
	}
	return ind
}

func Breed(p1, p2 *Individual) *Individual {
	g := len(p1.Genes)
	child := new(Individual)
	child.Genes = make([]float64, g)
	// Single crossover (for now).
	cx := rand.Intn(g)
	for i := range child.Genes {
		if i < cx {
			child.Genes[i] = p1.Genes[i]
		} else {
			child.Genes[i] = p2.Genes[i]
		}
	}
	return child
}

func (gs *Individual) Mutate(stddev float64) {
	for i, g := range gs.Genes {
		g += rand.NormFloat64() * stddev
		gs.Genes[i] = math.Min(1.0, math.Max(0.0, g))
	}
}

type Population struct {
	Members []*Individual
	Fittest float64
	sigma float64
	age int
}

func NewPopulation(d, k int) *Population {
	pop := new(Population)
	pop.Members = make([]*Individual, k)
	for i := range pop.Members {
		pop.Members[i] = NewIndividual(d)
	}
	pop.Fittest = 10000000.0
	pop.sigma = 0.1
	return pop
}

type EvalFn func ([]float64) float64

func (pop *Population) Evaluate(f EvalFn) float64 {
	pop.age += 1
	fittest := f(pop.Members[0].Genes)
	for _, m := range pop.Members {
		m.Fitness = f(m.Genes)
		if m.Fitness < fittest {
			fittest = m.Fitness
			if m.Fitness < pop.Fittest {
				pop.Fittest = m.Fitness
				pop.age = 0
			}
		}
	}
	delta := 1 - fittest
	total := 0.0
	for _, m := range pop.Members {
		m.Fitness = 1 / (m.Fitness + delta)
		total += m.Fitness
	}
	for _, m := range pop.Members {
		m.Fitness /= total
	}
	sort.Sort(pop)
	return fittest
}

func (p *Population) Len() int {
	return len(p.Members)
}

func (p *Population) Swap(i, j int) {
	p.Members[i], p.Members[j] = p.Members[j], p.Members[i]
}

func (p *Population) Less(i, j int) bool {
	return p.Members[i].Fitness < p.Members[j].Fitness
}

func (p *Population) PrintOut() {
	for _, m := range p.Members {
		fmt.Println(m)
	}
}

type CumulativeDistribution []float64

func (p *Population) FitnessCurve() CumulativeDistribution {
	cum := 0.0
	var curve CumulativeDistribution
	for _, m := range p.Members {
		cum += m.Fitness
		curve = append(curve, cum)
	}
	return curve
}

func (c CumulativeDistribution) RouletteSelection() int {
	return sort.SearchFloat64s(c, rand.Float64())
}

func Epoch(pop *Population) {
	ms := pop.Members
	bred := int(0.6 * float64(len(ms)))
	novel := int(0.9 * float64(len(ms)))
	curve := pop.FitnessCurve()
	var children Population
	if (pop.age > 3) {
		pop.sigma *= 0.99
		pop.age = 0
		fmt.Println("Reducing sigma to", pop.sigma)
	}
	for i := 0; i < bred; i++ {
		a, b := curve.RouletteSelection(), curve.RouletteSelection()
		child := Breed(ms[a], ms[b])
		child.Mutate(pop.sigma)
		children.Members = append(children.Members, child)
	}
	for i, c := range children.Members {
		ms[i] = c
	}
	d := len(ms[0].Genes)
	for i := bred; i < novel; i++ {
		ms[i] = NewIndividual(d)
	}
}
