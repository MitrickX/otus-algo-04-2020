<h1>Basic containers: compare performance of different array implementations</h1>

<blockquote>
1 задание. Динамические массивы.
Написать метод добавления и удаления элементов:
void add(T item, int index);
T remove(int index); // возвращает удаляемый элемент
по индексу во все варианты динамических массивов:
SingleArray, VectorArray, FactorArray, MatrixArray.

2 задание. Таблица сравнения производительности.
Сравнить время выполнения разных операций
для разных массивов с разным порядком значений.
* сделать обёртку над ArrayList() и тоже сравнить.
Составить таблицу и приложить её скриншот.
Сделать выводы и сформулировать их в нескольких предложениях.
</blockquote>

<h2>Benchmarking</h2>

Benchmarking by build-in golang tool 
``` 
go test -bench=. -benchmem -benchtime=10s -timeout=10m ./...
```

First column is name of test (algorithm)<br> 
Second column is number of operations (method calls) that runs during measuring time (10 seconds for current case)<br>
Third column is time cost of one call in nanoseconds<br>
Forth column is heap memory usage per one operation (method call)<br>
Fifth column is heap memory allocation one operation (method call)<br>

<h3>Grouped by implementation type</h3>   

***Single***
```
BenchmarkAppend100-4           	  306968	     38394 ns/op	   85712 B/op	     201 allocs/op
BenchmarkAdd100-4              	  253224	     46608 ns/op	   87512 B/op	     203 allocs/op
BenchmarkSet100-4              	 2181444	      5543 ns/op	     800 B/op	     100 allocs/op
BenchmarkGet100-4              	 5153558	      2318 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove100-4     	  171607	     68956 ns/op	  128080 B/op	     251 allocs/op
BenchmarkAddRemove100-4        	  157177	     84078 ns/op	  130008 B/op	     253 allocs/op
BenchmarkAppend1000-4          	    4051	   3616967 ns/op	 8478497 B/op	    2001 allocs/op
BenchmarkAdd1000-4             	    4209	   2995177 ns/op	 8494877 B/op	    2003 allocs/op
BenchmarkSet1000-4             	  192181	     55520 ns/op	    8044 B/op	    1000 allocs/op
BenchmarkGet1000-4             	  510429	     22774 ns/op	      16 B/op	       0 allocs/op
BenchmarkAppendRemove1000-4    	    2529	   4859275 ns/op	12713621 B/op	    2501 allocs/op
BenchmarkAddRemove1000-4       	    2437	   4888466 ns/op	12730141 B/op	    2503 allocs/op
BenchmarkAppend10000-4         	      31	 362701403 ns/op	834970412 B/op	   20043 allocs/op
BenchmarkAdd10000-4            	      32	 384402486 ns/op	835133891 B/op	   20041 allocs/op
BenchmarkSet10000-4            	   19498	    604597 ns/op	  122823 B/op	   10001 allocs/op
BenchmarkGet10000-4            	   47577	    246197 ns/op	   17549 B/op	       0 allocs/op
BenchmarkAppendRemove10000-4   	      15	 720061613 ns/op	1252414727 B/op	   25063 allocs/op
BenchmarkAddRemove10000-4      	      15	 681929412 ns/op	1252577908 B/op	   25054 allocs/op
```

***Vector***

