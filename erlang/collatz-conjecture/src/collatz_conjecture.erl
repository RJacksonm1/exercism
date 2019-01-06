-module(collatz_conjecture).

-export([steps/1]).

steps(N) when is_integer(N), N > 0 -> steps(N, 0);
steps(_) ->
    {error, "Only positive numbers are allowed"}.

steps(1, Iterations) -> Iterations;
steps(N, Iterations) when N rem 2 == 0 ->
    steps(N div 2, Iterations + 1);
steps(N, Iterations) ->
    steps(3 * N + 1, Iterations + 1).
