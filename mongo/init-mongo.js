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

// insert user seed data
// where id is objectID, name is string
db.user.insertOne({
    id: ObjectId("5f9b3b3b3b3b3b3b3b3b3b3c"),
    name: "user1"
});

// insert todo seed data
// where id is objectID, text is string, done is boolean, userId is objectID
db.todo.insertOne({
    id: ObjectId("5f9b3b3b3b3b3b3b3b3b3b3b"),
    text: "todo1",
    done: false,
    userId: ObjectId("5f9b3b3b3b3b3b3b3b3b3b3c")
}, {
    id: ObjectId("5f9b3b3b3b3b3b3b3b3b3b3d"),
    text: "todo2",
    done: false,
    userId: ObjectId("5f9b3b3b3b3b3b3b3b3b3b3c")
});
