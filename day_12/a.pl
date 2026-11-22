:- use_module('../io/reader', [read_file/2]).
:- use_module('../io/writer', [print_list/1]).

lines(Lines) :-
	read_file('input.txt', RawLines),
	maplist(string_codes, Lines, RawLines),
	print_list(Lines).
