Optimization benchmarks
=======================

Benchmark problems for evaluating non-linear optimization
techniques. Most of these problems are multimodal and non-separable
and many are non-differentiable. Since these are intended to evaluate
optimization techniques that would apply to noisy functions and
measurements, each problem includes a noisy version (where repeated
evaluation of f(x*) produces different results).

The problems are taken from [1], which in turn modified the benchmark
problems used by [2]. I have scaled them so that all problems are
attempting to minimize an objective function defined over [-1, 1]^d;
in most but not all cases the minimum value of f(x*) = 0.

[1] Ortiz-Boyer, Domingo, César Hervás-Martınez, and Nicolás
Garcıa-Pedrajas. "Cixl2: A crossover operator for evolutionary
algorithms based on population features." Journal of Artificial
Intelligence Research 24.1 (2005): 1-48.

[2] Eiben, Agoston E., and Thomas Bäck. "Empirical investigation of
multiparent recombination operators in evolution strategies."
Evolutionary Computation 5.3 (1997): 347-365.
