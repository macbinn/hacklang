dao = require(`src/dao`)

web.route(`^/$`, ctx => {
  ctx.static(`static/index.html`)
}, ctx => {

})

web.route(`^/api/blogs/$`, ctx => {
  ctx.json(dao.blogs.all())
}, ctx => {
  dao.blogs.new(ctx.jsonBody)
  ctx.json(true)
})

web.run(`:8080`)