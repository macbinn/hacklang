db = sql.db(`sqlite3`, `example/db.sqlite3`)

user = db.users

print(user.all())
