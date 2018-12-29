-module(grains).

-export([square/1, total/0]).

%% How many grains are on Square N on a chessboard, where each square
%% contains double the previous square's amount of grains.
%% We could compute this recursively, but this happens to be power of
%% X series. We _could_ use bitshifting here, but let's keep it simple
%% for future devs to grok and use math:pow instead

square(N) when N < 1; N > 64 ->
    {error, "square must be between 1 and 64"};
square(N) -> trunc(math:pow(2, N - 1)).

%% As each tiles doubles the previous, the sum of all previous values
%% would be 1 less than the next (65th) theoretical tile
total() -> trunc(math:pow(2, 64)) - 1.
