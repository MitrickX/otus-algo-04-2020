<h1>Prime numbers calculation algorithms</h1>

<blockquote>
На первой строчке записано целое число N >= 1. 
Найти количество простых чисел от 1 до N.

Решить задачу разными способами.
1. Через перебор делителей.
2. Несколько оптимизаций перебора делителей
3. Решето Эратосфена со сложностью O(n log log n).
4. Решето Эратосфена с оптимизацией памяти: битовая матрица, по 32 значения в одном int
5. Решето Эратосфена со сложностью O(n)
</blockquote>

<h2>Benchmarking</h2>

Benchmarking by build-in golang tool 
``` 
go test -bench=. -benchmem -benchtime=10s
```

First column is name of test (algorithm)<br> 
Second column is number of operations (method calls) that runs during measuring time (10 seconds for current case)<br>
Third column is time cost of one call in nanoseconds<br>
Forth column is heap memory usage per one operation (method call)<br>
Fifth column is heap memory allocation one operation (method call)<br>   

*** For input number n = 10^3 ***<br>
```
BenchmarkCountBruteForce1000-4                         	    2822	   4144680 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountBruteForceWithBreak1000-4                	   16971	    707029 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSqrtN1000-4                              	   70402	    170461 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSqrtNWithBreak1000-4                     	  227848	     52224 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSieve1000-4                              	 8315396	      1446 ns/op	    1024 B/op	       1 allocs/op
BenchmarkCountSieveWithMemoryOptimization1000-4        	 2336920	      5133 ns/op	     128 B/op	       1 allocs/op
BenchmarkCountFastSieve1000-4                          	 2079508	      5921 ns/op	   11648 B/op	       3 allocs/op
```

*** For input number n = 10^4 ***<br>
```
BenchmarkCountBruteForce10000-4                        	      27	 413675705 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountBruteForceWithBreak10000-4               	     229	  52059652 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSqrtN10000-4                             	    2175	   5495210 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSqrtNWithBreak10000-4                    	   10000	   1124045 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSieve10000-4                             	  864714	     14109 ns/op	   10240 B/op	       1 allocs/op
BenchmarkCountSieveWithMemoryOptimization10000-4       	  215004	     55602 ns/op	    1280 B/op	       1 allocs/op
BenchmarkCountFastSieve10000-4                         	  205532	     59657 ns/op	  102272 B/op	       3 allocs/op
```

*** For input number n = 10^5 ***<br>
```
BenchmarkCountBruteForceWithBreak100000-4              	       3	4090477755 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSqrtN100000-4                            	      67	 174502949 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSqrtNWithBreak100000-4                   	     471	  25400041 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSieve100000-4                            	   59814	    197985 ns/op	  106496 B/op	       1 allocs/op
BenchmarkCountSieveWithMemoryOptimization100000-4      	   22542	    534465 ns/op	   13568 B/op	       1 allocs/op
BenchmarkCountFastSieve100000-4                        	   21032	    572349 ns/op	  966656 B/op	       3 allocs/op
```

*** For input number n = 10^6 ***<br>
```
BenchmarkCountSqrtN1000000-4                           	       2	5515773140 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSqrtNWithBreak1000000-4                  	      19	 617669349 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountSieve1000000-4                           	    4766	   2421093 ns/op	 1007616 B/op	       1 allocs/op
BenchmarkCountSieveWithMemoryOptimization1000000-4     	    2229	   5357511 ns/op	  131072 B/op	       1 allocs/op
BenchmarkCountFastSieve1000000-4                       	    1759	   6727169 ns/op	 9314304 B/op	       3 allocs/op
```

*** For input number n = 10^7 ***<br>
```
BenchmarkCountSieve10000000-4                          	     314	  38266200 ns/op	10002434 B/op	       1 allocs/op
BenchmarkCountSieveWithMemoryOptimization10000000-4    	     216	  55421925 ns/op	 1253376 B/op	       1 allocs/op
BenchmarkCountFastSieve10000000-4                      	     150	  79764534 ns/op	91176965 B/op	       3 allocs/op
```

*** For input number n = 10^8 ***<br>
```
BenchmarkCountSieve100000000-4                         	      18	 626536663 ns/op	100007936 B/op	       1 allocs/op
BenchmarkCountSieveWithMemoryOptimization100000000-4   	      15	 722203081 ns/op	12500992 B/op	       1 allocs/op
BenchmarkCountFastSieve100000000-4                     	      14	 809040242 ns/op	897728512 B/op	       3 allocs/op
```