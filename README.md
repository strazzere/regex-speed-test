# regex-speed-test
Simple application just to test some fundamental concepts on Golang regex structures.

## What?
While working on some code reviews I was curious how efficient regular expressions
in Go are. Then I was stuck on an airplane shortly after, with nothing else to do,
so this code was created. Specifically I had a rather large list of exact known
items I would want to filter against (IP addresses and domains). My assumption was
that both a "smart" regular expression and a "dumb" regular expression would end up
being the same after a compilation.

A "dumb" regular expression would be something like `known1|known2|known3|known4`.
This might seem silly, but would be extremely easy to automatically generate from
a [Salt](https://github.com/saltstack/salt) file.

The "smart" regular expression for the above example would be `(known)(1|2|3|4)`.
This is a non-greedy expression which will match only the known items.

## Results

```
diff@trash-boat:~/repo/go-regex-timing $ make
go build regex-speed-test.go
diff@trash-boat:~/repo/go-regex-timing $ ./regex-speed-test
Running test 100000 times.
Smart Run -- compilation, fail, success, total
37528 17635 9903 65573
Dumb Run -- compilation, fail, success, total
29586 35771 3847 69420
diff@trash-boat:~/repo/go-regex-timing $ ./regex-speed-test
Running test 100000 times.
Smart Run -- compilation, fail, success, total
36129 16874 9715 63220
Dumb Run -- compilation, fail, success, total
28615 35020 3725 67573
```
_*Note:* that the above times are all represented in nanoseconds._

The above is an example of the output and a decent data point we can learn off of;
 - The compilation of the "smart" regex is higher than a "dumb" one, this is as expected.
   The compilation of the regex does not significantly make a difference though depending
   on how the structure of the code works, as the hit will be paid up from during execution,
   in most cases.
 - Returning a fail, asserting that something _does not_ match, is faster in a "smart"
   regex.
 - Returning a success, asserting that something _does_ match, is faster in
  the "dumb" regex.
 - The total times are not necessarily significant, though interesting to watch.

My original thought was that the different would not matter - since I was assuming
the compilation of the regular expression would result in a similar simplified tree
being used. So a "dumb" regular expression may have actually compiled slower - this
ended up not being true.

## Conclusion

The results show that if you are attempting to use a regular expression for a blacklist,
basically optimizing for misses, a smarter regular expression may benefit your
continued execution time. If you are attempting to implement a whitelist, or optimizing
for hits, you may actually benefit from a dumb regular expression. In the end,
unless you are performing constant checks at an extremely high number of times,
you likely will not see any difference since there have all been measured in nanoseconds.


## License

    Copyright 2017 Tim 'diff' Strazzere <strazz@gmail.com>

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
