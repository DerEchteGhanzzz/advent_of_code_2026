:- module(reader, [
    read_file/2
]).

read_file(FileName, Lines):-
	open(FileName, read, Str),
	read_lines(Str, Lines),
	close(Str).

read_lines(Stream, []):-
	at_end_of_stream(Stream).

read_lines(Stream, [Line|Rest]):-
	 \+ at_end_of_stream(Stream),
	read_line(Stream, Line),
	read_lines(Stream, Rest).

read_line(InStream, Chars) :-
    get_code(InStream, Char),
    check_char_and_read_rest(Char, Chars, InStream).

% Newline
check_char_and_read_rest(10, [], _):-
	!.

% End of Stream
check_char_and_read_rest(-1, [], _):-
	!.

% End of File
check_char_and_read_rest(end_of_file, [], _):-
	!.

% Ignore \r
check_char_and_read_rest(13, Chars, InStream):-
	!,
	get_code(InStream, NextChar),
	check_char_and_read_rest(NextChar, Chars, InStream).

check_char_and_read_rest(Char, [Char|Chars], InStream):-
	get_code(InStream, NextChar),
	check_char_and_read_rest(NextChar, Chars, InStream).