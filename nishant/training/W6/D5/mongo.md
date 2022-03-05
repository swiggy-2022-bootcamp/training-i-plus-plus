


start shell
```bash
$ mongo
```


DBs
```bash
show dbs //show all dbs

use crud_test  // use or create  db
```

Read
```bash
db.doctors.find({})
```

insert
```bash
db.doctors.insertOne({
    name:"nishant",
	qualification:"md"
})
```

update
```bash
db.doctors.updateOne(
    {name:"nishant"},
    {   
        $set:{
            qualification:"mbbs"
        }
    }
)
```

delete
```bash
db.doctors.deleteOne(
    {name:"nishant"}
)
```
