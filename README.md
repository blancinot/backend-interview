# backend-interview
This repository contains instructions and draft project for powder backend interview

## requirements

- go
- docker
- docker-compose
- make
- psql

## how to setup

- Setup local postgres db
```sh
> docker-compose up -d
```

- Create db schema + Populate with fake data
```sh
> make migrate
> make populate
```

- Build API binary and run locally
```sh
make api && ./bin/interview_api config/api/local.json
```

## instructions

You inherited this project and there are issues you need to fix.

Usecase: `user` has 1+ `account` and each `account` register its own `transactions`.

User must be able to insert new transactions and get his current total.

For this test, we don't care about security flaws or float64 approximations.

You can fork this project and add your code/answers into it.

## project architecture

```
cmd_|_ # main
    |
    |__api # binary
    |
pkg_|_ # domain
    |
    |__user # domain name
      |
      |__sql # sql STORE implementation
      |__app # APP implementation, domain logic + store logic
      |__dto # DataTransferObject for external objects
```

### 0.0

- `GetUser` API route always returns a total = 0. Fix it to return approximated sum (due to float64). Hint: `account.FetchManyAccount` SQL method already exists in draft code (`handler.account` store).

- Write a new API route `CreateTransaction` to add a new transaction. Hint: `account.InsertTransaction` SQL method already exists in draft code (`handler.account` store).

- Add some minimal tests on those routes to ensure at least 1 success path. (any kind of test is ok)

### 0.1

Questions (text only):

- **You are running this service in production under ~100 req/s, what are your main concerns about scaling and stability ?**

>
```
First we must find reasons in our software architecture and in our code which introduces our latency and our inadequacy to manage an high traffic.

Breaking up the code into dedicated microservices will be a good practice. Now, considering the very low complexity of the code, it is overkill but a lot feature will certainly arrive.
I would recommend using Domain-Driven-Development to parse our microservices by business domain.
This will then allow you to accurately scale each microservices, each domain.
If the high request traffic is based mainly on transaction management for example, it might be useful to scale up just that particular microservice.

Each of these microservices must be autonomous and independent to scale without concurrent data writings risk. Lock mechanisms or write access isolation can be expected.
Going even further, we could even imagine auto-scalling according to the network traffic or time or other defined smart thing.
A load-balancer according to a predefined logic is also expected.

The database can also quickly become a bottleneck.
Several solutions wich can be benchmarked and combined:
- Scale the database (master/slave)
- Improve database settings
- Use better scalable distributed SQL database like [cockroach] (https://github.com/cockroachdb/cockroach).
- Use caching method in addition to your database or change postgres for another faster (read) database like nosql or k/v store.


We can hide high traffic problems by making asynchronous processes behind the requests. This idea doesn't solve the problem but it brings a better user experience. The lag will always be there but the user will not see timeouts. However, error handling becomes much more complicated. It's a design choice to do.


Of course, you have to add monitoring to control and be the most responsive when high traffic occurs.
With monitoring, We can anticipate or at least escalate the problem to solve it quickly.
```


- **We want to get rid of `account` intermediary table and attach directly transactions to `user`. Write up a database migration plan (+ add some example queries).**

>
- Migration plan:
```sql
ALTER table users ADD COLUMN account_total float DEFAULT 0.0;


UPDATE
    users u
SET
    account_total = account_total + a.total
FROM
    account a
WHERE
    a.user_id = u.id;


ALTER table transaction ADD COLUMN user_id VARCHAR(264);


UPDATE
    transaction t
SET
    user_id = a.user_id
FROM
    account a
WHERE
    t.account_id = a.id;


ALTER TABLE transaction
    DROP COLUMN account_id;


DROP TABLE account;
```


### 0.2 (bonus)

- (bonus) Add a new rule where a transaction is not accepted if there is not enough money on account.

- (bonus) User now wants to know his largest expense (`transaction`) between 2 dates. Create a new route `GetMaxTransaction` which takes 2 timestamps in parameters.

- ~~(bonus) Create and implement a mock on a store or app (of your choice). Using this mock, write up a benchmark comparing a route (of your choice) with and without mocks.~~
~~Mock implementation will be check in the code part.~~
~~For the benchmark part you need to provide a code part (no verification on clean/maintenance for this one, it's benchmark code) + benchmark results (in easy to read format plz)~~
Not enough time to do this one
