-module(isogram).

-export([is_isogram/1]).

is_isogram(Phrase) -> is_isogram(string:lowercase(Phrase), []).

%% Dashes and spaces are allowed to repeat
is_isogram([First | Phrase], SeenChars) when [First] == "-"; [First] == " " ->
    is_isogram(Phrase, [First | SeenChars]);

%% Otherwise, if we've seen a character before return false
is_isogram([First | Phrase], SeenChars) ->
    case lists:member(First, SeenChars) of
        true -> false;
        false -> is_isogram(Phrase, [First | SeenChars])
    end;

%% And if we get through the entire phrase without returning false, we have an isogram!
is_isogram([], _SeenChars) ->
    true.
