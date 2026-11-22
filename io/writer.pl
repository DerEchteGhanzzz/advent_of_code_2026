:- module(writer, [
    print_list/1
]).

print_list([]):-
	nl.
print_list([X|L]):-
	write(X),
	nl,
	print_list(L).
