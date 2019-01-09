-module(as2).
-export([divisors/1,primes/1,join/2,hailseq/1,mergesort/1,fib/1,pythagorean/1,main/0]).


-spec divisors(integer()) -> [integer()].

divisors(X) ->
    [ I || I <- lists:seq(2,X div 2), X rem I == 0].
    
-spec primes(integer()) -> [integer()].
primes(X) ->
    [ I || I <- lists:seq(2,X), divisors(I) == []].
    

-spec join(string(), [string()]) -> string().
join(_,Y) when length(Y) == 0 -> "";
join(_,Y) when length(Y) == 1 -> lists:nth(1,Y);
join(X,Y) when length(Y) > 1 ->
    join(X,[lists:nth(1,Y) ++ X ++ lists:nth(2,Y)] ++ lists:sublist(Y,3,length(Y)-2)).
    
    
pythagorean(X) ->
    [{I,J,K}|| I <- lists:seq(1,X),J <- lists:seq(1,X),K <- lists:seq(5,X),math:pow(I,2)+math:pow(J,2)==math:pow(K,2),I=<J].
    

-spec hailstone(integer()) -> integer().
hailstone(X) when X rem 2 == 0 ->
    X div 2;
hailstone(X) when X rem 2 == 1 ->
    3 * X + 1.




-spec hailnext([integer()], integer()) -> [integer()].
hailnext(X,1) ->
    X ++ [1];
hailnext(X,Y) ->
    hailnext(X ++ [Y], hailstone(Y)).

-spec hailseq(integer()) -> [integer()].
hailseq(X) -> 
    hailnext([],X).
    

merge(C,A,B) when (A == []) and (B == []) -> C;
merge(C,[L|A],B) when B == [] ->
    merge(C ++ [L], A, B);
merge(C,A,[R|B]) when A == [] ->
    merge(C ++ [R], A, B);
merge(C,[L|A],[R|B]) ->
    if
        L =< R  ->
            merge(C ++ [L], A, [R] ++ B);
        true    ->
            merge(C ++ [R], [L] ++ A, B)
    end.
    
mergesort(X) when length(X) == 2 ->
    merge([],[lists:nth(1,X)],[lists:nth(2,X)]);
mergesort(X) when length(X) > 2 ->
    Part = lists:split(length(X) div 2,X),
    merge([],mergesort(element(1,Part)), mergesort(element(2,Part)));
mergesort(X) -> X.
    


fb(X,Y,C) ->
    if
        C == 0 ->
            X;
        true ->
            fb(Y, X+Y, C-1)
    end.

fib(X) -> fb(0,1,X).


main() ->
    {divisors(30),
    primes(7),
    join(", ", ["one","two","three"]),
    pythagorean(30),
    hailseq(6),
    mergesort([1,9,3,2,7,6,4,8,5]),
    fib(38)}.
    