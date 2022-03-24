#Version Locking Node
FROM node:17.4

COPY . .

EXPOSE 3000 

RUN yarn install --frozen-lockfile

CMD ["yarn", "start"]