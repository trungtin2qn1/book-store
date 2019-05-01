db.createUser(
  {
    user: "admin",
    pwd: "bs_1234",
    roles: ["readWrite", "dbAdmin"]
  }
);

db.createCollection("customers");
db.createCollection("books");
db.createCollection("orders");