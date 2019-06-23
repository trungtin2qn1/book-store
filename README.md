# book-store

Web app for book store

***Build with source code:***

Source code has been written by golang, html, css and js and use mongodb for database

So if you want to run without so many environment and config you can try run with docker below

1. Setup-env:

- Make local db:

  1. Setup docker and docker-compose
  2. Run command: `docker-compose -f docker-compose-db.yml up -d`

- Setup glide:

2. Run with source code:

Run command: `make run`

***Build with docker:***

1. Setup docker:

You can install docker by this turotial: 

> 

2. Setup local db:
   - First you need to setup mongodb:
   You can follow this tutorial for installing: 
   > 
   Mongodb need to be in port: 27017
   - Create a database with name: `book_store`
   - Create a user with name: `admin` (In `book_store` db)
   - And password of user will be: `bs_1234`

3. Pull image from dockerhub:

Pull image from dockerhub:

> 

4. Build and run docker image with `book-store` image

> 
