-module(otherprograms).
-export([concurrent/1, fibslow/1,client/0,main/0,listminus/2]).

% This is a function to use concurrent computing to divide task into 4 subtasks 
% and let 4 child processes to run it and send result to the main process.
% Then the main process will send message to child processes to let them end.
% The concurrent is the main function which generate 4 child processes and send the tasks to them.
% Then use the main process to run server function to listen the result.

concurrent(X) ->
    Pid0 = spawn(otherprograms,client,[]),
    Pid0 ! { dofib, X, self()},
    Pid1 = spawn(otherprograms,client,[]),
    Pid1 ! { dofib, X+1, self()},
    Pid2 = spawn(otherprograms,client,[]),
    Pid2 ! { dofib, X+2, self()},
    Pid3 = spawn(otherprograms,client,[]),
    Pid3 ! { dofib, X+3, self()},
    Pids=[Pid0,Pid1,Pid2,Pid3],
    server(4,Pids).

% The server function will listen to the data sent from child processes and put them in a list.
% When the server have received all results which X is equal 0, it send finished to child processes.
server(0,Pids) -> 
    [Pid0,Pid1,Pid2,Pid3] = Pids,
    Pid0 ! finished,
    Pid1 ! finished,
    Pid2 ! finished,
    Pid3 ! finished,
    [];    
server(X,Pids) ->
    receive
        {fibdone, Result} ->
            server(X-1,Pids) ++ [Result]
            % Using recursion to listen 4 results from 4 child processes.
    
    end.
    
% It is the client processes to receive task and do the task and send the result by Pid ! message.
client() ->
    receive
        finished -> done;
        { dofib , FibNumber, Pid} ->
                Result = fibslow(FibNumber),
                Pid ! {fibdone, Result},
                client()
                % Using recursion to keep this process to wait for finished message to end process.
    end.



% Use the slow fib function to test.
fibslow(0) -> 0;
fibslow(1) -> 1;
fibslow(X) ->
    fibslow(X-1) + fibslow(X-2).
    
% Show the list deletion that Haskell do not have.
listminus(X,Y) ->
    X -- Y.

% Use the main function to run the program.
main() ->
    {concurrent(36),
    listminus([3,4,5,6],[3,6])}.