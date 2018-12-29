-module(difference_of_squares).

-export([sum_of_squares/1, square_of_sum/1, difference_of_squares/1, test_version/0]).

sum_of_squares(N) ->
  sum_of_squares(0, N, 0).

sum_of_squares(A, N, Running_total) when A =< N ->
  sum_of_squares(A + 1, N, Running_total + (A*A));

sum_of_squares(_, _, Running_total) ->
  Running_total.

square_of_sum(N) ->
  Sum = sum_of_naturals(0, N, 0),
  Sum * Sum.

sum_of_naturals(A, N, Running_total) when A =< N ->
  sum_of_naturals(A + 1, N, Running_total + A);

sum_of_naturals(_, _, Running_total) ->
  Running_total.

difference_of_squares(N) ->
  square_of_sum(N) - sum_of_squares(N).

test_version() -> 1.
