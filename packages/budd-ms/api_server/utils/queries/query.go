package queries

const TABLE_EXISTS = `SELECT EXISTS (
    SELECT 1 
    FROM information_schema.tables 
    WHERE table_schema = 'public' 
    AND table_name = 'your_table_name'
);
`

// TODO change this into a more optimised query
const ROW_EXISTS = `SELECT EXISTS (
    SELECT 1
    FROM $1
    WHERE your_column_name = $2
);
`
