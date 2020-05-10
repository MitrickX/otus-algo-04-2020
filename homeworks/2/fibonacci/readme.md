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
BenchmarkRecursion3-4      	1000000000	         9.87 ns/op
BenchmarkIteration3-4      	1000000000	         3.08 ns/op
BenchmarkGoldenRatio3-4    	358540016	        33.4 ns/op
BenchmarkMatrix3-4         	417119046	        28.8 ns/op
```

*** For input number n = 7 ***<br>
```
BenchmarkRecursion7-4      	140499646	        85.3 ns/op
BenchmarkIteration7-4      	1000000000	         5.53 ns/op
BenchmarkGoldenRatio7-4    	337825282	        35.5 ns/op
BenchmarkMatrix7-4         	288566896	        41.6 ns/op
```

*** For input number n = 15 ***<br>
```
BenchmarkRecursion15-4     	 2857064	      4251 ns/op
BenchmarkIteration15-4     	1000000000	        11.0 ns/op
BenchmarkGoldenRatio15-4   	306485979	        39.6 ns/op
BenchmarkMatrix15-4        	210470968	        57.1 ns/op
```

*** For input number n = 20 ***<br>
```
BenchmarkRecursion20-4     	  244970	     48823 ns/op
BenchmarkIteration20-4     	832726554	        14.1 ns/op
BenchmarkGoldenRatio20-4   	294445242	        40.9 ns/op
BenchmarkMatrix20-4        	176016022	        68.2 ns/op
```

*** For input number n = 40 ***<br>
```
BenchmarkRecursion40-4     	      15	 738446394 ns/op
BenchmarkIteration40-4     	382516539	        31.1 ns/op
BenchmarkGoldenRatio40-4   	274999706	        43.7 ns/op
BenchmarkMatrix40-4        	146647626	        82.2 ns/op
```

*** For input number n = 93 ***<br>
```
BenchmarkIteration93-4     	221186491	        54.6 ns/op
BenchmarkGoldenRatio93-4   	268371609	        44.8 ns/op
BenchmarkMatrix93-4        	132825297	        91.1 ns/op
```

*** For input number n = 184 ***<br>
```
BenchmarkIteration184-4    	 2060546	      5956 ns/op
BenchmarkMatrix184-4       	 1478784	      8023 ns/op
```

*** For input number n = 300 ***<br>
```
BenchmarkIteration300-4    	 1000000	     10078 ns/op
BenchmarkMatrix300-4       	 1419962	      8479 ns/op
```

*** For input number n = 501 ***<br>
```
BenchmarkIteration501-4    	  668965	     17704 ns/op
BenchmarkMatrix501-4       	 1251012	      9602 ns/op
```

*** For input number n = 999 ***<br>
```
BenchmarkIteration999-4    	  303046	     39618 ns/op
BenchmarkMatrix999-4       	  948970	     13453 ns/op
```

*** For input number n = 1001 ***<br>
```
BenchmarkIteration1001-4   	  294744	     39443 ns/op
BenchmarkMatrix1001-4      	 1000000	     11879 ns/op
```

*** For input number n = 10000 ***<br>
```
BenchmarkIteration10000-4     	   14773	    811433 ns/op
BenchmarkMatrix10000-4        	   98689	    122551 ns/op
```

*** For input number n = 100000 ***<br>
```
BenchmarkIteration100000-4    	     268	  44672866 ns/op
BenchmarkMatrix100000-4       	    3231	   3670449 ns/op
```

*** For input number n = 1000000 ***<br>
```
BenchmarkIteration1000000-4   	       2	7244338728 ns/op
BenchmarkMatrix1000000-4      	      87	 132377505 ns/op
```