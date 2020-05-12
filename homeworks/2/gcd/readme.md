<h1>GCD calculation algorithms</h1>

<blockquote>
На двух строчках записаны два целых числа.
Вывести их наибольший общий делитель.

Решить задачу разными способами.
1. Алгоритм Евклида через вычитание
2. Алгоритм Евклида через остаток
3. Алгоритм Стейнца через битовые операции
</blockquote>

<h2>Benchmarking</h2>

Benchmarking by build-in golang tool 
``` 
go test -bench=. -benchmem -benchtime=10s
```

First column is name of test (algorithm)<br> 
Second column is number of operations (method calls) that runs during measuring time (10 seconds for current case)<br>
Third column is time cost of one call in nanoseconds<br>

*** For input number x = 1, y = 100 ***<br>
```
BenchmarkSubtract_1_100-4                                                                                                                             	 3154558	      3890 ns/op
BenchmarkRest_1_100-4                                                                                                                                 	48139710	       234 ns/op
BenchmarkBinary_1_100-4                                                                                                                               	20194246	       568 ns/op
```

*** For input number x = 20, y = 30 ***<br>
```
BenchmarkSubtract_20_30-4                                                                                                                             	62767459	       180 ns/op
BenchmarkRest_20_30-4                                                                                                                                 	36719152	       311 ns/op
BenchmarkBinary_20_30-4                                                                                                                               	25695373	       466 ns/op
```

*** For input number x = 2, y = 123456789 ***<br>
```
BenchmarkSubtract_2_123456789-4                                                                                                                       	       5	2154724227 ns/op
BenchmarkRest_2_123456789-4                                                                                                                           	40027131	       300 ns/op
BenchmarkBinary_2_123456789-4                                                                                                                         	 5625045	      2133 ns/op
```

*** For input number x = 12, y = 1234567890 ***<br>
```
BenchmarkSubtract_12_1234567890-4                                                                                                                     	       3	3585204144 ns/op
BenchmarkRest_12_1234567890-4                                                                                                                         	38918703	       309 ns/op
BenchmarkBinary_12_1234567890-4                                                                                                                       	 5385480	      2221 ns/op
```

*** For input number x = 123456789, y = 9876543210 ***<br>
```
BenchmarkSubtract_123456789_9876543210-4                                                                                                              	     244	  48672735 ns/op
BenchmarkRest_123456789_9876543210-4                                                                                                                  	30615730	       395 ns/op
BenchmarkBinary_123456789_9876543210-4                                                                                                                	 4093018	      2933 ns/op
```

*** For input number x = 997, y = 1073676287 ***<br>
```
BenchmarkSubtract_997_1073676287-4                                                                                                                    	     318	  37635692 ns/op
BenchmarkRest_997_1073676287-4                                                                                                                        	19117738	       625 ns/op
BenchmarkBinary_997_1073676287-4                                                                                                                      	 4718510	      2540 ns/op
```

*** For input number x = 8564657687654654657, y = 1298074214633706835075030044377087 ***<br>
```
BenchmarkRest_8564657687654654657_1298074214633706835075030044377087-4                                                                                	 3742844	      3184 ns/op
BenchmarkBinary_8564657687654654657_1298074214633706835075030044377087-4                                                                              	 1000000	     10520 ns/op
```

*** For input number x = 162259276829213363391578010288127, y = 123426017006182806728593424683999798008235734137469123231828679 ***<br>
```
BenchmarkRest_162259276829213363391578010288127_123426017006182806728593424683999798008235734137469123231828679-4                                     	 1000000	     11207 ns/op
BenchmarkBinary_162259276829213363391578010288127_123426017006182806728593424683999798008235734137469123231828679-4                                   	  615214	     19162 ns/op
```

*** For input number x = 162259276829213363391578010288126, y = 123426017006182806728593424683999798008235734137469123231828678 ***<br>
```
BenchmarkRest_162259276829213363391578010288126_123426017006182806728593424683999798008235734137469123231828678-4                                     	 1000000	     11095 ns/op
BenchmarkBinary_162259276829213363391578010288126_123426017006182806728593424683999798008235734137469123231828678-4                                   	  664572	     17868 ns/op
```

*** For input number x = 608281864034267560872252163321295376887552831379210240000000000, y = 30414093201713378043612608166064768844377641568960512000000000000 ***<br>
```
BenchmarkRest_608281864034267560872252163321295376887552831379210240000000000_30414093201713378043612608166064768844377641568960512000000000000-4     	24187141	       489 ns/op
BenchmarkBinary_608281864034267560872252163321295376887552831379210240000000000_30414093201713378043612608166064768844377641568960512000000000000-4   	 2489271	      4836 ns/op
```