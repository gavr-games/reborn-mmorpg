import { createClient } from "redis";

const getRedisClient = async () => {
  return await createClient({
    url: process.env.REDIS_URL
  })
  .on('error', err => console.log('Redis Client Error', err))
  .connect()
}

export async function GET(
  request: Request,
  { params }: { params: { id: string } }
) {
  const client = await getRedisClient()
  const id = params.id
  const gameObject = await client.get(id)
  await client.disconnect()
  
  if (gameObject === null) {
    return Response.json({error: "Game Object not found"}, {status: 400})
  } else {
    return Response.json(JSON.parse(gameObject))
  }
}

export async function POST(
  request: Request,
  { params }: { params: { id: string } }
) {
  const res = await request.json()
  const client = await getRedisClient()
  const id = params.id

  await client.set(id, JSON.stringify(res))
  await client.disconnect()

  return Response.json({ msg: "Game Object updated." })
}

export async function DELETE(
  request: Request,
  { params }: { params: { id: string } }
) {
  const client = await getRedisClient()
  const id = params.id
  const gameObject = await client.del(id)
  await client.disconnect()
  
  if (gameObject === null) {
    return Response.json({error: "Game Object not found"}, {status: 400})
  } else {
    return Response.json({ msg: "Game Object deleted." })
  }
}
