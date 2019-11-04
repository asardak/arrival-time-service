db.routes.createIndex({location: "2dsphere"});
db.routes.createIndex({expireAt: 1}, {expireAfterSeconds: 0});