```
BenchmarkAppend100-4           	 1639062	      7582 ns/op	   10128 B/op	     111 allocs/op
BenchmarkAdd100-4              	  256513	     48920 ns/op	  100728 B/op	     194 allocs/op
BenchmarkSet100-4              	 2171811	      5485 ns/op	     800 B/op	     100 allocs/op
BenchmarkGet100-4              	 5115669	      2329 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove100-4     	  553922	     23381 ns/op	   14352 B/op	     116 allocs/op
BenchmarkAddRemove100-4        	  172459	     70292 ns/op	  105000 B/op	     199 allocs/op
BenchmarkAppend1000-4          	   35061	    366665 ns/op	  862387 B/op	    1101 allocs/op
BenchmarkAdd1000-4             	    4279	   3214114 ns/op	 8639424 B/op	    1994 allocs/op
BenchmarkSet1000-4             	  211940	     57995 ns/op	    8004 B/op	    1000 allocs/op
BenchmarkGet1000-4             	  513828	     22741 ns/op	       1 B/op	       0 allocs/op
BenchmarkAppendRemove1000-4    	    8641	   1296487 ns/op	 1285425 B/op	    1151 allocs/op
BenchmarkAddRemove1000-4       	    3160	   3775913 ns/op	 9062504 B/op	    2044 allocs/op
BenchmarkAppend10000-4         	     327	  36670345 ns/op	83636107 B/op	   11004 allocs/op
BenchmarkAdd10000-4            	      30	 375552882 ns/op	836605509 B/op	   20033 allocs/op
BenchmarkSet10000-4            	   20640	    578657 ns/op	   84052 B/op	   10000 allocs/op
BenchmarkGet10000-4            	   51872	    227872 ns/op	    1612 B/op	       0 allocs/op
BenchmarkAppendRemove10000-4   	      91	 132895871 ns/op	125380986 B/op	   11505 allocs/op
BenchmarkAddRemove10000-4      	      25	 476088627 ns/op	878350190 B/op	   20529 allocs/op
```

***Factor***

```
BenchmarkAppend100-4            	 2287796	      5269 ns/op	    4848 B/op	     108 allocs/op
BenchmarkAdd100-4               	  173504	     68819 ns/op	  174168 B/op	     203 allocs/op
BenchmarkSet100-4               	 2169901	      5519 ns/op	     800 B/op	     100 allocs/op
BenchmarkGet100-4               	 5146022	      2315 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove100-4      	  668562	     17994 ns/op	    6176 B/op	     112 allocs/op
BenchmarkAddRemove100-4         	  147787	     82362 ns/op	  175336 B/op	     207 allocs/op
BenchmarkAppend1000-4           	  265969	     45338 ns/op	   40720 B/op	    1011 allocs/op
BenchmarkAdd1000-4              	    2425	   4954604 ns/op	16927607 B/op	    2004 allocs/op
BenchmarkSet1000-4              	  211927	     56425 ns/op	    8000 B/op	    1000 allocs/op
BenchmarkGet1000-4              	  508872	     23159 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove1000-4     	   14227	    842046 ns/op	   51600 B/op	    1016 allocs/op
BenchmarkAddRemove1000-4        	    2007	   5927296 ns/op	16938491 B/op	    2009 allocs/op
BenchmarkAppend10000-4          	   19996	    584472 ns/op	  604241 B/op	   10015 allocs/op
BenchmarkAdd10000-4             	      16	 709617265 ns/op	1638143835 B/op	   20128 allocs/op
BenchmarkSet10000-4             	   20172	    594363 ns/op	   80030 B/op	   10000 allocs/op
BenchmarkGet10000-4             	   51592	    231457 ns/op	      11 B/op	       0 allocs/op
BenchmarkAppendRemove10000-4    	     159	  75036042 ns/op	  778961 B/op	   10022 allocs/op
BenchmarkAddRemove10000-4       	      14	 785899294 ns/op	1638252988 B/op	   20132 allocs/op
BenchmarkAppend100000-4         	    1903	   6207191 ns/op	 4994267 B/op	  100018 allocs/op
BenchmarkAdd100000-4            	       1	131493583214 ns/op	160410870856 B/op	  201595 allocs/op
BenchmarkSet100000-4            	    1720	   6943346 ns/op	  802905 B/op	  100058 allocs/op
BenchmarkGet100000-4            	    5008	   2312491 ns/op	     997 B/op	      19 allocs/op
BenchmarkAppendRemove100000-4   	       2	7328582295 ns/op	 6392328 B/op	  100027 allocs/op
BenchmarkAddRemove100000-4      	       1	138985683309 ns/op	160411962576 B/op	  201701 allocs/op
```

