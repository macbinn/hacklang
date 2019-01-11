dao = require(`src/dao`)

web.route(`^/api/users/$`, ctx => {

}, ctx => {
  user = {}
  user.name = ctx.jsonBody.name
  user.passwordHash = hash.bcrypt.generate(ctx.jsonBody.password)
  dao.users.new(user)
  ctx.json(true)
})

web.route(`^/api/signIn$`, ctx => {

}, ctx => {
  query = {}
  query.name = ctx.jsonBody.name
  user = dao.users.find(query)
  ok = hash.bcrypt.compare(user.passwordHash, ctx.jsonBody.password)
  ctx.json(ok)
})