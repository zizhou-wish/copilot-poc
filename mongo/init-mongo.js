db.createUser({
    user: 'user',
    pwd: 'pass',
    roles: [
        {
            role: 'readWrite',
            db: 'mydb',
        },
    ],
});

// create collection
db.createCollection("user");
db.createCollection("todo");
