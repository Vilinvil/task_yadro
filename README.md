Ильинский Владислав. 

### Задача на позицию «Инженер-стажёр по разработке ПО для систем хранения данных на Go»

тг: https://t.me/Vilin0

Условие задачи лежит в файле [task.md](task.md)

### Для запуска решения
Можно воспользоваться make или скопировать соответствующие команды из [Makefile](Makefile)
```shell
make run
```

### Для запуска тестов
```shell
make test
```

### Для запуска линтера
```shell
make lint
```

### Объяснение решения
Все три метода, приведенные мной основываются на одной идее. Количество шаров в каждом контейнере из-за свапов не 
меняется. Подсчитываем количество шаров в каждом контейнере и количество шаров каждого цвета. Если для каждого
контейнера не найдется соответствующие количество шаров определенного цвета, то не можем отсортировать свапами,
иначе можем.

Три подтипа CanMatrixSortWithCycles, CanMatrixSortWithMap, CanMatrixSortWithSort отличаются по способу определения 
равенства. Считывание всех входных значений занимает O(n*n) времени.

1. CanMatrixSortWithCycles. Сравнивает на циклах кол-во итераций n + n-1 + n-2 + ... + 1 = n * n/2.

Время: O(n*n + n*n/2) (оставлял везде с константами для сравнения)

Память: O(2*n)

2. CanMatrixSortWithMap. Записывает в мапу ключ - количество шаров в контейнере,
значение - количество контейнеров с одинаковым количеством шаров.

Время: O(n*n + 1) (В амортизированном случае), O(n*n + n*n) (В худшем случае, если все записи в одном бакете)

Память: O(n + n / Коэффициент заполненности мапы)

3. CanMatrixSortWithSort. Сортирует количество шаров в контейнерах и сравнивает сортированные массивы. 
Если их все элементы равны, то результат верный.

Использовал slices.SortFunc вместо sort.Sort, т.к. в документации к последней написано: 
"in many situations, the newer [slices.SortFunc] function is more ergonomic and runs faster."

Судя по документации в go используется pdqsort на основе quicksort, который в худшем случае работает за O(n*k). Где k - 
количество уникальных элементов. В аппроксимированном случае O(n*log(n)). Не использует доп. память. Тогда

Время: O(n*n + n*log(n)) (В амортизированном случае), O(n*n + n*k) (В худшем случае)

Память: O(2*n)

### Результаты 
Написал [бенчмарк](can_matrix_sort_test.go), запустить можно так:
```shell
make bench
```

Бенчмарк прогоняется на одном варианте входных данных, результаты представлены в таблице: 

| Название                            | количество прогонов  | время работы на операцию  | выделенная память на операцию  | кол-во аллокаций памяти на операцию  |
|-------------------------------------|----------------------|---------------------------|--------------------------------|--------------------------------------|
| BenchmarkCanMatrixSortWithCycles-16 | 875329               | 1834 ns/op                | 152 B/op                       | 15 allocs/op                         |
| BenchmarkCanMatrixSortWithMap-16    | 755131               | 1802 ns/op                | 128 B/op                       | 14 allocs/op                         |
| BenchmarkCanMatrixSortWithSort-16   | 901951               | 1784 ns/op                | 152 B/op                       | 15 allocs/op                         |

Результаты разных вариантов сопоставимы по времени работы, вариант CanMatrixSortWithMap требует меньше памяти.
В зависимости от задачи можно было бы на реальных входных данных и на машинах приближенным к рабочим прогнать бенчмарки.
Т.к. в реальном мире есть кеши процессора, оптимизации компилятора, разная работа инструкций на разных процессорах и
т.д. бенчмарки лучше бы показывали картину под конкретно наши условия. Я оставил в мейне вариант с циклами, т.к. он
самый простой и всегда работает за одно и то же время.
