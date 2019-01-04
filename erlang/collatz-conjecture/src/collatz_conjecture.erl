-module(collatz_conjecture).

-export([steps/1]).

steps(N) -> steps(N, 0).

steps(N, _) when N =< 0 ->
    {error, "Only positive numbers are allowed"};
steps(1, Iterations) -> Iterations;
steps(N, Iterations) when N rem 2 == 0 ->
    steps(round(N / 2), Iterations + 1);
steps(N, Iterations) when N rem 2 == 1 ->
    steps(3 * N + 1, Iterations + 1);
steps(N, Iterations) ->
    io:format("wtf ~w ~w ~n", [N, Iterations]).