***Matrix***
```
BenchmarkAppend100-4            	 1837200	      6531 ns/op	    2720 B/op	     105 allocs/op
BenchmarkAdd100-4               	  170934	     70193 ns/op	    4600 B/op	     109 allocs/op
BenchmarkSet100-4               	 1560444	      7705 ns/op	     800 B/op	     100 allocs/op
BenchmarkGet100-4               	 2338707	      5112 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove100-4      	  165130	     72832 ns/op	    2720 B/op	     105 allocs/op
BenchmarkAddRemove100-4         	   87514	    137487 ns/op	    4600 B/op	     109 allocs/op
BenchmarkAppend1000-4           	  182242	     65459 ns/op	   26736 B/op	    1026 allocs/op
BenchmarkAdd1000-4              	    2066	   5680531 ns/op	   28568 B/op	    1029 allocs/op
BenchmarkSet1000-4              	  154704	     77704 ns/op	    8000 B/op	    1000 allocs/op
BenchmarkGet1000-4              	  230043	     51110 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove1000-4     	    2023	   5932911 ns/op	   26736 B/op	    1026 allocs/op
BenchmarkAddRemove1000-4        	    1032	  11532237 ns/op	   28568 B/op	    1029 allocs/op
BenchmarkAppend10000-4          	   18289	    656788 ns/op	  266480 B/op	   10209 allocs/op
BenchmarkAdd10000-4             	      20	 554365380 ns/op	  268312 B/op	   10212 allocs/op
BenchmarkSet10000-4             	   14719	    812878 ns/op	   80018 B/op	   10000 allocs/op
BenchmarkGet10000-4             	   23416	    512865 ns/op	      11 B/op	       0 allocs/op
BenchmarkAppendRemove10000-4    	      19	 583508480 ns/op	  266480 B/op	   10209 allocs/op
BenchmarkAddRemove10000-4       	       9	1136236852 ns/op	  268312 B/op	   10212 allocs/op
BenchmarkAppend100000-4         	    1669	   7232139 ns/op	 2656760 B/op	  102012 allocs/op
BenchmarkAdd100000-4            	       1	55022687617 ns/op	 2658592 B/op	  102015 allocs/op
BenchmarkSet100000-4            	    1291	   9228598 ns/op	  802059 B/op	  100079 allocs/op
BenchmarkGet100000-4            	    2133	   5444658 ns/op	    1245 B/op	      47 allocs/op
BenchmarkAppendRemove100000-4   	       1	59713999670 ns/op	 2656752 B/op	  102012 allocs/op
BenchmarkAddRemove100000-4      	       1	114822470889 ns/op	 2658592 B/op	  102015 allocs/op
```

