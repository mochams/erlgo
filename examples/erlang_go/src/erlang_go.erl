-module(erlang_go).

-export([handle_call/0]).

-define(SERVER, ?MODULE).

handle_call() ->
    PrivDir = code:priv_dir(erlang_go),
    PortPath = filename:join([PrivDir, "go", "hello_world"]),
    Port = open_port({spawn, PortPath}, [
        {packet, 4}, binary, exit_status, use_stdio
    ]),
    Request = json:encode(#{action => <<"foobar">>, arguments => [<<"foo">>, <<"bar">>]}),
    Port ! {self(), {command, Request}},
    loop(Port).

loop(Port) ->
    receive
        {Port, {data, Response}} ->
            io:format("Received response from Go: ~p~n", [Response]),
            io:format("Format response: ~p~n", [json:decode(Response)]),
            loop(Port);
        {Port, {exit_status, Status}} ->
            io:format("Go process exited with status: ~p~n", [Status]),
            close_port(Port)
    after 5000 ->
        io:format("Timeout while waiting for response~n"),
        close_port(Port),
        {error, timeout}
    end.

close_port(Port) ->
    Port ! {self(), close}.
