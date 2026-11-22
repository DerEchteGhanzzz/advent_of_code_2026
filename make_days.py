import os

content = """:- use_module('../io/reader', [read_file/2]).
:- use_module('../io/writer', [print_list/1]).

lines(Lines) :-
	read_file('input.txt', RawLines),
	maplist(string_codes, Lines, RawLines),
	print_list(Lines).
"""

for i in range(1, 13):
    dir = f'day_{i}'
    os.mkdir(dir)
    with open(f'{dir}/a.pl', 'w') as f:
        f.write(content)
    with open(f'{dir}/b.pl', 'w') as f:
        f.write(content)
    with open(f'{dir}/input.txt', 'w') as f:
        f.write('')