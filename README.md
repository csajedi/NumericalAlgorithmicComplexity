
# Overview

The goal of this repository is to provide numerical tools to evaluate the algorithmic complexity of short strings. This effort reproduces in part the [work](https://arxiv.org/pdf/1101.4795.pdf) 

"Numerical Evaluation of Algorithmic Complexity for
Short Strings: A Glance Into the Innermost Structure of
Randomness"

by Jean-Paul Delahaye and Hector Zenil.

The basic computational task involves computing the entire space of a given n-state turing machine and retaining the resulting output tape. This tape is then analyzed for patterns in substring reproduction. This is defined as the empirical universal distribution. This effort cannot take full advantage of the existing busy beaver problem capabilites due to the requirement to retain output strings, further discussion of optimization strategies can be found [here](https://journals.plos.org/plosone/article/file?id=10.1371/journal.pone.0096223&type=printable)

## Concepts

### Formalism of the Turing Machine

## Methodology

For this experiment we are mainly interested in comparing implementation approaches and capturing the essence of the task. Details provided for exploiting symmetry to reduce the search space are left as a future exercise. And we provide a cutoff for machines at 500 steps to avoid non-halting machines and reproduce previous results.