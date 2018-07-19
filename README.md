# Scatterfeed
I am building a web platform for news content similar to news.google.com.
The most significant feature of my webapp will be the ability for users to
post comments externally on the content items and rate the content across
a variety of metrics. Users should be able to engage with news media. The
current polarization of our media signals the need for a stronger feedback
loop between media publications and their readers.

## content-creator
Go microservice for receiving user-generated content. Designed to run on Google
App Engine, a proprietary runtime. This application receives user comments and
ratings and saves them to Datastore.