***List***
```
BenchmarkAppend100-4            	 2170803	      5555 ns/op	    4912 B/op	     109 allocs/op
BenchmarkAdd100-4               	  726964	     16341 ns/op	    4920 B/op	     110 allocs/op
BenchmarkSet100-4               	 2084235	      5787 ns/op	     800 B/op	     100 allocs/op
BenchmarkGet100-4               	 5159324	      2314 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove100-4      	  697274	     16931 ns/op	    4912 B/op	     109 allocs/op
BenchmarkAddRemove100-4         	  427010	     27817 ns/op	    4920 B/op	     110 allocs/op
BenchmarkAppend1000-4           	  260896	     47705 ns/op	   40784 B/op	    1012 allocs/op
BenchmarkAdd1000-4              	   14788	    812560 ns/op	   40792 B/op	    1013 allocs/op
BenchmarkSet1000-4              	  201110	     59270 ns/op	    8000 B/op	    1000 allocs/op
BenchmarkGet1000-4              	  509500	     23147 ns/op	       0 B/op	       0 allocs/op
BenchmarkAppendRemove1000-4     	   14652	    817045 ns/op	   40784 B/op	    1012 allocs/op
BenchmarkAddRemove1000-4        	    7525	   1579288 ns/op	   40792 B/op	    1013 allocs/op
BenchmarkAppend10000-4          	   16746	    722019 ns/op	  906001 B/op	   10021 allocs/op
BenchmarkAdd10000-4             	     160	  74176327 ns/op	  906010 B/op	   10022 allocs/op
BenchmarkSet10000-4             	   19035	    636633 ns/op	   80047 B/op	   10000 allocs/op
BenchmarkGet10000-4             	   51576	    231247 ns/op	      17 B/op	       0 allocs/op
BenchmarkAppendRemove10000-4    	     160	  74192201 ns/op	  906002 B/op	   10021 allocs/op
BenchmarkAddRemove10000-4       	      81	 146781201 ns/op	  906011 B/op	   10022 allocs/op
BenchmarkAppend100000-4         	    1030	  11488449 ns/op	10047390 B/op	  100031 allocs/op
BenchmarkAdd100000-4            	       2	7263427086 ns/op	10047408 B/op	  100032 allocs/op
BenchmarkSet100000-4            	    1616	   7303653 ns/op	  806219 B/op	  100061 allocs/op
BenchmarkGet100000-4            	    4938	   2314705 ns/op	    2034 B/op	      20 allocs/op
BenchmarkAppendRemove100000-4   	       2	7239980040 ns/op	10047400 B/op	  100031 allocs/op
BenchmarkAddRemove100000-4      	       1	14561455656 ns/op	10047408 B/op	  100032 allocs/op
``` 

<h3>Grouped by tests</h3>

***Append of N=10^2***
```
Single BenchmarkAppend100-4           	  306968	     38394 ns/op	   85712 B/op	     201 allocs/op
Vector BenchmarkAppend100-4           	 1639062	      7582 ns/op	   10128 B/op	     111 allocs/op
Factor BenchmarkAppend100-4            	 2287796	      5269 ns/op	    4848 B/op	     108 allocs/op
Matrix BenchmarkAppend100-4            	 1837200	      6531 ns/op	    2720 B/op	     105 allocs/op
List   BenchmarkAppend100-4            	 2170803	      5555 ns/op	    4912 B/op	     109 allocs/op
```

***Append of N=10^3***
```
Single BenchmarkAppend1000-4          	    4051	   3616967 ns/op	 8478497 B/op	    2001 allocs/op
Vector BenchmarkAppend1000-4          	   35061	    366665 ns/op	  862387 B/op	    1101 allocs/op
Factor BenchmarkAppend1000-4           	  265969	     45338 ns/op	   40720 B/op	    1011 allocs/op
Matrix BenchmarkAppend1000-4           	  182242	     65459 ns/op	   26736 B/op	    1026 allocs/op
List   BenchmarkAppend1000-4           	  260896	     47705 ns/op	   40784 B/op	    1012 allocs/op
```

***Append of N=10^4***
```
Single BenchmarkAppend10000-4         	      31	 362701403 ns/op	834970412 B/op	   20043 allocs/op
Vector BenchmarkAppend10000-4         	     327	  36670345 ns/op	83636107 B/op	   11004 allocs/op
Factor BenchmarkAppend10000-4          	   19996	    584472 ns/op	  604241 B/op	   10015 allocs/op
Matrix BenchmarkAppend10000-4          	   18289	    656788 ns/op	  266480 B/op	   10209 allocs/op
List   BenchmarkAppend10000-4          	   16746	    722019 ns/op	  906001 B/op	   10021 allocs/op
```

***Append of N=10^5***
```
Factor BenchmarkAppend100000-4         	    1903	   6207191 ns/op	 4994267 B/op	  100018 allocs/op
Matrix BenchmarkAppend100000-4         	    1669	   7232139 ns/op	 2656760 B/op	  102012 allocs/op
List   BenchmarkAppend100000-4         	    1030	  11488449 ns/op	10047390 B/op	  100031 allocs/op
```


