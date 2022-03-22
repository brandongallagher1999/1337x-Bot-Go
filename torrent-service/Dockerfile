#Version Locking Node
FROM node:17.4

COPY . .

EXPOSE 80
EXPOSE 443

RUN yarn install --frozen-lockfile
RUN yarn global add nest

CMD ["yarn", "start"]