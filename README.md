# go-bot

Description:

When this application is deployed to the cloud or run locally, it functions as a versatile system. It can efficiently manage incoming messages from Telegram, provide responses, and act as an API endpoint for an online store. Customers can use this application to make purchases and receive a thank you message in return.

The Telegram bot serves as a platform where users can leave reviews. It operates in two different modes. In the first mode, hardcoded mode (activated when the USE_GPT flag is set to false), the bot simply responds with a thank you message. However, when the USE_GPT flag is set to true, the user's messages are analyzed by Chat GPT. Based on this analysis, the application determines whether the user is content with their purchase, has concerns, or is dissatisfied. It then provides an appropriate response in the Telegram chat.

To start locally run this command:

USE_GPT={select true or false} API_KEY={enter your telegram API key} GPT_KEY={enter Chat GPT API key} PORT=8081 ./main

Application also deployed to the https://railway.app/ where above env vars set up in railway