***Add of N=10^2*** 
```
Single BenchmarkAdd100-4              	  253224	     46608 ns/op	   87512 B/op	     203 allocs/op
Vector BenchmarkAdd100-4              	  256513	     48920 ns/op	  100728 B/op	     194 allocs/op
Factor BenchmarkAdd100-4               	  173504	     68819 ns/op	  174168 B/op	     203 allocs/op
Matrix BenchmarkAdd100-4               	  170934	     70193 ns/op	    4600 B/op	     109 allocs/op
List   BenchmarkAdd100-4               	  726964	     16341 ns/op	    4920 B/op	     110 allocs/op
```

***Add of N=10^3*** 
```
Single BenchmarkAdd1000-4             	    4209	   2995177 ns/op	 8494877 B/op	    2003 allocs/op
Vector BenchmarkAdd1000-4             	    4279	   3214114 ns/op	 8639424 B/op	    1994 allocs/op
Factor BenchmarkAdd1000-4              	    2425	   4954604 ns/op	16927607 B/op	    2004 allocs/op
Matrix BenchmarkAdd1000-4              	    2066	   5680531 ns/op	   28568 B/op	    1029 allocs/op
List   BenchmarkAdd1000-4              	   14788	    812560 ns/op	   40792 B/op	    1013 allocs/op
```

***Add of N=10^4*** 
```
Single BenchmarkAdd10000-4            	      32	 384402486 ns/op	835133891 B/op	   20041 allocs/op
Vector BenchmarkAdd10000-4            	      30	 375552882 ns/op	836605509 B/op	   20033 allocs/op
Factor BenchmarkAdd10000-4             	      16	 709617265 ns/op	1638143835 B/op	   20128 allocs/op
Matrix BenchmarkAdd10000-4             	      20	 554365380 ns/op	  268312 B/op	   10212 allocs/op
List   BenchmarkAdd10000-4             	     160	  74176327 ns/op	  906010 B/op	   10022 allocs/op
```

***Add of N=10^5***
```
Factor BenchmarkAdd100000-4            	       1	131493583214 ns/op	160410870856 B/op	  201595 allocs/op
Matrix BenchmarkAdd100000-4            	       1	 55022687617 ns/op	 2658592     B/op	  102015 allocs/op
List   BenchmarkAdd100000-4            	       2	  7263427086 ns/op	10047408     B/op	  100032 allocs/op
``` 

***Set of N=10^2*** 
```
Single BenchmarkSet100-4              	 2181444	      5543 ns/op	     800 B/op	     100 allocs/op
Vector BenchmarkSet100-4              	 2171811	      5485 ns/op	     800 B/op	     100 allocs/op
Factor BenchmarkSet100-4               	 2169901	      5519 ns/op	     800 B/op	     100 allocs/op
Matrix BenchmarkSet100-4               	 1560444	      7705 ns/op	     800 B/op	     100 allocs/op
List   BenchmarkSet100-4               	 2084235	      5787 ns/op	     800 B/op	     100 allocs/op
```

***Set of N=10^3*** 
```
Single BenchmarkSet1000-4             	  192181	     55520 ns/op	    8044 B/op	    1000 allocs/op
Vector BenchmarkSet1000-4             	  211940	     57995 ns/op	    8004 B/op	    1000 allocs/op
Factor BenchmarkSet1000-4              	  211927	     56425 ns/op	    8000 B/op	    1000 allocs/op
Matrix BenchmarkSet1000-4              	  154704	     77704 ns/op	    8000 B/op	    1000 allocs/op
List   BenchmarkSet1000-4              	  201110	     59270 ns/op	    8000 B/op	    1000 allocs/op
```

***Set of N=10^4*** 
```
Single BenchmarkSet10000-4            	   19498	    604597 ns/op	  122823 B/op	   10001 allocs/op
Vector BenchmarkSet10000-4            	   20640	    578657 ns/op	   84052 B/op	   10000 allocs/op
Factor BenchmarkSet10000-4             	   20172	    594363 ns/op	   80030 B/op	   10000 allocs/op
Matrix BenchmarkSet10000-4             	   14719	    812878 ns/op	   80018 B/op	   10000 allocs/op
List   BenchmarkSet10000-4             	   19035	    636633 ns/op	   80047 B/op	   10000 allocs/op
```

