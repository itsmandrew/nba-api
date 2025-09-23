-- PostgreSQL CSV Import Script for Players Data
-- Save this as import_players.sql and run with: psql -d your_database_name -f import_players.sql

-- Enable error reporting
\set ON_ERROR_STOP on

-- Display current timestamp for logging
SELECT 'Starting CSV import at: ' || NOW();

-- Check if the players table exists
SELECT 'Players table exists, proceeding with import...' 
WHERE EXISTS (
    SELECT 1 FROM information_schema.tables 
    WHERE table_name = 'players' AND table_schema = 'public'
);

-- Import the CSV data using \copy (for client-side files)
-- This approach imports to a temp table first, then inserts with proper defaults
CREATE TEMP TABLE temp_players (
    name TEXT,
    year_start TEXT,
    year_end TEXT,
    position TEXT,
    height TEXT,
    weight TEXT,
    birth_date TEXT,
    college TEXT
);

-- UPDATE THIS LINE: Replace 'path/to/your/file.csv' with the actual path to your CSV file
\copy temp_players FROM 'data/player_data.csv' WITH (FORMAT csv, HEADER true, DELIMITER ',');

-- Insert from temp table with proper handling of empty values
INSERT INTO players(name, year_start, year_end, position, height, weight, birth_date, college)
SELECT 
    name,
    CASE WHEN year_start = '' OR year_start IS NULL THEN 0 ELSE year_start::INTEGER END,
    CASE WHEN year_end = '' OR year_end IS NULL THEN 0 ELSE year_end::INTEGER END,
    COALESCE(NULLIF(position, ''), ''),
    COALESCE(NULLIF(height, ''), ''),
    CASE WHEN weight = '' OR weight IS NULL THEN 0 ELSE weight::INTEGER END,
    CASE WHEN birth_date = '' OR birth_date IS NULL THEN '1900-01-01'::DATE 
         ELSE TO_DATE(birth_date, 'Month DD, YYYY') END,
    COALESCE(NULLIF(college, ''), '')
FROM temp_players;


-- Verify the import
SELECT 'Total records imported: ' || COUNT(*) FROM players;

-- Show sample data
SELECT 'Sample of imported data:';
SELECT 
    id,
    name,
    year_start,
    year_end,
    position,
    height,
    weight,
    birth_date,
    college,
    created_at
FROM players 
ORDER BY id 
LIMIT 5;

-- Show some statistics
SELECT 'Import statistics:';
SELECT 
    COUNT(*) as total_records,
    COUNT(DISTINCT name) as unique_names,
    MIN(year_start) as earliest_year,
    MAX(year_end) as latest_year,
    COUNT(*) FILTER (WHERE birth_date IS NOT NULL) as records_with_birth_date,
    COUNT(*) FILTER (WHERE college IS NOT NULL) as records_with_college
FROM players;

-- Display completion message
SELECT 'CSV import completed successfully at: ' || NOW();
