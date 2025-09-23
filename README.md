# NBA Players Database Setup

Quick setup guide for importing NBA players CSV data using Goose migrations and PostgreSQL.

## Prerequisites

- PostgreSQL running
- Goose installed: `go install github.com/pressly/goose/v3/cmd/goose@latest`
- Your CSV file ready

## Setup Steps

### 1. Create Database

```bash
createdb nba
```

### 2. Run Migration

```bash
goose -dir sql/migrations postgres "<db_url_string>" up
```

### 3. Update Import Script

In `sql/scripts/import_players.sql`, change this line:

```sql
\copy temp_players FROM '/absolute/path/to/your/players.csv' WITH (FORMAT csv, HEADER true, DELIMITER ',');
```

### 4. Import Data

```bash
psql -d nba -f sql/scripts/import_players.sql
```

### 5. Verify

```bash
psql -d nba -c "SELECT COUNT(*) FROM players;"
```

## That's it

Your CSV columns should be: `name,year_start,year_end,position,height,weight,birth_date,college`

**Note**: Update the file path in step 4 with your actual CSV location.
