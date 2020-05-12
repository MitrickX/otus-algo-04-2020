<h1>Fibonacci calculation algorithms</h1>

Task sounds like this:
<blockquote> 
Алгоритмы поиска чисел Фибоначчи макс

* Через рекурсию
* Через итерацию
* По формуле золотого сечения
* Через умножение матриц
* Составить сравнительную таблицу времени работы алгоритмов для разных начальных данных.
</blockquote> 

<h2>Benchmarking</h2>

Benchmarking by build-in golang tool 
``` 
go test -bench=. -benchtime=10s
```

First column is name of test (algorithm)<br> 
Second column is number of operations (method calls) that runs during measuring time (10 seconds for current case)<br>
Third column is time cost of one call in nanoseconds<br>  

*** For input number n = 3 ***<br>
```
BenchmarkRecursion3-4         	1000000000	        10.0 ns/op
BenchmarkIteration3-4         	1000000000	         3.21 ns/op
BenchmarkGoldenRatio3-4       	348888536	        34.4 ns/op
BenchmarkMatrix3-4            	457728854	        26.2 ns/op
```

*** For input number n = 7 ***<br>
```
BenchmarkRecursion7-4         	138505834	        86.6 ns/op
BenchmarkIteration7-4         	1000000000	         5.62 ns/op
BenchmarkGoldenRatio7-4       	331497628	        36.2 ns/op
BenchmarkMatrix7-4            	310064804	        38.6 ns/op
```

*** For input number n = 15 ***<br>
```
BenchmarkRecursion15-4        	 2811890	      4262 ns/op
BenchmarkIteration15-4        	1000000000	        10.6 ns/op
BenchmarkGoldenRatio15-4      	315183194	        38.0 ns/op
BenchmarkMatrix15-4           	233502001	        51.4 ns/op
```

*** For input number n = 20 ***<br>
```
BenchmarkRecursion20-4        	  251836	     47311 ns/op
BenchmarkIteration20-4        	875289663	        13.7 ns/op
BenchmarkGoldenRatio20-4      	302741571	        39.6 ns/op
BenchmarkMatrix20-4           	238603214	        50.3 ns/op
```

*** For input number n = 40 ***<br>
```
BenchmarkRecursion40-4        	      15	 716104790 ns/op
BenchmarkIteration40-4        	387212206	        30.2 ns/op
BenchmarkGoldenRatio40-4      	285399190	        42.0 ns/op
BenchmarkMatrix40-4           	190397982	        63.1 ns/op
```

*** For input number n = 93 ***<br>
```
BenchmarkIteration93-4        	227668830	        53.1 ns/op
BenchmarkGoldenRatio93-4      	276103447	        43.5 ns/op
BenchmarkMatrix93-4           	166518230	        72.0 ns/op
```

*** For input number n = 184 ***<br>
```
BenchmarkIteration184-4       	 2059435	      5843 ns/op
BenchmarkMatrix184-4          	 1724370	      6948 ns/op
```

*** For input number n = 300 ***<br>
```
BenchmarkIteration300-4       	 1214179	      9869 ns/op
BenchmarkMatrix300-4          	 1675186	      7227 ns/op
```

*** For input number n = 501 ***<br>
```
BenchmarkIteration501-4       	  684914	     17354 ns/op
BenchmarkMatrix501-4          	 1462509	      8222 ns/op
```

*** For input number n = 999 ***<br>
```
BenchmarkIteration999-4       	  306033	     38630 ns/op
BenchmarkMatrix999-4          	 1000000	     11046 ns/op
```

*** For input number n = 1001 ***<br>
```
BenchmarkIteration1001-4      	  307359	     38602 ns/op
BenchmarkMatrix1001-4         	 1000000	     10554 ns/op
```

*** For input number n = 10^4 ***<br>
```
BenchmarkIteration10000-4     	   14499	    825826 ns/op
BenchmarkMatrix10000-4        	  118215	    102476 ns/op
```

*** For input number n = 10^5 ***<br>
```
BenchmarkIteration100000-4    	     261	  45499350 ns/op
BenchmarkMatrix100000-4       	    3847	   3113876 ns/op
```

*** For input number n = 10^6 ***<br>
```
BenchmarkIteration1000000-4   	       2	7448070736 ns/op
BenchmarkMatrix1000000-4      	     100	 117631223 ns/op
```

*** For input number = 10^7 ***<br>
```
BenchmarkMatrix10000000-4   	       2	5392057885 ns/op
```