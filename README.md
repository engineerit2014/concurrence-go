# Go advanced: concurrence-go

This project shows the use of concurrence in Golang. For this purpose, some common problems and their solutions using go are shown.

## Table of Contents

- [Concurrence problems](#concurrence-problems)
    - [Deposit & Withdraw](#deposit-&-withdraw)

## Concurrence problems

### Deposit & Withdraw

Deposit and withdrawal problem, commonly occurs in the management of transactions when you have N people depositing and N people consulting or withdrawing, if not handled correctly the balance generated in an account can be affected in any way. To solve this problem called race condition, we can use Go and its concurrency management tools. 

See the examples for more details.