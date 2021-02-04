To quickly delete multiple lines in a file, use the global command:
:g/pattern/d

You can use :s to count occurrence of a pattern.
To count how many times "donut" occur in a file, run:
:%s/donut//gn

