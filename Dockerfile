FROM golang:latest as builder
WORKDIR /root/
COPY . .

# Install NodeJS (https://github.com/nodesource/distributions#installation-instructions)
RUN apt-get update
RUN apt-get install -y ca-certificates curl gnupg build-essential
RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg

RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_18.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list

RUN apt-get update
RUN apt-get -y install nodejs

# Install yarn
RUN npm install -g yarn
RUN npm install -g typescript
RUN npm i -g tsx

# Build site
RUN cd embedg-site && yarn install && yarn build && cd ..

# Build app
RUN cd embedg-app && yarn install && yarn build && cd ..

RUN cd embedg-server && go build --tags "embedg-app embedg-site" && cd .. 

RUN apt-get update
RUN apt-get install -y ca-certificates gnupg build-essential
RUN chmod +x ./embedg-server
EXPOSE 8080
CMD ./embedg-server migrate postgres up; ./embedg-server server 
