const { PrismaClient } = require ('@prisma/client')
const bcrypt = require('bcryptjs');

const prisma = new PrismaClient()

async function main() {
  // Seed Admin User
  const hashedPassword = await bcrypt.hash(process.env.ADMIN_USER_PASSWORD, 10);
  const upsertUser = await prisma.user.upsert({
    where: {
      email: 'admin@reborn.game',
    },
    update: {},
    create: {
      email: 'admin@reborn.game',
      name: 'Admin',
      password: hashedPassword,
    },
  })
  console.log(`Created Admin User ${upsertUser.email}/${process.env.ADMIN_USER_PASSWORD}.`)
}

main()
  .then(async () => {
    await prisma.$disconnect()
  })
  .catch(async (e) => {
    console.error(e)
    await prisma.$disconnect()
    process.exit(1)
  })