***Set of N=10^5***
```
Factor BenchmarkSet100000-4            	    1720	   6943346 ns/op	  802905 B/op	  100058 allocs/op
Matrix BenchmarkSet100000-4            	    1291	   9228598 ns/op	  802059 B/op	  100079 allocs/op
List   BenchmarkSet100000-4            	    1616	   7303653 ns/op	  806219 B/op	  100061 allocs/op
``` 

***Get of N=10^2*** 
```
Single BenchmarkGet100-4              	 5153558	      2318 ns/op	       0 B/op	       0 allocs/op
Vector BenchmarkGet100-4              	 5115669	      2329 ns/op	       0 B/op	       0 allocs/op
Factor BenchmarkGet100-4               	 5146022	      2315 ns/op	       0 B/op	       0 allocs/op
Matrix BenchmarkGet100-4               	 2338707	      5112 ns/op	       0 B/op	       0 allocs/op
List   BenchmarkGet100-4               	 5159324	      2314 ns/op	       0 B/op	       0 allocs/op
```

***Get of N=10^3*** 
```

Single BenchmarkGet1000-4             	  510429	     22774 ns/op	      16 B/op	       0 allocs/op
Vector BenchmarkGet1000-4             	  513828	     22741 ns/op	       1 B/op	       0 allocs/op
Factor BenchmarkGet1000-4              	  508872	     23159 ns/op	       0 B/op	       0 allocs/op
Matrix BenchmarkGet1000-4              	  230043	     51110 ns/op	       0 B/op	       0 allocs/op
List   BenchmarkGet1000-4              	  509500	     23147 ns/op	       0 B/op	       0 allocs/op

```

***Get of N=10^4*** 
```
Single BenchmarkGet10000-4            	   47577	    246197 ns/op	   17549 B/op	       0 allocs/op
Vector BenchmarkGet10000-4            	   51872	    227872 ns/op	    1612 B/op	       0 allocs/op
Factor BenchmarkGet10000-4             	   51592	    231457 ns/op	      11 B/op	       0 allocs/op
Matrix BenchmarkGet10000-4             	   23416	    512865 ns/op	      11 B/op	       0 allocs/op
List   BenchmarkGet10000-4             	   51576	    231247 ns/op	      17 B/op	       0 allocs/op
```

***Get of N=10^5***
```
Factor BenchmarkGet100000-4            	    5008	   2312491 ns/op	     997 B/op	      19 allocs/op
Matrix BenchmarkGet100000-4            	    2133	   5444658 ns/op	    1245 B/op	      47 allocs/op
List   BenchmarkGet100000-4            	    4938	   2314705 ns/op	    2034 B/op	      20 allocs/op
``` 

***Append&Remove of N=10^2*** 
```
Single BenchmarkAppendRemove100-4     	  171607	     68956 ns/op	  128080 B/op	     251 allocs/op
Vector BenchmarkAppendRemove100-4     	  553922	     23381 ns/op	   14352 B/op	     116 allocs/op
Factor BenchmarkAppendRemove100-4      	  668562	     17994 ns/op	    6176 B/op	     112 allocs/op
Matrix BenchmarkAppendRemove100-4      	  165130	     72832 ns/op	    2720 B/op	     105 allocs/op
List   BenchmarkAppendRemove100-4      	  697274	     16931 ns/op	    4912 B/op	     109 allocs/op
```

***Append&Remove of N=10^3*** 
```
Single BenchmarkAppendRemove1000-4    	    2529	   4859275 ns/op	12713621 B/op	    2501 allocs/op
Vector BenchmarkAppendRemove1000-4    	    8641	   1296487 ns/op	 1285425 B/op	    1151 allocs/op
Factor BenchmarkAppendRemove1000-4     	   14227	    842046 ns/op	   51600 B/op	    1016 allocs/op
Matrix BenchmarkAppendRemove1000-4     	    2023	   5932911 ns/op	   26736 B/op	    1026 allocs/op
List   BenchmarkAppendRemove1000-4     	   14652	    817045 ns/op	   40784 B/op	    1012 allocs/op
```

