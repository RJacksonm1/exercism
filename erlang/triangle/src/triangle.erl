-module(triangle).

-export([kind/3, test_version/0]).

kind(A, B, C) when A =< 0 orelse B =< 0 orelse C =< 0 ->
    {error, "all side lengths must be positive"};
  
kind(A, B, C) when A + B =< C ->
    {error, "side lengths violate triangle inequality"};
  
kind(A, B, C) when A + C =< B ->
    {error, "side lengths violate triangle inequality"};
  
kind(A, B, C) when B + C =< A ->
    {error, "side lengths violate triangle inequality"};
  
kind(A, B, C) when A == B andalso B == C -> equilateral;
kind(A, B, _C) when A == B -> isosceles;
kind(A, _B, C) when A == C -> isosceles;
kind(_A, B, C) when B == C -> isosceles;
kind(_A, _B, _C) -> scalene.

test_version() -> 1.
