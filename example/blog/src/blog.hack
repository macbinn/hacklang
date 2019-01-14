dao = require(`src/dao`)

web.route({
  url: `^/$`,
  get: ctx => {
    ctx.static(`static/index.html`)
  }
})

web.route({
  url: `^/api/blogs/$`,
  get: ctx => {
    ctx.json(dao.blogs.all())
  },
  post: ctx => {
    if ctx.user {
      ctx.jsonBody.authorId = ctx.user.id
      dao.blogs.new(ctx.jsonBody)
      ctx.json(true)
    }
  }
})

web.route({
  url: `^/api/blogs/(\d+)$`,
  post: (ctx, id) => {
    if ctx.user {
      ok = dao.blogs.del(id)
      ctx.json(ok)
    }
  }
})

require(`src/user`)

web.run(`:8080`)