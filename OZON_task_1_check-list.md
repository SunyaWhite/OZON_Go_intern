# Чек-лист для тестирования консольного приложения

### Описание приложения:

На вход ожидаются данные. В переданных данных ПО заменяет все цифры на знак "*" и выводит результат. Работа должна осуществляться на операционных системах семейства unix.

### Допущения

Поскольку в условиях явно не указано, в каком формате данное приложение будет принимать данные и в каком возвращать, то я сделаю несколько допущений, которые также буду учитывать при составлении чек-листа:

 1. Базовым сценарием использования данного приложения будет передача входных данных без каких-либо опций. Данные представляют из себя символьные комбинации, разделенные пробелом. Результат выполнения программы будет выводиться в терминал через пробел.  
    + Если в программу был передан null-овый аргумент или пустая строка, то приложение должно уведомить об этом пользователя, т. е. вывести соответствующее сообщение.
    + Если в конце программы есть лишние пробелы, то программа должна их проигнорировать;
       
2. Приложение в качестве входных данных может принимать файл, откуда будет производиться считывание данных. Для этого консольное приложение будет использовать опцию -i или --input . Как параметр будет передаваться абсолютный путь к файлу. В независимости от формат, программа должна заменить все встречаемые символы цифр на символ «*».
    + Если файл пустой (т. е. отсутствуют символы), то поведение аналогично пустому вводу через терминал.
    + Если файла по заданному пути не существует или запускающий программу пользователь не имеет доступа к запрашиваемому файлу, то приложение должно уведомить об этом пользователя, т. е. вывести соответствующее сообщение с описание возникшей ошибки в терминал;

3. Приложение имеет возможность записывать выходные данные в файл. Для этого консольное приложение использовать опцию -o или --output. В качестве параметра используется абсолютный путь к существующему и создаваемому файлу. 
    + Также существует опция -f или --force, которая будет свидетельствовать о необходимости перезаписать файл с выходными данными, если по указанному пути уже существует файл.
    + В случае, если запускающий программу пользователь не имеет доступа к директории или не имеет права записи файлов, в которой должен быть создан файл, или перезаписывает файл, который считывает, то приложение должно уведомить об этом пользователя, т. е. вывести соответствующее сообщение с описание возникшей ошибки  в терминал.
    + В случае, если файл существует, но не была использована опция -f или --force, необходимо вывести сообщение в терминал для пользователя, что для перезаписи файла нужно использовать данную опцию.

### Чек-лист консольного приложения:

1. Проверка изменение входной одиночной строки. Через терминал запускаем программу с данными: **./name_of_console_app 123**. В результате выполнения в терминал должно быть выведено ***.
2. Проверка на множественный ввод. Через терминал запускаем программу с данными: **./name_of_console_app 123 456 789**. В результате выполнения в терминал должно быть выведено *** *** ***.
3. Проверка на пустой ввод. Через терминал запускаем программу с данными (пустой пробел): **./name_of_console_app** . В результате выполнения в терминал должно быть выведено сообщение об ошибке пустого ввода.
4. Проверка на лишний пробел. Через терминал запускаем программу с данными: **./name_of_console_app 123 456 789**  . В результате выполнения в терминал должно быть выведено *** *** ***.
5. Проверка на смешанный ввод на разных позициях. Через терминал запускаем программу с данными (в конце лишний пробел): **./name_of_console_app ab1 a2c 3bc**. В результате выполнения в терминал должно быть выведено ab* a*c *bс. Необходимо убедиться, что программа корректно отрабатывает при смешанном вводе.
6. Проверка замены избранных символов. Необходимо подготовить строку, которая будет включать в себя заглавные и прописные символы нескольких языков (например: русский, английский), цифры, специальные символы. Ввод производится через терминал. Необходимо убедиться, что только символы, представляющие десятичные цифры были заменены на  символ «*».
7. Проверка на входные данные в виде файла. Необходимо подготовить тестовый файл, который будет передаваться в качестве аргумента консольной программе. Далее необходимо запустить приложение:  **./name_of_console_app -i /path/to/test/file.txt** (расширение файла не имеет значения). Необходимо убедиться, что программа правильно считала файл и сделала замены в корректных местах. Повторить этот же пункт для опции **-- input**. Убедиться, что она также работает.
8. Проверка на пустой файл. Необходимо создать пустой файл в директории, куда пользователь имеет доступ. Далее необходимо запустить приложение: **./name_of_console_app -i /path/to/test/empty/file.txt**(Будем считать, что пункт 7 пройден успешно). В результате выполнения в терминал должно быть выведено сообщение об ошибке пустого ввода.
9. Проверка на ссылку на несуществующий файл. Необходимо запустить приложение: **./name_of_console_app -i /path/to/no/file.txt** (Будем считать, что пункт 7 пройден успешно). В результате выполнения в терминал должно быть выведено сообщение о ссылке на отсутствующий файл.
10. Проверка на запрещенный доступ входного файла. Создаем непустой файл в директории. Необходимо запустить приложение под учетной записью пользователя, который не имеет доступ к файлу: **./name_of_console_app -i /path/to/restricted/file.txt** (Будем считать, что пункт 7 пройден успешно). В результате выполнения в терминал должно быть выведено сообщение о запрещенном доступе к файлу.
11. Проверка на вывод результата в файл. Через терминал запускаем программу с данными: **./name_of_console_app ab1 a2c 3bc -o /path/to/output/file.txt** (расширение файла не имеет значение). В результате выполнения в файле должно быть записано ab* a*c *bс. Необходимо убедиться, что программа умеет выводить результат в файл. Провести эту же проверку с опцией --output.
12. Проверка на перезапись файла. Убедимся, что в директории **/path/to/output** уже есть файл **file.txt**. Через терминал запускаем программу с данными: **./name_of_console_app ab1 a2c 3bc -f -o /path/to/output/file.txt**. В результате выполнения в файле **file.txt** должно быть записано ab* a*c *bс. Необходимо убедиться, что программа умеет перезаписывать результат в уже существующем файле. Повторить пункт тестирования для опции **--force**.
13. Проверка на перезапись файла. Убедимся, что в директории /path/to/output уже есть файл file.txt. Через терминал запускаем программу с данными без опции **-f**: **./name_of_console_app ab1 a2c 3bc -o /path/to/output/file.txt**. В результате выполнения в в терминал должно быть выведено сообщение о невозможности перезаписать файл.
14. Проверка на запрещенный доступ выходного файла. Убедимся, что пользователь не имеет доступ к директории **/path/to/restricted/** . Через терминал запускаем программу с данными под учетной записью проверенного данного пользователя: **./name_of_console_app ab1 a2c 3bc -o /path/to/restricted/file.txt**. В результате выполнения в в терминал должно быть выведено сообщение о запрещенном доступе.
15. Проверка на работу с файлами. Необходимо убедиться, что пользователь имеет доступ к директориям **/path/to/output/ и /path/to/input/**. Через терминал запускаем программу с данными: **./name_of_console_app ab1 a2c 3bc -i /path/to/input/file.txt -o /path/to/output/file.txt**. Необходимо убедиться, что программа корректно отработала при чтении из файла и записи результата в файл. Проверить этот пункт при использовании опций --output и --input и перестановке опций с входными параметрами и выходными параметрами.
16. Проверка на перезапись входного файла. Необходимо убедиться, что пользователь имеет доступ к директории **/path/to/test/**. Через терминал запускаем программу с данными: **./name_of_console_app ab1 a2c 3bc -i /path/to/test/file.txt -o /path/to/test/file.txt**. Необходимо убедиться, что программа вывела сообщение об ошибке в терминал.
17. Пункты 1-16 необходимо протестировать на компьютерах под разными операционными системами семейства **Unix** на которых планируется использование данной программы. В качестве примера: **FreeBSD, Debian, Arch, Mac OS и тд..**