***Append&Remove of N=10^4*** 
```
Single BenchmarkAppendRemove10000-4   	      15	 720061613 ns/op	1252414727 B/op	   25063 allocs/op
Vector BenchmarkAppendRemove10000-4   	      91	 132895871 ns/op	125380986 B/op	   11505 allocs/op
Factor BenchmarkAppendRemove10000-4    	     159	  75036042 ns/op	  778961 B/op	   10022 allocs/op
Matrix BenchmarkAppendRemove10000-4    	      19	 583508480 ns/op	  266480 B/op	   10209 allocs/op
List   BenchmarkAppendRemove10000-4    	     160	  74192201 ns/op	  906002 B/op	   10021 allocs/op
```

***Append&Remove of N=10^5***
```
Factor BenchmarkAppendRemove100000-4   	       2	7328582295 ns/op	 6392328 B/op	  100027 allocs/op
Matrix BenchmarkAppendRemove100000-4   	       1	59713999670 ns/op	 2656752 B/op	  102012 allocs/op
List   BenchmarkAppendRemove100000-4   	       2	7239980040 ns/op	10047400 B/op	  100031 allocs/op
``` 

***Add&Remove of N=10^2*** 
```
Single BenchmarkAddRemove100-4        	  157177	     84078 ns/op	  130008 B/op	     253 allocs/op
Vector BenchmarkAddRemove100-4        	  172459	     70292 ns/op	  105000 B/op	     199 allocs/op
Factor BenchmarkAddRemove100-4         	  147787	     82362 ns/op	  175336 B/op	     207 allocs/op
Matrix BenchmarkAddRemove100-4         	   87514	    137487 ns/op	    4600 B/op	     109 allocs/op
List   BenchmarkAddRemove100-4         	  427010	     27817 ns/op	    4920 B/op	     110 allocs/op
```

***Add&Remove of N=10^3*** 
```
Single BenchmarkAddRemove1000-4       	    2437	   4888466 ns/op	12730141 B/op	    2503 allocs/op
Vector BenchmarkAddRemove1000-4       	    3160	   3775913 ns/op	 9062504 B/op	    2044 allocs/op
Factor BenchmarkAddRemove1000-4        	    2007	   5927296 ns/op	16938491 B/op	    2009 allocs/op
Matrix BenchmarkAddRemove1000-4        	    1032	  11532237 ns/op	   28568 B/op	    1029 allocs/op
List   BenchmarkAddRemove1000-4        	    7525	   1579288 ns/op	   40792 B/op	    1013 allocs/op
```

***Add&Remove of N=10^4*** 
```
Single BenchmarkAddRemove10000-4      	      15	 681929412 ns/op	1252577908 B/op	   25054 allocs/op
Vector BenchmarkAddRemove10000-4      	      25	 476088627 ns/op	878350190 B/op	   20529 allocs/op
Factor BenchmarkAddRemove10000-4       	      14	 785899294 ns/op	1638252988 B/op	   20132 allocs/op
Matrix BenchmarkAddRemove10000-4       	       9	1136236852 ns/op	  268312 B/op	   10212 allocs/op
List   BenchmarkAddRemove10000-4       	      81	 146781201 ns/op	  906011 B/op	   10022 allocs/op
```

***Add&Remove of N=10^5***
```
Factor BenchmarkAddRemove100000-4      	       1	138985683309 ns/op	160411962576 B/op	  201701 allocs/op
Matrix BenchmarkAddRemove100000-4      	       1	114822470889 ns/op	 2658592 B/op	  102015 allocs/op
List   BenchmarkAddRemove100000-4      	       1	 14561455656 ns/op	10047408 B/op	  100032 allocs/op
``` 