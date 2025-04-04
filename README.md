# Welcome to the CHILL-PI | a REST API written in GO
**And Yes! it's a wordplay on REST with a silent A** :)

## How things are connected

- `main.go`: Starts the server and initiates routing
- `routes/`: Defines which endpoints that calls which handlers
- `handlers`: Functions that "calls" requests
- `models/`: Datatypes & database-functions
- `database`: For SQLite or PostgreSQL

## Tech Stack & Structure

- **Web-Framework:** Gin
- **Orm:** GORM
- **DB:** SQLite (maybe migrate to PostgreSQL later)
- **Auth:** JWT

## Structure

### Why this Structure

- `handlers/`: A more GO typical practice that fits better with how it's
  usually done in GO opposed to `controllers`
- `middleware/`: Lets you scale with **JWT** on a later occations,
  without having to clean up things later
- `utils/`: A folder for collecting the small things without having
  to create a new folder for everything
- `logs/`: Perfect for the tmuxinator setup
- `test/`: Gives room to learn testing, without mixing it with production code

```md
go-notes-api/
├── main.go # Starter server
├── go.mod # Go modules
│
├── models/ # Datatyper (Note, User)
│ └── note.go
│
├── handlers/ # HTTP-handlere (business logic)
│ └── note_handler.go
│
├── routes/ # Setter opp routing
│ └── routes.go
│
├── database/ # Init og migrering
│ └── setup.go
│
├── middleware/ # F.eks. JWT-auth senere
│ └── auth.go
│
├── utils/ # Hjelpefunksjoner
│ └── validator.go
│
├── logs/ # For loggfiler (som du har i tmuxinator)
│ └── (server.log, db.log osv.)
│
├── tools/ # Verktøy og scripts
│ └── scanempty/main.go
│
└── test/ # Tester (kan være delt opp videre etter behov)
└── note_test.go
```
