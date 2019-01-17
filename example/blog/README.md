setup db
```bash
$ sqlite3 blog.sqlite3 < schema.sql
```

run blog server
```bash
$ ../../bin/hacklang src/blog.hack
```

open [http://127.0.0.1:8080/]()