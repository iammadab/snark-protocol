# snark-protocol

This is a basic implementation of the snark protocol describe in this [paper](https://arxiv.org/pdf/1906.07221.pdf).

In the paper, the protocol is described in stages, I follow the same format for my implementation.  

## Stages
### Hiding the evaluation point (homomorphic encryption)
In order to prevent the prover from generating fake proofs, the verifier first encrypts the evaluation point and sends those encrypted powers to the prover.  
No restriction is put in place to force the prover to use only those encrypted points. Hence, the protocol can still be broken.  

You can check out to this state by running:
```shell
    git checkout homomorphic-encryption
```

Then test the function that breaks the protocol with:
```shell
    go test --run TestBreakHE
```

(other stages) TBD