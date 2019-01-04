-module(accumulate).

-export([accumulate/2, test_version/0]).

accumulate(Fn, Ls) -> accumulate(Fn, Ls, []).

%% Loop through the input list, calling Fn for each item
accumulate(Fn, [Head | Rest], ResultLs) ->
    accumulate(Fn, Rest, [Fn(Head) | ResultLs]);

%% We've processed everything!
accumulate(_, [], ResultLs) ->
    %% We could've used ResultLs ++ Fn(Head) in the above accumulate/3
    %% definition, which I think would be being "head recursive"? But the
    %% Erlang docs say not to do that because its inefficient, as it'll be
    %% copying the accumulated list on every iteration. Instead we can keep
    %% it tail recursive (?) and reverse the accumulated list once the input
    %% has been fully processed.
    lists:reverse(ResultLs).

test_version() -> 1.
